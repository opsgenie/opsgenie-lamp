// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package command

import(
	"fmt"
	gcli "github.com/codegangsta/cli"
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"strings"
	// "errors"
	// "log"
	"os"
)

const MAX_ALERT_MSG_LENGTH int = 130

func isEmpty(args []string, c *gcli.Context) (bool, string) {
	for _, arg := range args {
		if !c.IsSet(arg) {
			return true, fmt.Sprintf("Argument '%s' is missing", arg)
		}
	}
	return false, "" 
}

func CreateAlertAction(c *gcli.Context) {

	// mandatory arguments: message, recipients (apiKey may be given by the configuration file)
	if empty, msg := isEmpty([]string{"message", "recipients"}, c); empty == true {
		cmdlog.Error(msg, "command", "createAlert")
		gcli.ShowCommandHelp(c, "createAlert")
		os.Exit(1)
	}
	// message can not be longer than MAX_ALERT_MSG_LENGTH chars
	if len(c.String("message")) > MAX_ALERT_MSG_LENGTH {
		cmdlog.Error( fmt.Sprintf("Alert message can not be longer than %d characters", MAX_ALERT_MSG_LENGTH), "command", "createAlert" )
		gcli.ShowCommandHelp(c, "createAlert")
		os.Exit(1)
	}

	recipientsArr := strings.Split( c.String("recipients"), "," )
	// get a client instance using an api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "createAlert")
		os.Exit(1)
	}
	// create the alert
	req := alerts.CreateAlertRequest{}
	req.Message = c.String("message")
	req.Recipients = recipientsArr
	// set the parameters
	if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("actions") {
		req.Actions = strings.Split(c.String("actions"), ",")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("tags") {
		req.Tags = strings.Split(c.String("tags"), ",")
	}
	if c.IsSet("description") {
		req.Description = c.String("description")
	}
	if c.IsSet("entity") {
		req.Entity = c.String("entity")
	}
	req.User = grabUsername(c);
	if c.IsSet("note") {
		req.Note = c.String("note")
	}
	// send the request
	resp, err := cli.Create(req)	
	if err != nil {
		cmdlog.Error("Unable to create the alert: " + err.Error() , "command", "createAlert")
		os.Exit(1)
	}	
	cmdlog.Info("Alert created with the ID: " + resp.AlertId, "command", "createAlert")
}


func GetAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "getAlert")
		gcli.ShowCommandHelp(c, "getAlert")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "getAlert")
		gcli.ShowCommandHelp(c, "getAlert")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "getAlert")
		os.Exit(1)
	}
	// build the get-alert request
	req := alerts.GetAlertRequest{}
	if c.IsSet("alertId") {
		req.Id = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	// send the request
	resp, err := cli.Get(req)
	if err != nil {
		cmdlog.Error("Unable to get the alert." + err.Error(), "command", "getAlert")
		os.Exit(1)
	}
	// output
	outputFormat := strings.ToLower(c.String("output-format"))
	switch outputFormat {
		case "yaml": 
			output, err := ResultToYaml(resp) 
			if err != nil {
				cmdlog.Error(err.Error(), "command", "getAlert")
				os.Exit(1)
			}
			cmdlog.Info( output )
		default:
			isPretty := c.IsSet("pretty")
			output, err := ResultToJson(resp, isPretty) 
			if err != nil {
				cmdlog.Error(err.Error())
				os.Exit(1)
			}
			cmdlog.Info( output , "command", "getAlert")
	}
}

func AttachFileAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, attachment (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "attachFile")
		gcli.ShowCommandHelp(c, "attachFile")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "attachFile")
		gcli.ShowCommandHelp(c, "attachFile")
		os.Exit(1)
	}
	if !c.IsSet("attachment") {
		cmdlog.Error("Attachment file must be given", "command", "attachFile")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "attachFile")
		os.Exit(1)
	}
	// build the attach-file request
	req := alerts.AttachFileAlertRequest{}
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("attachment") {
		req.Attachment = c.String("attachment")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AttachFile(req)
	if err != nil {
		cmdlog.Error( fmt.Sprintf("Unable to attach the file %s", c.String("attachment")) , "command", "attachFile")
		os.Exit(1)
	}	
	cmdlog.Info( fmt.Sprintf("%s attached successfuly", c.String("attachment")) , "command", "attachFile")
}

func AcknowledgeAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "acknowledge")
		gcli.ShowCommandHelp(c, "acknowledge")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "acknowledge")
		gcli.ShowCommandHelp(c, "acknowledge")
		os.Exit(1)
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "acknowledge")
		os.Exit(1)
	}
	// build the attach-file request
	req := alerts.AcknowledgeAlertRequest{}
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.Acknowledge(req)
	if err != nil {
		cmdlog.Error("Could not acknowledge the alert", "command", "acknowledge")
		os.Exit(1)
	}
	cmdlog.Info("Acknowledged successfuly", "command", "acknowledge")	
}


func RenotifyAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "renotify")
		gcli.ShowCommandHelp(c, "renotify")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "renotify")
		gcli.ShowCommandHelp(c, "renotify")
		os.Exit(1)
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error())
		os.Exit(1)
	}
	// build the renotify request
	req := alerts.RenotifyAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("recipients") {
		req.Recipients = strings.Split(c.String("recipients"), ",")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.Renotify(req)
	if err != nil {
		cmdlog.Error("Could not renotify the recipient(s)", "command", "renotify")
		os.Exit(1)
	}
	cmdlog.Info("Renotified successfuly", "command", "renotify")
}

func TakeOwnershipAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "takeOwnership")
		gcli.ShowCommandHelp(c, "takeOwnership")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "takeOwnership")
		gcli.ShowCommandHelp(c, "takeOwnership")
		os.Exit(1)
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error())
		os.Exit(1)
	}
	// build the renotify request
	req := alerts.TakeOwnershipAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.TakeOwnership(req)
	if err != nil {
		cmdlog.Error("Could not take the ownership", "command", "takeOwnership")
		os.Exit(1)
	}
	cmdlog.Info("Ownership taken successfuly", "command", "takeOwnership")
}

func AssignOwnerAction(c *gcli.Context) {
		// mandatory arguments: alertId/alias, owner (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "assign")
		gcli.ShowCommandHelp(c, "assign")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "assign")
		gcli.ShowCommandHelp(c, "assign")
		os.Exit(1)
	}	
	if !c.IsSet("owner") {
		cmdlog.Error("Owner should be provided, it can not be empty", "command", "assign")
		gcli.ShowCommandHelp(c, "assign")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "assign")
		os.Exit(1)
	}
	// build the renotify request
	req := alerts.AssignOwnerAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("owner") {
		req.Owner = c.String("owner")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AssignOwner(req)
	if err != nil {
		cmdlog.Error("Could not assign the ownership", "command", "assign")
		os.Exit(1)
	}
	cmdlog.Info("Ownership assigned successfuly", "command", "assign")
}

func AddTeamAction(c *gcli.Context) {
			// mandatory arguments: alertId/alias, team (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "addTeam")
		gcli.ShowCommandHelp(c, "addTeam")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "addTeam")
		gcli.ShowCommandHelp(c, "addTeam")
		os.Exit(1)
	}	
	if !c.IsSet("team") {
		cmdlog.Error("Team should be provided, it can not be empty", "command", "addTeam")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error())
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.AddTeamAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("team") {
		req.Team = c.String("team")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddTeam(req)
	if err != nil {
		cmdlog.Error("Could not add team", "command", "addTeam")
		os.Exit(1)
	}
	cmdlog.Info("Team added successfuly", "command", "addTeam")
}


func AddRecipientAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, recipient (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "addRecipient")
		gcli.ShowCommandHelp(c, "addRecipient")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "addRecipient")
		gcli.ShowCommandHelp(c, "addRecipient")
		os.Exit(1)
	}	
	if !c.IsSet("recipient") {
		cmdlog.Error("Recipient should be provided, it can not be empty", "command", "addRecipient")
		gcli.ShowCommandHelp(c, "addRecipient")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "addRecipient")
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.AddRecipientAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("recipient") {
		req.Recipient = c.String("recipient")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddRecipient(req)
	if err != nil {
		cmdlog.Error("Could not add recipient", "command", "addRecipient")
		os.Exit(1)
	}
	cmdlog.Info("Recipient added successfuly", "command", "addRecipient")	
}

func AddNoteAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, note (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "addNote")
		gcli.ShowCommandHelp(c, "addNote")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "addNote")
		gcli.ShowCommandHelp(c, "addNote")
		os.Exit(1)
	}	
	if !c.IsSet("note") {
		cmdlog.Error("Note argument should be provided, it can not be empty", "command", "addNote")
		gcli.ShowCommandHelp(c, "addNote")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error())
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.AddNoteAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddNote(req)
	if err != nil {
		cmdlog.Error("Could not add note", "command", "addNote")
		os.Exit(1)
	}
	cmdlog.Info("Note added successfuly", "command", "addNote")	
}

func ExecuteActionAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, action (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "executeAction")
		gcli.ShowCommandHelp(c, "executeAction")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "executeAction")
		gcli.ShowCommandHelp(c, "executeAction")
		os.Exit(1)
	}	
	if !c.IsSet("action") {
		cmdlog.Error("Note argument should be provided, it can not be empty", "command", "executeAction")
		gcli.ShowCommandHelp(c, "executeAction")
		os.Exit(1)
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "executeAction")
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.ExecuteActionAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("action") {
		req.Action = c.String("action")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.ExecuteAction(req)
	if err != nil {
		cmdlog.Error("Could not execute the action", "command", "executeAction")
		os.Exit(1)
	}
	cmdlog.Info(fmt.Sprintf("Action '%s' executed successfuly", c.String("action")), "command", "executeAction")	
}

func CloseAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		cmdlog.Error("Either alert id or alias must be provided, not both", "command", "closeAlert")
		gcli.ShowCommandHelp(c, "closeAlert")
		os.Exit(1)
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		cmdlog.Error("At least one of the alert id and alias must be provided", "command", "closeAlert")
		gcli.ShowCommandHelp(c, "closeAlert")
		os.Exit(1)
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "closeAlert")
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.CloseAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("notify") {
		req.Notify = strings.Split( c.String("notify"), "," )
	}	
	// send the request
	_, err = cli.Close(req)
	if err != nil {
		cmdlog.Error("Could not close the alert", "command", "closeAlert")
		os.Exit(1)
	}
	cmdlog.Info("Alert closed successfuly", "command", "closeAlert")	
}

func DeleteAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId (apiKey may be given by the configuration file)
	if !c.IsSet("alertId") {
		cmdlog.Error("Alert id must be provided", "command", "deleteAlert")
		gcli.ShowCommandHelp(c, "deleteAlert")
		os.Exit(1)
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		cmdlog.Error(err.Error(), "command", "deleteAlert")
		os.Exit(1)
	}
	// build the add-team request
	req := alerts.DeleteAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} 
	req.User = grabUsername(c);
	if c.IsSet("source") {
		req.Source = c.String("source")
	}	
	// send the request
	_, err = cli.Delete(req)
	if err != nil {
		cmdlog.Error("Could not delete the alert", "command", "deleteAlert")
		os.Exit(1)
	}
	cmdlog.Info( fmt.Sprintf("Alert with id of %s deleted successfuly", c.String("alertId")) , "command", "deleteAlert")	
}
