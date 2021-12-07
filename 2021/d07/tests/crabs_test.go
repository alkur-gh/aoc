package tests

import (
    "testing"
    "day7/crabs"
)

func TestFindPositionMinimizingFuel(t *testing.T) {
    tests := []struct {
        path string
        posWant int
        fuelWant int
    }{
        {"./files/handout.txt", 5, 168},
        {"./files/input.txt", 484, 93397632},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            positions := crabs.ReadPositions(tt.path)
            pos, fuel := crabs.FindPositionMinimizingFuel(positions)
            if pos != tt.posWant {
                t.Errorf("pos: got %d, want %d", pos, tt.posWant)
            }
            if fuel != tt.fuelWant {
                t.Errorf("fuel: got %d, want %d", fuel, tt.fuelWant)
            }
        })
    }
}
