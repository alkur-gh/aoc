package node

func MakePairNode(left, right *Node) *Node {
    height := left.height
    if right.height > height {
        height = right.height
    }
    return &Node{height + 1, 0, left, right, nil}
}

func MakeValueNode(value int) *Node {
    return &Node{0, value, nil, nil, nil}
}

func Parse(s string) *Node {
    stack := []*Node{}
    for _, r := range s {
        if (r >= '0') && (r <= '9') {
            num := int(r - '0')
            stack = append(stack, MakeValueNode(num))
        } else if r == ']' {
            n := len(stack)
            stack[n - 2] = MakePairNode(stack[n - 2], stack[n - 1])
            stack = stack[:n - 1]
        }
    }
    return stack[0]
}
