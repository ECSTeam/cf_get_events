#!/bin/bash

set -e

(cf uninstall-plugin "get-events" || true) && go build -o get-events-plugin main.go && cf install-plugin get-events-plugin
