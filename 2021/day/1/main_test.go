package main

import (
    "testing"
)

func TestCountIncreasements(t *testing.T) {
    var tests = []struct {
        path string
        want int
    }{
        {"./tests/empty.txt", 0},
        {"./tests/one.txt", 0},
        {"./tests/two_increasing.txt", 1},
        {"./tests/two_decreasing.txt", 0},
        {"./tests/handout.txt", 7},
        {"./tests/problem.txt", 1162},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            ans := countIncreasementsSlidingWindow(tt.path, 1)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func TestCountIncreasementsWithWindowSizeOfThree(t *testing.T) {
    var tests = []struct {
        path string
        want int
    }{
        {"./tests/handout.txt", 5},
        {"./tests/problem.txt", 1190},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            ans := countIncreasementsSlidingWindow(tt.path, 3)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
