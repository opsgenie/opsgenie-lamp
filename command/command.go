// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

/*
Package 'command' creates various OpsGenie API clients:
 	- Alert
 	- Heartbeat
 	- Integration
 	- Policy
*/
package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	gcli "github.com/codegangsta/cli"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/opsgenie/opsgenie-lamp/cfg"
	yaml "gopkg.in/yaml.v2"
)

var verbose = false

func printVerboseMessage(message string) {
	if verbose {
		fmt.Printf(message + "\n")
	}
}

// The 'Api key' is the most common parameter for all commands.
// It is provided either on command line or on the configuration file.
// The 'grabApiKey' function is used through all commands in purpose of
// creating OpsGenie clients.
func grabApiKey(c *gcli.Context) string {
	if val, success := getVal("apiKey", c); success {
		return val
	} else {
		apiKey := cfg.Get("apiKey")
		printVerboseMessage("apiKey flag is not set in the command, reading apiKey from config..")
		return apiKey
	}
}

func grabUsername(c *gcli.Context) string {
	if val, success := getVal("user", c); success {
		return val
	} else {
		return cfg.Get("user")
	}
}

func getVal(argName string, c *gcli.Context) (string, bool) {
	if c.IsSet(argName) {
		arg := c.String(argName)
		isEmpty(argName, arg, c)
		return arg, true
	}
	return "", false
}

func isEmpty(argName string, arg string, c *gcli.Context) bool {
	var prefix string
	for _, name := range c.FlagNames() {
		if len(name) == 1 {
			prefix = "-"
		} else {
			prefix = "--"
		}
		if strings.Contains(arg, prefix+name) {
			fmt.Printf("Value of argument '%s' is empty\n", argName)
			gcli.ShowCommandHelp(c, c.Command.Name)
			os.Exit(1)
		}
	}
	return false
}

func getProxyConf(host string, port int) (proxy *ogcli.ClientProxyConfiguration) {
	printVerboseMessage("Configuring proxy settings with host " + host + " and port " + strconv.Itoa(port))
	pc := new(ogcli.ClientProxyConfiguration)
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

func getConnectionConf() (connCfg *ogcli.HttpTransportSettings) {
	printVerboseMessage("Configuring connection settings..")
	cfg := new(ogcli.HttpTransportSettings)
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
	if strDuration := cfg.Get(key); strDuration != "" {
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
	} else {
		printVerboseMessage("Could not parse [" + key + "] with value [" + strDuration + "].")
		return 0
	}
}

func initialize(c *gcli.Context) *ogcli.OpsGenieClient {
	if c.IsSet("v") {
		verbose = true
		printVerboseMessage("Will execute command in verbose mode.")
	}
	readConfigFile(c)
	apiKey := grabApiKey(c)
	cli := new(ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	if apiUrl := cfg.Get("opsgenie.api.url"); apiUrl != "" {
		cli.SetOpsGenieApiUrl(apiUrl)
	}
	proxyHost := cfg.Get("proxyHost")
	proxyPort, err := strconv.Atoi(cfg.Get("proxyPort"))
	if err == nil && proxyPort != 0 && proxyHost != "" {
		cli.SetClientProxyConfiguration(getProxyConf(proxyHost, proxyPort))
	}
	cli.SetHttpTransportSettings(getConnectionConf())
	return cli
}

// In order to interact with the Alert API, one must handle an AlertClient.
// The 'NewAlertClient' function creates and returns an instance of that type.
func NewAlertClient(c *gcli.Context) (*ogcli.OpsGenieAlertClient, error) {
	cli := initialize(c)
	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		message := "Can not create the alert client. " + cliErr.Error()
		fmt.Printf(message + "\n")
		return nil, errors.New(message)
	}
	printVerboseMessage("Alert Client created..")
	return alertCli, nil
}

// In order to interact with the Heartbeat API, one must handle a HeartbeatClient.
// The 'NewHeartbeatClient' function creates and returns an instance of that type.
func NewHeartbeatClient(c *gcli.Context) (*ogcli.OpsGenieHeartbeatClient, error) {
	cli := initialize(c)
	hbCli, cliErr := cli.Heartbeat()

	if cliErr != nil {
		message := "Can not create the heartbeat client. " + cliErr.Error()
		fmt.Printf(message + "\n")
		return nil, errors.New(message)
	}
	printVerboseMessage("Heartbeat Client created..")
	return hbCli, nil
}

// In order to interact with the Integration API, one must handle an IntegrationClient.
// The 'NewIntegrationClient' function creates and returns an instance of that type.
func NewIntegrationClient(c *gcli.Context) (*ogcli.OpsGenieIntegrationClient, error) {
	cli := initialize(c)
	intCli, cliErr := cli.Integration()

	if cliErr != nil {
		message := "Can not create the integration client. " + cliErr.Error()
		fmt.Printf(message + "\n")
		return nil, errors.New(message)
	}
	printVerboseMessage("Integration Client created..")
	return intCli, nil
}

// In order to interact with the Policy API, one must handle a PolicyClient.
// The 'NewPolicyClient' function creates and returns an instance of that type.
func NewPolicyClient(c *gcli.Context) (*ogcli.OpsGeniePolicyClient, error) {
	cli := initialize(c)
	polCli, cliErr := cli.Policy()

	if cliErr != nil {
		message := "Can not create the policy client. " + cliErr.Error()
		fmt.Printf(message + "\n")
		return nil, errors.New(message)
	}
	printVerboseMessage("Policy Client created..")
	return polCli, nil
}

// The 'getAlert' command returns a GetAlertResponse object.
// That object has a type of struct and can easily be represented in Yaml format.
// The 'ResultToYaml' function is called whenever "output-format" parameter is
// set to yaml.
func ResultToYaml(data interface{}) (string, error) {
	d, err := yaml.Marshal(&data)
	if err != nil {
		return "", errors.New("Can not marshal the response into YAML format. " + err.Error())
	}
	return string(d), nil
}

// The 'getAlert' command returns a GetAlertResponse object.
// That object has a type of struct and can easily be represented in JSON format.
// The 'ResultToJson' function is called whenever "output-format" parameter is
// set to json or not provided. "getAlert" command defaults to JSON format.
// Pretty formating yields an indented style of representation. Pretty formating
// is on when the "pretty" flag is provided alongside.
func ResultToJson(data interface{}, pretty bool) (string, error) {
	if pretty {
		b, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", errors.New("Can not marshal the response into JSON format. " + err.Error())
		}
		return string(b), nil
	} else {
		b, err := json.Marshal(data)
		if err != nil {
			return "", errors.New("Can not marshal the response into JSON format" + err.Error())
		}
		return string(b), nil
	}
	return "", nil
}

func readConfigFile(c *gcli.Context) {
	cfg.Verbose = verbose
	if val, success := getVal("config", c); success {
		cfg.LoadConfigFromGivenPath(val)
	} else {
		cfg.LoadConfiguration()
	}
}
