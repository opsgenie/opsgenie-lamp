// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package lamp

import(
	"fmt"
	gcli "github.com/codegangsta/cli"
	"log"
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	// ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"strings"
	// "errors"

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
		log.Fatalln(msg)
	}
	// message can not be longer than MAX_ALERT_MSG_LENGTH chars
	if len(c.String("message")) > MAX_ALERT_MSG_LENGTH {
		log.Fatalln( fmt.Sprintf("Alert message can not be longer than %d characters", MAX_ALERT_MSG_LENGTH) )
	}

	recipientsArr := strings.Split( c.String("recipients"), "," )
	// get a client instance using an api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}
	// send the request
	resp, err := cli.Create(req)	
	if err != nil {
		log.Fatalln("Unable to create the alert: " + err.Error() )
	}	
	log.Println("Alert created with the ID: " + resp.AlertId)
}


func GetAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
		log.Fatalln("Unable to get the alert")
	}
	// output
	outputFormat := strings.ToLower(c.String("output-format"))
	switch outputFormat {
		case "yaml": 
			output, err := ResultToYaml(resp) 
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Println( output )
		default:
			isPretty := c.IsSet("pretty")
			output, err := ResultToJson(resp, isPretty) 
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Println( output )
	}
}

func AttachFileAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, attachment (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}
	if !c.IsSet("attachment") {
		log.Fatalln("Attachment file must be given")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AttachFile(req)
	if err != nil {
		log.Fatalln( fmt.Sprintf("Unable to attach the file %s", c.String("attachment")) )
	}	
	log.Println( fmt.Sprintf("%s attached successfuly", c.String("attachment")) )
}

func AcknowledgeAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the attach-file request
	req := alerts.AcknowledgeAlertRequest{}
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.Acknowledge(req)
	if err != nil {
		log.Fatalln("Could not acknowledge the alert")
	}
	log.Println("Acknowledged successfuly")
}


func RenotifyAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.Renotify(req)
	if err != nil {
		log.Fatalln("Could not renotify the recipient(s)")
	}
	log.Println("Renotified successfuly")
}

func TakeOwnershipAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the renotify request
	req := alerts.TakeOwnershipAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.TakeOwnership(req)
	if err != nil {
		log.Fatalln("Could not take the ownership")
	}
	log.Println("Ownership taken successfuly")
}

func AssignOwnerAction(c *gcli.Context) {
		// mandatory arguments: alertId/alias, owner (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	if !c.IsSet("owner") {
		log.Fatalln("Owner should be provided, it can not be empty")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AssignOwner(req)
	if err != nil {
		log.Fatalln("Could not assign the ownership")
	}
	log.Println("Ownership assigned successfuly")
}

func AddTeamAction(c *gcli.Context) {
			// mandatory arguments: alertId/alias, team (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	if !c.IsSet("team") {
		log.Fatalln("Team should be provided, it can not be empty")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddTeam(req)
	if err != nil {
		log.Fatalln("Could not add team")
	}
	log.Println("Team added successfuly")
}


func AddRecipientAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, recipient (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	if !c.IsSet("recipient") {
		log.Fatalln("Recipient should be provided, it can not be empty")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddRecipient(req)
	if err != nil {
		log.Fatalln("Could not add recipient")
	}
	log.Println("Recipient added successfuly")	
}

func AddNoteAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, note (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	if !c.IsSet("note") {
		log.Fatalln("Note argument should be provided, it can not be empty")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the add-team request
	req := alerts.AddNoteAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.AddNote(req)
	if err != nil {
		log.Fatalln("Could not add note")
	}
	log.Println("Note added successfuly")	
}

func ExecuteActionAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias, action (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	if !c.IsSet("action") {
		log.Fatalln("Note argument should be provided, it can not be empty")
	}
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
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
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("note") {
		req.Note = c.String("note")
	}	
	// send the request
	_, err = cli.ExecuteAction(req)
	if err != nil {
		log.Fatalln("Could not execute the action")
	}
	log.Println(fmt.Sprintf("Action '%s' executed successfuly", c.String("action")))	
}

func CloseAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId/alias (apiKey may be given by the configuration file)
	if c.IsSet("alertId") && c.IsSet("alias") {
		log.Fatalln("Either alert id or alias must be provided, not both")
	}
	if !c.IsSet("alertId") && !c.IsSet("alias") {
		log.Fatalln("At least one of the alert id and alias must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the add-team request
	req := alerts.CloseAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} else if c.IsSet("alias") {
		req.Alias = c.String("alias")
	}
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}
	if c.IsSet("notify") {
		req.Notify = strings.Split( c.String("notify"), "," )
	}	
	// send the request
	_, err = cli.Close(req)
	if err != nil {
		log.Fatalln("Could not close the alert")
	}
	log.Println("Alert closed successfuly")	
}

func DeleteAlertAction(c *gcli.Context) {
	// mandatory arguments: alertId (apiKey may be given by the configuration file)
	if !c.IsSet("alertId") {
		log.Fatalln("Alert id must be provided")
	}	
	// get a client instance using the api key
	cli, err := NewAlertClient( grabApiKey(c) )	
	if err != nil {
		log.Fatalln(err.Error())
	}
	// build the add-team request
	req := alerts.DeleteAlertRequest{}	
	if c.IsSet("alertId") {
		req.AlertId = c.String("alertId")		
	} 
	if c.IsSet("user") {
		req.User = c.String("user")
	}
	if c.IsSet("source") {
		req.Source = c.String("source")
	}	
	// send the request
	_, err = cli.Delete(req)
	if err != nil {
		log.Fatalln("Could not delete the alert")
	}
	log.Println( fmt.Sprintf("Alert with id of %s deleted successfuly", c.String("alertId")) )	
}
