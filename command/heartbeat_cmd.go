package command

import (
	"errors"
	"fmt"
	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
	"os"
)

func NewHeartbeatClient(c *gcli.Context) (*heartbeat.Client, error) {
	heartbeatCli, cliErr := heartbeat.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the heartbeat client. " + cliErr.Error()
		fmt.Printf("%s\n", message)
		return nil, errors.New(message)
	}
	printVerboseMessage("Heartbeat Client created.")
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

	printVerboseMessage("Heartbeat request prepared from flags, sending request to Opsgenie..")

	response, err := cli.Ping(nil, name)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Ping request has received. RequestID: " + response.RequestId)
}
