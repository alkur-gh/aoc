package main

import (
    "fmt"
    "strconv"
)

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

type Node interface {
    magnitude() int
    height() int
    add(Node) Node
    reduce() (Node, bool)
}

type PairNode struct {
    left, right Node
    h int
}

type ValueNode int

func MakePairNode(left, right Node) PairNode {
    height := 1 + max(left.height(), right.height())
    return PairNode{left, right, height}
}

func (p PairNode) magnitude() int {
    return 3 * p.left.magnitude() + 2 * p.right.magnitude()
}

func (p PairNode) height() int {
    return p.h
}

func (p PairNode) add(node Node) Node {
    return MakePairNode(p, node)
}

func (p PairNode) reduce() (Node, bool) {
    if p.height() <= 4 {
        return p, false
    }

    current := p
    parent_left := p
    parent_right := p
    for i := 4; i >= 1; i-- {
        if current.left.height() == i {
            parent_right = current
            current, _ = current.left.(PairNode)
        } else {
            parent_left = current
            current, _ = current.right.(PairNode)
        }
    }

    fmt.Println(parent_left)
    fmt.Println(parent_right)
    fmt.Println(current)

    return p, true
}

func (p PairNode) String() string {
    return fmt.Sprintf("[%s,%s]", p.left, p.right)
}

func (v ValueNode) magnitude() int {
    return int(v)
}

func (v ValueNode) height() int {
    return 0
}

func (v ValueNode) add(node Node) Node {
    return MakePairNode(v, node)
}

func (v ValueNode) reduce() (Node, bool) {
    panic("unexpected")
}

func (v ValueNode) String() string {
    return strconv.Itoa(int(v))
}

func ParseNode(s string) Node {
    stack := []Node{}
    for _, r := range s {
        if (r >= '0') && (r <= '9') {
            num := int(r - '0')
            stack = append(stack, ValueNode(num))
        } else if r == ']' {
            n := len(stack)
            stack[n - 2] = MakePairNode(stack[n - 2], stack[n - 1])
            stack = stack[:n - 1]
        }
    }
    return stack[0]
}

func main() {
    p1 := ParseNode("[[[[4,3],4],4],[7,[[8,4],9]]]")
    p2 := ParseNode("[1,1]")
    root := p1.add(p2)
//    fmt.Println(root)
//    fmt.Println(root.height())
//    fmt.Println(p1.reduce())
    fmt.Println(root.reduce())
//    fmt.Println(root.magnitude())
}
