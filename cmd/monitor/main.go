package main

import (
    "networkMonitor/pkg/capture"
    "networkMonitor/pkg/logger"
    "networkMonitor/pkg/logparser"
    "os"
)

func main() {
    // Initialize logger
    logFilePath := "logs/network_traffic.log"
    logger := logger.InitLogger(logFilePath)

    // Check if the program should run in capture mode or parse mode
    if len(os.Args) > 1 && os.Args[1] == "parse" {
        // Parse and display log file entries
        logparser.ParseLogFile(logFilePath)
    } else {
        // Start packet capture
        capture.StartCapture("en0", logger) // Use the correct network interface
    }
}
