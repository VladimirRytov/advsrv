#!/bin/bash

errorHandler() {
  echo ERROR: ${BASH_COMMAND} failed with error code $?
  exit 1
}

trap  errorHandler ERR

create_icon() {
gitvers=$(git tag --contains | sed -n 's/v//p')

go install github.com/tc-hib/go-winres@latest
~/go/bin/go-winres init
cp ./assets/icons/*.png ./winres
cp ./scripts/winres-template.json ./winres/winres.json
sed -i 's/${VERSION}'/"$gitvers"'/' ./winres/winres.json
~/go/bin/go-winres make
}

set -o errtrace
trap "exit 1" ERR
create_icon
mv ./*.syso ./cmd/advsrv/
CC=x86_64-w64-mingw32-gcc PKG_CONFIG_PATH=/usr/x86_64-w64-mingw32/sys-root/mingw/lib/pkgconfig GOOS=windows GOARCH=amd64 \
go build -C ./cmd/advsrv/ -v -ldflags "-w -s -linkmode=external -X main.version=$2" -o $1

echo "====> building for windows is done <===="