// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

package command

import (
	"fmt"
	"os"
	"strings"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/alerts"
)

func CreateAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alerts.CreateAlertRequest{}

	if val, success := getVal("message", c); success {
		req.Message = val
	}
	if val, success := getVal("teams", c); success {
		req.Teams = strings.Split(val, ",")
	}
	if val, success := getVal("recipients", c); success {
		req.Recipients = strings.Split(val, ",")
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("actions", c); success {
		req.Actions = strings.Split(val, ",")
	}
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("tags", c); success {
		req.Tags = strings.Split(val, ",")
	}
	if val, success := getVal("description", c); success {
		req.Description = val
	}
	if val, success := getVal("entity", c); success {
		req.Entity = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("note", c); success {
		req.Note = val
	}
	if c.IsSet("D") {
		req.Details = extractDetailsFromCommand(c)
	}

	printVerboseMessage("Create alert request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Create(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Alert created successfully.")
	fmt.Println("alertId=" + resp.AlertId)
}

func extractDetailsFromCommand(c *gcli.Context) map[string]string {
	details := make(map[string]string)
	extraProps := c.StringSlice("D")
	for i := 0; i < len(extraProps); i++ {
		prop := extraProps[i]
		if !isEmpty("D", prop, c) && strings.Contains(prop, "=") {
			p := strings.Split(prop, "=")
			details[p[0]] = strings.Join(p[1:], "=")
		} else {
			fmt.Sprintf("Dynamic parameters should have the value of the form a=b, but got:" + prop)
			gcli.ShowCommandHelp(c, c.Command.Name)
			os.Exit(1)
		}
	}

	return details
}

func GetAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alerts.GetAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	printVerboseMessage("Get alert request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Get(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Got Alert successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := ResultToYaml(resp)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := ResultToJson(resp, isPretty)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(output)
	}
}

func AttachFileAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AttachFileAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("attachment", c); success {
		f, err := os.Open(val)
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		req.Attachment = f
	}

	if val, success := getVal("indexFile", c); success {
		req.IndexFile = val
	}

	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Attach request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AttachFile(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printVerboseMessage("File attached to alert successfully.")
}

func AcknowledgeAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AcknowledgeAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Acknowledge alert request prepared from flags, sending request to OpsGenie..")

	_, err = cli.Acknowledge(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert acknowledged successfully.")
}

func RenotifyAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.RenotifyAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("recipients", c); success {
		req.Recipients = strings.Split(val, ",")
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Renotify request prepared from flags, sending request to OpsGenie..")

	_, err = cli.Renotify(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Renotified successfully.")
}

func TakeOwnershipAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.TakeOwnershipAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Take ownership request prepared from flags, sending request to OpsGenie..")

	_, err = cli.TakeOwnership(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Ownership taken successfully.")
}

func AssignOwnerAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AssignOwnerAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("owner", c); success {
		req.Owner = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Assign ownership request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AssignOwner(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Ownership assigned successfully.")
}

func AddTeamAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddTeamAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("team", c); success {
		req.Team = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Add team request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AddTeam(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Team added successfully.")
}

func AddRecipientAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddRecipientAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("recipient", c); success {
		req.Recipient = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Add recipient request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AddRecipient(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Recipient added successfully.")
}

func AddNoteAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddNoteAlertRequest{}

	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Add note request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AddNote(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Note added successfully.")
}

func ExecuteActionAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.ExecuteActionAlertRequest{}

	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	action, success := getVal("action", c)
	if success {
		req.Action = action
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Execute action request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ExecuteAction(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Action" + action + "executed successfully.")
	fmt.Println("Action [" + action + "] result=" + resp.Result)
}

func CloseAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.CloseAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}
	if val, success := getVal("notify", c); success {
		req.Notify = strings.Split(val, ",")
	}

	printVerboseMessage("Close alert request prepared from flags, sending request to OpsGenie..")

	_, err = cli.Close(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Alert closed successfully.")
}

func DeleteAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.DeleteAlertRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}

	printVerboseMessage("Delete alert request prepared from flags, sending request to OpsGenie..")

	_, err = cli.Delete(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert deleted successfully.")
}
