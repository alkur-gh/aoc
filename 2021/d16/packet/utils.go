package packet

import (
    "fmt"
)

func (p Packet) Print() {
    for i, _ := range p {
        fmt.Printf("%2d ", i)
    }
    fmt.Println()
    for _, b := range p {
        fmt.Printf("%2d ", b)
    }
    fmt.Println()
}

func (p Packet) Version() int {
    return int(4*p[0] + 2*p[1] + p[2])
}

func (p Packet) TypeId() int {
    return int(4*p[3] + 2*p[4] + p[5])
}

func (p Packet) LengthTypeId() int {
    return int(p[6])
}

func (p Packet) SubpacketsTotalLength() int {
    val := 0
    for i := 7; i < 7 + 15; i++ {
        val = 2*val + int(p[i])
    }
    return val
}

func (p Packet) NumberOfSubpackets() int {
    val := 0
    for i := 7; i < 7 + 11; i++ {
        val = 2*val + int(p[i])
    }
    return val
}

func (p Packet) End() int {
    typeId := p.TypeId()
    switch typeId {
    case LITERAL_VALUE_TYPE_ID:
        i := 6
        last := false
        for !last {
            last = (p[i] == 0)
            i += 5
        }
        return i
    default:
        lengthTypeId := p.LengthTypeId()
        switch lengthTypeId {
        case TOTAL_LENGTH_TYPE_ID:
            return 7 + 15 + p.SubpacketsTotalLength()
        case AMOUNT_LENGTH_TYPE_ID:
            end := 7 + 11
            n := p.NumberOfSubpackets()
            for i := 0; i < n; i++ {
                end += p[end:].End()
            }
            return end
        default:
            panic("enexpected")
        }
    }
}

func (p Packet) Subpackets() []Packet {
    packets := []Packet{}

    lengthTypeId := p.LengthTypeId()
    switch lengthTypeId {
    case TOTAL_LENGTH_TYPE_ID:
        start := 7 + 15
        tl := p.SubpacketsTotalLength()
        for start < 7 + 15 + tl {
            end := p[start:].End()
            packets = append(packets, p[start:start+end])
            start += end
        }
        return packets
    case AMOUNT_LENGTH_TYPE_ID:
        start := 7 + 11
        n := p.NumberOfSubpackets()
        for i := 0; i < n; i++ {
            end := p[start:].End()
            packets = append(packets, p[start:start+end])
            start += end
        }
        return packets
    }

    return packets
}
