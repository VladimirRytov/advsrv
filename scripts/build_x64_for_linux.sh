#!/bin/bash

trap "exit 1" ERR
cp ./assets/icons/file.png ./internal/handlers/filehandler
GOARCH=amd64 go build -C ./cmd/advsrv/ -v -ldflags "-w -s -linkmode=external -X github.com/VladimirRytov/advsrv/internal/front/cli/start.Version=$2" -o $1
echo "====> building for linux is done <===="