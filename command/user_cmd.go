package command

import (
	"bytes"
	"errors"
	"fmt"
	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"os"
	"strconv"
	"time"
)

var configurations *client.Config

func NewUserClient(c *gcli.Context) (*user.Client, error) {
	configurations = getConfigurations(c)
	userCli, cliErr := user.NewClient(configurations)
	if cliErr != nil {
		message := "Can not create the user client. " + cliErr.Error()
		fmt.Printf("%s\n", message)
		return nil, errors.New(message)
	}
	printVerboseMessage("User Client created.")
	return userCli, nil
}

// ListUsersAction retrieves users from Opsgenie.
func ExportUsersAction(c *gcli.Context) {
	cli, err := NewUserClient(c)
	if err != nil {
		os.Exit(1)
	}

	printVerboseMessage("List users request prepared from flags, sending request to Opsgenie..")

	var users []user.User
	var offset = 0

	req := generateListUsersRequest(c)
	for {
		req.Offset = offset
		resp, err := cli.List(nil, &req)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		users = append(users, resp.Users...)

		if len(resp.Users) < req.Limit {
			break
		} else {
			offset = offset + req.Limit
		}
	}
	writeCsv(c, users)
}

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		configurations.Logger.Fatal("Cannot create file", err)
	}
	return f
}

func generateListUsersRequest(c *gcli.Context) user.ListRequest {
	req := user.ListRequest{}
	req.Limit = 100

	if val, success := getVal("query", c); success {
		req.Query = val
		printVerboseMessage("Listing users with given query")
	}

	return req
}

func writeCsv(c *gcli.Context, users []user.User) {
	csv, err := createCsv(users)

	if err != nil {
		configurations.Logger.Fatal(err)
	} else {
		destinationPath := "."
		if val, success := getVal("destinationPath", c); success {
			destinationPath = val
			printVerboseMessage(fmt.Sprintf("Creating report file under: %s", destinationPath))
		} else {
			printVerboseMessage("Creating report file into current directory..")
		}
		file := createFile(destinationPath + "/result.csv")
		defer file.Close()

		_, err := file.Write(csv)

		if err != nil {
			configurations.Logger.Fatal(err)
		} else {
			printVerboseMessage("The output file named result.csv has just been created.")
		}
	}
}

func createCsv(users []user.User) ([]byte, error) {
	var buf bytes.Buffer
	headers := []string{"id", "blocked", "verified", "username", "fullname", "roleName", "timezone",
		"locale", "country", "state", "city", "line", "zipcode", "createdAt"}

	writeHeaders(&buf, headers)
	buf.WriteString("\n")

	for _, user := range users {
		extractFields(&buf, user)
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

func writeHeaders(buf *bytes.Buffer, headers []string) {
	for index, header := range headers {
		buf.WriteString(header)
		if index < len(headers)-1 {
			buf.WriteString(",")
		}
	}
}

func extractFields(buf *bytes.Buffer, user user.User) {
	buf.WriteString(user.Id)
	buf.WriteString(",")
	buf.WriteString(strconv.FormatBool(user.Blocked))
	buf.WriteString(",")
	buf.WriteString(strconv.FormatBool(user.Verified))
	buf.WriteString(",")
	buf.WriteString(user.Username)
	buf.WriteString(",")
	buf.WriteString(user.FullName)
	buf.WriteString(",")
	buf.WriteString(user.Role.RoleName)
	buf.WriteString(",")
	buf.WriteString(user.TimeZone)
	buf.WriteString(",")
	buf.WriteString(user.Locale)
	buf.WriteString(",")
	buf.WriteString(user.UserAddress.Country)
	buf.WriteString(",")
	buf.WriteString(user.UserAddress.State)
	buf.WriteString(",")
	buf.WriteString(user.UserAddress.City)
	buf.WriteString(",")
	buf.WriteString(user.UserAddress.Line)
	buf.WriteString(",")
	buf.WriteString(user.UserAddress.ZipCode)
	buf.WriteString(",")
	buf.WriteString(user.CreatedAt.Format(time.RFC822))
	buf.WriteString(",")
}
