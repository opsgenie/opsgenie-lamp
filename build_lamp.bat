@echo off
cd %GOPATH%\src\github.com\opsgenie\opsgenie-lamp
go get ...
cd ..
rmdir /s /q "opsgenie-go-sdk"
git clone https://github.com/opsgenie/opsgenie-go-sdk.git
cd opsgenie-go-sdk
git fetch
git checkout sdk_enhancements
cd ../opsgenie-lamp
SETLOCAL
set GOOS=windows
SETLOCAL
set GOARCH=386
go build

SETLOCAL
set GOOS=linux
SETLOCAL
set GOARCH=386
go build

