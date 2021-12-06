package tests

import (
    "fmt"
    "testing"
    lf "day6/lanternfish"
)

func TestCountFish(t *testing.T) {
    tests := []struct {
        path string
        days int
        want int
    }{
        {"./files/handout.txt", 80, 5934},
        {"./files/handout.txt", 256, 26984457539},
        {"./files/input.txt", 80, 396210},
        {"./files/input.txt", 256, 1770823541496},
    }

    for _, tt := range tests {
        t.Run(tt.path + ":" + fmt.Sprint(tt.days), func (t *testing.T) {
            state := lf.ReadInitialState(tt.path)
            for i := 0; i < tt.days; i++ {
                state.Update()
            }
            ans := state.CountFish()

            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
