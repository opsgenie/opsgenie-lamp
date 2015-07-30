// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-lamp/command"
)

const lampVersion string = "1.0.0"

var commonFlags = []gcli.Flag{
	gcli.BoolFlag{
		Name:  "v",
		Usage: "Execute commands in verbose mode",
	},
	gcli.StringFlag{
		Name:  "apiKey",
		Usage: "API key used for authenticating API requests. If not given, the api key in the conf file is used",
	},
	gcli.StringFlag{
		Name:  "user",
		Usage: "Owner of the execution",
	},
	gcli.StringFlag{
		Name:  "config",
		Usage: "Configuration file path",
	},
}

func createAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "message",
			Usage: "Alert text limited to 130 characters",
		},
		gcli.StringFlag{
			Name:  "recipients",
			Usage: "The user names of individual users or names of groups",
		},
		gcli.StringFlag{
			Name:  "teams",
			Usage: "A comma seperated list of teams",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "A user defined identifier for the alert and there can be only one alert with open status with the same alias.",
		},
		gcli.StringFlag{
			Name:  "actions",
			Usage: "A comma separated list of actions that can be executed",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Field to specify source of alert. By default, it will be assigned to IP address of incoming request",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Alert text in long form. Unlike the message field, not limited to 130 characters",
		},
		gcli.StringFlag{
			Name:  "entity",
			Usage: "The entity the alert is related to",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringSliceFlag{
			Name:  "D",
			Usage: "Additional alert properties.\n\tSyntax: -D key=value",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createAlert",
		Flags:  flags,
		Usage:  "Creates an alert at OpsGenie",
		Action: command.CreateAlertAction,
	}
	return cmd
}

func getAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
		gcli.BoolFlag{
			Name:  "pretty",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getAlert",
		Flags:  flags,
		Usage:  "Gets an alert content from OpsGenie",
		Action: command.GetAlertAction,
	}
	return cmd
}

func attachFileCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file will be attached. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the file will be attached. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "attachment",
			Usage: "Absolute or relative path to the file",
		},
		gcli.StringFlag{
			Name:  "indexFile",
			Usage: "",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "attachFile",
		Flags:  flags,
		Usage:  "Attaches files to an alert",
		Action: command.AttachFileAction,
	}
	return cmd
}

func acknowledgeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be acknowledged. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be acknowledged. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "acknowledge",
		Flags:  flags,
		Usage:  "Acknowledges an alert at OpsGenie",
		Action: command.AcknowledgeAction,
	}
	return cmd

}

func renotifyCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that recipient will be renotified for. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that recipient will be renotified for. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "recipients",
			Usage: "The user names of individual users or names of groups that will be renotified for alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "renotify",
		Flags:  flags,
		Usage:  "Renotifies recipients at OpsGenie.",
		Action: command.RenotifyAction,
	}
	return cmd
}

func takeOwnershipCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be owned. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be owned. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "takeOwnership",
		Flags:  flags,
		Usage:  "Takes the ownership of an alert at OpsGenie.",
		Action: command.TakeOwnershipAction,
	}
	return cmd
}

func assignOwnerCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be owned. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be owned. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "owner",
			Usage: "The users who will be the owner of the alert after the execution",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "assign",
		Flags:  flags,
		Usage:  "Assigns the ownership of an alert to the specified user.",
		Action: command.AssignOwnerAction,
	}
	return cmd
}

func addTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new team will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new team will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "team",
			Usage: "The team that will be added to the alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addTeam",
		Flags:  flags,
		Usage:  "Adds a new team to an alert.",
		Action: command.AddTeamAction,
	}
	return cmd
}

func addRecipientCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new recipient will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new recipient will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "recipient",
			Usage: "The recipient that will be added to the alert",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addRecipient",
		Flags:  flags,
		Usage:  "Adds a new recipient to an alert.",
		Action: command.AddRecipientAction,
	}
	return cmd
}

func addNoteCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be retrieved. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addNote",
		Flags:  flags,
		Usage:  "Adds a user comment for an alert.",
		Action: command.AddNoteAction,
	}
	return cmd
}

func addTagsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new tags will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the new tags will be added. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "A comma separated list of labels attached to the alert.",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addTags",
		Flags:  flags,
		Usage:  "Adds tags to an alert.",
		Action: command.AddTagsAction,
	}
	return cmd
}

func executeActionCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the action will be executed on. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that the action will be executed on. Either id or alias must be provided. Alias option can only be used open alerts",
		},
		gcli.StringFlag{
			Name:  "action",
			Usage: "Action to execute",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "executeAction",
		Flags:  flags,
		Usage:  "Executes alert actions at OpsGenie",
		Action: command.ExecuteActionAction,
	}
	return cmd
}

func closeAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId,id",
			Usage: "Id of the alert that will be closed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the alert that will be closed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "notify",
			Usage: "Comma separated list of user and groups which will be notified. Also special values \"all\", \"recipients\" and \"owner\" is accepted",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note text",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "closeAlert",
		Flags:  flags,
		Usage:  "Closes an alert at OpsGenie",
		Action: command.CloseAlertAction,
	}
	return cmd
}

func deleteAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be deleted",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteAlert",
		Flags:  flags,
		Usage:  "Deletes an alert at OpsGenie.",
		Action: command.DeleteAlertAction,
	}
	return cmd
}

func heartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat on OpsGenie",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "heartbeat",
		Flags:  flags,
		Usage:  "Sends heartbeat to OpsGenie",
		Action: command.HeartbeatAction,
	}
	return cmd
}

func enableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "enable",
		Flags:  flags,
		Usage:  "Enables OpsGenie Integration and Policy.",
		Action: command.EnableAction,
	}
	return cmd
}

func disableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be disabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the integration/policy that will be disabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "disable",
		Flags:  flags,
		Usage:  "Disables OpsGenie Integration and Policy.",
		Action: command.DisableAction,
	}
	return cmd
}

func initCommands(app *gcli.App) {
	app.Commands = []gcli.Command{
		createAlertCommand(),
		getAlertCommand(),
		attachFileCommand(),
		acknowledgeCommand(),
		renotifyCommand(),
		takeOwnershipCommand(),
		assignOwnerCommand(),
		addTeamCommand(),
		addRecipientCommand(),
		addTagsCommand(),
		addNoteCommand(),
		executeActionCommand(),
		closeAlertCommand(),
		deleteAlertCommand(),
		heartbeatCommand(),
		enableCommand(),
		disableCommand(),
	}
}

func main() {
	app := gcli.NewApp()
	app.Name = "lamp"
	app.Version = lampVersion
	app.Usage = "Command line interface for OpsGenie"
	app.Author = "OpsGenie"
	app.Action = func(c *gcli.Context) {
		fmt.Printf("Run 'lamp help' for the options\n")
	}
	initCommands(app)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error occured while executing command: %s\n", err.Error())
	}
}
