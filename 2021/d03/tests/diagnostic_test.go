package tests

import (
    "testing"
    "day3/diagnostic"
)

func TestPowerConsumption(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 198},
        {"./files/input.txt", 2003336},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            report := diagnostic.ReadReportFromFile(tt.path)
            ans := report.PowerConsumption()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func TestLifeSupportRating(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 230},
        {"./files/input.txt", 1877139},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            report := diagnostic.ReadReportFromFile(tt.path)
            ans := report.LifeSupportRating()
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
