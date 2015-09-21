// KabelDeutschland streaming proxy
// Author: andre@freshest.me
// Date: 23.04.2015
// Version: 1
package main

import (
	"bitbucket.org/gotamer/cfg"
	"flag"
	"fmt"
	"os"
	"kd/config"
	"kd"
	"path/filepath"
)

var (
	help = flag.Bool("h", false, "display help message")
	verbose = flag.Bool("v", false, "enable verbose mode to see more debug output.")
	version = flag.Bool("version", false, "shows the current version number.")
	config_file_param = flag.String("c", "", "specifiy the config.json location, if not next to binary")
	no_check_cert_param = flag.Bool("no-check-certificate", false, "disable root CA check for HTTP requests")
	no_cache = flag.Bool("no-cache", false, "disables playlist caching")
	Config  *config.Config
)

func main() {
	flag.Parse()

	// if environment is not set, print help
	if *help == true {
		fmt.Println("you need to set the following params:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if environment is not set, print help
	if *version == true {
		fmt.Println("KabelDeutschland streaming proxy, http://freshest.me")
		fmt.Println("0.1.3")
		os.Exit(1)
	}

	// load config
	var cfg_file string
	if *config_file_param != "" {
		cfg_file = *config_file_param
	} else {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		cfg_file = dir + "/config.json"
	}
	err := cfg.Load(cfg_file, &Config)
	if err != nil {
		cfg.Save(cfg_file, &Config)
		fmt.Println("\n\tPlease edit your configuration at: ", cfg_file, "\n")
		os.Exit(0)
	}

	// run service
	kd.Service(Config, verbose, no_check_cert_param, no_cache)
}