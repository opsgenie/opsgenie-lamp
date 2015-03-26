// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package command

import(
	// "fmt"
	gcli "github.com/codegangsta/cli"
	"log"
	hb "github.com/opsgenie/opsgenie-go-sdk/heartbeat"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	// "strings"
	// "errors"
)

func HeartbeatAction(c *gcli.Context) {
	// mandatory arguments: name (apiKey may be given by the configuration file)
	if !c.IsSet("name") {
		log.Fatalln("Name parameter is mandatory and must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewHeartbeatClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the renotify request
	req := hb.SendHeartbeatRequest{}	
	if c.IsSet("name") {
		req.Name = c.String("name")		
	}	
	// send the request
	_, err = cli.Send(req)
	if err != nil {
		log.Fatalln("Could not send heartbeat")
	}
	log.Println("Heartbeat sent successfuly")
}