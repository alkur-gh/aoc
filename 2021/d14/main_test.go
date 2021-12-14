package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		path  string
		steps int
		want  int
	}{
		{"./files/handout.txt", 10, 1588},
		{"./files/handout.txt", 40, 2188189693529},
		{"./files/input.txt", 10, 3284},
		{"./files/input.txt", 40, 4302675529689},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			template, insPairs := ReadInput(tt.path)
			ans := Solve(template, insPairs, tt.steps)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
