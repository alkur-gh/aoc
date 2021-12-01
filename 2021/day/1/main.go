package main

import (
    "fmt"
    "io/ioutil"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseNumber(bytes []byte) (num int) {
    num = 0
    for _, d := range bytes {
        num = num * 10 + int(d) - int('0')
    }
    return
}

func sum(nums ...int) (total int) {
    total = 0
    for _, num := range nums {
        total += num
    }
    return
}

func updateWindow(window []int, num int) ([]int, int, int) {
    prev := sum(window...)
    window = append(window[1:], num)
    curr := sum(window...)
    return window, prev, curr
}

func countIncreasementsSlidingWindow(path string, window_size int) (count int) {
    data, err := ioutil.ReadFile(path)
    check(err)

    window := make([]int, 0)
    var prev, curr int
    i := 0
    count = 0

    for j := 0; j < len(data); j++ {
        if data[j] == 10 {
            num := parseNumber(data[i:j])
            if len(window) < window_size {
                window = append(window, num)
            } else {
                window, prev, curr = updateWindow(window, num)
                if curr > prev {
                    count++
                }
            }
            i = j + 1
        }
    }

    if i != len(data) {
         num := parseNumber(data[i:])
         window, prev, curr = updateWindow(window, num)
         if curr > prev {
             count++
         }
    }

    return
}

func main() {
    path := "input.txt"
    count := countIncreasementsSlidingWindow(path, 3)
    fmt.Println("Count:", count)
}
