package logparser

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "time"
)

type PacketLog struct {
    LogTimestamp    time.Time
    PacketTimestamp time.Time
    SrcIP           string
    DstIP           string
}

func ParseLogFile(logFilePath string) {
    file, err := os.Open(logFilePath)
    if err != nil {
        log.Fatalf("Failed to open log file: %s", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ", ")
        if len(parts) < 4 {
            continue
        }

        logTime, err := time.Parse("2006/01/02 15:04:05", strings.Split(parts[0], " ")[0])
        if err != nil {
            log.Printf("Failed to parse log timestamp: %s", err)
            continue
        }

        packetTime, err := time.Parse(time.RFC3339, strings.Split(parts[1], " ")[1])
        if err != nil {
            log.Printf("Failed to parse packet timestamp: %s", err)
            continue
        }

        srcIP := strings.Split(parts[2], ": ")[1]
        dstIP := strings.Split(parts[3], ": ")[1]

        packetLog := PacketLog{
            LogTimestamp:    logTime,
            PacketTimestamp: packetTime,
            SrcIP:           srcIP,
            DstIP:           dstIP,
        }

        fmt.Printf("Log Time: %s, Packet Time: %s, SrcIP: %s, DstIP: %s\n",
            packetLog.LogTimestamp.Format(time.RFC3339),
            packetLog.PacketTimestamp.Format(time.RFC3339),
            packetLog.SrcIP,
            packetLog.DstIP)
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Failed to read log file: %s", err)
    }
}
