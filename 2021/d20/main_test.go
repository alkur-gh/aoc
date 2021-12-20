package main

import "testing"

func TestEnhanceMultipleTimesCountLit(t *testing.T) {
    tests := []struct {
        path string
        times int
        want int
    }{
        {"./files/handout.txt", 50, 3351},
        {"./files/input.txt", 50, 19638},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            algo, grid := ReadInput(tt.path)
            enh := grid.EnhanceMultipleTimes(algo, tt.times)
            ans := enh.CountLit()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
