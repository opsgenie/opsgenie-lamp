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
	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-lamp/cfg"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

var verbose = false

func printVerboseMessage(message string) {
	if verbose {
		fmt.Println(message)
	}
	logger := client.Config{}.Logger
	if logger != nil {
		logger.Debug(message)
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
	printVerboseMessage("apiKey flag is not set in the command, reading apiKey from config..")
	return apiKey
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

// TODO: comment out for proxy and connection timeout
/*func proxyConf(host string, port int) (proxy *ogcli.ProxyConfiguration) {
	printVerboseMessage("Configuring proxy settings with host " + host + " and port " + strconv.Itoa(port))
	pc := new(ogcli.ProxyConfiguration)
	pc.Protocol = cfg.Get("proxyProtocol")
	pc.Host = host
	pc.Port = port
	username := cfg.Get("proxyUsername")
	password := cfg.Get("proxyPassword")
	if username != "" && password != "" {
		pc.Username = username
		pc.Password = password
	}
	return pc
}

func connectionConf() (connCfg *ogcli.HTTPTransportSettings) {
	printVerboseMessage("Configuring connection settings..")
	cfg := new(ogcli.HTTPTransportSettings)
	reqTimeout := parseDuration("requestTimeout")
	if reqTimeout != 0 {
		cfg.RequestTimeout = reqTimeout
	}
	connTimeout := parseDuration("connectionTimeout")
	if connTimeout != 0 {
		cfg.ConnectionTimeout = connTimeout
	}
	return cfg
}

func parseDuration(key string) time.Duration {
	strDuration := cfg.Get(key)
	if strDuration != "" {
		printVerboseMessage("Will try to parse [" + key + "] with value [" + strDuration + "] from string to time duration in seconds..")
		var reqTimeout time.Duration
		var err error
		if strings.HasSuffix(strDuration, "s") {
			reqTimeout, err = time.ParseDuration(strDuration)
		} else {
			reqTimeout, err = time.ParseDuration(strDuration + "s")
		}
		if err != nil {
			printVerboseMessage("Could not parse " + strDuration + " from string to time duration, opsgenie client will use default value.")
			return 0
		}
		return reqTimeout
	}
	printVerboseMessage("Could not parse [" + key + "] with value [" + strDuration + "].")
	return 0
}*/

func getConfigurations(c *gcli.Context) *client.Config {
	if c.IsSet("v") {
		verbose = true
		printVerboseMessage("Will execute command in verbose mode.")
	}

	readConfigFile(c)
	apiKey := grabAPIKey(c)
	apiURL := cfg.Get("apiUrl")
	if apiURL == "" {
		apiURL = string(client.API_URL)
	}
	config := client.Config{
		ApiKey:         apiKey,
		OpsGenieAPIURL: client.ApiUrl(apiURL),
	}
	proxyHost := cfg.Get("proxyHost") // TODO missing proxy configurations
	proxyPort := cfg.Get("proxyPort")
	if proxyHost != "" {
		config.ProxyUrl = proxyHost + ":" + proxyPort
	}
	// TODO add timeout configurations
	return &config
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
