package node

type Node struct {
    height int
    value int
    left, right, parent *Node
}

func (node *Node) Magnitude() int {
    if node.height == 0 {
        return node.value
    } else {
        return 3*node.left.Magnitude() + 2*node.right.Magnitude()
    }
}

func (node *Node) Add(other *Node) *Node {
    res := MakePairNode(node, other).Copy()
    res.Reduce()
    return res
}

func (node *Node) Copy() *Node {
    if node.height == 0 {
        return MakeValueNode(node.value)
    } else {
        return MakePairNode(node.left.Copy(), node.right.Copy())
    }
}
