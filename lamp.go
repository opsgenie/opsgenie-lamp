// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"fmt"
	gcli "github.com/codegangsta/cli"
	command "github.com/opsgenie/opsgenie-lamp/command"
)

const LAMP_VERSION string = "1.0.0"

func getCreateAlertCommand() gcli.Command {
	cmd := gcli.Command{	Name: "createAlert",
							Flags: []gcli.Flag{
								gcli.StringFlag{
									Name: "message",
									Usage: "Alert text limited to 130 characters",
								},
								gcli.StringFlag{
									Name: "recipients",
									Usage: "The user names of individual users or names of groups",
								},
								gcli.StringFlag{
									Name: "alias",
									Usage: "A user defined identifier for the alert and there can be only one alert with open status with the same alias.",
								},
								gcli.StringFlag{
									Name: "actions",
									Usage: "A comma separated list of actions that can be executed",
								},
								gcli.StringFlag{
									Name: "source",
									Usage: "Field to specify source of alert. By default, it will be assigned to IP address of incoming request",
								},
								gcli.StringFlag{
									Name: "tags",
									Usage: "A comma separated list of labels attached to the alert",
								},				
								gcli.StringFlag{
									Name: "description",
									Usage:"Alert text in long form. Unlike the message field, not limited to 130 characters",
								},
								gcli.StringFlag{
									Name: "entity",
									Usage:"The entity the alert is related to",
								},
								gcli.StringFlag{
									Name: "user",
									Usage:"Owner of the execution",
								},					
								gcli.StringFlag{
									Name: "note",
									Usage:"Additional alert note",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
					},
					Usage: "Create alerts in OpsGenie",
					Action: command.CreateAlertAction,
				}
	return cmd
}


func getGetAlertCommand() gcli.Command {
	cmd := gcli.Command {	Name: "getAlert",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be retrieved. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be retrieved. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"output-format",
									Value: "json",
									Usage: "Prints the output in json or yaml formats",
								},		
								gcli.BoolFlag{
									Name:"pretty",
									Usage: "For more readable JSON output",
								},					
							},
							Usage: "Get alerts details from OpsGenie",
							Action: command.GetAlertAction,
						}
	return cmd
}

func getAttachFileCommand() gcli.Command {
	cmd := gcli.Command {	Name: "attachFile",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that the file will be attached. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that the file will be attached. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"attachment",
									Usage: "Absolute or relative path to the file",
								},	
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Attach file to alerts in OpsGenie",
							Action: command.AttachFileAction,
						}
	return cmd
}


func getAcknowledgeCommand() gcli.Command {
	cmd := gcli.Command {	Name: "acknowledge",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be acknowledged. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be acknowledged. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Acknowledge an alert in OpsGenie",
							Action: command.AcknowledgeAction,
						}
	return cmd
	
}

func getRenotifyCommand() gcli.Command {
	cmd := gcli.Command {	Name: "renotify",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that recipient will be renotified for. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that recipient will be renotified for. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"recipients",
									Usage: "The user names of individual users or names of groups that will be renotified for alert",
								},											
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Renotify alert recipients. If no recipient specified, all existing alert recipients will be renotified.",
							Action: command.RenotifyAction,
						}
	return cmd
}

func getTakeOwnershipCommand() gcli.Command {
	cmd := gcli.Command {	Name: "takeOwnership",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be owned. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be owned. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},										
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Take ownership of an alert in OpsGenie",
							Action: command.TakeOwnershipAction,
						}
	return cmd
}

func getAssignOwnerCommand() gcli.Command {
	cmd := gcli.Command {	Name: "assign",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be owned. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be owned. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},	
								gcli.StringFlag{
									Name:"owner",
									Usage: "The users who will be the owner of the alert after the execution",
								},																		
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Take ownership of an alert in OpsGenie",
							Action: command.AssignOwnerAction,
						}
	return cmd
}


func getAddTeamCommand() gcli.Command {
	cmd := gcli.Command {	Name: "addTeam",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that the new team will be added. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that the new team will be added. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"team",
									Usage: "The team that will be added to the alert",
								},											
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Add a new team to an alert in OpsGenie",
							Action: command.AddTeamAction,
						}
	return cmd
}

func getAddRecipientCommand() gcli.Command {
	cmd := gcli.Command {	Name: "addRecipient",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that the new recipient will be added. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that the new recipient will be added. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},
								gcli.StringFlag{
									Name:"recipient",
									Usage: "The recipient that will be added to the alert",
								},											
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Additional alert note",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Add a new recipient to an alert in OpsGenie",
							Action: command.AddRecipientAction,
						}
	return cmd
}

func getAddNoteCommand() gcli.Command {
	cmd := gcli.Command {	Name: "addNote",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be retrieved. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be retrieved. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},									
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Note text",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Add notes to an alert in OpsGenie",
							Action: command.AddNoteAction,
						}
	return cmd	
}

func getExecuteActionCommand() gcli.Command {
	cmd := gcli.Command {	Name: "executeAction",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that the action will be executed on. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that the action will be executed on. Either alertId or alias must be provided. Alias option can only be used open alerts",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},		
								gcli.StringFlag{
									Name:"action",
									Usage: "Action to execute",
								},																		
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Note text",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Execute a custom action on an alert in OpsGenie",
							Action: command.ExecuteActionAction,
						}
	return cmd	
}

func getCloseAlertCommand() gcli.Command {
	cmd := gcli.Command {	Name: "closeAlert",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be closed. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name:"alias",
									Usage: "Alias of the alert that will be closed. Either alertId or alias must be provided",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},		
								gcli.StringFlag{
									Name:"notify",
									Usage: "Comma separated list of user and groups which will be notified. Also special values \"all\", \"recipients\" and \"owner\" is accepted",
								},																		
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"note",
									Usage: "Note text",
								},	
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Execute a custom action on an alert in OpsGenie",
							Action: command.CloseAlertAction,
						}
	return cmd	
}

func getDeleteAlertCommand() gcli.Command {
	cmd := gcli.Command {	Name: "deleteAlert",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"alertId",
									Usage:"Id of the alert that will be deleted",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},														
								gcli.StringFlag{
									Name:"user",
									Usage: "Owner of execution",
								},			
								gcli.StringFlag{
									Name:"source",
									Usage: "Source of action",
								},	
							},
							Usage: "Execute a custom action on an alert in OpsGenie",
							Action: command.DeleteAlertAction,
						}
	return cmd	
}

func getHeartbeatCommand() gcli.Command {
		cmd := gcli.Command {	Name: "heartbeat",
							Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"name",
									Usage:"Name of the heartbeat on OpsGenie",
								},
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},									
							},
							Usage: "Send periodic heartbeat messages to OpsGenie",
							Action: command.HeartbeatAction,
						}
	return cmd	
}

func getEnableIntegrationCommand() gcli.Command {
		cmd := gcli.Command {	Name: "enable",
								Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"id",
									Usage:"Id of the integration/policy that will be enabled. Either id or name must be provided",
								},
								gcli.StringFlag{
									Name:"name",
									Usage:"Name of the integration/policy that will be enabled. Either id or name must be provided",
								},
								gcli.StringFlag{
									Name:"type",
									Usage:"integration or policy",
								},								
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},									
							},
							Usage: "Enable integrations and policies",
							Action: command.EnableIntegrationAction,
						}
	return cmd	
}

func getDisableIntegrationCommand() gcli.Command {
		cmd := gcli.Command {	Name: "disable",
								Flags: []gcli.Flag {
								gcli.StringFlag{
									Name:"id",
									Usage:"Id of the integration/policy that will be disabled. Either id or name must be provided",
								},
								gcli.StringFlag{
									Name:"name",
									Usage:"Name of the integration/policy that will be disabled. Either id or name must be provided",
								},
								gcli.StringFlag{
									Name:"type",
									Usage:"integration or policy",
								},									
								gcli.StringFlag{
									Name: "apiKey",
									Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
								},									
							},
							Usage: "Disable integrations and policies",
							Action: command.DisableIntegrationAction,
						}
	return cmd	
}


func initCommands(app *gcli.App) {
	app.Commands = []gcli.Command {
		// create alert command
		getCreateAlertCommand(),
		getGetAlertCommand(),
		getAttachFileCommand(),
		getAcknowledgeCommand(),
		getRenotifyCommand(),
		getTakeOwnershipCommand(),
		getAssignOwnerCommand(),
		getAddTeamCommand(),
		getAddRecipientCommand(),
		getAddNoteCommand(),
		getExecuteActionCommand(),
		getCloseAlertCommand(),
		getDeleteAlertCommand(),
		getHeartbeatCommand(),
		getEnableIntegrationCommand(),
		getDisableIntegrationCommand(),
	}
}


func main() {
	app := gcli.NewApp()
	app.Name = "lamp"
	app.Version = LAMP_VERSION
	app.Usage = "Command line interface for OpsGenie"
	app.Author = "OpsGenie"
	app.Action = func(c *gcli.Context) {
		fmt.Println("Run 'lamp help' for the options")
	}
	initCommands(app)
	app.Run(os.Args)
}

