package packet

import (
    "os"
    "bufio"
)

func ExpandHex(hex string) Packet {
    packet := []byte{}
    for _, r := range hex {
        packet = append(packet, HEX_TABLE[r]...)
    }
    return packet
}

func ExpandHexFromFile(path string) Packet {
    f, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    scanner.Scan()
    return ExpandHex(scanner.Text())
}
