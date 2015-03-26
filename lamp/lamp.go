// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package lamp

import (
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client" 
	"errors"
	gcfg "code.google.com/p/gcfg"
	"encoding/json"
	yaml "gopkg.in/yaml.v2"
	gcli "github.com/codegangsta/cli"
)

const CONF_FILE_LINUX string = "/etc/opsgenie/conf/opsgenie-integration.conf"

type LampConfig struct {
	Lamp struct {
		ApiKey string
	}
}

var lampCfg LampConfig

func grabApiKey(c *gcli.Context) string {
	if c.IsSet("apiKey") {
		return c.String("apiKey")
	} else {
		return lampCfg.Lamp.ApiKey
	}
	return ""
}

func NewAlertClient(apiKey string) (*ogcli.OpsGenieAlertClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	alertCli, cliErr := cli.Alert()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the alert client")
	}	
	return alertCli, nil
}

func NewHeartbeatClient(apiKey string) (*ogcli.OpsGenieHeartbeatClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	hbCli, cliErr := cli.Heartbeat()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the heartbeat client")
	}	
	return hbCli, nil
}

func NewIntegrationClient(apiKey string) (*ogcli.OpsGenieIntegrationClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	intCli, cliErr := cli.Integration()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the integration client")
	}	
	return intCli, nil
}

func NewPolicyClient(apiKey string) (*ogcli.OpsGeniePolicyClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	polCli, cliErr := cli.Policy()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the policy client")
	}	
	return polCli, nil
}


func ResultToYaml(data interface{}) (string, error) {
	// output in yaml
	d, err := yaml.Marshal(&data)
    if err != nil {
    	return "", errors.New("Can not marshal the response into YAML format")
   	}
   	return string(d), nil
}

func ResultToJson(data interface{}, pretty bool) (string, error){
	// output in json
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

func init() {
	err := gcfg.ReadFileInto(&lampCfg, CONF_FILE_LINUX)	
	if err != nil {
		panic(errors.New("Can not read the lamp configuration file!"))
	}
}
