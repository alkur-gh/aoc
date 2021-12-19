package node

func (root *Node) Reduce() {
    for root.Explode() || root.Split() {
    }
}

func (root *Node) Explode() bool {
    if root.height <= 4 {
        return false
    }

    var parent *Node
    parent = nil
    current := root

    for i := 4; i >= 1; i-- {
        current.parent = parent
        parent = current
        if current.left.height == i {
            current = current.left
        } else {
            current = current.right
        }
    }
    current.parent = parent

    prev := current.prev()
    if prev != nil {
        prev.value += current.left.value
    }

    next := current.next()
    if next != nil {
        next.value += current.right.value
    }

    current.value = 0
    current.left = nil
    current.right = nil

    root.updateHeights()

    return true
}

func (root *Node) Split() bool {
    node := root.findLeftmostBigValue()
    if node == nil {
        return false
    }
    node.left = MakeValueNode(node.value / 2)
    node.right = MakeValueNode((node.value + 1) / 2)
    root.updateHeights()
    return true
}

func (node *Node) findLeftmostBigValue() *Node {
    if node.height == 0 {
        if node.value >= 10 {
            return node
        } else {
            return nil
        }
    }
    left := node.left.findLeftmostBigValue()
    if left != nil {
        return left
    }
    return node.right.findLeftmostBigValue()
}

func (node *Node) updateHeights() {
    if node.left == nil {
        node.height = 0
    } else {
        node.left.updateHeights()
        node.right.updateHeights()
        node.height = max(node.left.height, node.right.height) + 1
    }
}

func (node *Node) prev() *Node {
    for (node.parent != nil) && (node.parent.left == node) {
        node = node.parent
    }
    if node.parent == nil {
        return nil
    }
    node = node.parent.left
    for node.right != nil {
        node = node.right
    }
    return node
}

func (node *Node) next() *Node {
    for (node.parent != nil) && (node.parent.right == node) {
        node = node.parent
    }
    if node.parent == nil {
        return nil
    }
    node = node.parent.right
    for node.left != nil {
        node = node.left
    }
    return node
}

