#!/bin/bash

set -e

(cf uninstall-plugin "get-events" || true) && go build -o get-events-plugin cf_get_events.go && cf install-plugin get-events-plugin
