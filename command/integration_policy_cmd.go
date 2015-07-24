package command

import (
	"fmt"
	"os"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/integration"
	"github.com/opsgenie/opsgenie-go-sdk/policy"
)

// EnableAction enables an integration/policy according to the --type parameter at OpsGenie.
func EnableAction(c *gcli.Context) {
	val, _ := getVal("type", c)
	switch val {
	case "policy":
		cli, err := NewPolicyClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := policy.EnablePolicyRequest{}
			if val, success := getVal("id", c); success {
				req.ID = val
		}
			if val, success := getVal("name", c); success {
			req.Name = val
		}
		printVerboseMessage("Enable policy request prepared from flags, sending request to OpsGenie..")
		_, err = cli.Enable(req)
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			os.Exit(1)
		}
		fmt.Printf("Policy enabled successfuly\n")

	case "integration":
		cli, err := NewIntegrationClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := integration.EnableIntegrationRequest{}
			if val, success := getVal("id", c); success {
				req.ID = val
		}
			if val, success := getVal("name", c); success {
			req.Name = val
		}
		printVerboseMessage("Enable integration request prepared from flags, sending request to OpsGenie..")
		_, err = cli.Enable(req)
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			os.Exit(1)
		}
		fmt.Printf("Integration enabled successfuly\n")
	default:
		fmt.Printf("Invalid type option %s, specify either integration or policy\n", val)
		gcli.ShowCommandHelp(c, "enable")
		os.Exit(1)
	}
}

// DisableAction disables an integration/policy according to the --type parameter at OpsGenie.
func DisableAction(c *gcli.Context) {
	val, _ := getVal("type", c)
	switch val {
	case "policy":
		cli, err := NewPolicyClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := policy.DisablePolicyRequest{}
		if val, success := getVal("id", c); success {
			req.ID = val
		}
		if val, success := getVal("name", c); success {
			req.Name = val
		}
		printVerboseMessage("Disable policy request prepared from flags, sending request to OpsGenie..")
		_, err = cli.Disable(req)
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			os.Exit(1)
		}
		fmt.Printf("Policy disabled successfuly\n")

	case "integration":
		cli, err := NewIntegrationClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := integration.DisableIntegrationRequest{}
		if val, success := getVal("id", c); success {
			req.ID = val
		}
		if val, success := getVal("name", c); success {
			req.Name = val
		}
		printVerboseMessage("Disable integration request prepared from flags, sending request to OpsGenie..")
		_, err = cli.Disable(req)
		if err != nil {
			fmt.Printf(err.Error() + "\n")
			os.Exit(1)
		}
		fmt.Printf("Integration disabled successfuly\n")
	default:
		fmt.Printf("Invalid type option %s, specify either integration or policy\n", val)
		gcli.ShowCommandHelp(c, "disable")
		os.Exit(1)
	}
}
