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
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client" 
	"errors"
	gcfg "code.google.com/p/gcfg"
	"encoding/json"
	yaml "gopkg.in/yaml.v2"
	gcli "github.com/codegangsta/cli"
	"log"
)
// The configuration file used by the client
// TODO Read from the environment variable LAMP_HOME
// TODO Same configuration for the Windows is required
const CONF_FILE_LINUX string = "/etc/opsgenie/conf/opsgenie-integration.conf"

// Configuration is parsed from an 'ini' style file.
// The key-value pairs are stored inside a struct data type.
// TODO logging properties to be read
type LampConfig struct {
	Lamp struct {
		ApiKey string
	}
}

var lampCfg LampConfig
// The 'Api key' is the most common parameter for all commands.
// It is provided either on command line or on the configuration file.
// The 'grabApiKey' function is used through all commands in purpose of
// creating OpsGenie clients.
func grabApiKey(c *gcli.Context) string {
	if c.IsSet("apiKey") {
		return c.String("apiKey")
	} else {
		return lampCfg.Lamp.ApiKey
	}
	return ""
}
// In order to interact with the Alert API, one must handle an AlertClient.
// The 'NewAlertClient' function creates and returns an instance of that type.
func NewAlertClient(apiKey string) (*ogcli.OpsGenieAlertClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	alertCli, cliErr := cli.Alert()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the alert client")
	}	
	return alertCli, nil
}
// In order to interact with the Heartbeat API, one must handle a HeartbeatClient.
// The 'NewHeartbeatClient' function creates and returns an instance of that type.
func NewHeartbeatClient(apiKey string) (*ogcli.OpsGenieHeartbeatClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	hbCli, cliErr := cli.Heartbeat()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the heartbeat client")
	}	
	return hbCli, nil
}
// In order to interact with the Integration API, one must handle an IntegrationClient.
// The 'NewIntegrationClient' function creates and returns an instance of that type.
func NewIntegrationClient(apiKey string) (*ogcli.OpsGenieIntegrationClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	intCli, cliErr := cli.Integration()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the integration client")
	}	
	return intCli, nil
}
// In order to interact with the Policy API, one must handle a PolicyClient.
// The 'NewPolicyClient' function creates and returns an instance of that type.
func NewPolicyClient(apiKey string) (*ogcli.OpsGeniePolicyClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	polCli, cliErr := cli.Policy()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the policy client")
	}	
	return polCli, nil
}
// The 'getAlert' command returns a GetAlertResponse object. 
// That object has a type of struct and can easily be represented in Yaml format.
// The 'ResultToYaml' function is called whenever "output-format" parameter is
// set to yaml.
func ResultToYaml(data interface{}) (string, error) {
	d, err := yaml.Marshal(&data)
    if err != nil {
    	return "", errors.New("Can not marshal the response into YAML format")
   	}
   	return string(d), nil
}
// The 'getAlert' command returns a GetAlertResponse object. 
// That object has a type of struct and can easily be represented in JSON format.
// The 'ResultToJson' function is called whenever "output-format" parameter is
// set to json or not provided. "getAlert" command defaults to JSON format.
// Pretty formating yields an indented style of representation. Pretty formating 
// is on when the "pretty" flag is provided alongside.
func ResultToJson(data interface{}, pretty bool) (string, error){
	if pretty {
		b, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			return "", errors.New("Can not marshal the response into JSON format")
		}		
		return string(b), nil
	} else {
		b, err := json.Marshal(data)
		if err != nil {
			return "", errors.New("Can not marshal the response into JSON format")
		}		
		return string(b), nil
	}
	return "", nil
}
// "init" is a special function that loads in whenever the 'command' package is
// first allocated in memory. Therefore, it has the lines of instructions to
// initialize the program. Here, it is responsible for reading the configuration 
// into the configuration struct data.
func init() {
	err := gcfg.ReadFileInto(&lampCfg, CONF_FILE_LINUX)	
	if err != nil {
		log.Fatalln("Can not read the lamp configuration file!")
	}
}
