# Lamp: OpsGenie Command line interface

## Introduction
Lamp is a command line utility to interact with [OpsGenie](http://www.opsgenie.com) service. It allows users to create and close alerts, attach files, etc. 
Lamp is used to integrate any management tool that can execute a shell script with OpsGenie. Lamp interacts with the OpsGenie service through the 
(RESTful Web API)[https://www.opsgenie.com/docs/api-and-client-libraries/web-api].

Lamp has a built in contextual help system for obtaining information on available commands, and available options for their use.  If you invoke lamp with the `--help` option, 
you will see the available list of commands. If you invoke lamp with the `--help` option with a specific command, you will see the options for that command.

For ease of use the *apiKey* should be set in the `opsgenie-integration.conf` file and for some flexible use cases `--apiKey` parameter can also be used 
when executing lamp commands.

