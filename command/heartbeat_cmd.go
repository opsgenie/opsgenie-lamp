package command

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func NewHeartbeatClient(c *gcli.Context) (*heartbeat.Client, error) {
	heartbeatCli, cliErr := heartbeat.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the heartbeat client. " + cliErr.Error()
		printMessage(INFO, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Heartbeat Client created.")
	return heartbeatCli, nil
}

// HeartbeatAction sends an Heartbeat signal to Opsgenie.
func PingHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	var name string
	if val, success := getVal("name", c); success {
		name = val
	}

	printMessage(DEBUG,"Heartbeat request prepared from flags, sending request to Opsgenie..")

	response, err := cli.Ping(nil, name)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(DEBUG,"Ping request has received. RequestID: " + response.RequestId)
}

func CreateHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	addRequest := heartbeat.AddRequest{}

	if val, success := getVal("name", c); success {
		addRequest.Name = val
	}

	if val, success := getVal("description", c); success {
		addRequest.Description = val
	}

	if val, success := getVal("alertMessage", c); success {
		addRequest.AlertMessage = val
	}

	if val, success := getVal("alertTag", c); success {
		addRequest.AlertTag = strings.Split(val, ",")
	}

	if val, success := getVal("alertPriority", c); success {
		addRequest.AlertPriority = val
	}

	if val, success := getVal("ownerTeam", c); success {
		addRequest.OwnerTeam = og.OwnerTeam{
			Name: val,
		}
	}

	if val, success := getVal("interval", c); success {
		addRequest.Interval, err = strconv.Atoi(val)

		if err != nil {
			printMessage(ERROR, "Please provide a valid integer for interval.")
			os.Exit(1)
		}
	}

	if val, success := getVal("intervalType", c); success {
		if val == "m" {
			addRequest.IntervalUnit = heartbeat.Minutes
		} else if val == "h" {
			addRequest.IntervalUnit = heartbeat.Hours
		} else if val == "d" {
			addRequest.IntervalUnit = heartbeat.Days
		} else {
			printMessage(ERROR, "Please provide a valid interval unit.")
			os.Exit(1)
		}
	}

	enabled := c.IsSet("enabled")
	addRequest.Enabled = &enabled

	printMessage(DEBUG, "Heartbeat create request created from flags. Sedning to Opsgenie...")

	response, err := cli.Add(nil, &addRequest)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(DEBUG,"Heartbeat will be created " + response.RequestId)
}

func DeleteHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	var name string
	if val, success := getVal("name", c); success {
		name = val
	}

	printMessage(DEBUG, "Heartbeat delete request created from flags. Sending to Opsgenie...")

	response, err := cli.Delete(nil, name)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(DEBUG,"Heartbeat will be deleted.")
	printMessage(INFO, response.RequestId)
}

func DisableHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	var name string
	if val, success := getVal("name", c); success {
		name = val
	}

	printMessage(DEBUG, "Heartbeat disable request created from flags. Sending to Opsgenie...")

	response, err := cli.Disable(nil, name)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(DEBUG,"Heartbeat will be disabled.")
	printMessage(INFO, response.RequestId)
}

func EnableHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	var name string
	if val, success := getVal("name", c); success {
		name = val
	}

	printMessage(DEBUG, "Heartbeat enable request created from flags. Sending to Opsgenie...")

	response, err := cli.Enable(nil, name)
	if err != nil {
		printMessage(ERROR,err.Error())
	}
	printMessage(DEBUG, "Heartbeat will be enabled")
	printMessage(INFO, response.RequestId)
}

func ListHeartbeatAction(c *gcli.Context) {
	cli, err := NewHeartbeatClient(c)
	if err != nil {
		os.Exit(1)
	}

	response, err := cli.List(nil)
	if err != nil {
		printMessage(ERROR,err.Error())
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printMessage(DEBUG,"Heartbeats listed successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(response.Heartbeats)
		if err != nil {
			printMessage(ERROR,err.Error())
			os.Exit(1)
		}
		printMessage(INFO, output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(response.Heartbeats, isPretty)
		if err != nil {
			printMessage(ERROR,err.Error())
			os.Exit(1)
		}
		printMessage(INFO, output)
	}
}
