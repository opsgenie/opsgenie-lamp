// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

/*
Package command creates various OpsGenie API clients:
 	- Alert
 	- Heartbeat
 	- Integration
 	- Policy
And contains command action implementations that uses OpsGenie API clients mentioned above. Commands use OpsGenie Go SDK to send requests to OpsGenie.
*/
package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-lamp/cfg"
	gcli "github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var verbose = false

var logFilePath  = ""

type LogLevel string

const (
	INFO LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
	DEBUG LogLevel = "DEBUG"
)

func printMessage(logLevel LogLevel, message string) {
	if logLevel == DEBUG {
		if verbose {
			log.Println(string(logLevel) + " : " + message)
			logToFile(DEBUG, message)
		}
		logger := client.Config{}.Logger
		if logger != nil {
			logger.Debug(message)
		}
	} else {
		log.Println(string(logLevel) + " : " + message)
		logToFile(logLevel, message)
	}
}

func logToFile(logLevel LogLevel, message string) {
	if logFilePath != "" {
		file, fileError := os.OpenFile(grabLogPath(),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if fileError != nil {
			log.Println(fileError)
		}

		defer file.Close()
		logger := log.New(file, string(logLevel+" : "), log.LstdFlags)
		logger.Println(message)
	}
}

/*
The 'API Key' is the most common parameter for all commands.
It is provided either on command line or on the configuration file.
*/
func grabAPIKey(c *gcli.Context) string {
	if val, success := getVal("apiKey", c); success {
		return val
	}
	apiKey := cfg.Get("apiKey")
	printMessage(DEBUG,"apiKey flag is not set in the command, reading apiKey from config..")
	return apiKey
}

func grabLogPath() string {
	logPath := cfg.Get("logPath")
	date := time.Now().Format("2006-01-02")
	if logPath != "" {
		return logPath +  "/" + date + "-Lamp.log"
	}
	return ""
}

func grabUsername(c *gcli.Context) string {
	if val, success := getVal("user", c); success {
		return val
	}
	return cfg.Get("user")
}

// getVal method returns the given argument names value if command context contains it.
func getVal(argName string, c *gcli.Context) (string, bool) {
	if c.IsSet(argName) {
		arg := c.String(argName)
		isEmpty(argName, arg, c)
		return arg, true
	}
	return "", false
}

// isEmpty method check is the given argument is empty or not. Parameter with empty values are not allowed in opsgenie-lamp
func isEmpty(argName string, arg string, c *gcli.Context) bool {
	var prefix string
	for _, name := range c.FlagNames() {
		if len(name) == 1 {
			prefix = "-"
		} else {
			prefix = "--"
		}
		if strings.EqualFold(arg, prefix+name) {
			fmt.Printf("Value of argument '%s' is empty\n", argName)
			gcli.ShowCommandHelp(c, c.Command.Name)
			os.Exit(1)
		}
	}
	return false
}

func getConfigurations(c *gcli.Context) *client.Config {
	if c.IsSet("v") {
		verbose = true
		printMessage(DEBUG,"Will execute command in verbose mode.")
	}

	readConfigFile(c)
	apiKey := grabAPIKey(c)
	apiURL := cfg.Get("apiUrl")
	if apiURL == "" {
		apiURL = string(client.API_URL)
	}
	logFilePath = grabLogPath()
	if logFilePath == "" {
		log.Println(string(INFO) + ": Logging to file is disabled, To enable Logging to file Please specify logPath in configuration")
	}
	config := client.Config{
		ApiKey:         apiKey,
		OpsGenieAPIURL: client.ApiUrl(apiURL),
	}
	proxyHost := cfg.Get("proxyHost")
	var port = 0
	var err error
	if cfg.Get("proxyPort") != "" {
		port, err = strconv.Atoi(cfg.Get("proxyPort"))
		if err != nil {
			printMessage(DEBUG,"Invalid proxy port.")
		}
	}

	if proxyHost != "" {
		printMessage(DEBUG,"Configuring proxy settings with host " + proxyHost)
		config.ProxyConfiguration = &client.ProxyConfiguration{
			Username: cfg.Get("proxyUsername"),
			Password: cfg.Get("proxyPassword"),
			Host:     proxyHost,
			Protocol: proxyProtocol(cfg.Get("proxyProtocol")),
			Port:     port,
		}
	}
	config.ConfigureLogLevel(cfg.Get("lamp.log.level"))
	if cfg.Get("requestTimeout") != "" {
		timeout, err := strconv.Atoi(cfg.Get("requestTimeout"))
		if err != nil {
			timeout = 0
			printMessage(DEBUG,"Invalid requestTimeout value. Will use default requestTimeout value")
		}
		if timeout != 0 {
			config.RequestTimeout = time.Second * time.Duration(timeout)
		}
	}
	return &config
}

func proxyProtocol(protocol string) client.Protocol {
	switch protocol {
	case "http":
		return client.Http
	case "socks5":
		return client.Socks5
	default:
		return client.Https
	}
}

/*
The 'getAlert' command returns a GetAlertResponse object.
The 'ResultToYaml' function is called whenever "output-format" parameter is
set to yaml.
*/
func resultToYAML(data interface{}) (string, error) {
	d, err := yaml.Marshal(&data)
	if err != nil {
		return "", errors.New("Can not marshal the response into YAML format. " + err.Error())
	}
	return string(d), nil
}

/*
The 'getAlert' command returns a GetAlertResponse object.
The 'ResultToJson' function is called whenever "output-format" parameter is
set to json or not provided. "getAlert" command defaults to JSON format.
Pretty formatting yields an indented style of representation. Pretty formatting
is on when the "pretty" flag is provided alongside.
*/
func resultToJSON(data interface{}, pretty bool) (string, error) {
	if pretty {
		b, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", errors.New("Can not marshal the response into JSON format. " + err.Error())
		}
		return string(b), nil
	}
	b, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("Can not marshal the response into JSON format" + err.Error())
	}
	return string(b), nil
}

func readConfigFile(c *gcli.Context) {
	cfg.Verbose = verbose
	if val, success := getVal("config", c); success {
		cfg.LoadConfigFromGivenPath(val)
	} else {
		cfg.LoadConfiguration()
	}
}

func getListLogsCommandDefaultSize() int{
	return cfg.GetlistLogCommandDefaultBucketSize()
}

