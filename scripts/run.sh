#!/bin/sh
export CAPTURE_INTERFACE=eth0
export LOG_FILE=logs/network_traffic.log

go run cmd/monitor/main.go