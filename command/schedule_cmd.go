package command

import (
	"context"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
	"time"
)

func NewScheduleClient(c *gcli.Context) *schedule.Client {
	scheduleCli, cliErr := schedule.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the schedule client. " + cliErr.Error()
		printMessage(INFO, message)
		os.Exit(1)
	}
	printMessage(DEBUG,"Schedule Client created.")

	return scheduleCli
}

func CreateScheduleAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.CreateRequest{}

	if name, ok := getVal("name", c); ok {
		req.Name = name
	}
	if description, ok := getVal("description", c); ok {
		req.Description = description
	}
	if timezone, ok := getVal("tz", c); ok {
		req.Timezone = timezone
	}
	enabled := c.IsSet("enabled")
	req.Enabled = &enabled

	if teamName, ok := getVal("team", c); ok {
		req.OwnerTeam = &og.OwnerTeam{
			Name: teamName,
		}
	}

	printMessage(DEBUG,"Create schedule request prepared, sending request to Opsgenie..")

	resp, err := cli.Create(nil, &req)
	renderResponse(c, resp, err)
}

func GetScheduleAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.GetRequest{}

	req.IdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.IdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.IdentifierValue = id
	}

	printMessage(DEBUG,"Get schedule request prepared, sending request to Opsgenie..")

	resp, err := cli.Get(nil, &req)

	renderResponse(c, resp, err)
}

func ListScheduleAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.ListRequest{}
	expand := c.IsSet("expand")
	req.Expand = &expand
	printMessage(DEBUG,"List schedules request prepared, sending request to Opsgenie..")
	resp, err := cli.List(nil, &req)
	renderResponse(c, resp, err)
}

func UpdateScheduleAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.UpdateRequest{}

	req.IdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.IdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.IdentifierValue = id
	}

	if name, ok := getVal("name", c); ok {
		req.Name = name
	}
	if description, ok := getVal("description", c); ok {
		req.Description = description
	}
	if timezone, ok := getVal("tz", c); ok {
		req.Timezone = timezone
	}
	enabled := c.IsSet("enabled")
	req.Enabled = &enabled

	if teamName, ok := getVal("team", c); ok {
		req.OwnerTeam = &og.OwnerTeam{
			Name: teamName,
		}
	}

	printMessage(DEBUG,"Update schedule request prepared, sending request to Opsgenie..")

	resp, err := cli.Update(nil, &req)

	renderResponse(c, resp, err)
}

func DeleteScheduleAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.DeleteRequest{}

	req.IdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.IdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.IdentifierValue = id
	}

	printMessage(DEBUG,"Delete schedule request prepared, sending request to Opsgenie..")

	resp, err := cli.Delete(nil, &req)

	renderResponse(c, resp, err)
}

func GetScheduleTimelineAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.GetTimelineRequest{}

	req.IdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.IdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.IdentifierValue = id
	}

	if expand, ok := getVal("expand", c); ok {
		switch expand {
		case "base":
			req.Expands = []schedule.ExpandType{schedule.Base}
		case "forwarding":
			req.Expands = []schedule.ExpandType{schedule.Forwarding}
		case "override":
			req.Expands = []schedule.ExpandType{schedule.Override}
		}
	}

	req.IntervalUnit = schedule.Weeks

	if intervalUnit, ok := getVal("intervalUnit", c); ok {
		switch intervalUnit {
		case "days":
			req.IntervalUnit = schedule.Days
		case "months":
			req.IntervalUnit = schedule.Months
		default:
			req.IntervalUnit = schedule.Weeks
		}
	}

	if interval, ok := getVal("interval", c); ok {
		intervalInt, err := strconv.Atoi(interval)
		exitOnErr(err)
		req.Interval = intervalInt
	}

	if dateStr, ok := getVal("date", c); ok {
		date, err := time.Parse(time.RFC3339, dateStr)
		exitOnErr(err)
		req.Date = &date
	}

	printMessage(DEBUG,"Get schedule timeline request prepared, sending request to Opsgenie..")
	resp, err := cli.GetTimeline(nil, &req)

	renderResponse(c, resp, err)

}

func CreateScheduleRotationAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.CreateRotationRequest{}
	rotation := og.Rotation{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifierValue = id
	}

	if name, ok := getVal("name", c); ok {
		rotation.Name = name
	}

	if rotationType, ok := getVal("type", c); ok {
		switch rotationType {
		case "hourly":
			rotation.Type = og.Hourly
		case "daily":
			rotation.Type = og.Daily
		case "weekly":
			rotation.Type = og.Weekly
		}
	}

	if startDate, ok := getVal("startDate", c); ok {
		date, err := time.Parse(time.RFC3339, startDate)
		exitOnErr(err)
		rotation.StartDate = &date
	}

	if endDate, ok := getVal("endDate", c); ok {
		date, err := time.Parse(time.RFC3339, endDate)
		exitOnErr(err)
		rotation.EndDate = &date
	}

	if length, ok := getVal("length", c); ok {
		length, err := strconv.ParseUint(length, 0, 32)
		exitOnErr(err)
		rotation.Length = uint32(length)
	}

	if participantsStr, ok := getVal("participants", c); ok {
		participants := []og.Participant{}

		for _, value := range strings.Split(participantsStr, ",") {
			participantData := strings.Split(strings.TrimSpace(value), ":")

			switch participantData[0] {
			case "user":
				participants = append(participants, og.Participant{
					Type:     og.User,
					Username: participantData[1],
				})
			case "team":
				participants = append(participants, og.Participant{
					Type: og.Team,
					Name: participantData[1],
				})
			case "none":
				participants = append(participants, og.Participant{
					Type: og.None,
				})
			}
		}

		rotation.Participants = participants
	}

	req.Rotation = &rotation

	printMessage(DEBUG,"Create schedule rotations request prepared, sending request to Opsgenie..")

	resp, err := cli.CreateRotation(nil, &req)

	renderResponse(c, resp, err)
}

func GetScheduleRotationAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.GetRotationRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifierValue = id
	}

	if rotationId, ok := getVal("rotation-id", c); ok {
		req.RotationId = rotationId
	}

	printMessage(DEBUG,"Get schedule rotation request prepared, sending request to Opsgenie..")

	resp, err := cli.GetRotation(nil, &req)

	renderResponse(c, resp, err)
}

func ListScheduleRotationsAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.ListRotationsRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifierValue = id
	}

	printMessage(DEBUG,"List schedule rotations request prepared, sending request to Opsgenie..")

	resp, err := cli.ListRotations(nil, &req)

	renderResponse(c, resp, err)
}

func UpdateScheduleRotationAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.UpdateRotationRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifierValue = id
	}

	if rotationId, ok := getVal("rotation-id", c); ok {
		req.RotationId = rotationId
	}

	rotation := og.Rotation{}

	if name, ok := getVal("name", c); ok {
		rotation.Name = name
	}

	if rotationType, ok := getVal("type", c); ok {
		switch rotationType {
		case "hourly":
			rotation.Type = og.Hourly
		case "daily":
			rotation.Type = og.Daily
		case "weekly":
			rotation.Type = og.Weekly
		}
	}

	if startDate, ok := getVal("startDate", c); ok {
		date, err := time.Parse(time.RFC3339, startDate)
		exitOnErr(err)
		rotation.StartDate = &date
	}

	if endDate, ok := getVal("endDate", c); ok {
		date, err := time.Parse(time.RFC3339, endDate)
		exitOnErr(err)
		rotation.EndDate = &date
	}

	if length, ok := getVal("length", c); ok {
		length, err := strconv.ParseUint(length, 0, 32)
		exitOnErr(err)
		rotation.Length = uint32(length)
	}

	if participantsStr, ok := getVal("participants", c); ok {
		participants := []og.Participant{}

		for _, value := range strings.Split(participantsStr, ",") {
			participantData := strings.Split(strings.TrimSpace(value), ":")

			switch participantData[0] {
			case "user":
				participants = append(participants, og.Participant{
					Type:     og.User,
					Username: participantData[1],
				})
			case "team":
				participants = append(participants, og.Participant{
					Type: og.Team,
					Name: participantData[1],
				})
			case "none":
				participants = append(participants, og.Participant{
					Type: og.None,
				})
			}
		}

		rotation.Participants = participants
	}

	req.Rotation = &rotation
	printMessage(DEBUG,"Update schedule rotation request prepared, sending request to Opsgenie..")

	resp, err := cli.UpdateRotation(nil, &req)

	renderResponse(c, resp, err)
}

func DeleteScheduleRotationAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.DeleteRotationRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifierValue = id
	}

	if rotationId, ok := getVal("rotation-id", c); ok {
		req.RotationId = rotationId
	}

	printMessage(DEBUG,"Delete schedule rotation request prepared, sending request to Opsgenie..")

	resp, err := cli.DeleteRotation(nil, &req)

	renderResponse(c, resp, err)
}

func CreateScheduleOverrideAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.CreateScheduleOverrideRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifier = id
	}

	if overrideAlias, ok := getVal("alias", c); ok {
		req.Alias = overrideAlias
	}

	if startDate, ok := getVal("startDate", c); ok {
		date, err := time.Parse(time.RFC3339, startDate)
		exitOnErr(err)
		req.StartDate = date
	}

	if endDate, ok := getVal("endDate", c); ok {
		date, err := time.Parse(time.RFC3339, endDate)
		exitOnErr(err)
		req.EndDate = date
	}

	if responderStr, ok := getVal("responder", c); ok {
		responderData := strings.Split(strings.TrimSpace(responderStr), ":")
		switch responderData[0] {
		case "user":
			req.User = schedule.Responder{
				Type:     schedule.UserResponderType,
				Username: responderData[1],
			}
		case "team":
			req.User = schedule.Responder{
				Type: schedule.TeamResponderType,
				Name: responderData[1],
			}
		case "escalation":
			req.User = schedule.Responder{
				Type: schedule.EscalationResponderType,
				Name: responderData[1],
			}
		}
	}

	if rotationsStr, ok := getVal("rotations", c); ok {
		rotations := []schedule.RotationIdentifier{}
		for _, rotationId := range strings.Split(rotationsStr, ",") {
			rotations = append(rotations, schedule.RotationIdentifier{Id: rotationId})
		}
		req.Rotations = rotations
	}
	printMessage(DEBUG,"Create schedule override request prepared, sending request to Opsgenie..")

	resp, err := cli.CreateScheduleOverride(nil, &req)

	renderResponse(c, resp, err)
}

func ListScheduleOverridesAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.ListScheduleOverrideRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifier = id
	}

	printMessage(DEBUG,"List schedule overrides request prepared, sending request to Opsgenie..")

	resp, err := cli.ListScheduleOverride(nil, &req)

	renderResponse(c, resp, err)
}

func GetScheduleOverrideAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.GetScheduleOverrideRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifier = id
	}

	if overrideAlias, ok := getVal("alias", c); ok {
		req.Alias = overrideAlias
	}
	printMessage(DEBUG,"Get schedule override request prepared, sending request to Opsgenie..")

	resp, err := cli.GetScheduleOverride(nil, &req)

	renderResponse(c, resp, err)
}

func UpdateScheduleOverrideAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.UpdateScheduleOverrideRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifier = id
	}

	if overrideAlias, ok := getVal("alias", c); ok {
		req.Alias = overrideAlias
	}

	if startDate, ok := getVal("startDate", c); ok {
		date, err := time.Parse(time.RFC3339, startDate)
		exitOnErr(err)
		req.StartDate = date
	}

	if endDate, ok := getVal("endDate", c); ok {
		date, err := time.Parse(time.RFC3339, endDate)
		exitOnErr(err)
		req.EndDate = date
	}

	if responderStr, ok := getVal("responder", c); ok {
		responderData := strings.Split(strings.TrimSpace(responderStr), ":")
		switch responderData[0] {
		case "user":
			req.User = schedule.Responder{
				Type:     schedule.UserResponderType,
				Username: responderData[1],
			}
		case "team":
			req.User = schedule.Responder{
				Type: schedule.TeamResponderType,
				Name: responderData[1],
			}
		case "escalation":
			req.User = schedule.Responder{
				Type: schedule.EscalationResponderType,
				Name: responderData[1],
			}
		}
	}

	if rotationsStr, ok := getVal("rotations", c); ok {
		rotations := []schedule.RotationIdentifier{}
		for _, rotationId := range strings.Split(rotationsStr, ",") {
			rotations = append(rotations, schedule.RotationIdentifier{Id: rotationId})
		}
		req.Rotations = rotations
	}

	printMessage(DEBUG,"Update schedule override request prepared, sending request to Opsgenie..")

	resp, err := cli.UpdateScheduleOverride(nil, &req)

	renderResponse(c, resp, err)
}

func DeleteScheduleOverrideAction(c *gcli.Context) {
	cli := NewScheduleClient(c)
	req := schedule.DeleteScheduleOverrideRequest{}

	req.ScheduleIdentifierType = schedule.Id
	if identifier, ok := getVal("identifierType", c); ok {
		if identifier == "name" {
			req.ScheduleIdentifierType = schedule.Name
		}
	}
	if id, ok := getVal("id", c); ok {
		req.ScheduleIdentifier = id
	}

	if overrideAlias, ok := getVal("alias", c); ok {
		req.Alias = overrideAlias
	}
	printMessage(DEBUG,"Delete schedule override request prepared, sending request to Opsgenie..")

	resp, err := cli.DeleteScheduleOverride(nil, &req)

	renderResponse(c, resp, err)
}

func GetOnCallsAction(c *gcli.Context){
	cli := NewScheduleClient(c)
	req := &schedule.GetOnCallsRequest{}
	if scheduleName, ok := getVal("name",c); ok {
		req.ScheduleIdentifierType = schedule.Name
		req.ScheduleIdentifier = scheduleName
	} else if scheduleID, ok := getVal("id",c);ok {
		req.ScheduleIdentifierType = schedule.Id
		req.ScheduleIdentifier = scheduleID
	}
	if atTime, ok := getVal("atTime", c); ok {
		t, err := time.Parse(time.RFC3339, atTime)
		exitOnErr(err)
		req.Date = &t
	}
	flat := c.IsSet("flat")
	req.Flat = &flat

	resp, err := cli.GetOnCalls(context.Background(), req)
	renderResponse(c, resp, err)
}

func GetNextOnCallAction(c *gcli.Context){
	cli := NewScheduleClient(c)
	req := &schedule.GetNextOnCallsRequest{}
	if scheduleName, ok := getVal("name",c); ok {
		req.ScheduleIdentifierType = schedule.Name
		req.ScheduleIdentifier = scheduleName
	} else if scheduleID, ok := getVal("id",c);ok {
		req.ScheduleIdentifierType = schedule.Id
		req.ScheduleIdentifier = scheduleID
	}
	if atTime, ok := getVal("atTime", c); ok {
		t, err := time.Parse(time.RFC3339, atTime)
		exitOnErr(err)
		req.Date = &t
	}
	flat := c.IsSet("flat")
	req.Flat = &flat

	resp, err := cli.GetNextOnCall(context.Background(), req)
	renderResponse(c, resp, err)
}

func ExportOnCallsAction(c *gcli.Context){
	cli := NewScheduleClient(c)
	req := &schedule.ExportOnCallUserRequest{}
	if userName, ok := getVal("userName",c); ok {
		req.UserIdentifier = userName
	} else if userID, ok := getVal("userId",c);ok {
		req.UserIdentifier = userID
	}
	if exportTo, ok := getVal("exportTo",c);ok {
		req.ExportedFilePath = exportTo
	}
	icsFile, err := cli.ExportOnCallUser(context.Background(), req)
	exitOnErr(err)
	printMessage(INFO,"Downloaded file "+ icsFile.Name())
}


func exitOnErr(err error) {
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
}

func renderResponse(c *gcli.Context, resp interface{}, err error) {
	exitOnErr(err)

	outputFormat := strings.ToLower(c.String("output-format"))
	printMessage(DEBUG,"Got response successfully, rendering in " + outputFormat + " format")

	switch outputFormat {
	case "yaml":
		output, err := resultToYAML(resp)
		exitOnErr(err)
		printMessage(INFO,output)
	default:
		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp, isPretty)
		exitOnErr(err)
		printMessage(INFO, output)
	}
}
