package bingo

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

type Board struct {
    size int
    positions map[int]int
}

func ReadNumbersAndBoardsFromFile(path string, size int) ([]int, []Board) {
    f, err := os.Open(path)
    check(err)
    scanner := bufio.NewScanner(f)

    var numbers []int

    scanner.Scan()
    for _, s := range strings.Split(scanner.Text(), ",") {
        num, e := strconv.Atoi(s)
        check(e)
        numbers = append(numbers, num)
    }

    var boards []Board

    for scanner.Scan() {
        board := scanBoard(scanner, size)
        boards = append(boards, board)
    }

    return numbers, boards
}

func scanBoard(scanner *bufio.Scanner, size int) Board {
    positions := make(map[int]int)
    k := 0
    for i := 0; i < size; i++ {
        scanner.Scan()
        for _, s := range strings.Fields(scanner.Text()) {
            num, e := strconv.Atoi(s)
            check(e)
            positions[num] = k
            k++
        }
    }
    return Board{size, positions}
}

func GetBestAndWorstBoards(numbers []int, boards []Board) (int, int, int, int) {
    minSteps, minStepsScore := len(numbers) + 1, 0
    maxSteps, maxStepsScore := 0, 0
    for _, board := range boards {
        steps, score := board.score(numbers)
        if steps < minSteps {
            minSteps = steps
            minStepsScore = score
        }
        if steps > maxSteps {
            maxSteps = steps
            maxStepsScore = score
        }
    }
    return minSteps, minStepsScore, maxSteps, maxStepsScore
}

func (board *Board) score(seq []int) (int, int) {
    hits := make([]bool, board.size * board.size)
    count := 0
    for _, num := range seq {
        count++
        pos, prs := board.positions[num]
        if prs {
            hits[pos] = true
        }
        if checkWin(hits, board.size) {
            return count, num * sumOfNotHit(hits, board.positions)
        }
    }
    return count, 0
}

func sumOfNotHit(hits []bool, positions map[int]int) int {
    sum := 0
    for num, pos := range positions {
        if !hits[pos] {
            sum += num
        }
    }
    return sum
}

func checkWin(hits []bool, size int) bool {
    for i := 0; i < size; i++ {
        hstreak := 0
        vstreak := 0

        for j := 0; j < size; j++ {
            if hits[i * size + j] {
                hstreak++
            }
            if hits[j * size + i] {
                vstreak++
            }
        }

        if hstreak == size || vstreak == size {
            return true
        }
    }
    return false
}

func check(e error) {
    if (e != nil) {
        panic(e)
    }
}
