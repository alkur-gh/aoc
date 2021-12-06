package lanternfish

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

func check(e error) {
    if (e != nil) {
        panic(e)
    }
}

type State struct {
    lifetime map[int]int
}

func ReadInitialState(path string) *State {
    f, err := os.Open(path)
    check(err)
    scanner := bufio.NewScanner(f)
    if (!scanner.Scan()) {
        panic("expected line")
    }
    line := scanner.Text()
    parts := strings.Split(line, ",")
    lifetime := make(map[int]int)
    for i := 0; i < len(parts); i++ {
        num, err := strconv.Atoi(parts[i])
        check(err)
        lifetime[num]++
    }
    return &State{lifetime}
}

func (state *State) Update() {
    next := make(map[int]int)
    for k, v := range state.lifetime {
        if k == 0 {
            next[8] += v
            next[6] += v
        } else {
            next[k - 1] += v
        }
    }
    state.lifetime = next
}

func (state *State) CountFish() int {
    count := 0
    for _, v := range state.lifetime {
        count += v
    }
    return count
}
