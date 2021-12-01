package main

import (
    "fmt"
    "os"
    "bufio"
    "io"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func sum(nums []int) (total int) {
    total = 0
    for _, num := range nums {
        total += num
    }
    return
}

func countIncreasementsSlidingWindow(path string, window_size int) (count int) {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    reader := bufio.NewReader(f)

    var num int
    window := make([]int, 0, window_size + 1)
    count = 0

    for {
        if _, err := fmt.Fscanln(reader, &num); err == io.EOF {
            break
        } else {
            check(err)
        }

        window = append(window, num)
        if len(window) > window_size {
            if sum(window[1:]) > sum(window[:window_size]) {
                count++
            }
            /* GC won't collect first element
             * yet after recreation of array as a result of `append`
             * there won't be any garbage, prayge
             */
            window = window[1:]
        }
    }

    return
}

func main() {
    path := "tests/problem.txt"
    count := countIncreasementsSlidingWindow(path, 3)
    fmt.Println("Count:", count)
}
