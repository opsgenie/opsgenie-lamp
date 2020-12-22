package command

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
	gcli "github.com/urfave/cli"
	"os"
)

func NewHeartbeatClient(c *gcli.Context) (*heartbeat.Client, error) {
	heartbeatCli, cliErr := heartbeat.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the heartbeat client. " + cliErr.Error()
		printMessage(INFO, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Heartbeat Client created.")
	return heartbeatCli, nil
}

// HeartbeatAction sends an Heartbeat signal to Opsgenie.
func HeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	var name string
	if val, success := getVal("name", c); success {
		name = val
	}

	printMessage(DEBUG,"Heartbeat request prepared from flags, sending request to Opsgenie..")

	response, err := cli.Ping(nil, name)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(DEBUG,"Ping request has received. RequestID: " + response.RequestId)
}
