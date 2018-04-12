package command

import (
	gcli "github.com/codegangsta/cli"
	"github.com/opsgenie/opsgenie-go-sdk/userv2"
	"fmt"
	"os"
	"log"
	"bytes"
	"strconv"
	"time"
)

// ListUsersAction retrieves users from OpsGenie.
func ExportUsersAction(c *gcli.Context) {
	cli, err := NewUserClient(c)
	if err != nil {
		os.Exit(1)
	}

	printVerboseMessage("List users request prepared from flags, sending request to OpsGenie..")

	users := []userv2.User{}
	var offset int = 0

	req := generateListUsersRequest(c)
	for {
		req.Offset = offset
		resp, err := cli.List(req)

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

func createFile(p string) *os.File{
	f, err := os.Create(p)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	return f
}

func getDestinationPath(c *gcli.Context) string {
	var destinationPath string = "."
	val, success := getVal("destinationPath", c)

	if success {
		destinationPath = val
	} else {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		destinationPath = dir
	}
	return destinationPath
}

func generateListUsersRequest(c *gcli.Context) (userv2.ListUsersRequest) {
	req := userv2.ListUsersRequest{}
	req.Limit = 100

	if val, success := getVal("query", c); success {
		req.Query = val;
		printVerboseMessage("Listing users with given query.")
	}

	return req
}

func writeCsv(c *gcli.Context, users []userv2.User) {
	csv, err := createCsv(users)

	if err != nil {
		log.Fatal(err)
	} else {
		destinationPath := getDestinationPath(c)

		file := createFile(destinationPath + "/result.csv")
		defer file.Close()

		_, err := file.Write(csv)

		if err != nil {
			log.Fatal(err)
		} else {
			printVerboseMessage("The output file named result.csv has just been created.")
		}
	}
}

func createCsv(users []userv2.User) ([]byte, error){
	var buf bytes.Buffer
	headers := []string {"id", "blocked", "verified", "username", "fullname", "roleId", "roleName", "timezone",
							"locale", "country", "state", "city", "line", "zipcode", "createdAt", "mutedUntil"}

	writeHeaders(&buf, headers)
	buf.WriteString("\n")

	for _, user := range users {
		extractFields(&buf, user)
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

func writeHeaders(buf *bytes.Buffer, headers []string){
	for index, header := range headers {
		buf.WriteString(header)
		if index < len(headers) -1 {
			buf.WriteString(",")
		}
	}
}

func extractFields(buf *bytes.Buffer, user userv2.User){
	buf.WriteString(user.ID)
	buf.WriteString(",")
	buf.WriteString(strconv.FormatBool(user.Blocked))
	buf.WriteString(",")
	buf.WriteString(strconv.FormatBool(user.Verified))
	buf.WriteString(",")
	buf.WriteString(user.Username)
	buf.WriteString(",")
	buf.WriteString(user.FullName)
	buf.WriteString(",")
	buf.WriteString(user.Role.ID)
	buf.WriteString(",")
	buf.WriteString(user.Role.Name)
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
	if !user.MutedUntil.IsZero(){
		buf.WriteString(user.MutedUntil.Format(time.RFC822))
	} else {
	buf.WriteString("")
	}
}


