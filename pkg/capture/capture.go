package capture

import (
    "encoding/hex"
    "encoding/json"
    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
    "log"
    "time"
    "unicode/utf8"
)

type PacketData struct {
    SrcIP        string
    DstIP        string
    Protocol     string
    Length       int
    SrcPort      string
    DstPort      string
    EthernetType string
    TCPFlags     string
    TTL          uint8
    Payload      string
    HumanPayload string
    Timestamp    time.Time
}

func isPrintableASCII(s string) bool {
    for _, c := range s {
        if c < 32 || c > 126 {
            return false
        }
    }
    return true
}

func toHumanReadable(payload []byte) string {
    if utf8.Valid(payload) && isPrintableASCII(string(payload)) {
        return string(payload)
    }
    return hex.EncodeToString(payload)
}

func StartCapture(interfaceName string, logger *log.Logger) {
    handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
    if err != nil {
        log.Fatalf("Failed to open device for capturing: %s", err)
    }
    defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        networkLayer := packet.NetworkLayer()
        if networkLayer == nil {
            continue
        }

        transportLayer := packet.TransportLayer()
        applicationLayer := packet.ApplicationLayer()
        ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
        var srcPort, dstPort, tcpFlags, payload, humanPayload string
        var ttl uint8
        var ethernetType string

        if tcpLayer, ok := transportLayer.(*layers.TCP); ok {
            srcPort = tcpLayer.SrcPort.String()
            dstPort = tcpLayer.DstPort.String()
            if tcpLayer.SYN {
                tcpFlags += "SYN "
            }
            if tcpLayer.ACK {
                tcpFlags += "ACK "
            }
            if tcpLayer.FIN {
                tcpFlags += "FIN "
            }
            if tcpLayer.RST {
                tcpFlags += "RST "
            }
            if tcpLayer.PSH {
                tcpFlags += "PSH "
            }
            if tcpLayer.URG {
                tcpFlags += "URG "
            }
        } else if udpLayer, ok := transportLayer.(*layers.UDP); ok {
            srcPort = udpLayer.SrcPort.String()
            dstPort = udpLayer.DstPort.String()
        }

        if ipLayer, ok := networkLayer.(*layers.IPv4); ok {
            ttl = ipLayer.TTL
        }

        if ethLayer, ok := ethernetLayer.(*layers.Ethernet); ok {
            ethernetType = ethLayer.EthernetType.String()
        }

        if applicationLayer != nil {
            payload = hex.EncodeToString(applicationLayer.Payload())
            humanPayload = toHumanReadable(applicationLayer.Payload())
        }

        packetData := PacketData{
            SrcIP:        networkLayer.NetworkFlow().Src().String(),
            DstIP:        networkLayer.NetworkFlow().Dst().String(),
            Protocol:     networkLayer.LayerType().String(),
            Length:       len(packet.Data()),
            SrcPort:      srcPort,
            DstPort:      dstPort,
            EthernetType: ethernetType,
            TCPFlags:     tcpFlags,
            TTL:          ttl,
            Payload:      payload,
            HumanPayload: humanPayload,
            Timestamp:    time.Now(),
        }

        logEntry, _ := json.Marshal(packetData)
        logger.Printf(string(logEntry))
    }
}
