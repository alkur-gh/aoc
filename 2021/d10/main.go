package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

var MATCHING = map[rune]rune{'}': '{', ']': '[', ')': '(', '>': '<'}
var POINTS = map[rune]int{'(': 1, '[': 2, '{': 3, '<': 4}

func getLineIncompletenessScore(line string) (int, bool) {
    stack := make([]rune, 0)
    for _, ch := range line {
        match, prs := MATCHING[ch]
        if prs {
            if len(stack) == 0 || stack[len(stack) - 1] != match {
                // corrupted line
                return 0, true
            }
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, ch)
        }
    }
    score := 0
    for i := len(stack) - 1; i >= 0; i-- {
        score = 5 * score + POINTS[stack[i]]
    }
    return score, false
}

func ScoreIncompleteLines(path string) int {

    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    scores := make([]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
        score, corrupted := getLineIncompletenessScore(line)
        if !corrupted {
            scores = append(scores, score)
        }
    }

    sort.Ints(scores)
    return scores[len(scores) / 2]
}

func main() {
    path := "./files/handout.txt"
    fmt.Println(ScoreIncompleteLines(path))
}
