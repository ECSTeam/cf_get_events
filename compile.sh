#!/bin/bash

GOOS=darwin go build -o get-events-plugin-osx
#GOOS=linux go build -o get-events-plugin-linux
#GOOS=windows GOARCH=amd64 go build -o get-events-plugin.exe
if [ $? != 0 ]; then
   printf "Error when executing compile\n"
   exit 1
fi
cf uninstall-plugin get-events
cf install-plugin -f ./get-events-plugin-osx
