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

setlocal
set GOOS=windows
set GOARCH=386
echo building for %GOOS%/%GOARCH%
go build

set GOOS=linux
set GOARCH=386
echo building for %GOOS%/%GOARCH%
go build

endlocal
