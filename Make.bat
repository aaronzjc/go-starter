@echo off

set GO111MODULE=on
if "%1" == "dev" (
    go build -o ./bin/api.exe ./cmd/api/main.go
    start ./bin/api.exe -c ./conf/dev.yml
)