package command

import (
	gcli "github.com/codegangsta/cli"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	"fmt"
	"os"
)

// HeartbeatAction sends an Heartbeat signal to OpsGenie.
func HeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := hb.PingHeartbeatRequest{}
	if val, success := getVal("name", c); success {
		req.Name = val
	}

	printVerboseMessage("Heartbeat request prepared from flags, sending request to OpsGenie..")

	response, err := cli.Ping(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Ping request has recived. RequestID: " + response.RequestID)
}
