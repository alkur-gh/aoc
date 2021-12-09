package main

import "testing"

func TestGetBasinScore(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/handout.txt", 1134},
        {"./files/input.txt", 856716},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            heightmap := ReadHeightmap(tt.path)
            ans := GetBasinScore(heightmap)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
