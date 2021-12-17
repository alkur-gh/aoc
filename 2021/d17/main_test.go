package main

import (
    "fmt"
    "testing"
)

func TestCountHittingVelocities(t *testing.T) {
    tests := []struct {
        r Rectangle
        want int
    }{
        {Rectangle{Point{20, -10}, Point{30, -5}}, 112},
        {Rectangle{Point{211, -124}, Point{232, -69}}, 2032},
    }

    for _, tt := range tests {
        t.Run(fmt.Sprintf("%s", tt.r), func (t *testing.T) {
            ans := tt.r.CountHittingVelocities(Point{0, 0})
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
