package command

import (
	"errors"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

func NewIncidentClient(c *gcli.Context) (*incident.Client, error) {
	incidentcli, cliErr := incident.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the incident client. " + cliErr.Error()
		fmt.Printf("%s\n", message)
		return nil, errors.New(message)
	}
	printVerboseMessage("Incident Client created.")
	return incidentcli, nil
}

func CreateIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req :=incident.CreateRequest{}
	if val, success := getVal("message", c); success {
		req.Message =  val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	req.Responders = grabIncidentResponders(c)

	if val, success := getVal("tags", c); success {
		req.Tags = strings.Split(val, ",")
	}

	req.Details = grabIncidentDetails(c)

	req.Priority = grabIncidentPriority(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	if val, success := getVal("serviceId", c); success {
		req.ServiceId =  val
	}

	notifyStakeHolders := c.IsSet("notifyStakeHolders")

	req.StatusPageEntity = grabStatusPageEntity(c)

	req.NotifyStakeholders = &notifyStakeHolders

	printVerboseMessage("Create Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.Create(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Creating Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func DeleteIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.DeleteRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	printVerboseMessage("Delete Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.Delete(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Deleting Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func GetIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.GetRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	printVerboseMessage("Get Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.Get(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Fetching Incident. RequestID: " + resp.RequestId)
	output, err := resultToJSON(resp, true)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Incident : " + output)
}

func ListIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.ListRequest{}

	if val, success := getVal("limit", c); success {
		req.Limit, _ =  strconv.Atoi(val)
	}

	req.Sort = grabSortField(c)

	if val, success := getVal("offset", c); success {
		req.Offset, _ =  strconv.Atoi(val)
	}

	req.Order = getOrderIncident(c)

	if val, success := getVal("query", c); success {
		req.Query =  val
	}

	printVerboseMessage("List Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.List(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Fetching List of Incident. RequestID: " + resp.RequestId)
	output, err := resultToJSON(resp, true)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Incident List : " + output)

}

func CloseIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.CloseRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	printVerboseMessage("Close Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.Close(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Cosing Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func AddNoteIncidentAction(c *gcli.Context) {

	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.AddNoteRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	printVerboseMessage("Add Note to Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.AddNote(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Adding Note to Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func AddResponderIncidentAction(c *gcli.Context) {

	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.AddResponderRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	req.Responders = grabIncidentResponders(c)

	printVerboseMessage("Add Responder to Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.AddResponder(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Adding Responder to Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)

}

func AddTagsIncidentAction(c *gcli.Context) {


	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.AddTagsRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	if val, success := getVal("tags", c); success {
		req.Tags =  strings.Split(val, ",")
	}


	printVerboseMessage("Add Tags to Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.AddTags(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Adding Tags to Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func RemoveTagsIncidentAction(c *gcli.Context) {

	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.RemoveTagsRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	if val, success := getVal("tags", c); success {
		req.Tags =  strings.Split(val, ",")
	}

	printVerboseMessage("Remove Tags from Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.RemoveTags(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Removing Tags from Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func AddDetailsIncidentAction(c *gcli.Context) {

	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.AddDetailsRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	req.Details = grabIncidentDetails(c)

	printVerboseMessage("Add Details to Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.AddDetails(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Adding details to Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func RemoveDetailsIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.RemoveDetailsRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("note", c); success {
		req.Note =  val
	}

	if val, success := getVal("keys", c); success {
		req.Keys =  strings.Split(val, ",")
	}

	printVerboseMessage("Remove Details from Incident Request Created. Sending to Opsgenie...")

	resp, err := cli.RemoveDetails(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Removing Details from Incident. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func UpdatePriorityIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.UpdatePriorityRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	req.Priority =  grabIncidentPriority(c)

	printVerboseMessage("Update Incident Priority Request Created. Sending to Opsgenie...")

	resp, err := cli.UpdatePriority(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Incident Priority is being updated. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func UpdateMessageIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.UpdateMessageRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("message", c); success {
		req.Message =  val
	}

	printVerboseMessage("Update Incident Message Request Created. Sending to Opsgenie...")

	resp, err := cli.UpdateMessage(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Incident message is being updated. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func UpdateDescriptionIncidentAction(c *gcli.Context) {
	cli, err := NewIncidentClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := incident.UpdateDescriptionRequest{}

	if val, success := getVal("identifier", c); success {
		req.Id =  val
	}

	req.Identifier = grabIncidentIdentifierType(c)

	if val, success := getVal("description", c); success {
		req.Description =  val
	}

	printVerboseMessage("Update Incident Description Request Created. Sending to Opsgenie...")

	resp, err := cli.UpdateDescription(nil, &req)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	printVerboseMessage("Incident Description is being updated. RequestID: " + resp.RequestId)
	fmt.Println("RequestID: " + resp.RequestId)
}

func grabIncidentIdentifierType(c *gcli.Context) incident.IdentifierType {
	if val, success := getVal("identifierType", c); success {
		if val == "tiny" {
			return incident.Tiny
		}
	}
	return incident.Id
}

func grabSortField(c *gcli.Context) incident.SortField {
	if val, success := getVal("sortField", c); success {
		if val == "createdAt" {
			return incident.CreatedAt
		}else if val == "tinyId" {
			return incident.TinyId
		}else if val == "message" {
			return incident.Message
		}else if val == "status" {
			return incident.Status
		}else if val == "isSeen" {
			return incident.IsSeen
		}else {
			return incident.Owner
		}
	}
	return incident.CreatedAt
}

func getOrderIncident(c *gcli.Context) incident.Order {
	if val, success := getVal("order", c); success {
		if val == "asc" {
			return incident.Asc
		}
	}
	return incident.Desc
}

func grabIncidentResponders(c *gcli.Context) []incident.Responder {

	var responders []incident.Responder

	responderTypes := grabResponderTypes(c)
	var responderIdentifiers []string

	if val, success := getVal("responder", c); success {
		responderIdentifiers =  strings.Split(val, ",")
	}

	if len(responderTypes) != len(responderIdentifiers) {
		fmt.Println("type and responders should have equal values")
		os.Exit(1)
	}

	for index,_ := range responderTypes {
		if responderTypes[index] == incident.User {
			responders = append(responders, incident.Responder{
				Type: responderTypes[index],
				Id:   responderIdentifiers[index],
			})
		} else {
			responders = append(responders, incident.Responder{
				Type: responderTypes[index],
				Name:   responderIdentifiers[index],
			})
		}
	}

	return responders
}

func grabResponderTypes(c *gcli.Context) []incident.ResponderType {
	if val, success := getVal("type", c); success {

		var types []incident.ResponderType

		responderType := strings.Split(val, ",")
		for _,value := range responderType {
			if value == "user" {
				types = append(types, incident.User)
			} else {
				types = append(types, incident.Team)
			}
		}
		return types
	}
	return nil
}

func grabIncidentDetails(c *gcli.Context) map[string]string {

	var detailKeys []string
	var detailValues []string
	var details = make(map[string]string)

	if val, success := getVal("detailKeys", c); success {
		detailKeys = strings.Split(val, ",")
	}

	if val, success := getVal("detailValues", c); success {
		detailValues = strings.Split(val, ",")
	}

	if len(detailKeys) != len(detailValues) {
		fmt.Println("detailKeys and detailValues should have equal values")
		os.Exit(1)
	}

	for index,value := range detailKeys {
		details[value] = detailValues[index]
	}
	return details
}

func grabIncidentPriority(c *gcli.Context) incident.Priority {

	if val, success := getVal("priority", c); success {
		if val == "P1" {
			return incident.P1
		} else if val == "P2" {
			return incident.P2
		} else if val == "P3" {
			return incident.P3
		} else if val == "P4" {
			return incident.P4
		} else if val == "P5"{
			return incident.P5
		}
	}

	fmt.Println("Please add correct Priority")
	os.Exit(1)
	return incident.P3
}

func grabStatusPageEntity(c *gcli.Context) *incident.StatusPageEntity {

	var statusPageEntity incident.StatusPageEntity

	if val, success := getVal("statusPageEntityTitle", c); success {
		statusPageEntity.Title =  val
	}

	if val, success := getVal("statusPageEntityDescription", c); success {
		statusPageEntity.Description = val
	}

	return &statusPageEntity
}

