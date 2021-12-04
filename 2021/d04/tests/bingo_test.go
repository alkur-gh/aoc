package tests

import (
    "testing"
    "day4/bingo"
)

func TestGetBestAndWorstBoards(t *testing.T) {
    tests := []struct {
        path string
        size int
        minStepsWant int
        minStepsScoreWant int
        maxStepsWant int
        maxStepsScoreWant int
    }{
        {"./files/handout.txt", 5, 12, 4512, 15, 1924},
        {"./files/input.txt", 5, 27, 35711, 86, 5586},
    }

    for _, tt := range tests {
        t.Run(tt.path, func (t *testing.T) {
            numbers, boards :=
                bingo.ReadNumbersAndBoardsFromFile(tt.path, tt.size)

            minStepsAns, minStepsScoreAns, maxStepsAns, maxStepsScoreAns :=
                bingo.GetBestAndWorstBoards(numbers, boards)

            if minStepsAns != tt.minStepsWant {
                t.Errorf("min steps: got %d, want %d",
                    minStepsAns, tt.minStepsWant)
            }
            if minStepsScoreAns != tt.minStepsScoreWant {
                t.Errorf("min steps score: got %d, want %d",
                    minStepsScoreAns, tt.minStepsScoreWant)
            }
            if maxStepsAns != tt.maxStepsWant {
                t.Errorf("min steps: got %d, want %d",
                    maxStepsAns, tt.maxStepsWant)
            }
            if maxStepsScoreAns != tt.maxStepsScoreWant {
                t.Errorf("min steps score: got %d, want %d",
                    maxStepsScoreAns, tt.maxStepsScoreWant)
            }
        })
    }
}
