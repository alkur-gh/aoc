package main

import "testing"

func TestRunMovements(t *testing.T) {
    var tests = []struct {
        path string
        want int
    }{
        {"./tests/input.txt", 1956047400},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            ans := runMovements(tt.path)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
