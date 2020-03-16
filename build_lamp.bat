@echo off
cd %GOPATH%\src\github.com\opsgenie\opsgenie-lamp
go get .

setlocal
set GOOS=windows
set GOARCH=386
echo building for %GOOS%/%GOARCH%
go build -mod=vendor

set GOOS=linux
set GOARCH=386
echo building for %GOOS%/%GOARCH%
go build -mod=vendor

set GOOS=darwin
set GOARCH=386
echo building for %GOOS%/%GOARCH%
go build -mod=vendor -o opsgenie-lamp-mac

endlocal
