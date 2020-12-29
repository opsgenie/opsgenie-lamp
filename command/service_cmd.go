package command

import (
	"github.com/opsgenie/opsgenie-go-sdk-v2/service"
	gcli "github.com/urfave/cli"
	"os"
	"strconv"
)

func NewServiceClient(c *gcli.Context) (*service.Client){
	serviceCli, cliErr := service.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the Service client. " + cliErr.Error()
		printMessage(INFO, message)
		os.Exit(1)
	}
	printMessage(DEBUG,"Service Client created.")
	return serviceCli
}

// CreateServiceAction creates a new service in OpsGenie
func CreateServiceAction(c *gcli.Context) {
	cli := NewServiceClient(c)
	
	req := service.CreateRequest{}
	if val, success := getVal("name", c); success {
		req.Name = val
	}

	if val, success := getVal("teamId", c); success {
		req.TeamId = val
	}

	if val, success := getVal("visibility", c); success {
		req.Visibility = service.Visibility(val)
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	printMessage(DEBUG,"Create Service Request Created. Sending to Opsgenie...")

	resp, err := cli.Create(nil, &req)
	exitOnErr(err)

	printMessage(DEBUG,"Creating Service. RequestID: " + resp.RequestId)
	printMessage(INFO,"RequestID: " + resp.RequestId)
}

// UpdateServiceAction updates a service in OpsGenie
func UpdateServiceAction(c *gcli.Context) {
	cli := NewServiceClient(c)

	req := service.UpdateRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}

	if val, success := getVal("name", c); success {
		req.Name = val
	}

	if val, success := getVal("description", c); success {
		req.Description = val
	}

	if val, success := getVal("visibility", c); success {
		req.Visibility = service.Visibility(val)
	}

	printMessage(DEBUG,"Update Service Request Created. Sending to Opsgenie...")

	resp, err := cli.Update(nil, &req)
	exitOnErr(err)

	printMessage(DEBUG,"Updating Service. RequestID: " + resp.RequestId)
	printMessage(INFO,"RequestID: " + resp.RequestId)
}

// DeleteServiceAction updates a service in OpsGenie
func DeleteServiceAction(c *gcli.Context) {
	cli := NewServiceClient(c)

	req := service.DeleteRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}

	printMessage(DEBUG,"Delete Service Request Created. Sending to Opsgenie...")

	resp, err := cli.Delete(nil, &req)
	exitOnErr(err)

	printMessage(DEBUG,"Deleting Service. RequestID: " + resp.RequestId)
	printMessage(INFO,"RequestID: " + resp.RequestId)
}

// GetServiceAction updates a service in OpsGenie
func GetServiceAction(c *gcli.Context) {
	cli := NewServiceClient(c)

	req := service.GetRequest{}
	if val, success := getVal("id", c); success {
		req.Id = val
	}

	printMessage(DEBUG,"Get Service Request Created. Sending to Opsgenie...")

	resp, err := cli.Get(nil, &req)
	renderResponse(c, resp, err)
}

// ListServiceAction updates a service in OpsGenie
func ListServiceAction(c *gcli.Context) {
	cli := NewServiceClient(c)

	req := service.ListRequest{}
	if val, success := getVal("limit", c); success {
		limit, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			os.Exit(2)
		}
		req.Limit = int(limit)
	}

	if val, success := getVal("offset", c); success {
		offset, err := strconv.Atoi(val)
		if err != nil {
			os.Exit(2)
		}
		req.Offset = offset
	}

	printMessage(DEBUG,"Get Service List Request Created. Sending to Opsgenie...")

	resp, err := cli.List(nil, &req)
	renderResponse(c, resp, err)

}

