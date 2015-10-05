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
	help            = flag.Bool("h", false, "display help message")
	version         = flag.Bool("version", false, "shows the current version number.")
	configFileParam = flag.String("c", "", "specifiy the config.json location, if not next to binary")
	Config          = &config.Config{}
)

const (
	Version = "0.1.3"
)

func main() {
	flag.Parse()

	// if environment is not set, print help
	if *help {
		fmt.Println("you need to set the following params:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// if environment is not set, print help
	if *version {
		fmt.Println("KabelDeutschland streaming proxy, http://freshest.me")
		fmt.Println(Version)
		os.Exit(1)
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
		os.Exit(0)
	}

	// run service
	kd.Service(Config)
}
