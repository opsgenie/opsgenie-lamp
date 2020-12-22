package command

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration"
	"github.com/opsgenie/opsgenie-go-sdk-v2/policy"
	gcli "github.com/urfave/cli"
	"os"
)

func NewIntegrationClient(c *gcli.Context) (*integration.Client, error) {
	integrationCli, cliErr := integration.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the integration client. " + cliErr.Error()
		printMessage(ERROR, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Integration Client created.")
	return integrationCli, nil
}

func NewPolicyClient(c *gcli.Context) (*policy.Client, error) {
	policyCli, cliErr := policy.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the policy client. " + cliErr.Error()
		printMessage(ERROR, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Policy Client created.")
	return policyCli, nil
}

// EnableAction enables an integration/policy according to the --type parameter at Opsgenie.
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
			req.Id = val
		}
		if val, success := getVal("teamId", c); success {
			req.TeamId = val
		}
		if val, success := getVal("policyType", c); success {
			req.Type = policy.PolicyType(val)
		}
		printMessage(DEBUG,"Enable policy request prepared from flags, sending request to Opsgenie..")
		_, err = cli.EnablePolicy(nil, &req)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}
		printMessage(INFO,"Policy enabled successfuly")

	case "integration":
		cli, err := NewIntegrationClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := integration.EnableIntegrationRequest{}
		if val, success := getVal("id", c); success {
			req.Id = val
		}
		printMessage(DEBUG,"Enable integration request prepared from flags, sending request to Opsgenie..")
		_, err = cli.Enable(nil, &req)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}
		printMessage(INFO,"Integration enabled successfuly")
	default:
		printMessage(INFO,"Invalid type option " + val + ", specify either integration or policy")
		gcli.ShowCommandHelp(c, "enable")
		os.Exit(1)
	}
}

// DisableAction disables an integration/policy according to the --type parameter at Opsgenie.
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
			req.Id = val
		}
		if val, success := getVal("teamId", c); success {
			req.TeamId = val
		}
		if val, success := getVal("policyType", c); success {
			req.Type = policy.PolicyType(val)
		}
		printMessage(DEBUG,"Disable policy request prepared from flags, sending request to Opsgenie..")
		_, err = cli.DisablePolicy(nil, &req)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}
		printMessage(INFO, "Policy disabled successfuly")

	case "integration":
		cli, err := NewIntegrationClient(c)
		if err != nil {
			os.Exit(1)
		}

		req := integration.DisableIntegrationRequest{}
		if val, success := getVal("id", c); success {
			req.Id = val
		}
		printMessage(DEBUG,"Disable integration request prepared from flags, sending request to Opsgenie..")
		_, err = cli.Disable(nil, &req)
		if err != nil {
			printMessage(ERROR,err.Error())
			os.Exit(1)
		}
		printMessage(INFO,"Integration disabled successfuly")
	default:
		printMessage(ERROR,"Invalid type option " + val + ", specify either integration or policy")
		gcli.ShowCommandHelp(c, "disable")
		os.Exit(1)
	}
}
