package lamp

import (
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client" 
	"errors"
	gcfg "code.google.com/p/gcfg"
	"encoding/json"
	yaml "gopkg.in/yaml.v2"
)

const CONF_FILE_LINUX string = "/etc/opsgenie/conf/opsgenie-integration.conf"

type LampConfig struct {
	Lamp struct {
		ApiKey string
	}
}

var lampCfg LampConfig


func NewAlertClient(apiKey string) (*ogcli.OpsGenieAlertClient, error) {
	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(apiKey)
	
	alertCli, cliErr := cli.Alert()
	
	if cliErr != nil {
		return nil, errors.New("Can not create the alert client")
	}	
	return alertCli, nil
}

func ResultToYaml(data interface{}) (string, error) {
	// output in yaml
	d, err := yaml.Marshal(&data)
    if err != nil {
    	return "", errors.New("Can not marshal the response into YAML format")
   	}
   	return string(d), nil
}

func ResultToJson(data interface{}) (string, error){
	// output in json
	b, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("Can not marshal the response into JSON format")
	}
	return string(b), nil
}



func init() {
	err := gcfg.ReadFileInto(&lampCfg, CONF_FILE_LINUX)	
	if err != nil {
		panic(errors.New("Can not read the lamp configuration file!"))
	}
}
