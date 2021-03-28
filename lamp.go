// Copyright 2015 Opsgenie. All rights reserved.
// Use of this source code is governed by a Apache Software
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/opsgenie/opsgenie-lamp/command"
	gcli "github.com/urfave/cli"
	"os"
)

const lampVersion string = "3.2.0"

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

var renderingFlags = []gcli.Flag{
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

func createAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "message",
			Usage: "Alert text limited to 130 characters",
		},
		gcli.StringFlag{
			Name:  "teams",
			Usage: "A comma separated list of teams",
		},
		gcli.StringFlag{
			Name:  "users",
			Usage: "A comma separated list of users",
		},
		gcli.StringFlag{
			Name:  "escalations",
			Usage: "A comma separated list of escalations",
		},
		gcli.StringFlag{
			Name:  "schedules",
			Usage: "A comma separated list of schedules",
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
		gcli.StringFlag{
			Name:  "priority",
			Usage: "The priority of alert. Values: P1, P2, P3, P4, P5 default is P3",
		},
		gcli.StringSliceFlag{
			Name:  "D",
			Usage: "Additional alert properties.\n\tSyntax: -D key=value",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createAlert",
		Flags: flags,
		Usage: "Creates an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CreateAlertAction(c)
			return nil
		},
	}
	return cmd
}

func getAlertCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getAlert",
		Flags: flags,
		Usage: "Gets an alert content from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.GetAlertAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "query",
			Usage: "Search query to apply while filtering the alerts",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 20. Max value for this parameter is 100",
		},
		gcli.StringFlag{
			Name:  "offset",
			Usage: "Start index of the result set (to apply pagination). Minimum value (and also default value) is 0",
		},
		gcli.StringFlag{
			Name:  "sortBy",
			Usage: "Name of the field that result set will be sorted by. Default value is createdAt.",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default: desc",
		},
		gcli.StringFlag{
			Name:  "searchIdentifier",
			Usage: "Identifier of the saved search query to apply while filtering the alerts",
		},
		gcli.StringFlag{
			Name: "searchIdentifierType",
			Usage: "Identifier type of the value at searchIdentifier, which can be id or name. Default value is id." +
				" If searchIdentifier is not provided, this value is ignored.",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlerts",
		Flags: flags,
		Usage: "Lists alerts contents from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertsAction(c)
			return nil
		},
	}
	return cmd
}

func countAlertsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "query",
			Usage: "Search query to apply while filtering the alerts. If it is given, createdAfter, createdBefore, updatedAfter, updatedBefore, status and tags will be ignored",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 20. Max value for this parameter is 100",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "countAlerts",
		Flags: flags,
		Usage: "Counts alerts at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CountAlertsAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertNotesCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 100.",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default : desc",
		},
		gcli.StringFlag{
			Name:  "offset",
			Usage: "Starting value of the offset property.",
		},
		gcli.StringFlag{
			Name:  "direction",
			Usage: "Page direction to apply for the given offset. Possible values are next and prev. Default value is `next`",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertNotes",
		Flags: flags,
		Usage: "Lists alert notes from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertNotesAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertLogsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Page size. Default is 100.",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "asc/desc, default : desc",
		},
		gcli.StringFlag{
			Name:  "offset",
			Usage: "Starting value of the offset property.",
		},
		gcli.StringFlag{
			Name:  "direction",
			Usage: "Page direction to apply for the given offset. Possible values are next and prev. Default value is next.",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertLogs",
		Flags: flags,
		Usage: "Lists alert logs from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertLogsAction(c)
			return nil
		},
	}
	return cmd
}

func listAlertRecipientsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAlertRecipients",
		Flags: flags,
		Usage: "Lists alert recipients from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListAlertRecipientsAction(c)
			return nil
		},
	}
	return cmd
}

func unAcknowledgeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be unacknowledged. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
	cmd := gcli.Command{Name: "unacknowledge",
		Flags: flags,
		Usage: "UnAcknowledges an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.UnAcknowledgeAction(c)
			return nil
		},
	}
	return cmd

}

func snoozeCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be snoozed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "The date in ISO8601 format snooze will end",
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
	cmd := gcli.Command{Name: "snooze",
		Flags: flags,
		Usage: "Snoozes an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.SnoozeAction(c)
			return nil
		},
	}
	return cmd

}

func removeTagsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the tags will be removed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
	cmd := gcli.Command{Name: "removeTags",
		Flags: flags,
		Usage: "Removes tags from an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.RemoveTagsAction(c)
			return nil
		}}
	return cmd
}

func addDetailsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new details will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Additional alert note",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
		gcli.StringSliceFlag{
			Name:  "D",
			Usage: "Additional alert properties.\n\tSyntax: -D key=value",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addDetails",
		Flags: flags,
		Usage: "Adds details to an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.AddDetailsAction(c)
			return nil
		}}
	return cmd
}

func removeDetailsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the details will be removed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "keys",
			Usage: "Set of properties to be removed from alert details",
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
	cmd := gcli.Command{Name: "removeDetails",
		Flags: flags,
		Usage: "Removes details from an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.RemoveDetailsAction(c)
			return nil
		}}
	return cmd
}

func escalateToNextActionCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the next escalation will be processed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "escalationId",
			Usage: "Id of the escalation that will be escalated to the next level. Either escalationName or escalationId must be provided.",
		},
		gcli.StringFlag{
			Name:  "escalationName",
			Usage: "Name of the escalation that will be escalated to the next level. Either escalationName or escalationId must be provided.",
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
	cmd := gcli.Command{Name: "escalateToNext",
		Flags: flags,
		Usage: "Escalates to the next rule in the specified escalation at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.EscalateToNextAction(c)
			return nil
		}}
	return cmd
}

func attachFileCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file will be attached. Either id, alias or tinyId must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "filePath",
			Usage: "Absolute or relative path to the file",
		},
		gcli.StringFlag{
			Name:  "fileName",
			Usage: "",
		},
		gcli.StringFlag{
			Name:  "indexFile",
			Usage: "",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "attachFile",
		Flags: flags,
		Usage: "Attaches files to an alert",
		Action: func(c *gcli.Context) error {
			command.AttachFileAction(c)
			return nil
		},
	}
	return cmd
}

func getAttachmentCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file was attached. Either id, alias or tinyId must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "attachmentId",
			Usage: "Id of the alert attachment",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getAttachment",
		Flags: flags,
		Usage: "Gets the attachment download link for specified alert attachment",
		Action: func(c *gcli.Context) error {
			command.GetAttachmentAction(c)
			return nil
		},
	}
	return cmd
}

func downloadAttachmentCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file was attached. Either id, alias or tinyId must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "attachmentId",
			Usage: "Id of the alert attachment",
		},
		gcli.StringFlag{
			Name:  "destinationPath",
			Usage: "Destination path to download file to",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "downloadAttachment",
		Flags: flags,
		Usage: "Downloads the attachment for specified alert attachment",
		Action: func(c *gcli.Context) error {
			command.DownloadAttachmentAction(c)
			return nil
		},
	}
	return cmd
}

func listAttachmentsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file was attached. Either id, alias or tinyId must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "output-format",
			Value: "json",
			Usage: "Prints the output in json or yaml formats",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listAttachments",
		Flags: flags,
		Usage: "List the attachment meta information for specified alert",
		Action: func(c *gcli.Context) error {
			command.ListAlertAttachmentsAction(c)
			return nil
		},
	}
	return cmd
}

func deleteAttachmentCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the file was attached. Either id, alias or tinyId must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "attachmentId",
			Usage: "Id of the alert attachment",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteAttachment",
		Flags: flags,
		Usage: "Delete the attachment with given id for specified alert",
		Action: func(c *gcli.Context) error {
			command.DeleteAlertAttachmentAction(c)
			return nil
		},
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
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Acknowledges an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.AcknowledgeAction(c)
			return nil
		},
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
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Assigns the ownership of an alert to the specified user.",
		Action: func(c *gcli.Context) error {
			command.AssignOwnerAction(c)
			return nil
		}}
	return cmd
}

func addTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new team will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Adds a new team to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddTeamAction(c)
			return nil
		}}
	return cmd
}

func addResponderCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new recipient will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "The responder type that which is provided at responder.",
		},
		gcli.StringFlag{
			Name:  "responder",
			Usage: "The responder that will be added to the alert",
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
	cmd := gcli.Command{Name: "addResponder",
		Flags: flags,
		Usage: "Adds a new responder to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddResponderAction(c)
			return nil
		}}
	return cmd
}

func addNoteCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be retrieved. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Adds a user comment for an alert.",
		Action: func(c *gcli.Context) error {
			command.AddNoteAction(c)
			return nil
		}}
	return cmd
}

func addTagsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the new tags will be added. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Adds tags to an alert.",
		Action: func(c *gcli.Context) error {
			command.AddTagsAction(c)
			return nil
		}}
	return cmd
}

func executeActionCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that the action will be executed on. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Executes alert actions at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ExecuteActionAction(c)
			return nil
		}}
	return cmd
}

func closeAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId,id",
			Usage: "Id of the alert that will be closed. Either id or alias must be provided",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
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
		Flags: flags,
		Usage: "Closes an alert at Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CloseAlertAction(c)
			return nil
		}}
	return cmd
}

func deleteAlertCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "alertId, id",
			Usage: "Id of the alert that will be deleted",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier type of the specified id, which can be id, tiny or alias. Default value = id",
		},
		gcli.StringFlag{
			Name:  "source",
			Usage: "Source of the action",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteAlert",
		Flags: flags,
		Usage: "Deletes an alert at Opsgenie.",
		Action: func(c *gcli.Context) error {
			command.DeleteAlertAction(c)
			return nil
		}}
	return cmd
}

func pingHeartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat on Opsgenie",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "pingHeartbeat",
		Flags: flags,
		Usage: "Sends heartbeat to Opsgenie",
		Action: func(c *gcli.Context) error {
			command.PingHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func createHeartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name: "name",
			Usage: "Name of the heartbeat to be created",
		},
		gcli.StringFlag{
			Name: "description",
			Usage: "Description for the heartbeat to be created",
		},
		gcli.IntFlag{
			Name: "interval",
			Usage: "Interval after which hearbeat will expire",
		},
		gcli.StringFlag{
			Name: "intervalType",
			Usage: "Type of interval : 'm' (minute), 'h' (hours), 'd' (days)",
		},
		gcli.BoolFlag{
			Name: "enabled",
			Usage: "If present, heartbaat will be enabled",
		},
		gcli.StringFlag{
			Name: "ownerTeam",
			Usage: "Owner team for the heartbeat",
		},
		gcli.StringFlag{
			Name: "alertMessage",
			Usage: "Heartbeat alert nessage",
		},
		gcli.StringFlag{
			Name: "alertTag",
			Usage: "Tag for the heartbeat alert",
		},
		gcli.StringFlag{
			Name: "alertPriority",
			Usage: "Priority of the alert created by Heartbeat",
		},
	}

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createHeartbeat",
		Flags: flags,
		Usage: "Creates a new opsgenie heartbeat",
		Action: func(c *gcli.Context) error {
			command.CreateHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func deleteHeartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat to be deleted",
		},
	}

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteHeartbeat",
		Flags: flags,
		Usage: "Deletes opsgenie heartbeat",
		Action: func(c *gcli.Context) error {
			command.DeleteHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func disableHeartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat to be disabled",
		},
	}

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "disableHeartbeat",
		Flags: flags,
		Usage: "Disabled opsgenie heartbeat",
		Action: func(c *gcli.Context) error {
			command.DisableHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func enableHeartbeatCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the heartbeat to be Enabled",
		},
	}

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "enableHeartbeat",
		Flags: flags,
		Usage: "Enable opsgenie heartbeat",
		Action: func(c *gcli.Context) error {
			command.EnableHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func listHeartbeatCommand() gcli.Command {
	flags := append(commonFlags, renderingFlags...)
	cmd := gcli.Command{Name: "listHeartbeat",
		Flags: flags,
		Usage: "List opsgenie heartbeats",
		Action: func(c *gcli.Context) error {
			command.ListHeartbeatAction(c)
			return nil
		}}
	return cmd
}

func enableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
		gcli.StringFlag{
			Name:  "policyType",
			Usage: "Policy type should be one of alert or notification",
		},
		gcli.StringFlag{
			Name:  "teamId",
			Usage: "Team Id for policies which are created on a team",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "enable",
		Flags: flags,
		Usage: "Enables Opsgenie Integration and Policy.",
		Action: func(c *gcli.Context) error {
			command.EnableAction(c)
			return nil
		}}
	return cmd
}

func disableCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Id of the integration/policy that will be enabled. Either id or name must be provided",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "integration or policy",
		},
		gcli.StringFlag{
			Name:  "policyType",
			Usage: "Policy type should be one of alert or notification",
		},
		gcli.StringFlag{
			Name:  "teamId",
			Usage: "Team Id for team policies",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "disable",
		Flags: flags,
		Usage: "Disables Opsgenie Integration and Policy.",
		Action: func(c *gcli.Context) error {
			command.DisableAction(c)
			return nil
		}}
	return cmd
}
func downloadLogsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "start",
			Usage: "Log files starting this date",
		},
		gcli.StringFlag{
			Name:  "end",
			Usage: "Log files before this date",
		},
		gcli.StringFlag{
			Name:  "path",
			Usage: "Directory path, log files will be downloaded under this directory",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "downloadLogs",
		Flags: flags,
		Usage: "Download Logs.",
		Action: func(c *gcli.Context) error {
			command.DownloadLogs(c)
			return nil
		}}
	return cmd
}

func exportUsersCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "query",
			Usage: "Search query to apply while filtering the users",
		},
		gcli.StringFlag{
			Name:  "destinationPath",
			Usage: "Destination path to download file to",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "exportUsers",
		Flags: flags,
		Usage: "Exports user list from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ExportUsersAction(c)
			return nil
		},
	}
	return cmd
}

func fetchEscalationCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "To identify whether type of identifier it will send",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Unique Identifier for escalation",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getEscalation",
		Flags: flags,
		Usage: "get Escalation with identifier",
		Action: func(c *gcli.Context) error {
			command.GetEscalationAction(c)
			return nil
		},
	}
	return cmd
}

func deleteEscalationCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "To identify whether type of identifier it will send",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Unique Identifier for escalation",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteEscalation",
		Flags: flags,
		Usage: "Delete Escalation with identifier",
		Action: func(c *gcli.Context) error {
			command.DeleteEscalationAction(c)
			return nil
		},
	}
	return cmd
}

func CreateEscalationCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the escalation which is being created",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Description of the escalation",
		},
		gcli.StringFlag{
			Name:  "escalationCondition",
			Usage: "Condition of rules which is needed in Escalation",
		},
		gcli.StringFlag{
			Name:  "notifyType",
			Usage: "Type of group(user,team,schedule) which needs to be notified",
		},
		gcli.StringFlag{
			Name:  "particpantType",
			Usage: "Participant Type which will be notified",
		},
		gcli.StringFlag{
			Name:  "particpantName",
			Usage: "Name of participant whose participant type in mentioned in order",
		},
		gcli.StringFlag{
			Name:  "delay",
			Usage: "Delay in escalation rule after which notify will take place",
		},
		gcli.StringFlag{
			Name:  "teamName",
			Usage: "Owner Team Name of escalation",
		},
		gcli.StringFlag{
			Name:  "waitInterval",
			Usage: "Wait Interval in Repeat request for escalation",
		},
		gcli.StringFlag{
			Name:  "count",
			Usage: "Count in Repeat request before next escalation and Repeat works",
		},
		gcli.BoolFlag{
			Name:  "recipientStatus",
			Usage: "Flag for recipient status",
		},
		gcli.BoolFlag{
			Name:  "closeAlertAfterAll",
			Usage: "Check for if escaltion repeats are completed, close alert automatically",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createEscalation",
		Flags: flags,
		Usage: "Create Escalation",
		Action: func(c *gcli.Context) error {
			command.CreateEscalationAction(c)
			return nil
		},
	}
	return cmd
}

func UpdateEscalationCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the escalation which is being updated",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Description of the escalation",
		},
		gcli.StringFlag{
			Name:  "escalationCondition",
			Usage: "Condition of rules which is needed in Escalation",
		},
		gcli.StringFlag{
			Name:  "notifyType",
			Usage: "Type of group(user,team,schedule) which needs to be notified",
		},
		gcli.StringFlag{
			Name:  "particpantType",
			Usage: "Participant Type which will be notified",
		},
		gcli.StringFlag{
			Name:  "particpantName",
			Usage: "Name of participant whose participant type in mentioned in order",
		},
		gcli.StringFlag{
			Name:  "delay",
			Usage: "Delay in escalation rule after which notify will take place",
		},
		gcli.StringFlag{
			Name:  "teamName",
			Usage: "Owner Team Name of escalation",
		},
		gcli.StringFlag{
			Name:  "waitInterval",
			Usage: "Wait Interval in Repeat request for escalation",
		},
		gcli.StringFlag{
			Name:  "count",
			Usage: "Count in Repeat request before next escalation and Repeat works",
		},
		gcli.BoolFlag{
			Name:  "recipientStatus",
			Usage: "Flag for recipient status",
		},
		gcli.BoolFlag{
			Name:  "closeAlertAfterAll",
			Usage: "Check for if escalation repeats are completed, close alert automatically",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "To identify whether type of identifier it will send",
		},
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Unique Identifier for escalation",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateEscalation",
		Flags: flags,
		Usage: "Update Escalation",
		Action: func(c *gcli.Context) error {
			command.UpdateEscalationAction(c)
			return nil
		},
	}
	return cmd
}

func createTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:     "name, n",
			Usage:    "Team Name",
			Required: true,
		},
		gcli.StringFlag{
			Name:  "desc, d",
			Usage: "Description",
		},
		gcli.StringFlag{
			Name:  "userName",
			Usage: "User Name",
		},
		gcli.StringFlag{
			Name:  "userId",
			Usage: "User Id",
		},
		gcli.StringFlag{
			Name:  "role",
			Usage: "User Role",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createTeam",
		Flags: flags,
		Usage: "Create a team in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CreateTeamAction(c)
			return nil
		},
	}
	return cmd
}

func listTeamLogsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "offset, o",
			Usage: "Logs offset",
		},
		gcli.StringFlag{
			Name:  "limit, l",
			Usage: "Logs limit",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "Order of logs, asc/desc",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listTeamLogs",
		Flags: flags,
		Usage: "List Team logs in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListTeamLogsAction(c)
			return nil
		},
	}
	return cmd
}

func updateTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "desc, d",
			Usage: "Description",
		},
		gcli.StringFlag{
			Name:  "userName",
			Usage: "User Name",
		},
		gcli.StringFlag{
			Name:  "userId",
			Usage: "User Id",
		},
		gcli.StringFlag{
			Name:  "role",
			Usage: "User Role",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateTeam",
		Flags: flags,
		Usage: "Update a team in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.UpdateTeamAction(c)
			return nil
		},
	}
	return cmd
}

func getTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getTeam",
		Flags: flags,
		Usage: "Get team info in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.GetTeamAction(c)
			return nil
		},
	}
	return cmd
}

func deleteTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteTeam",
		Flags: flags,
		Usage: "Delete a team in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteTeamAction(c)
			return nil
		},
	}
	return cmd
}

func listTeamCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listTeams",
		Flags: flags,
		Usage: "List all teams in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListTeamsAction(c)
			return nil
		},
	}
	return cmd
}

func listRolesCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "listRoles",
		Flags: flags,
		Usage: "List team roles Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListRolesAction(c)
			return nil
		},
	}
	return cmd
}

func listTeamRoutingRulesCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "listRoutingRules",
		Flags: flags,
		Usage: "List team routing rules Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListTeamRoutingRulesAction(c)
			return nil
		},
	}
	return cmd
}

func deleteTeamRoutingRulesCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "ruleId",
			Usage: "Rule Id to deleted",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "deleteRoutingRule",
		Flags: flags,
		Usage: "Delete team routing rule in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteTeamRoutingRuleAction(c)
			return nil
		},
	}
	return cmd
}

func createRoleCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "id, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:     "roleName",
			Usage:    "Role Name",
			Required: true,
		},
		gcli.StringFlag{
			Name:     "rights",
			Usage:    "Role rights",
			Required: true,
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "createRole",
		Flags: flags,
		Usage: "Create a member role in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CreateRoleAction(c)
			return nil
		},
	}
	return cmd
}

func listAllRoleRightsCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "listRoleRights",
		Flags: flags,
		Usage: "Lists all available role rights in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListRoleRightsAction(c)
			return nil
		},
	}
	return cmd
}

func getRoleCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "teamName, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "teamId, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "roleName",
			Usage: "Role Name",
		},
		gcli.StringFlag{
			Name:  "roleId",
			Usage: "Role Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "getRole",
		Flags: flags,
		Usage: "Get a member role in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.GetTeamRoleAction(c)
			return nil
		},
	}
	return cmd
}

func deleteRoleCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "teamName, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "teamId, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "roleName",
			Usage: "Role Name",
		},
		gcli.StringFlag{
			Name:  "roleId",
			Usage: "Role Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "deleteRole",
		Flags: flags,
		Usage: "Delete a member role in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteTeamRoleAction(c)
			return nil
		},
	}
	return cmd
}

func addMemberCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "teamName, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "teamId, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "role",
			Usage: "Role Name",
		},
		gcli.StringFlag{
			Name:  "userId",
			Usage: "User Id",
		},
		gcli.StringFlag{
			Name:  "userName",
			Usage: "Username (email)",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "addMember",
		Flags: flags,
		Usage: "Add a member to a team in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.AddMemberAction(c)
			return nil
		},
	}
	return cmd
}

func removeMemberCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "teamName, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "teamId, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "userId",
			Usage: "User Id",
		},
		gcli.StringFlag{
			Name:  "userName",
			Usage: "Username (email)",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "removeMember",
		Flags: flags,
		Usage: "Remove a member to a team in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.RemoveMemberAction(c)
			return nil
		},
	}
	return cmd
}

func getRoutingRuleCommand() gcli.Command {
	commandFlags := []gcli.Flag{
		gcli.StringFlag{
			Name:  "teamName, n",
			Usage: "Team Name",
		},
		gcli.StringFlag{
			Name:  "teamId, i",
			Usage: "Team Id",
		},
		gcli.StringFlag{
			Name:  "ruleId",
			Usage: "Rule Id",
		},
		gcli.BoolFlag{
			Name:  "pretty, p",
			Usage: "For more readable JSON output",
		},
	}
	flags := append(commonFlags, commandFlags...)

	cmd := gcli.Command{Name: "getRoutingRule",
		Flags: flags,
		Usage: "Get a routing rule in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.GetRoutingRuleAction(c)
			return nil
		},
	}
	return cmd
}

func createScheduleCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "name, n",
			Usage: "Name of the schedule to be created",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Description of the schedule to be created",
		},
		gcli.StringFlag{
			Name:  "tz",
			Usage: "Timezone for the schedule",
		},
		gcli.StringFlag{
			Name:  "team",
			Usage: "Name of the team under which the schedule is",
		},
		gcli.BoolFlag{
			Name:  "enabled",
			Usage: "Enable the created schedule",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createSchedule",
		Flags: flags,
		Usage: "Creates a schedule on Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CreateScheduleAction(c)
			return nil
		},
	}
	return cmd
}

func getScheduleCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getSchedule",
		Flags: flags,
		Usage: "Gets a schedule's content from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.GetScheduleAction(c)
			return nil
		},
	}
	return cmd
}

func updateScheduleCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the schedule",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Description of the schedule to be created",
		},
		gcli.StringFlag{
			Name:  "tz",
			Usage: "Timezone for the schedule",
		},
		gcli.StringFlag{
			Name:  "team",
			Usage: "Name of the team under which the schedule is",
		},
		gcli.BoolFlag{
			Name:  "enabled",
			Usage: "Enable the schedule",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateSchedule",
		Flags: flags,
		Usage: "Updates a schedule on Opsgenie",
		Action: func(c *gcli.Context) error {
			command.UpdateScheduleAction(c)
			return nil
		},
	}
	return cmd
}

func deleteScheduleCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteSchedule",
		Flags: flags,
		Usage: "Delete a schedule's content from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteScheduleAction(c)
			return nil
		},
	}
	return cmd
}

func listSchedulesCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.BoolFlag{
			Name:  "expand",
			Usage: "Get more detailed response by expanding it",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listSchedules",
		Flags: flags,
		Usage: "List schedules in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListScheduleAction(c)
			return nil
		},
	}
	return cmd
}

func getScheduleTimelineCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "interval",
			Usage: "Length of time as integer in intervalUnits",
		},
		gcli.StringFlag{
			Name:  "intervalUnit",
			Usage: "Unit of the time to retrieve the timeline. Possible values: days, weeks and months",
		},
		gcli.StringFlag{
			Name:  "expand",
			Usage: "Get a detailed response. Possible values: base, forwarding, override",
		},
		gcli.StringFlag{
			Name:  "date",
			Usage: "Time (in ISO8601) to return future date on-call participants. Default is now",
		},
	}, renderingFlags...)

	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getScheduleTimeline",
		Flags: flags,
		Usage: "Get the timeline of the given schedule",
		Action: func(c *gcli.Context) error {
			command.GetScheduleTimelineAction(c)
			return nil
		},
	}
	return cmd
}

func createScheduleRotationCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "startDate",
			Usage: "Schedule Start Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "Schedule End Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "Type of rotation. Available values: daily, weekly and hourly",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the rotation",
		},
		gcli.StringFlag{
			Name:  "participants",
			Usage: "A comma separated list of participants. Example: user:<user-email>, team:<team-name>, none",
		},
		gcli.StringFlag{
			Name:  "length",
			Usage: "Length of the rotation. Default: 1",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createScheduleRotation",
		Flags: flags,
		Usage: "List the rotations for a schedule",
		Action: func(c *gcli.Context) error {
			command.CreateScheduleRotationAction(c)
			return nil
		},
	}
	return cmd
}

func getScheduleRotationCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "rotation-id",
			Usage: "ID of the rotation",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getScheduleRotation",
		Flags: flags,
		Usage: "Get a rotation from the schedule",
		Action: func(c *gcli.Context) error {
			command.GetScheduleRotationAction(c)
			return nil
		},
	}
	return cmd
}

func listScheduleRotationsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listScheduleRotations",
		Flags: flags,
		Usage: "List the rotations for a schedule",
		Action: func(c *gcli.Context) error {
			command.ListScheduleRotationsAction(c)
			return nil
		},
	}
	return cmd
}

func updateScheduleRotationCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "rotation-id",
			Usage: "ID of the rotation",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the rotation",
		},
		gcli.StringFlag{
			Name:  "startDate",
			Usage: "Schedule Start Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "Schedule End Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "Type of rotation. Available values: daily, weekly and hourly",
		},
		gcli.StringFlag{
			Name:  "participants",
			Usage: "A comma separated list of participants. Example: user:<user-email>, team:<team-name>, none",
		},
		gcli.StringFlag{
			Name:  "length",
			Usage: "Length of the rotation. Default: 1",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateScheduleRotation",
		Flags: flags,
		Usage: "Update a rotation from the schedule",
		Action: func(c *gcli.Context) error {
			command.UpdateScheduleRotationAction(c)
			return nil
		},
	}
	return cmd
}

func deleteScheduleRotationCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "rotation-id",
			Usage: "ID of the rotation",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteScheduleRotation",
		Flags: flags,
		Usage: "Delete a rotation from the schedule",
		Action: func(c *gcli.Context) error {
			command.DeleteScheduleRotationAction(c)
			return nil
		},
	}
	return cmd
}

func createScheduleOverrideCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the schedule override",
		},
		gcli.StringFlag{
			Name:  "startDate",
			Usage: "Schedule override Start Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "Schedule override End Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "responder",
			Usage: "Responder type and name. Examples: user:<user-email>, team:<team-name>, escalation:<escalation-name>",
		},
		gcli.StringFlag{
			Name:  "rotations",
			Usage: "Comma separated list of rotation ids",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createScheduleOverride",
		Flags: flags,
		Usage: "Create a schedule override",
		Action: func(c *gcli.Context) error {
			command.CreateScheduleOverrideAction(c)
			return nil
		},
	}
	return cmd
}

func listScheduleOverridesCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listScheduleOverrides",
		Flags: flags,
		Usage: "List the overrides for a schedule",
		Action: func(c *gcli.Context) error {
			command.ListScheduleOverridesAction(c)
			return nil
		},
	}
	return cmd
}

func getScheduleOverrideCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the schedule override",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getScheduleOverride",
		Flags: flags,
		Usage: "Get details for a schedule override",
		Action: func(c *gcli.Context) error {
			command.GetScheduleOverrideAction(c)
			return nil
		},
	}
	return cmd
}

func updateScheduleOverrideCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the schedule override",
		},
		gcli.StringFlag{
			Name:  "startDate",
			Usage: "Schedule override Start Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "endDate",
			Usage: "Schedule override End Date in ISO8601 format",
		},
		gcli.StringFlag{
			Name:  "responder",
			Usage: "Responder type and name. Examples: user:<user-email>, team:<team-name>, escalation:<escalation-name>",
		},
		gcli.StringFlag{
			Name:  "rotations",
			Usage: "Comma separated list of rotation ids",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateScheduleOverride",
		Flags: flags,
		Usage: "Update a schedule override.",
		Action: func(c *gcli.Context) error {
			command.UpdateScheduleOverrideAction(c)
			return nil
		},
	}
	return cmd
}

func deleteScheduleOverrideCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the specified schedule id, which can be id, name. Default value = id",
		},
		gcli.StringFlag{
			Name:  "alias",
			Usage: "Alias of the schedule override",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteScheduleOverride",
		Flags: flags,
		Usage: "Delete a schedule override.",
		Action: func(c *gcli.Context) error {
			command.DeleteScheduleOverrideAction(c)
			return nil
		},
	}
	return cmd
}

func getOnCallsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the schedule",
		},
		gcli.BoolFlag{
			Name: "flat, f",
			Usage: "If the response should be flat",
		},
		gcli.StringFlag{
			Name: "atTime",
			Usage: "Time at which to be checked, default is current time",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getOncall",
		Flags: flags,
		Usage: "Get On-Calls ",
		Action: func(c *gcli.Context) error {
			command.GetOnCallsAction(c)
			return nil
		},
	}
	return cmd
}

func getNextOnCallCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the schedule",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the schedule",
		},
		gcli.BoolFlag{
			Name: "flat, f",
			Usage: "If the response should be flat",
		},
		gcli.StringFlag{
			Name: "atTime",
			Usage: "Time at which to be checked, default is current time",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getNextOncall",
		Flags: flags,
		Usage: "Get next On-Call",
		Action: func(c *gcli.Context) error {
			command.GetNextOnCallAction(c)
			return nil
		},
	}
	return cmd
}
func exportUserOnCallsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "userId",
			Usage: "ID of the user",
		},
		gcli.StringFlag{
			Name:  "userName",
			Usage: "username(email) of the user",
		},
		gcli.StringFlag{
			Name:  "exportTo",
			Usage: "file destination to export to",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "exportUserOncalls",
		Flags: flags,
		Usage: "Export User Oncalls as ics file",
		Action: func(c *gcli.Context) error {
			command.ExportOnCallsAction(c)
			return nil
		},
	}
	return cmd
}

func createIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "message",
			Usage: "Message of the incident which is being created",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Description of the incident which is being created",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "Comma seperated list of responder types",
		},
		gcli.StringFlag{
			Name:  "responder",
			Usage: "Comma seperated list of responder name/value applicable as per type specified",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "Comma seperated list of tags which will be added to the incident",
		},
		gcli.StringFlag{
			Name:  "detailKeys",
			Usage: "Details of incident are stored in Key/value pair. Key of details which needs to be added",
		},
		gcli.StringFlag{
			Name:  "detailValues",
			Usage: "Value for each key specified in details key in same order",
		},
		gcli.StringFlag{
			Name:  "priority",
			Usage: "Priority of the incident",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note which needs to be added to incident",
		},
		gcli.StringFlag{
			Name:  "serviceId",
			Usage: "Service Id for the incident",
		},
		gcli.BoolFlag{
			Name:  "notifyStakeHolders",
			Usage: "Bool flag to show if stakeholders needs to be notified",
		},
		gcli.StringFlag{
			Name:  "statusPageEntityTitle",
			Usage: "Title of Status Page Entity",
		},
		gcli.StringFlag{
			Name:  "statusPageEntityDescription",
			Usage: "Description of Status Page Entity",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createIncident",
		Flags: flags,
		Usage: "Create Incident in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.CreateIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func deleteIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which need to be deleted",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteIncident",
		Flags: flags,
		Usage: "Deletes incident in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func getIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident whose information we want",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getIncident",
		Flags: flags,
		Usage: "Fetches incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.GetIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func getIncidentListCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Number of incidents required in first list",
		},
		gcli.StringFlag{
			Name:  "sortField",
			Usage: "Parameter according to which list needs to sorted",
		},
		gcli.StringFlag{
			Name:  "offset",
			Usage: "Offset for list of incidents in case of Pagination",
		},
		gcli.StringFlag{
			Name:  "order",
			Usage: "Order by which incident needs to be sorted {asc/desc}",
		},
		gcli.StringFlag{
			Name:  "query",
			Usage: "Query to be executed for finding list of incidents",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getIncidentList",
		Flags: flags,
		Usage: "List all incident in Opsgenie",
		Action: func(c *gcli.Context) error {
			command.ListIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func closeIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which need to be closed",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note in the incident before closing it",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "closeIncident",
		Flags: flags,
		Usage: "Closes incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.CloseIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func addNoteIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which note need to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addNoteToIncident",
		Flags: flags,
		Usage: "Add Notes to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.AddNoteIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func addResponderIncidentCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which responders need to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "type",
			Usage: "Comma seperated list of responder types",
		},
		gcli.StringFlag{
			Name:  "responder",
			Usage: "Comma seperated list of responder name/value applicable as per type specified",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addIncidentResponders",
		Flags: flags,
		Usage: "Add Responders to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.AddResponderIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func addIncidentTagsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which tags need to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "Comma seperated Tags to be added to the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addIncidentTags",
		Flags: flags,
		Usage: "Add Tags to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.AddTagsIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func removeIncidentTagsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident whose tags need to be deleted",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "tags",
			Usage: "Comma seperated Tags to be removed to the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "removeIncidentTags",
		Flags: flags,
		Usage: "Remove Tags to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.RemoveTagsIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func addIncidentDetailsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which details need to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "detailKeys",
			Usage: "Details are stored in Key/Value Pair. Comma seperated Keys in details to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "detailValues",
			Usage: "Comma seperated Values corresponding to keys in details to be added to the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "addIncidentDetails",
		Flags: flags,
		Usage: "Add Details to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.AddDetailsIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func removeIncidentDetailsCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident which details need to be removed",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "note",
			Usage: "Note to be added to the incident",
		},
		gcli.StringFlag{
			Name:  "keys",
			Usage: "Comma seperated detail Keys to be removed to the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "removeIncidentDetails",
		Flags: flags,
		Usage: "Remove Tags to incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.RemoveDetailsIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func updateIncidentPriorityCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident whose priority needs to be updated",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "priority",
			Usage: "Updated Priority of the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateIncidentPriority",
		Flags: flags,
		Usage: "Update Priority of incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.UpdatePriorityIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func updateIncidentMessageCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident whose message needs to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "message",
			Usage: "Updated Message of the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateIncidentMessage",
		Flags: flags,
		Usage: "Updated Message of incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.UpdateMessageIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func updateIncidentDescriptionCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "identifier",
			Usage: "Identifier of the incident whose decription needs to be added",
		},
		gcli.StringFlag{
			Name:  "identifierType",
			Usage: "Identifier type of the incident {id,tiny}",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Updated Description of the incident",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateIncidentDescription",
		Flags: flags,
		Usage: "Updated Description of incident in Opsgenie with Id",
		Action: func(c *gcli.Context) error {
			command.UpdateDescriptionIncidentAction(c)
			return nil
		},
	}
	return cmd
}

func createServiceCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the new Service",
		},
		gcli.StringFlag{
			Name:  "teamId",
			Usage: "ID of the team",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "Service Description",
		},
		gcli.StringFlag{
			Name:  "visibility",
			Usage: "Service Visibility",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "createService",
		Flags: flags,
		Usage: "Create a new service",
		Action: func(c *gcli.Context) error {
			command.CreateServiceAction(c)
			return nil
		},
	}
	return cmd
}
func updateServiceCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the service",
		},
		gcli.StringFlag{
			Name:  "name",
			Usage: "Name of the service",
		},
		gcli.StringFlag{
			Name:  "description",
			Usage: "service Description",
		},
		gcli.StringFlag{
			Name:  "visibility",
			Usage: "Service Visibility",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "updateService",
		Flags: flags,
		Usage: "Update a service",
		Action: func(c *gcli.Context) error {
			command.UpdateServiceAction(c)
			return nil
		},
	}
	return cmd
}
func deleteServiceCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the  service",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "deleteService",
		Flags: flags,
		Usage: "Deletes service from Opsgenie",
		Action: func(c *gcli.Context) error {
			command.DeleteServiceAction(c)
			return nil
		},
	}
	return cmd
}
func getServiceCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "id",
			Usage: "ID of the service",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "getService",
		Flags: flags,
		Usage: "Get service details",
		Action: func(c *gcli.Context) error {
			command.GetServiceAction(c)
			return nil
		},
	}
	return cmd
}
func listServiceCommand() gcli.Command {
	commandFlags := append([]gcli.Flag{
		gcli.StringFlag{
			Name:  "limit",
			Usage: "Number of results per page",
		},
		gcli.StringFlag{
			Name:  "offset",
			Usage: "Pagination offset",
		},
	}, renderingFlags...)
	flags := append(commonFlags, commandFlags...)
	cmd := gcli.Command{Name: "listServices",
		Flags: flags,
		Usage: "List the registered services",
		Action: func(c *gcli.Context) error {
			command.ListServiceAction(c)
			return nil
		},
	}
	return cmd
}

func initCommands(app *gcli.App) {
	app.Commands = []gcli.Command{
		createAlertCommand(),
		getAlertCommand(),
		attachFileCommand(),
		getAttachmentCommand(),
		downloadAttachmentCommand(),
		listAttachmentsCommand(),
		deleteAttachmentCommand(),
		acknowledgeCommand(),
		assignOwnerCommand(),
		addTeamCommand(),
		addTagsCommand(),
		addResponderCommand(),
		addNoteCommand(),
		executeActionCommand(),
		closeAlertCommand(),
		deleteAlertCommand(),
		pingHeartbeatCommand(),
		createHeartbeatCommand(),
		deleteHeartbeatCommand(),
		disableHeartbeatCommand(),
		enableHeartbeatCommand(),
		listHeartbeatCommand(),
		enableCommand(),
		disableCommand(),
		listAlertsCommand(),
		countAlertsCommand(),
		listAlertNotesCommand(),
		listAlertLogsCommand(),
		listAlertRecipientsCommand(),
		unAcknowledgeCommand(),
		snoozeCommand(),
		removeTagsCommand(),
		addDetailsCommand(),
		removeDetailsCommand(),
		escalateToNextActionCommand(),
		exportUsersCommand(),
		downloadLogsCommand(),
		createTeamCommand(),
		getTeamCommand(),
		updateTeamCommand(),
		deleteTeamCommand(),
		listTeamCommand(),
		addMemberCommand(),
		removeMemberCommand(),
		createRoleCommand(),
		listAllRoleRightsCommand(),
		getRoleCommand(),
		deleteRoleCommand(),
		listRolesCommand(),
		listTeamRoutingRulesCommand(),
		listTeamLogsCommand(),
		deleteTeamRoutingRulesCommand(),
		getRoutingRuleCommand(),
		CreateEscalationCommand(),
		fetchEscalationCommand(),
		deleteEscalationCommand(),
		UpdateEscalationCommand(),
		createScheduleCommand(),
		getScheduleCommand(),
		listSchedulesCommand(),
		updateScheduleCommand(),
		deleteScheduleCommand(),
		getScheduleTimelineCommand(),
		createScheduleRotationCommand(),
		getScheduleRotationCommand(),
		listScheduleRotationsCommand(),
		updateScheduleRotationCommand(),
		deleteScheduleRotationCommand(),
		createScheduleOverrideCommand(),
		listScheduleOverridesCommand(),
		getScheduleOverrideCommand(),
		updateScheduleOverrideCommand(),
		deleteScheduleOverrideCommand(),
		getOnCallsCommand(),
		getNextOnCallCommand(),
		exportUserOnCallsCommand(),
		createIncidentCommand(),
		deleteIncidentCommand(),
		getIncidentCommand(),
		getIncidentListCommand(),
		closeIncidentCommand(),
		addNoteIncidentCommand(),
		addResponderIncidentCommand(),
		addIncidentTagsCommand(),
		removeIncidentTagsCommand(),
		addIncidentDetailsCommand(),
		removeIncidentDetailsCommand(),
		updateIncidentPriorityCommand(),
		updateIncidentMessageCommand(),
		updateIncidentDescriptionCommand(),
		createServiceCommand(),
		updateServiceCommand(),
		deleteServiceCommand(),
		getServiceCommand(),
		listServiceCommand(),

	}
}

func main() {
	app := gcli.NewApp()
	app.Name = "lamp"
	app.Version = lampVersion
	app.Usage = "Command line interface for Opsgenie"
	app.Author = "Opsgenie"
	app.Action = func(c *gcli.Context) error {
		fmt.Printf("Run 'lamp help' for the options\n")
		return nil
	}
	initCommands(app)
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error occured while executing command: %s\n", err.Error())
	}
}
