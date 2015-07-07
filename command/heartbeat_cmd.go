// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

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
		fmt.Printf(err.Error() + "\n")
		os.Exit(1)
	}
	printVerboseMessage("Send heartbeat successfully")
	fmt.Printf("heartbeat=%d\n", response.Heartbeat)
}
