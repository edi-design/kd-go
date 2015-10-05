// KabelDeutschland streaming proxy
// Author: andre@freshest.me
// Date: 23.04.2015
// Version: 1
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/edi-design/kd-go/kd"
	"github.com/edi-design/kd-go/kd/config"

	"bitbucket.org/gotamer/cfg"
)

var (
	version         = flag.Bool("version", false, "shows the current version number.")
	configFileParam = flag.String("c", "", "specifiy the config.json location, if not next to binary")
	Config          = &config.Config{}
)

const (
	Version = "0.1.3"
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println("KabelDeutschland streaming proxy, http://freshest.me")
		fmt.Println(Version)
		return
	}

	// load config
	var cfgFile string
	if *configFileParam != "" {
		cfgFile = *configFileParam
	} else {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		cfgFile = dir + "/config.json"
	}
	err := cfg.Load(cfgFile, Config)
	if err != nil {
		cfg.Save(cfgFile, Config)
		fmt.Println("\n\tPlease edit your configuration at: ", cfgFile, "\n")
		return
	}

	// run service
	kd.Service(Config)
}
