package test

import (
    "testing"
    "day5/vents"
)

func TestCountIntersects(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 12},
        {"./files/input.txt", 16518},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            rows, cols, segments := vents.ReadSegments(tt.path)
            ventsMap := vents.NewVentsMap(rows, cols)
            for _, segment := range segments {
                ventsMap.AddSegment(segment)
            }

            ans := ventsMap.CountIntersections()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
