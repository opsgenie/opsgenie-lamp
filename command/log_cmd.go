package command

import (
	"errors"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/logs"
	gcli "github.com/urfave/cli"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func NewCustomerLogClient(c *gcli.Context) (*logs.Client, error) {
	logsCli, cliErr := logs.NewClient(getConfigurations(c))
	if cliErr != nil {
		message := "Can not create the logs client. " + cliErr.Error()
		printMessage(INFO, message)
		return nil, errors.New(message)
	}
	printMessage(DEBUG,"Logs Client created.")
	return logsCli, nil
}

func DownloadLogs(c *gcli.Context) {
	cli, err := NewCustomerLogClient(c)
	if err != nil {
		os.Exit(1)
	}

	req := logs.ListLogFilesRequest{}
	if val, success := getVal("start", c); success {
		req.Marker = val
		req.Limit = 1000
	}

	filePath := "."
	if val, success := getVal("path", c); success {
		filePath = val
		printMessage(DEBUG,fmt.Sprintf("Downloading log files under: %s", filePath))
	} else {
		printMessage(DEBUG,"Downloading log files into current directory..")
	}

	endDate := ""
	if val, success := getVal("end", c); success {
		endDate = val
	}
	printMessage(DEBUG,"List Downloadable Logs request prepared from flags, sending request to Opsgenie..")
	for {
		response, err := cli.ListLogFiles(nil, &req)
		if err != nil {
			printMessage(ERROR, err.Error())
			os.Exit(1)
		}
		if response.Marker == "" {
			printMessage(DEBUG,"Successfully downloaded all the files")
			break
		}
		req.Marker = getLinksAndDownloadTheFile(response.Logs, endDate, filePath, cli)
		if req.Marker == "" {
			printMessage(DEBUG,"Successfully downloaded all the files")
			break
		}
	}
}

func getLinksAndDownloadTheFile(receivedLogs []logs.Log, endDate string, filePath string, cli *logs.Client) string {
	currentFileDate := ""
	for _, log := range receivedLogs {
		downloadResponse, err := cli.GenerateLogFileDownloadLink(nil, &logs.GenerateLogFileDownloadLinkRequest{
			FileName: log.FileName,
		})
		time.Sleep(time.Duration(500 * time.Millisecond))
		if err != nil {
			printMessage(DEBUG,fmt.Sprintf("Error: %s while downloading log file: %s, but proceding rest of the log files", err.Error(), log.FileName))
			continue
		}
		currentFileDate = log.FileName[:len(log.FileName)-5]
		if endDate == "" || checkDate(endDate, currentFileDate) {
			err := downloadFile(filePath+fmt.Sprintf("/%s", log.FileName), downloadResponse.LogFileDownloadLink)
			if err != nil {
				printMessage(ERROR,err.Error())
				os.Exit(1)
			}
			printMessage(DEBUG,fmt.Sprintf("Successfully downloaded file: %s", log.FileName))
		} else {
			currentFileDate = ""
			break
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

func checkDate(endDate string, currentFileDate string) bool {
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
