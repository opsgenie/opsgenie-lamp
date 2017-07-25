package command

import (
	"fmt"
	"os"
	"strings"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/alerts"
	"strconv"
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	"time"
	"io"
	"net/http"
)

// CreateAlertAction creates an alert at OpsGenie.
func CreateAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alertsv2.CreateAlertRequest{}

	if val, success := getVal("message", c); success {
		req.Message = val
	}
	if val, success := getVal("teams", c); success {
		teamNames := strings.Split(val, ",")

		var teams []alertsv2.TeamRecipient

		for _, name := range teamNames {
			teams = append(teams, &alertsv2.Team{Name: name})
		}
		req.Teams = teams
	}

	if _, success := getVal("recipients", c); success {
		printWarningMessage("WARNING: recipients param is deprecated and the value is ignoring")
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

	if val, success := getVal("priority", c); success {
		req.Priority = alertsv2.Priority(val)
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
	printVerboseMessage("Alert will be created.")
	fmt.Printf("requestId=%s\n", resp.RequestID)
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
	req := alertsv2.GetAlertRequest{
		Identifier: &alertsv2.Identifier{},
	}

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
		output, err := resultToYAML(resp.Alert)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp.Alert, isPretty)
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

	req := alertsv2.AddAlertAttachmentRequest{
		AttachmentAlertIdentifier: &alertsv2.AttachmentAlertIdentifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("tinyId", c); success {
		req.TinyID = val
	}

	if val, success := getVal("attachment", c); success {
		req.AttachmentFilePath = val
	}

	if val, success := getVal("indexFile", c); success {
		req.IndexFile = val
	}

	req.User = grabUsername(c)

	printVerboseMessage("Attach request prepared from flags, sending request to OpsGenie..")

	response, err := cli.AttachFile(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("File attached to alert successfully.")
	fmt.Printf("Result : %s\n", response.Result)
}

// GetAttachmentAction retrieves a download link to specified alert attachment
func GetAttachmentAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)

	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.GetAlertAttachmentRequest{
		AttachmentAlertIdentifier: &alertsv2.AttachmentAlertIdentifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("tinyId", c); success {
		req.TinyID = val
	}

	if val, success := getVal("attachmentId", c); success {
		req.AttachmentId = val
	}

	printVerboseMessage("Get alert attachment request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.GetAttachmentFile(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Got Alert Attachment successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.Attachment)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		output, err := CustomJsonMarshaller(resp.Attachment)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", output)
	}
}

// DownloadAttachmentAction downloads the attachment specified with attachmentId for given alert
func DownloadAttachmentAction(c *gcli.Context) {
	var destinationPath string
	cli, err := NewAlertClient(c)

	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.GetAlertAttachmentRequest{
		AttachmentAlertIdentifier: &alertsv2.AttachmentAlertIdentifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("tinyId", c); success {
		req.TinyID = val
	}

	if val, success := getVal("attachmentId", c); success {
		req.AttachmentId = val
	}

	if val, success := getVal("destinationPath", c); success {
		destinationPath = val
	}

	printVerboseMessage("Download alert attachment request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.GetAttachmentFile(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	fileName := resp.Attachment.Name
	downloadLink := resp.Attachment.DownloadLink

	var output *os.File

	if destinationPath != "" {
		output, err = os.Create(destinationPath + "/" + fileName)
	} else {
		output, err = os.Create(fileName)
	}

	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(downloadLink)
	if err != nil {
		fmt.Println("Error while downloading", fileName, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)

	if err != nil {
		fmt.Println("Error while downloading", fileName, "-", err)
		return
	}
}

// ListAlertAttachmentsAction returns a list of attachment meta information for specified alert
func ListAlertAttachmentsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)

	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.ListAlertAttachmentRequest{
		AttachmentAlertIdentifier: &alertsv2.AttachmentAlertIdentifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("tinyId", c); success {
		req.TinyID = val
	}

	printVerboseMessage("List alert attachments request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ListAlertAttachments(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("List Alert Attachment successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.AlertAttachments)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		output, err := CustomJsonMarshaller(resp.AlertAttachments)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("%s\n", output)
	}
}

// DeleteAlertAttachmentAction deletes the specified alert attachment from alert
func DeleteAlertAttachmentAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.DeleteAlertAttachmentRequest{
		AttachmentAlertIdentifier: &alertsv2.AttachmentAlertIdentifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("tinyId", c); success {
		req.TinyID = val
	}

	if val, success := getVal("attachmentId", c); success {
		req.AttachmentId = val
	}


	printVerboseMessage("Delete alert attachment request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.DeleteAttachment(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert attachment will be deleted. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
	fmt.Println("Result: " + resp.Result)
}

// AcknowledgeAction acknowledges an alert at OpsGenie.
func AcknowledgeAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.AcknowledgeRequest{
		Identifier: &alertsv2.Identifier{},
	}

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

	resp, err := cli.Acknowledge(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Acknowledge request will be processed. RequestID " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// RenotifyAction re-notifies recipients at OpsGenie.
func RenotifyAction(c *gcli.Context) {
	printWarningMessage("WARNING: `Renotify` is deprecated and will be removed!")
	cli, err := OldAlertClient(c)
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
	printWarningMessage("WARNING: `Take Ownership` is deprecated and will be removed!")
	cli, err := OldAlertClient(c)
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

	req := alertsv2.AssignAlertRequest{
		Identifier: &alertsv2.Identifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("owner", c); success {
		req.Owner = alertsv2.User{Username: val}
	}

	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Assign ownership request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Assign(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Ownership assignment request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// AddTeamAction adds a team to an alert at OpsGenie.
func AddTeamAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alertsv2.AddTeamToAlertRequest{
		Identifier: &alertsv2.Identifier{},
	}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("team", c); success {
		req.Team = alertsv2.Team{Name: val}
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Add team request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.AddTeamToAlert(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Add team request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// AddRecipientAction adds recipient to an alert at OpsGenie.
func AddRecipientAction(c *gcli.Context) {
	printWarningMessage("WARNING: `Add recipient` feature is deprecated and will be removed!")

	cli, err := OldAlertClient(c)
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

	req := alertsv2.AddTagsToAlertRequest{
		Identifier: &alertsv2.Identifier{},
	}
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

	resp, err := cli.AddTags(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Add tags request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// AddNoteAction adds a note to an alert at OpsGenie.
func AddNoteAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.AddNoteRequest{
		Identifier: &alertsv2.Identifier{},
	}

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

	resp, err := cli.AddNote(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Add note request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// ExecuteActionAction executes a custom action on an alert at OpsGenie.
func ExecuteActionAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.ExecuteCustomActionRequest{
		Identifier: &alertsv2.Identifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	action, success := getVal("action", c)
	if success {
		req.ActionName = action
	}

	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Execute action request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ExecuteCustomAction(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Execute custom action request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// CloseAlertAction closes an alert at OpsGenie.
func CloseAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.CloseRequest{
		Identifier: &alertsv2.Identifier{},
	}
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
	if _, success := getVal("notify", c); success {
		printWarningMessage("WARNING: notify is deprecated for removal and ignoring")
	}

	printVerboseMessage("Close alert request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Close(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Alert will be closed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// DeleteAlertAction deletes an alert at OpsGenie.
func DeleteAlertAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.DeleteAlertRequest{
		Identifier: &alertsv2.Identifier{},
	}
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

	resp, err := cli.Delete(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert will be deleted. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// ListAlertsAction retrieves alert details from OpsGenie.
func ListAlertsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := generateListAlertRequest(c)

	printVerboseMessage("List alerts request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.List(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Got Alerts successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.Alerts)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp.Alerts, isPretty)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	}
}

func generateListAlertRequest(c *gcli.Context) (alertsv2.ListAlertRequest) {
	req := alertsv2.ListAlertRequest{}

	if val, success := getVal("limit", c); success {
		limit, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		req.Limit = int(limit)
	}
	if val, success := getVal("sortBy", c); success {
		req.Sort = alertsv2.SortField(val)
	}
	if val, success := getVal("order", c); success {
		req.Order = alertsv2.Order(val)
	}
	if val, success := getVal("serchId", c); success {
		req.SearchIdentifier = val
		req.SearchIdentifierType = "id"
	}

	if val, success := getVal("searchName", c); success {
		req.SearchIdentifier = val
		req.SearchIdentifierType = "name"
	}

	if val, success := getVal("offset", c); success {
		offset, err := strconv.Atoi(val)
		if err != nil {
			os.Exit(2)
		}
		req.Offset = offset
	}

	if val, success := getVal("query", c); success {
		req.Query = val;
		printVerboseMessage("query is given other fields is ignoring")
	} else {
		generateQueryUsingOldStyleParams(c, &req)
	}

	return req
}
func generateQueryUsingOldStyleParams(c *gcli.Context, req *alertsv2.ListAlertRequest) {
	var queries []string
	if val, success := getVal("createdAfter", c); success {
		createdAfter, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		queries = append(queries, "createdAt > "+strconv.FormatUint(createdAfter, 10))
	}
	if val, success := getVal("createdBefore", c); success {
		createdBefore, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		queries = append(queries, "createdAt < "+strconv.FormatUint(createdBefore, 10))
	}
	if val, success := getVal("updatedAfter", c); success {
		updatedAfter, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		queries = append(queries, "updatedAt > "+strconv.FormatUint(updatedAfter, 10))
	}
	if val, success := getVal("updatedBefore", c); success {
		updatedBefore, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		queries = append(queries, "updatedAt < "+strconv.FormatUint(updatedBefore, 10))
	}
	if val, success := getVal("status", c); success {
		queries = append(queries, "status: "+val)
	}
	if val, success := getVal("teams", c); success {
		for _, teamName := range strings.Split(val, ",") {
			queries = append(queries, "teams: "+teamName)

		}
	}
	if val, success := getVal("tags", c); success {
		var tags []string
		operator := "AND"

		if val, success := getVal("tagsOperator", c); success {
			operator = val
		}

		for _, tag := range strings.Split(val, ",") {
			tags = append(tags, tag)
		}

		tagsPart := "tag: (" + strings.Join(tags, " "+operator+" ") + ")"
		queries = append(queries, tagsPart)
	}
	if len(queries) != 0 {
		req.Query = strings.Join(queries, " AND ");
	}
}

// CountAlertsAction retrieves number of alerts from OpsGenie.
func CountAlertsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := generateListAlertRequest(c)

	printVerboseMessage("Count alerts request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.List(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d\n", len(resp.Alerts))
}

// ListAlertNotesAction retrieves specified alert notes from OpsGenie.
func ListAlertNotesAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.ListAlertNotesRequest{
		Identifier: &alertsv2.Identifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("limit", c); success {
		limit, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		req.Limit = int(limit)
	}

	if val, success := getVal("order", c); success {
		req.Order = alertsv2.Order(val)
	}

	if val, success := getVal("direction", c); success {
		req.Direction = alertsv2.Direction(val)
	}

	if val, success := getVal("offset", c); success {
		req.Offset = val;
	}

	if val, success := getVal("lastKey", c); success && req.Offset == "" {
		req.Offset = val
		printWarningMessage("WARNING: lastKey param is deprecated for removal")
	}
	printVerboseMessage("List alert notes request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ListAlertNotes(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Alert notes listed successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.AlertNotes)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp.AlertNotes, isPretty)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	}
}

// ListAlertLogsAction retrieves specified alert logs from OpsGenie.
func ListAlertLogsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alertsv2.ListAlertLogsRequest{
		Identifier: &alertsv2.Identifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}

	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	if val, success := getVal("limit", c); success {
		limit, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		req.Limit = int(limit)
	}

	if val, success := getVal("order", c); success {
		req.Order = alertsv2.Order(val)
	}

	if val, success := getVal("direction", c); success {
		req.Direction = alertsv2.Direction(val)
	}

	if val, success := getVal("offset", c); success {
		req.Offset = val
	}

	if val, success := getVal("lastKey", c); success && req.Offset == "" {
		req.Offset = val
		printWarningMessage("WARNING: lastKey param is deprecated for removal")
	}
	printVerboseMessage("List alert notes request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ListAlertLogs(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Alert notes listed successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.AlertLogs)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp.AlertLogs, isPretty)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	}
}

// ListAlertRecipientsAction retrieves specified alert recipients from OpsGenie.
func ListAlertRecipientsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}
	req := alertsv2.ListAlertRecipientsRequest{
		Identifier: &alertsv2.Identifier{},
	}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}

	printVerboseMessage("List alert recipients request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.ListAlertRecipients(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	outputFormat := strings.ToLower(c.String("output-format"))
	printVerboseMessage("Alert recipients listed successfully, and will print as " + outputFormat)
	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp.Recipients)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp.Recipients, isPretty)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", output)
	}
}

// UnAcknowledgeAction unacknowledges an alert at OpsGenie.
func UnAcknowledgeAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.UnacknowledgeRequest{
		Identifier: &alertsv2.Identifier{},

	}
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

	printVerboseMessage("Unacknowledge alert request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Unacknowledge(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Alert will be unacknowledged. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// SnoozeAction snoozes an alert at OpsGenie.
func SnoozeAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.SnoozeRequest{}
	req.Identifier = &alertsv2.Identifier{}

	if _, success := getVal("timezone", c); success {
		printWarningMessage("ERROR: timezone is deprecated and ignoring please use ISO8601 format for `endDate` param to define timezone")
		os.Exit(1)
	}

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

	if val, success := getVal("endDate", c); success {

		endTime, err := time.Parse(time.RFC3339, val)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}

		req.EndTime = endTime
	}
	printVerboseMessage("Snooze request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.Snooze(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("will be snoozed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// RemoveTagsAction removes tags from an alert at OpsGenie.
func RemoveTagsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.RemoveTagsRequest{
		Identifier: &alertsv2.Identifier{},
	}
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

	printVerboseMessage("Remove tags request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.RemoveTags(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Tags will be removed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// AddDetailsAction adds details to an alert at OpsGenie.
func AddDetailsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.AddDetailsRequest{
		Identifier: &alertsv2.Identifier{},
	}
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
	if c.IsSet("D") {
		req.Details = extractDetailsFromCommand(c)
	}
	printVerboseMessage("Add details request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.AddDetails(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Details will be added. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// RemoveDetailsAction removes details from an alert at OpsGenie.
func RemoveDetailsAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.RemoveDetailsRequest{
		Identifier: &alertsv2.Identifier{},
	}
	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("keys", c); success {
		req.Keys = strings.Split(val, ",")
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Remove details request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.RemoveDetails(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Details will be removed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}

// EscalateToNextAction processes the next available rule in the specified escalation.
func EscalateToNextAction(c *gcli.Context) {
	cli, err := NewAlertClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := alertsv2.EscalateToNextRequest{
		Identifier: &alertsv2.Identifier{},
	}

	if val, success := getVal("id", c); success {
		req.ID = val
	}
	if val, success := getVal("alias", c); success {
		req.Alias = val
	}
	if val, success := getVal("escalationId", c); success {
		req.Escalation.ID = val
	}
	if val, success := getVal("escalationName", c); success {
		req.Escalation.Name = val
	}
	req.User = grabUsername(c)
	if val, success := getVal("source", c); success {
		req.Source = val
	}
	if val, success := getVal("note", c); success {
		req.Note = val
	}

	printVerboseMessage("Escalate to next request prepared from flags, sending request to OpsGenie..")

	resp, err := cli.EscalateToNext(req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	printVerboseMessage("Escalated to next request will be processed. RequestID: " + resp.RequestID)
	fmt.Println("RequestID: " + resp.RequestID)
}
