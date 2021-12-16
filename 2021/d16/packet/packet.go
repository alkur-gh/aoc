package packet

func (p Packet) LiteralValue() int {
    typeId := p.TypeId()
    switch typeId {
    case LITERAL_VALUE_TYPE_ID:
        res, i := 0, 6
        last := false
        for !last {
            last = (p[i] == 0)
            res = 16*res + int(8*p[i+1] + 4*p[i+2] + 2*p[i+3] + p[i+4])
            i += 5
        }
        return res
    case SUM_TYPE_ID:
        packets := p.Subpackets()
        res := 0
        for _, sb := range packets {
            res += sb.LiteralValue()
        }
        return res
    case PRODUCT_TYPE_ID:
        packets := p.Subpackets()
        res := 1
        for _, sb := range packets {
            res *= sb.LiteralValue()
        }
        return res
    case MINIMUM_TYPE_ID:
        packets := p.Subpackets()
        n := len(packets)
        if n < 1 {
            panic("expected at least one subpacket")
        }
        res := packets[0].LiteralValue()
        for i := 1; i < n; i++ {
            val := packets[i].LiteralValue()
            if val < res {
                res = val
            }
        }
        return res
    case MAXIMUM_TYPE_ID:
        packets := p.Subpackets()
        n := len(packets)
        if n < 1 {
            panic("expected at least one subpacket")
        }
        res := packets[0].LiteralValue()
        for i := 1; i < n; i++ {
            val := packets[i].LiteralValue()
            if val > res {
                res = val
            }
        }
        return res
    case GREATER_THAN_TYPE_ID:
        packets := p.Subpackets()
        if len(packets) != 2 {
            panic("expected exactly two subpackets")
        }
        res := 0
        if packets[0].LiteralValue() > packets[1].LiteralValue() {
            res = 1
        }
        return res
    case LESS_THAN_TYPE_ID:
        packets := p.Subpackets()
        if len(packets) != 2 {
            panic("expected exactly two subpackets")
        }
        res := 0
        if packets[0].LiteralValue() < packets[1].LiteralValue() {
            res = 1
        }
        return res
    case EQUAL_TO_TYPE_ID:
        packets := p.Subpackets()
        if len(packets) != 2 {
            panic("expected exactly two subpackets")
        }
        res := 0
        if packets[0].LiteralValue() == packets[1].LiteralValue() {
            res = 1
        }
        return res
    default:
        panic("unexpected type id")
    }
    panic("unexpected")
}
