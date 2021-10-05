package command

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/integration_v3"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"os"
	gcli "github.com/urfave/cli"
	"strings"
)

func NewIntegrationV3Client(c *gcli.Context) (*integration_v3.Client, error) {
	integrationV3, cliErr := integration_v3.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the alert client. " + cliErr.Error()
		printMessage(ERROR,message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Alert Client created.")
	return integrationV3, nil
}

func GetIntegrationAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.GetRequest{}

	if val, success := getVal("id", c); success {
		req.Id = val
	}

	printMessage(DEBUG,"Get integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.Get(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func GetIntegrationListAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.ListRequest{}

	if val, success := getVal("teamId", c); success {
		req.TeamId = val
	}

	if val, success := getVal("type", c); success {
		req.IntegrationType = val
	}

	printMessage(DEBUG,"List integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.List(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func DeleteIntegrationAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.DeleteIntegrationRequest{}

	if val, success := getVal("id", c); success {
		req.Id = val
	}

	printMessage(DEBUG,"Delete integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.Delete(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func AuthenticateIntegrationAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.AuthenticateIntegrationRequest{}

	if val, success := getVal("type", c); success {
		req.Type = val
	}

	printMessage(DEBUG,"Authenticate integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.Authenticate(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func CreateIntegrationAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.CreateIntegrationRequest{}

	if val, success := getVal("name", c); success {
		req.Name = val
	}

	if val, success := getVal("type", c); success {
		req.Type = val
	}

	if val, success := getVal("teamId", c); success {
		req.TeamId = val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	if _, success := getVal("enabled", c); success {
		req.Enabled = c.IsSet("enabled")
	}

	if val, success := getVal("typeSpecificProperties", c); success {
		req.TypeSpecificProperties = getTypeSpecificProperties(val)
	}

	printMessage(DEBUG,"Create integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.Create(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func UpdateIntegrationAction(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.UpdateIntegrationRequest{}

	if val, success := getVal("Name", c); success {
		req.Name = val
	}

	if val, success := getVal("teamId", c); success {
		req.TeamId = val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	if _, success := getVal("enabled", c); success {
		req.Enabled = c.IsSet("enabled")
	}

	if val, success := getVal("typeSpecificProperties", c); success {
		req.TypeSpecificProperties = getTypeSpecificProperties(val)
	}

		printMessage(DEBUG, "Update integration request prepared from flags, sending request to Opsgenie...")

		resp, err := cli.Update(nil, &req)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}

		isPretty := c.IsSet("pretty")
		output, err := resultToJSON(resp, isPretty)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}
		printMessage(INFO, output)
}

func GetIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.GetIntegrationActionsRequest{}

	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("actionId", c); success {
		req.ActionId = val
	}

	printMessage(DEBUG,"Get integration Action request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.GetAction(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func ListIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.ListIntegrationActionsRequest{}

	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("direction", c); success {
		req.Direction = val
	}

	if val, success := getVal("integrationType", c); success {
		req.IntegrationType = val
	}

	if val, success := getVal("domain", c); success {
		req.Domain = val
	}

	printMessage(DEBUG,"List integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.ListIntegrationAction(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func DeleteIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.DeleteIntegrationActionsRequest{}


	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("actionId", c); success {
		req.ActionId = val
	}

	printMessage(DEBUG,"Delete integration Action request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.DeleteAction(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	output, err := resultToJSON(resp, true)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func ReorderIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.ReOrderIntegrationActionsRequest{}


	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("actionId", c); success {
		req.ActionId = val
	}

	if val, success := getVal("successorId", c); success {
		req.SuccessorId = val
	}

	printMessage(DEBUG,"Reorder integration Action request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.ReorderAction(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	output, err := resultToJSON(resp, true)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func CreateIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.CreateIntegrationActionsRequest{}

	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("name", c); success {
		req.Name = val
	}

	if val, success := getVal("type", c); success {
		req.Type = getActionType(val)
	}

	if val, success := getVal("direction", c); success {
		req.Direction = val
	}

	if val, success := getVal("domain", c); success {
		req.Domain = val
	}

	if val, success := getVal("actionGroupId", c); success {
		req.ActionGroupId = val
	}

	if val, success := getVal("actionMappingType", c); success {
		req.ActionMapping.Type = getActionType(val)
	}

	if val, success := getVal("actionMappingParameter", c); success {
		req.ActionMapping.Parameter = val
	}

	if _, success := getVal("enabled", c); success {
		enabled := c.IsSet("enabled")
		req.Enabled = &enabled
	}

	if val, success := getVal("fieldMappingUser", c); success {
		req.Mapping.User = val
	}

	if val, success := getVal("fieldMappingNote", c); success {
		req.Mapping.Note = val
	}

	if val, success := getVal("fieldMappingAlias", c); success {
		req.Mapping.Alias = val
	}

	if val, success := getVal("fieldMappingSource", c); success {
		req.Mapping.Source = val
	}

	if val, success := getVal("fieldMappingMessage", c); success {
		req.Mapping.Message = val
	}

	if val, success := getVal("fieldMappingDescription", c); success {
		req.Mapping.Description = val
	}

	if val, success := getVal("fieldMappingEntity", c); success {
		req.Mapping.Entity = val
	}

	if val, success := getVal("typeSpecificProperties", c); success {
		req.TypeSpecificProperties = getTypeSpecificProperties(val)
	}

	responders := generateIntegrationResponders(c, integration_v3.Team, "teams")
	responders = append(responders, generateIntegrationResponders(c, integration_v3.User, "users")...)
	responders = append(responders, generateIntegrationResponders(c, integration_v3.Escalation, "escalations")...)
	responders = append(responders, generateIntegrationResponders(c, integration_v3.Schedule, "schedules")...)

	req.Mapping.Responders = responders

	if val, success := getVal("tags", c); success {
		req.Mapping.Tags = strings.Split(val, ",")
	}

	if val, success := getVal("alertActions", c); success {
		req.Mapping.AlertActions = strings.Split(val, ",")
	}

	if val, success := getVal("filterConditionMatchType", c); success {
		req.Filter.ConditionMatchType = getFilterConditionMatchType(val)
	}

	if val, success := getVal("filterCondition", c); success {
		req.Filter.Conditions = getFilterCondition(val)
	}

	printMessage(DEBUG,"Create integration action request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.CreateAction(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func UpdateIntegrationActionCommand(c *gcli.Context) {
	cli, err := NewIntegrationV3Client(c)
	if err != nil {
		os.Exit(1)
	}
	req := integration_v3.UpdateIntegrationActionsRequest{}

	if val, success := getVal("integrationId", c); success {
		req.IntegrationId = val
	}

	if val, success := getVal("name", c); success {
		req.Name = val
	}

	if val, success := getVal("type", c); success {
		req.Type = getActionType(val)
	}

	if val, success := getVal("actionMappingType", c); success {
		req.ActionMapping.Type = getActionType(val)
	}

	if val, success := getVal("actionMappingParameter", c); success {
		req.ActionMapping.Parameter = val
	}

	if _, success := getVal("enabled", c); success {
		enabled := c.IsSet("enabled")
		req.Enabled = &enabled
	}

	if val, success := getVal("fieldMappingUser", c); success {
		req.Mapping.User = val
	}

	if val, success := getVal("fieldMappingNote", c); success {
		req.Mapping.Note = val
	}

	if val, success := getVal("fieldMappingAlias", c); success {
		req.Mapping.Alias = val
	}

	if val, success := getVal("fieldMappingSource", c); success {
		req.Mapping.Source = val
	}

	if val, success := getVal("fieldMappingMessage", c); success {
		req.Mapping.Message = val
	}

	if val, success := getVal("fieldMappingDescription", c); success {
		req.Mapping.Description = val
	}

	if val, success := getVal("fieldMappingEntity", c); success {
		req.Mapping.Entity = val
	}

	if val, success := getVal("typeSpecificProperties", c); success {
		req.TypeSpecificProperties = getTypeSpecificProperties(val)
	}

	responders := generateIntegrationResponders(c, integration_v3.Team, "teams")
	responders = append(responders, generateIntegrationResponders(c, integration_v3.User, "users")...)
	responders = append(responders, generateIntegrationResponders(c, integration_v3.Escalation, "escalations")...)
	responders = append(responders, generateIntegrationResponders(c, integration_v3.Schedule, "schedules")...)

	req.Mapping.Responders = responders

	if val, success := getVal("tags", c); success {
		req.Mapping.Tags = strings.Split(val, ",")
	}

	if val, success := getVal("alertActions", c); success {
		req.Mapping.AlertActions = strings.Split(val, ",")
	}

	if val, success := getVal("filterConditionMatchType", c); success {
		req.Filter.ConditionMatchType = getFilterConditionMatchType(val)
	}

	if val, success := getVal("filterCondition", c); success {
		req.Filter.Conditions = getFilterCondition(val)
	}

	printMessage(DEBUG, "Update integration request prepared from flags, sending request to Opsgenie...")

	resp, err := cli.UpdateAction(nil, &req)
	if err != nil {
		printMessage(ERROR, err.Error())
		os.Exit(1)
	}

	isPretty := c.IsSet("pretty")
	output, err := resultToJSON(resp, isPretty)
	if err != nil {
		printMessage(ERROR, err.Error())
		os.Exit(1)
	}
	printMessage(INFO, output)
}

func getTypeSpecificProperties(typeSpecificString string) map[string]string {
	parameters := strings.Split(typeSpecificString, ",")

	details := make(map[string]string)

	for i := 0; i < len(parameters); i++ {
		prop := parameters[i]
		if len(prop)>0 && strings.Contains(prop, ":") {
			p := strings.Split(prop, ":")
			details[p[0]] = p[1]
		} else {
			printMessage(ERROR, "Type Specific parameters should have the value of the form a:b, but got: " + prop + "\n")
			os.Exit(1)
		}
	}
	return details
}

func getActionType(value string)  integration_v3.ActionType {
	if value == "create" {
		return integration_v3.Create
	} else if value == "close" {
		return integration_v3.Close
	} else if value == "acknowledge" {
		return integration_v3.Acknowledge
	} else if value == "AddNote" {
		return integration_v3.AddNote
	} else if value == "ignore" {
		return integration_v3.Ignore
	} else {
		printMessage(ERROR, " actionType can only be close, create, acknowledge, AddNote and ignore \n")
		os.Exit(1)
	}
	return integration_v3.Ignore
}

func getFilterConditionMatchType(value string)  og.ConditionMatchType {
	if value == "match-all" {
		return og.MatchAll
	} else if value == "match-any-condition" {
		return og.MatchAnyCondition
	} else if value == "match-all-conditions" {
		return og.MatchAllConditions
	} else {
		printMessage(ERROR, " Condition Type can only be match-all, match-any-condition and match-all-conditions \n")
		os.Exit(1)
	}
	return og.MatchAllConditions
}

func generateIntegrationResponders(c *gcli.Context, responderType integration_v3.ResponderType, parameter string) []integration_v3.Responder {
	if val, success := getVal(parameter, c); success {
		responderNames := strings.Split(val, ",")

		var responders []integration_v3.Responder

		for _, name := range responderNames {
			responders = append(responders, integration_v3.Responder{
				Name:     name,
				Username: name,
				Type:     responderType,
			})
		}
		return responders
	}
	return nil
}

func getFilterCondition(value string) []og.Condition {

	parameters := strings.Split(value, ",")

	var details []og.Condition

	for i := 0; i < len(parameters); i++ {
		prop := parameters[i]
		if len(prop)>0 && strings.Contains(prop, ":") {
			p := strings.Split(prop, ":")
			if len(p)>0 {
				var condition og.Condition;
				isNot := p[1] == "true"
				condition.Field = getConditionFieldType(p[0])
				condition.IsNot = &isNot
				condition.Operation = getConditionOperation(p[2])
				condition.ExpectedValue = p[3]
				condition.Key = p[4]
				details = append(details, condition)
			}
		} else {
			printMessage(ERROR, "Type Specific parameters should have the value of the form a:b, but got: " + prop + "\n")
			os.Exit(1)
		}
	}
	return details

}

func getConditionFieldType(value string)  og.ConditionFieldType {
	if value == "message" {
		return og.Message
	} else if value == "alias" {
		return og.Alias
	} else if value == "description" {
		return og.Description
	} else if value == "source" {
		return og.Source
	} else if value == "entity" {
		return og.Entity
	} else if value == "eventType" {
		return og.EventType
	} else if value == "tags" {
		return og.Tags
	} else if value == "actions" {
		return og.Actions
	} else if value == "details" {
		return og.Details
	} else if value == "extra-properties" {
		return og.ExtraProperties
	} else if value == "recipients" {
		return og.Recipients
	} else if value == "teams" {
		return og.Teams
	} else if value == "priority" {
		return og.Priority
	} else if value == "conversationSubject" {
		return og.ConversationSub
	} else if value == "from_address" {
		return og.FromAddress
	} else if value == "from_name" {
		return og.FromName
	} else if value == "subject" {
		return og.Subject
	}  else {
		printMessage(ERROR, " Condition Field Type can only be message, alias, description, source, entity, eventType, tags," +
			"actions, details, extra-properties, recipients, teams, priority, conversationSubject, from_address, from_name and subject \n")
		os.Exit(1)
	}
	return og.Message
}

func getConditionOperation(value string)  og.ConditionOperation {
	if value == "matches" {
		return og.Matches
	} else if value == "contains" {
		return og.Contains
	} else if value == "starts-with" {
		return og.StartsWith
	} else if value == "ends-with" {
		return og.EndsWith
	} else if value == "equals" {
		return og.Equals
	} else if value == "contains-key" {
		return og.ContainsKey
	} else if value == "contains-value" {
		return og.ContainsValue
	} else if value == "greater-than" {
		return og.GreaterThan
	} else if value == "less-than" {
		return og.LessThan
	} else if value == "is-empty" {
		return og.IsEmpty
	} else if value == "equals-ignore-whitespace" {
		return og.EqualsIgnoreWhitespcae
	}  else {
		printMessage(ERROR, " Condition Operation can only be matches, contains, starts-with, ends-with, equals, contains-key, contains-value," +
			"greater-than, less-than, is-empty and equals-ignore-whitespac \n")
		os.Exit(1)
	}
	return og.Matches
}

