package command

import (
	"fmt"
	gcli "github.com/codegangsta/cli"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	cl "github.com/opsgenie/opsgenie-go-sdk/log"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func DownloadLogs(c *gcli.Context) {
	cli, err := NewCustomerLogClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := cl.ListLogFilesRequest{}
	if val, success := getVal("start", c); success {
		req.Marker = val
	}

	filePath := "."
	if val, success := getVal("path", c); success {
		filePath = val
		printVerboseMessage(fmt.Sprintf("Downloading log files under: %s", filePath))
	} else {
		printVerboseMessage("Downloading log files into current directory..")
	}

	endDate := ""
	if val, success := getVal("end", c); success {
		endDate = val
	}

	printVerboseMessage("List Downloadable Logs request prepared from flags, sending request to OpsGenie..")

	for {
		response, err := cli.ListLogFiles(req)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			os.Exit(1)
		}
		if response.Marker == "" {
			printVerboseMessage("Successfully downloaded all the files")
			break
		}
		req.Marker = getLinksAndDownloadTheFile(response.Logs, endDate, filePath, cli)
	}
}

func getLinksAndDownloadTheFile(logs []cl.Log, endDate string, filePath string, cli *ogcli.OpsGenieLogClient) (string) {
	currentFileDate := ""
	for _, logg := range logs {
		downloadResponse, err := cli.LogFileDownloadLink(cl.GenerateLogFileDownloadRequest{
			Filename: logg.Filename,
		})
		time.Sleep(time.Duration(500 * time.Millisecond))
		if err != nil {
			printVerboseMessage(fmt.Sprintf("Error: %s while downloading log file: %s, but proceding rest of the log files", err.Error(), logg.Filename))
			continue
		}
		currentFileDate = logg.Filename[:len(logg.Filename)-5]
		if endDate == "" || checkDate(endDate, currentFileDate) {
			downloadFile(filePath+fmt.Sprintf("/%s", logg.Filename), downloadResponse)
			printVerboseMessage(fmt.Sprintf("Successfully downloaded file: %s", logg.Filename))
		}
	}
	return currentFileDate
}

func downloadFile(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func checkDate(endDate string, currentFileDate string) (bool) {
	a := strings.Split(endDate, "-")
	b := strings.Split(currentFileDate, "-")
	for i, s := range a {
		var ai, bi int
		fmt.Sscanf(s, "%d", &ai)
		fmt.Sscanf(b[i], "%d", &bi)
		if ai > bi {
			return true
		}
		if bi > ai {
			return false
		}
	}
	return true
}