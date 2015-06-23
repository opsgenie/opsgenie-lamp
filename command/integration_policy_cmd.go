// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package command

import(
	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/integration"
	"github.com/opsgenie/opsgenie-go-sdk/policy"
	"os"
	"fmt"
)

func EnableAction(c *gcli.Context) {
	if val, success := getVal("type", c); success{
		switch val {
		case "policy":
			cli, err := NewPolicyClient( c )
			if err != nil {
				os.Exit(1)
			}

			req := policy.EnablePolicyRequest{}
			if val, success := getVal("id", c); success{
				req.Id = val
			}
			if val, success := getVal("name", c); success{
				req.Name = val
			}
			printVerboseMessage("Enable policy request prepared from flags, sending request to OpsGenie..")
			_, err = cli.Enable(req)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Policy enabled successfuly")

		case "integration":
			cli, err := NewIntegrationClient( c )
			if err != nil {
				os.Exit(1)
			}

			req := integration.EnableIntegrationRequest{}
			if val, success := getVal("id", c); success{
				req.Id = val
			}
			if val, success := getVal("name", c); success{
				req.Name = val
			}
			printVerboseMessage("Enable integration request prepared from flags, sending request to OpsGenie..")
			_, err = cli.Enable(req)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Integration enabled successfuly")
		default:
			fmt.Sprintf("Invalid type option %s, specify either integration or policy", val)
			gcli.ShowCommandHelp(c, "enable")
			os.Exit(1)
		}
	}
}

func DisableAction(c *gcli.Context) {
	if val, success := getVal("type", c); success{
		switch val {
		case "policy":
			cli, err := NewPolicyClient( c )
			if err != nil {
				os.Exit(1)
			}

			req := policy.DisablePolicyRequest{}
			if val, success := getVal("id", c); success{
				req.Id = val
			}
			if val, success := getVal("name", c); success{
				req.Name = val
			}
			printVerboseMessage("Disable policy request prepared from flags, sending request to OpsGenie..")
			_, err = cli.Disable(req)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Policy disabled successfuly")

		case "integration":
			cli, err := NewIntegrationClient( c )
			if err != nil {
				os.Exit(1)
			}

			req := integration.DisableIntegrationRequest{}
			if val, success := getVal("id", c); success{
				req.Id = val
			}
			if val, success := getVal("name", c); success{
				req.Name = val
			}
			printVerboseMessage("Disable integration request prepared from flags, sending request to OpsGenie..")
			_, err = cli.Disable(req)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println("Integration disabled successfuly")
		default:
			fmt.Sprintf("Invalid type option %s, specify either integration or policy", val)
			gcli.ShowCommandHelp(c, "disable")
			os.Exit(1)
		}
	}
}
