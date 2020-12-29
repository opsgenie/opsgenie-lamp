package command

import (
	"errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/escalation"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func NewEscalationClient(c *gcli.Context) (*escalation.Client, error) {
	escalationcli, cliErr := escalation.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the escalation client. " + cliErr.Error()
		printMessage(INFO, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Escalation Client created.")
	return escalationcli, nil
}

// CreateEscalationAction creates an escalation at Opsgenie.
func CreateEscalationAction(c *gcli.Context) {
	cli, err := NewEscalationClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := escalation.CreateRequest{}
	if val, success := getVal("name", c); success {
		req.Name =  val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	ruleRequest := generateRuleRequest(c)
	req.Rules = ruleRequest

	if val, success := getVal("teamName", c); success {
		req.OwnerTeam = &og.OwnerTeam{
			Name: val,
		}
	}

	repeatRequest := generateRepeatRequest(c)

	req.Repeat = &repeatRequest

	printMessage(DEBUG,"Find Escalation Request Created. Sending to Opsgenie...")

	resp, err := cli.Create(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	printMessage(DEBUG,"Fetching Escaltion. RequestID: " + resp.RequestId)
	printMessage(INFO, "RequestID: " + resp.RequestId)
}

func generateRuleRequest(c *gcli.Context) []escalation.RuleRequest {

	var rules []escalation.RuleRequest

	escalationCondition := grabEscalationCondition(c)
	notifyTypes := grabNotifyType(c)
	participantTypes := grabParticipantType(c)
	participantNames := grabParticipantName(c)
	delay := grabEscalationDelayRequest(c)

	if !(len(escalationCondition) == len(notifyTypes) && len(participantNames) == len(participantTypes) && len(delay) == len(escalationCondition) && len(delay) == len(participantTypes)) {
		printMessage(ERROR,"escalationCondition, notifyTypes, participantTypes, participantNames, delay should have equal number of values")
		os.Exit(1)
	}

	for index,_ := range escalationCondition {
		rules = append(rules, escalation.RuleRequest{
			Condition: escalationCondition[index],
			NotifyType: notifyTypes[index],
			Recipient: og.Participant{
				Type: participantTypes[index],
				Name: participantNames[index],
			},
			Delay: escalation.EscalationDelayRequest{
				TimeAmount: delay[index],
			},
		})
	}

	return rules
}

func generateRepeatRequest(c *gcli.Context) escalation.RepeatRequest {

	var repeatRequest escalation.RepeatRequest

	if val, success := getVal("waitInterval", c); success {
		temp_waitInterval,_ := strconv.ParseUint(val,10,32)
		repeatRequest.WaitInterval = uint32(temp_waitInterval)
	}

	if val, success := getVal("count", c); success {
		temp_count,_ := strconv.ParseUint(val,10,32)
		repeatRequest.Count = uint32(temp_count)
	}

	tempRecipientStates := c.IsSet("recipientStatus")
	repeatRequest.ResetRecipientStates = &tempRecipientStates

	tempCloseAlert := c.IsSet("closeAlertAfterAll")
	repeatRequest.CloseAlertAfterAll = &tempCloseAlert

	return repeatRequest
}

// GetEscalationAction fetches an escalation at Opsgenie.
func GetEscalationAction(c *gcli.Context) {
	cli, err := NewEscalationClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := escalation.GetRequest{}
	if _, success := getVal("identifierType", c); success {
		req.IdentifierType = grabEscalationIdentifier(c)
	}
	if val, success := getVal("identifier", c); success {
		req.Identifier = val
	}

	printMessage(DEBUG,"Find Escalation Request Created. Sending to Opsgenie...")

	resp, err := cli.Get(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	printMessage(DEBUG,"Fetching Escaltion. RequestID: " + resp.RequestId)
	printMessage(INFO, "RequestID: " + resp.RequestId)
}

// UpdateEscalationAction updates an escalation at Opsgenie.
func UpdateEscalationAction(c *gcli.Context) {
	cli, err := NewEscalationClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := escalation.UpdateRequest{}
	if val, success := getVal("name", c); success {
		req.Name =  val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	ruleRequest := generateRuleRequest(c)
	req.Rules = ruleRequest

	if val, success := getVal("teamName", c); success {
		req.OwnerTeam = &og.OwnerTeam{
			Name: val,
		}
	}

	repeatRequest := generateRepeatRequest(c)

	req.Repeat = &repeatRequest

	if _, success := getVal("identifierType", c); success {
		req.IdentifierType = grabEscalationIdentifier(c)
	}
	if val, success := getVal("identifier", c); success {
		req.Identifier = val
	}

	printMessage(DEBUG,"Find Escalation Request Created. Sending to Opsgenie...")

	resp, err := cli.Update(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	printMessage(DEBUG,"Fetching Escaltion. RequestID: " + resp.RequestId)
	printMessage(INFO,"RequestID: " + resp.RequestId)
}

// DeleteEscalationAction deletes an escalation at Opsgenie.
func DeleteEscalationAction(c *gcli.Context) {
	cli, err := NewEscalationClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := escalation.DeleteRequest{}
	if _, success := getVal("identifierType", c); success {
		req.IdentifierType = grabEscalationIdentifier(c)
	}
	if val, success := getVal("identifier", c); success {
		req.Identifier = val
	}

	printMessage(DEBUG,"Delete Escalation Request Created. Sending to Opsgenie...")

	resp, err := cli.Delete(nil, &req)
	if err != nil {
		printMessage(ERROR,err.Error())
		os.Exit(1)
	}

	printMessage(DEBUG,"Deleting Escaltion. RequestID: " + resp.RequestId)
	printMessage(INFO,"RequestID: " + resp.RequestId)
}

func grabEscalationIdentifier(c *gcli.Context) escalation.Identifier {
	if val, success := getVal("escalationConditiom", c); success {
		if val == "id" {
			return escalation.Id
		} else  {
			return escalation.Name
		}
	}
	return escalation.Name
}

func grabEscalationCondition(c *gcli.Context) []og.EscalationCondition {
		if val, success := getVal("escalationCondition", c); success {

			var conditions []og.EscalationCondition

			escalationCondition := strings.Split(val, ",")
			for _,value := range escalationCondition {
				if value == "if-not-acked" {
					conditions = append(conditions, og.IfNotAcked)
				} else {
					conditions = append(conditions, og.IfNotClosed)
				}
			}
			return conditions
		}
		return nil
}

func grabNotifyType(c *gcli.Context) []og.NotifyType {
	if val, success := getVal("notifyType", c); success {

		var notifyTypes []og.NotifyType

		tempNotifyTypes := strings.Split(val, ",")
		for _,value := range tempNotifyTypes {
			switch value {
			case "next":
				notifyTypes = append(notifyTypes, og.Next)
			case "previous":
				notifyTypes = append(notifyTypes, og.Previous)
			case "default":
				notifyTypes = append(notifyTypes, og.Default)
			case "users":
				notifyTypes = append(notifyTypes, og.Users)
			case "admins":
				notifyTypes = append(notifyTypes, og.Admins)
			case "random":
				notifyTypes = append(notifyTypes, og.Random)
			default:
				notifyTypes = append(notifyTypes, og.All)
			}
		}
		return notifyTypes
	}
	return nil
}

func grabParticipantType(c *gcli.Context) []og.ParticipantType {
	if val, success := getVal("particpantType", c); success {

		var paticipantTypes []og.ParticipantType

		tempParticipantType := strings.Split(val, ",")
		for _,value := range tempParticipantType {
			switch value {
			case "user":
				paticipantTypes = append(paticipantTypes, og.User)
			case "team":
				paticipantTypes = append(paticipantTypes, og.Team)
			case "escalation":
				paticipantTypes = append(paticipantTypes, og.Escalation)
			case "schedule":
				paticipantTypes = append(paticipantTypes, og.Schedule)
			case "none":
				paticipantTypes = append(paticipantTypes, og.None)
			}
		}
		return paticipantTypes
	}
	return nil
}

func grabParticipantName(c *gcli.Context) []string {
	if val, success := getVal("particpantName", c); success {
		return strings.Split(val, ",")
	}
	return nil
}

func grabEscalationDelayRequest(c *gcli.Context) []uint32 {
	if val, success := getVal("delay", c); success {

		var Delays []uint32

		tempDelays := strings.Split(val, ",")
		for _,value := range tempDelays{
			tempUIntDelay,_ :=  strconv.ParseUint(value, 10,32)
			Delays = append(Delays, uint32(tempUIntDelay))
		}
		return Delays
	}
	return nil
}
