#!/bin/bash

echo "please input os code (1:linux, 2:darwin, 3:windows):"
read oscode

date=`date +%Y%m%d`
if [ ${oscode} = 1 ]; then
  osname="linux"
  out="go-working-calendar-linux"
fi
if [ ${oscode} = 2 ]; then
  osname="darwin"
  out="go-working-calendar-darwin"
fi
if [ ${oscode} = 3 ]; then
  osname="windows"
  out="go-working-calendar-windows.exe"
fi


echo "select build os: ${osname}"
echo "please wait for a while ..."
CGO_ENABLED=0 GOOS=${osname} GOARCH=amd64 go build -o dist/${date}/${out}
cp ./config.toml dist/${date}/config.toml
cp ./calendar.json dist/${date}/calendar.json
cp ./go-working-calendar.service dist/${date}/go-working-calendar.service
cp ./install-service.sh dist/${date}/install-service.sh
cp ./README.md dist/${date}/README.md

echo "build finish!"