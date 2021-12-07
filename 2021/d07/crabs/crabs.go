package crabs

import (
//    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
    "math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadPositions(path string) []int {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    if !scanner.Scan() {
        panic("expected line")
    }
    line := scanner.Text()
    var positions []int
    for _, s := range strings.Split(line, ",") {
        num, err := strconv.Atoi(s)
        check(err)
        positions = append(positions, num)
    }
    return positions
}

func FuelCost(positions []int, dest int) int {
    ret := 0
    for _, pos := range positions {
        diff := pos - dest
        if diff < 0 {
            diff = -diff
        }
        ret += diff * (diff + 1) / 2
    }
    return ret
}

func FindPositionMinimizingFuel(positions []int) (int, int) {
    sort.Ints(positions)

//    f, err := os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE, 0755);
//    check(err)
//    defer f.Close()
//    writer := bufio.NewWriter(f)

    best, bestFuel := 0, math.MaxInt64
    min, max := positions[0], positions[len(positions) - 1]

    for i := min; i <= max; i++ {
        fuel := FuelCost(positions, i)
//        grad := (FuelCost(positions, i - 1) + FuelCost(positions, i + 1)) / 2
//
//        writer.WriteString(fmt.Sprintf("%d,%d,%d\n",i, fuel, grad))

        if fuel < bestFuel {
            bestFuel = fuel
            best = i
        }
    }

//    writer.Flush()
//    f.Sync()

    return best, bestFuel
}

