package kd

import (
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/edi-design/kd-go/kd/config"
	"github.com/gorilla/mux"
)

// struct of json configuration
type MainConfig struct {
	Service struct {
		Username string
		Password string
	}
}

var (
	Config           *config.Config
	verbose          = flag.Bool("v", false, "enable verbose mode to see more debug output.")
	noCheckCertParam = flag.Bool("no-check-certificate", false, "disable root CA check for HTTP requests")
	noCache          = flag.Bool("no-cache", false, "disables playlist caching")
)

// main
func Service(ObjConfig *config.Config, verbose *bool) {
	// write config to environment vars
	Config = ObjConfig

	// check credentials
	signIn()

	// init router
	serv := mux.NewRouter()

	subroute := serv.PathPrefix("/").Subrouter()
	subroute.HandleFunc("/", channelHandler).Methods("GET")
	subroute.HandleFunc("/{quality}", channelHandler).Methods("GET")
	subroute.HandleFunc("/{quality}/{format}", channelHandler).Methods("GET")

	// not found handler. fallback if given path is not set up.
	subroute.HandleFunc("/{path:.*}", notFoundHandler)

	// start http-handle
	http.Handle("/", serv)

	fmt.Println("== Listening ...")
	printInterfaces()
	http.ListenAndServe(Config.Service.Listen, nil)
}

// Default route-handler if no configured endpoint matches.
func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	path := params["path"]

	err := errors.New("use known subroutes")
	fmt.Println(err)
	fmt.Printf("path requested: %s:", path)

	writer.WriteHeader(http.StatusNotFound)
}

// Handles the root directory requests.
func channelHandler(writer http.ResponseWriter, request *http.Request) {
	// init vars
	var result config.ChannelList
	var data string

	// debug output
	fmt.Println("== Get channellist")

	// get params
	params := mux.Vars(request)
	format := params["format"]
	quality := params["quality"]

	cache_file, quality_playlist := getQualityInformations(quality)

	request_url := getUrl(config.METHOD_CHANNELLIST)
	body := "{\"initObj\":" + getInitObj() + "," + config.CHANNEL_OBJECT + "}"
	err := httpRequest("POST", request_url, body, &result)

	if err != nil {
		fmt.Printf("could not fetch: %v", err)
	}

	// read cache
	cache_stat, err_cache := os.Stat(cache_file)
	if err_cache == nil && (time.Now().Unix()-cache_stat.ModTime().Unix() <= config.CACHE_LIFETIME) {
		cached_data, _ := ioutil.ReadFile(cache_file)
		data = string(cached_data[:])
	} else {
		// call backend
		data = data + config.M3U_HEAD
		for _, channel := range result {
			link, err_link := getLicensedLink(channel.Files[0].FileID, channel.Files[0].URL, quality_playlist)
			if err_link != nil {
				fmt.Println(err_link.Error())
				data = "This works only if you are using a KabelDeutschland Internet connection.\n" + err_link.Error()
				break
			}
			data = data + fmt.Sprintf(config.M3U_LINE, channel.MediaName, link)
		}

		// write cache file
		if !*noCache {
			ioutil.WriteFile(cache_file, []byte(data), 0644)
		}
	}

	// set header
	if format == "txt" {
		writer.Header().Set("Content-Type", "text/plain")
	} else {
		writer.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
	}

	writer.Header().Set("Status", "200 OK")
	writer.Header().Set("Content-Disposition", "inline; filename=\"playlist.m3u\"")
	writer.Header().Set("Cache-Control", "no-cache, must-revalidate")
	writer.Header().Set("Expies", "Sat, 26 Jul 1997 05:00:00 GMT")

	writer.Write([]byte(data))
}

// get playlist according to requested quality
func getQualityInformations(quality string) (string, string) {
	var quality_file string
	var quality_playlist string

	switch quality {
	case "low":
		quality_file = fmt.Sprintf(config.CACHE_FILE, quality)
		quality_playlist = config.QUALITY_LOW
	case "high":
		quality_file = fmt.Sprintf(config.CACHE_FILE, quality)
		quality_playlist = config.QUALITY_HIGH
	default:
		quality_file = fmt.Sprintf(config.CACHE_FILE, "medium")
		quality_playlist = config.QUALITY_MEDIUM
	}

	return quality_file, quality_playlist
}

// request a link with a valid session
func getLicensedLink(id string, link string, playlist string) (string, error) {
	var result config.LicensedLink

	request_url := getUrl(config.METHOD_LICENSED_LINK)
	body := "{\"initObj\":" + getInitObj() + ",\"mediaFileId\":" + id + ",\"baseLink\":\"" + string(link[:]) + "\"}"
	err := httpRequest("POST", request_url, body, &result)

	if err != nil {
		fmt.Printf("could not fetch: %v", err)
		return "", errors.New("no link")
	}

	resp, err_get := http.Get(result.MainUrl)

	if err_get != nil {
		return "", err_get
	}

	url := resp.Request.URL.String()
	i := strings.LastIndex(url, "/")

	url = url[:i] + "/" + playlist

	return url, nil
}

// concats params to return a valid API url
func getUrl(method string) string {
	return fmt.Sprintf("%s?m=%s&iOSv=%s&Appv=%s", config.GATEWAY, method, config.IOS_VERSION, config.APP_VERSION)
}

// check credentials
func signIn() {
	fmt.Println("== Checking credentials")

	var result config.SignIn

	request_url := getUrl(config.METHOD_SIGNIN)

	body :=
		"{\"initObj\":" +
			getInitObj() +
			",\"userName\":\"" + Config.Service.Username + "\"" +
			",\"password\":\"" + Config.Service.Password + "\"" +
			",\"providerID\":0" +
			"}"

	handleError(fmt.Sprint(body))
	err := httpRequest("POST", request_url, body, &result)

	switch {
	case err != nil, result.LoginStatus != 0:
		handleError(fmt.Sprint("Returned result: %v", result))
		fmt.Println("Credentials are wrong")
		os.Exit(1)
	}

	fmt.Println("done")
}

// print interfaces to know where the proxy is listening
func printInterfaces() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("Can't get interfaces. You have to have at least one network connection.")
		log.Fatal("No interface found")
	}

	for _, addr := range addrs {

		var ip net.IP
		switch v := addr.(type) {
		case *net.IPAddr:
		case *net.IPNet:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		fmt.Println("http://" + ip.String() + Config.Service.Listen)
	}
}

// main helper to call any http request.
func httpRequest(method string, url string, body string, result interface{}) error {
	var (
		req *http.Request
		err error
	)

	// init client, skip cert check, because of some problems with env without root-ca
	tr := &http.Transport{}
	if *noCheckCertParam {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		handleError("= certificate check disabled")
	}
	client := &http.Client{Transport: tr}

	switch method {
	case "GET":
		req, err = http.NewRequest(method, url, nil)
	case "POST":
		req, err = http.NewRequest(method, url, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	default:
		return errors.New(method + " is not a valid method.")
	}

	if err != nil {
		handleError(fmt.Sprintf("could not stat request: %v", err))
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		handleError(fmt.Sprintf("could not fetch: %v", err))
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		handleError(fmt.Sprintf("could not decode response: %v", err))
	}

	return err
}

// handle verbose mode otuput
func handleError(message string) {
	if *verbose {
		fmt.Println(message)
	}
}

// init obj
func getInitObj() string {
	initObject, _ := b64.StdEncoding.DecodeString(config.INIT_OBJECT)
	return string(initObject)
}
