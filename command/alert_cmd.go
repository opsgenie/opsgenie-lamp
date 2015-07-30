package command

import (
	"fmt"
	"os"
	"strings"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/alerts"
)

// CreateAlertAction creates an alert at OpsGenie.
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Alert created successfully.")
	fmt.Printf("alertId=%s\n", resp.AlertID)
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
			fmt.Printf("Dynamic parameters should have the value of the form a=b, but got: %s\n", prop)
			gcli.ShowCommandHelp(c, c.Command.Name)
			os.Exit(1)
		}
	}

	return details
}

// GetAlertAction retrieves specified alert details from OpsGenie.
func GetAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alerts.GetAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	printVerboseMessage("Get alert request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Get(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Got Alert successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp, isPretty)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	}
}

// AttachFileAction attaches a file to an alert at OpsGenie.
func AttachFileAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AttachFileAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("attachment", c); success {
		f, err := os.Open(val)
		defer f.Close()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("File attached to alert successfully.")
}

// AcknowledgeAction acknowledges an alert at OpsGenie.
func AcknowledgeAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AcknowledgeAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert acknowledged successfully.")
}

// RenotifyAction re-notifies recipients at OpsGenie.
func RenotifyAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.RenotifyAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Renotified successfully.")
}

// TakeOwnershipAction takes the ownership of an alert at OpsGenie.
func TakeOwnershipAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.TakeOwnershipAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Ownership taken successfully.")
}

// AssignOwnerAction assigns the specified user as the owner of the alert at OpsGenie.
func AssignOwnerAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AssignOwnerAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Ownership assigned successfully.")
}

// AddTeamAction adds a team to an alert at OpsGenie.
func AddTeamAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddTeamAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Team added successfully.")
}

// AddRecipientAction adds recipient to an alert at OpsGenie.
func AddRecipientAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddRecipientAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Recipient added successfully.")
}

// AddTagsAction adds tags to an alert at OpsGenie.
func AddTagsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddTagsAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("tags", c); success {
		req.Tags = strings.Split(val, ",")
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Add tag request prepared from flags, sending request to OpsGenie..")

	_, err = cli.AddTags(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Tags added successfully.")
}

// AddNoteAction adds a note to an alert at OpsGenie.
func AddNoteAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.AddNoteAlertRequest{}

	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Note added successfully.")
}

// ExecuteActionAction executes a custom action on an alert at OpsGenie.
func ExecuteActionAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.ExecuteActionAlertRequest{}

	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Action" + action + "executed successfully.")
	fmt.Printf("result=%s\n", resp.Result)
}

// CloseAlertAction closes an alert at OpsGenie.
func CloseAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.CloseAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Alert closed successfully.")
}

// DeleteAlertAction deletes an alert at OpsGenie.
func DeleteAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alerts.DeleteAlertRequest{}
	if val, success := getVal("id", c); success {
		req.ID = val
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
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert deleted successfully.")
}
