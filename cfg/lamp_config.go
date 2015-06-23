package cfg
import (
	"github.com/ccding/go-config-reader/config"
	"os"
	"path/filepath"
	"fmt"
	"strings"
	"github.com/cihub/seelog"
	"github.com/opsgenie/opsgenie-go-sdk/logging"
)

const (
	CONF_PATH = "LAMP_CONF_PATH"
	LOG_DIR= "LAMP_LOGS_DIR"
	SEP string = string(filepath.Separator)
)

var lampConfig *config.Config
var Verbose = false

func printVerboseMessage(message string){
	if Verbose {
		fmt.Println(message)
	}
}

func getLampHome() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error occurred while getting lamp home path" + err.Error())
		return ""
	}
	return dir + SEP
}

func LoadConfigFromGivenPath(confPath string){
	printVerboseMessage("Will read configuration from: \n--config " + confPath)
	load(confPath)
}

func LoadConfiguration(){
	confPath := os.Getenv(CONF_PATH)
	if confPath == "" {
		confPath = getLampHome()  + ".." + SEP +"conf" + SEP + "opsgenie-integration.conf"
		printVerboseMessage("LAMP_CONF_PATH environment variable is not set. Will try to read config from: \n" + confPath)
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		confPath = getLampHome()  + "conf" + SEP +"lamp.conf"
		printVerboseMessage("Could not find the file specified. Will try to read config from: \n" + confPath)
	}
	load(confPath)
}

func load(confPath string){
	if _, err := os.Stat(confPath); !os.IsNotExist(err) {
		conf := config.NewConfig(confPath)
		conf.Read()
		lampConfig = conf
		configureLog()
	} else {
		printVerboseMessage("Could not read config file: " + err.Error())
	}
}

func Get(key string) string{
	if lampConfig != nil{
		return lampConfig.Get("",key)
	}
	return ""
}

func configureLog(){
	level := lampConfig.Get("", "lamp.log.level")
	if level == ""{
		level = "warn"
		printVerboseMessage("Could not get log level from configuration, will use default \"warn\".")
	}

	logDir := os.Getenv(LOG_DIR)

	var outPath string
	logFile := lampConfig.Get("", "lamp.log.file")
	if logFile == ""{
		logFile = "lamp.log"
		printVerboseMessage("Could not get log filename from configuration. \"lamp.log\" will be used as log filename.")
	}
	if logDir != ""{
		outPath = logDir + SEP + logFile
		printVerboseMessage("Will write logs to: \n" + outPath)
	}else{
		outPath = getLampHome() + "logs" + SEP + logFile
		printVerboseMessage("LAMP_LOGS_DIR environment variable is not set. Will write logs to: \n" + outPath)
	}

	logConfig := getTemplate(outPath, level)
	logger, err := seelog.LoggerFromConfigAsBytes([] byte(logConfig))
	if err != nil {
		fmt.Println("Error occured while configuring logger: " + err.Error())
	}
	logging.UseLogger(logger)
}

func getTemplate(outPath string, level string) string {
	return `
<seelog type="sync" minlevel="`+ strings.ToLower(level)+`">
	<outputs formatid="main">
		<rollingfile formatid="main" type="date" filename="`+ outPath +`" datepattern="02-01-2006"/>
	</outputs>
	<formats>
		<format id="main" format="%Date(06/01/02 15:04:05.000) [%Level] %Msg%n"/>
	</formats>
</seelog>`
}
