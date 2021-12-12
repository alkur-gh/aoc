package main

import "testing"

func TestCountPaths(t *testing.T) {
    tests := []struct {
        path string
        want int
    }{
        {"./files/simple.txt", 36},
        {"./files/handout.txt", 3509},
        {"./files/input.txt", 117095},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            system := ReadCaveSystem(tt.path)
            ans := CountPaths(system)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

func BenchmarkCountPaths(b *testing.B) {
    path := "./files/input.txt"
    for i := 0; i < b.N; i++ {
        system := ReadCaveSystem(path)
        CountPaths(system)
    }
}
