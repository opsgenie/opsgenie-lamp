# Lamp: OpsGenie Command line interface

## Introduction
[![](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/opsgenie/opsgenie-lamp/command)

Lamp is a command line utility to interact with [OpsGenie](http://www.opsgenie.com) service. It allows users to create and close alerts, attach files, etc. 
Lamp is used to integrate any management tool that can execute a shell script with OpsGenie. Lamp interacts with the OpsGenie service through the 
[RESTful Web API](https://www.opsgenie.com/docs/api-and-client-libraries/web-api).

Lamp has a built in contextual help system for obtaining information on available commands, and available options for their use.  If you invoke lamp with the `--help` option, 
you will see the available list of commands. If you invoke lamp with the `--help` option with a specific command, you will see the options for that command.

For ease of use *apiKey* should be set in conf file that lamp will use, for some flexible use cases --apiKey parameter can also be used when executing lamp commands.

## Pre-requisites
* The API is built using Go 1.4.2. Some features may not be
available or supported unless you have installed a relevant version of Go.
Please click [https://golang.org/dl/](https://golang.org/dl/) to download and
get more information about installing Go on your computer.
* Make sure you have properly set both `GOROOT` and `GOPATH`
environment variables.
* Before you can begin, you need to sign up [OpsGenie](http://www.opsgenie.com) if you
don't have a valid account yet. Create an API Integration and get your API key.

## Installation
To download all packages in the repo with their dependencies, simply run

`go get github.com/opsgenie/opsgenie-lamp/...`

The command will automatically download and install necessary package files and dependencies under the src directory of the GOPATH.

Then cd to GOPATH/src/github.com/opsgenie/opsgenie-lamp and run:

`go install`

## Configuration
You can make configurations via Lamp configuration file. It's located under

`LAMP_HOME/conf/lamp.conf`

If you want to use a configuration file located in some custom location, you can define it in your commands:

`opsgenie-lamp createAlert --message "host down" --config "/opt/conf/myConfigurationFile.conf"`

## Usage
After run `go install` you can start executing commands using OpsGenie Lamp.

You can create an alert OpsGenie by executing the command below:

`lamp createAlert --message "appserver1 down" --recipients john.smith@acme.com --apiKey your_api_key`

For more information and command samples about OpsGenie Lamp, please refer to [OpsGenie Lamp](http://www.opsgenie.com/docs/lamp/lamp-command-line-interface-for-opsgenie)

