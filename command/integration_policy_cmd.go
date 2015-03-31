// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package command

import(
	"fmt"
	gcli "github.com/codegangsta/cli"
	integration "github.com/opsgenie/opsgenie-go-sdk/integration"
	policy "github.com/opsgenie/opsgenie-go-sdk/policy"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	// "strings"
	// "errors"
	"os"
)

func EnableIntegrationAction(c *gcli.Context) {
	// mandatory arguments: id/name (apiKey may be given by the configuration file)
	if c.IsSet("id") && c.IsSet("name") {
		cmdlog.Error("Either alert id or name must be provided, not both")
		gcli.ShowCommandHelp(c, "enable")
		os.Exit(1)
	}
	if !c.IsSet("id") && !c.IsSet("name") {
		cmdlog.Error("At least one of the alert id and name must be provided")
		gcli.ShowCommandHelp(c, "enable")
		os.Exit(1)
	}
	switch c.String("type") {
	case "policy":		
		// get a client instance using the api key
		cli, err := NewPolicyClient( grabApiKey(c) )	
		if err != nil {
			cmdlog.Error(err.Error())
			os.Exit(1)
		}
		// build the enable-policy request
		req := policy.EnablePolicyRequest{}	
		if c.IsSet("id") {
			req.Id = c.String("id")		
		} else if c.IsSet("name") {
			req.Name = c.String("name")
		}	
		// send the request
		_, err = cli.Enable(req)
		if err != nil {
			cmdlog.Error("Could not enable the policy")
			os.Exit(1)
		}
		cmdlog.Info("Policy enabled successfuly")

	case "integration":
		// get a client instance using the api key
		cli, err := NewIntegrationClient( grabApiKey(c) )	
		if err != nil {
			cmdlog.Error(err.Error())
			os.Exit(1)
		}
		// build the enable-integration request
		req := integration.EnableIntegrationRequest{}	
		if c.IsSet("id") {
			req.Id = c.String("id")		
		} else if c.IsSet("name") {
			req.Name = c.String("name")
		}	
		// send the request
		_, err = cli.Enable(req)
		if err != nil {
			cmdlog.Error("Could not enable the integration")
			os.Exit(1)
		}
		cmdlog.Info("Integration enabled successfuly")
	default:
		cmdlog.Error(fmt.Sprintf("Invalid type option %s, specify either integration or policy", c.String("type")))
		gcli.ShowCommandHelp(c, "enable")
		os.Exit(1)
	}

}

func DisableIntegrationAction(c *gcli.Context) {
	// mandatory arguments: id/name (apiKey may be given by the configuration file)
	if c.IsSet("id") && c.IsSet("name") {
		cmdlog.Error("Either alert id or name must be provided, not both")
		gcli.ShowCommandHelp(c, "disable")
		os.Exit(1)
	}
	if !c.IsSet("id") && !c.IsSet("name") {
		cmdlog.Error("At least one of the alert id and name must be provided")
		gcli.ShowCommandHelp(c, "disable")
		os.Exit(1)
	}

	switch c.String("type") {
	case "policy":
		// get a client instance using the api key
		cli, err := NewPolicyClient( grabApiKey(c) )	
		if err != nil {
			cmdlog.Error(err.Error())
			os.Exit(1)
		}
		// build the disable-policy request
		req := policy.DisablePolicyRequest{}	
		if c.IsSet("id") {
			req.Id = c.String("id")		
		} else if c.IsSet("name") {
			req.Name = c.String("name")
		}	
		// send the request
		_, err = cli.Disable(req)
		if err != nil {
			cmdlog.Error("Could not disable the policy")
			os.Exit(1)
		}
		cmdlog.Info("Policy disabled successfuly")		

	case "integration":
		// get a client instance using the api key
		cli, err := NewIntegrationClient( grabApiKey(c) )	
		if err != nil {
			cmdlog.Error(err.Error())
			os.Exit(1)
		}
		// build the disable-integration request
		req := integration.DisableIntegrationRequest{}	
		if c.IsSet("id") {
			req.Id = c.String("id")		
		} else if c.IsSet("name") {
			req.Name = c.String("name")
		}	
		// send the request
		_, err = cli.Disable(req)
		if err != nil {
			cmdlog.Error("Could not disable the integration")
			os.Exit(1)
		}
		cmdlog.Info("Integration disabled successfuly")
	default:
		cmdlog.Error(fmt.Sprintf("Invalid type option %s, specify either integration or policy", c.String("type")))
		gcli.ShowCommandHelp(c, "disable")
		os.Exit(1)
	}
}
