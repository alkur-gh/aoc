package node

import (
    "fmt"
    "strconv"
)

func (node *Node) String() string {
    if node.height == 0 {
        return strconv.Itoa(node.value)
    } else {
        return fmt.Sprintf("[%s,%s]", node.left, node.right)
    }
}
