// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package command

import(
	"fmt"
	gcli "github.com/codegangsta/cli"
	"log"
	integration "github.com/opsgenie/opsgenie-go-sdk/integration"
	policy "github.com/opsgenie/opsgenie-go-sdk/policy"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	// "strings"
	// "errors"
)

func EnableIntegrationAction(c *gcli.Context) {
	// mandatory arguments: id/name (apiKey may be given by the configuration file)
	if c.IsSet("id") && c.IsSet("name") {
		log.Fatalln("Either alert id or name must be provided, not both")
	}
	if !c.IsSet("id") && !c.IsSet("name") {
		log.Fatalln("At least one of the alert id and name must be provided")
	}
	switch c.String("type") {
	case "policy":		
		// get a client instance using the api key
		cli, err := NewPolicyClient( grabApiKey(c) )	
		if err != nil {
			log.Fatalln(err.Error())
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
			log.Fatalln("Could not enable the policy")
		}
		log.Println("Policy enabled successfuly")

	case "integration":
		// get a client instance using the api key
		cli, err := NewIntegrationClient( grabApiKey(c) )	
		if err != nil {
			log.Fatalln(err.Error())
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
			log.Fatalln("Could not enable the integration")
		}
		log.Println("Integration enabled successfuly")
	default:
		log.Fatalln(fmt.Sprintf("Invalid type option %s, specify either integration or policy", c.String("type")))
	}

}

func DisableIntegrationAction(c *gcli.Context) {
	// mandatory arguments: id/name (apiKey may be given by the configuration file)
	if c.IsSet("id") && c.IsSet("name") {
		log.Fatalln("Either alert id or name must be provided, not both")
	}
	if !c.IsSet("id") && !c.IsSet("name") {
		log.Fatalln("At least one of the alert id and name must be provided")
	}

	switch c.String("type") {
	case "policy":
		// get a client instance using the api key
		cli, err := NewPolicyClient( grabApiKey(c) )	
		if err != nil {
			log.Fatalln(err.Error())
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
			log.Fatalln("Could not disable the policy")
		}
		log.Println("Policy disabled successfuly")		

	case "integration":
		// get a client instance using the api key
		cli, err := NewIntegrationClient( grabApiKey(c) )	
		if err != nil {
			log.Fatalln(err.Error())
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
			log.Fatalln("Could not disable the integration")
		}
		log.Println("Integration disabled successfuly")
	default:
		log.Fatalln(fmt.Sprintf("Invalid type option %s, specify either integration or policy", c.String("type")))
	}
}
