package main

import (
    "fmt"
    "os"
    "bufio"
    "day18/node"
)

func ReadNodes(path string) []*node.Node {
    f, _ := os.Open(path)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    var nodes []*node.Node
    for scanner.Scan() {
        nodes = append(nodes, node.Parse(scanner.Text()))
    }
    return nodes
}

func AddUpNodes(nodes ...*node.Node) *node.Node {
    res := nodes[0]
    n := len(nodes)
    for i := 1; i < n; i++ {
        res = res.Add(nodes[i])
    }
    return res
}

func BestTwoSum(nodes []*node.Node) int {
    n := len(nodes)
    max := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i != j {
                res := AddUpNodes(nodes[i], nodes[j])
                magnitude := res.Magnitude()
                if magnitude > max {
                    max = magnitude
                }
            }
        }
    }
    return max
}

func main() {
    path := "./files/handout.txt"
    if len(os.Args) > 1 {
        path = os.Args[1]
    }
    nodes := ReadNodes(path)
    max := BestTwoSum(nodes)
    fmt.Println(max)
}
