package command

import (
	// "fmt"
	gcli "github.com/codegangsta/cli"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	// "strings"
	// "errors"
	"fmt"
	"os"
)

// HeartbeatAction sends an Heartbeat signal to OpsGenie.
func HeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := hb.SendHeartbeatRequest{}
	if val, success := getVal("name", c); success {
		req.Name = val
	}

	printVerboseMessage("Heartbeat request prepared from flags, sending request to OpsGenie..")

	response, err := cli.Send(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Send heartbeat successfully")
	fmt.Printf("heartbeat=%d\n", response.Heartbeat)
}
