package tests

import (
    "testing"
    "day16/packet"
)

func TestLiteralValue(t *testing.T) {
    tests := []struct {
        hex string
        want int
    }{
        {"D2FE28", 2021},
        {"04005AC33890", 54}, // sum of 1, 2
        {"880086C3E88112", 7}, // minimum of 7, 8, 9
        {"CE00C43D881120", 9}, // maximum of 7, 8, 9
        {"D8005AC2A8F0", 1}, // 5 < 15
        {"F600BC2D8F", 0}, // 5 > 15
        {"9C005AC2F8F0", 0}, // 5 == 15
        {"9C0141080250320F1802104A08", 1}, // 1 + 3 == 2 * 2
    }

    for _, tt := range tests {
        t.Run(tt.hex, func (t *testing.T) {
            p := packet.ExpandHex(tt.hex)
            ans := p.LiteralValue()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func TestLiteralValueFromFile(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/input.txt", 1675198555015},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            p := packet.ExpandHexFromFile(tt.path)
            ans := p.LiteralValue()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
