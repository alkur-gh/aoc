package main

import "testing"

func TestScoreIncompleteLines(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 288957},
        {"./files/input.txt", 4361305341},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            ans := ScoreIncompleteLines(tt.path)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
