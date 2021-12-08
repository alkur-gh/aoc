package main

import "testing"

func TestSumOfAllOutputs(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./tests/files/handout.txt", 61229},
        {"./tests/files/input.txt", 946346},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            ch := make(chan []string, 2)
            go IteratePatternsAndOutputs(tt.path, ch)
            ans := SumOfAllOutputs(ch)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
