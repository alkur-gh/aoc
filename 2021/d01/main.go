package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func countIncreasementsSlidingWindow(path string, window_size int) (count int) {
    f, err := os.Open(path)
    check(err)
    defer f.Close()

    window := make([]int, 0, window_size)
    front_idx := 0
    count = 0

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        check(err)

        if len(window) < window_size {
            window = append(window, num)
        } else {
            if num > window[front_idx] {
                count++
            }
            window[front_idx] = num
            front_idx = (front_idx + 1) % window_size
        }
    }

    return
}

func main() {
    path := "tests/problem.txt"
    count := countIncreasementsSlidingWindow(path, 3)
    fmt.Println("Count:", count)
}
