package main

import "testing"

func TestMagnitude(t *testing.T) {
    tests := []struct {
        repr string
        want int
    }{
        {"[9, 1]", 29},
        {"[[9, 1], [1, 9]]", 129},
    }

    for _, tt := range tests {
        t.Run(tt.repr, func (t *testing.T) {
            root := ParseNode(tt.repr)
            ans := root.magnitude()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

