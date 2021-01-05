/*
Copyright 2015 OpsGenie. All rights reserved.
Use of this source code is governed by a Apache Software
license that can be found in the LICENSE file.
*/

//Package cfg reads configurations and provides configuration props to commands.
package cfg

import (
	"fmt"
	"github.com/ccding/go-config-reader/config"
	"os"
	"path/filepath"
)

const (
	confPath        = "LAMP_CONF_PATH"
	sep      string = string(filepath.Separator)
)

var lampConfig *config.Config

// Verbose is an exported variable to determine command is executing verbose mode or not.
var Verbose = false

func printVerboseMessage(message string) {
	if Verbose {
		fmt.Printf("%s\n", message)
	}
}

func lampHome() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Error occurred while getting lamp home path: %s\n", err.Error())
		return ""
	}
	return dir + sep
}

// LoadConfigFromGivenPath method reads configuration file from the given path.
func LoadConfigFromGivenPath(confPath string) {
	printVerboseMessage("Will read configuration from: \n--config " + confPath)
	load(confPath)
}

// LoadConfiguration method tries to find and read the configuration file some specific paths.
func LoadConfiguration() {
	confPath := os.Getenv(confPath)
	if confPath == "" {
		confPath = lampHome() + ".." + sep + "conf" + sep + "opsgenie-integration.conf"
		printVerboseMessage("LAMP_CONF_PATH environment variable is not set. Will try to read config from: \n" + confPath)
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		userHomeDir, _ := os.UserHomeDir()
		confPath = userHomeDir + sep + ".config" + sep + "lamp.conf"
		printVerboseMessage("Could not find the file specified. Will try to read config from: \n" + confPath)
	}
	load(confPath)
}

func load(confPath string) {
	if _, err := os.Stat(confPath); !os.IsNotExist(err) {
		conf := config.NewConfig(confPath)
		conf.Read()
		lampConfig = conf
		configureLog()
	} else {
		printVerboseMessage("Could not read config file: " + err.Error())
	}
}

// Get method returns the configuration properties value according to the key.
func Get(key string) string {
	if lampConfig != nil {
		return lampConfig.Get("", key)
	}
	return ""
}

func configureLog() {
	level := Get("lamp.log.level")
	if level == "" {
		level = "warn"
		printVerboseMessage("Could not get log level from configuration, will use default \"warn\".")
	}
}
