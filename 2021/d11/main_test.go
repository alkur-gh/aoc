package main

import "testing"

func TestFindStepWhenAllFlashed(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 195},
        {"./files/input.txt", 276},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            levels := ReadEnergyLevels(tt.path)
            ans := FindStepWhenAllFlashed(levels)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
