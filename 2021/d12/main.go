package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "math"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadCaveSystem(path string) map[string][]string {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    system := make(map[string][]string)
    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), "-")
        src, dest := parts[0], parts[1]
        system[src] = append(system[src], dest)
        system[dest] = append(system[dest], src)
    }
    return system
}

func isUpper(str string) bool {
    return str == strings.ToUpper(str)
}

func generateAllowance(system map[string][]string,
        toTry string) map[string]int {
    allowance := map[string]int{toTry: 2}
    for node, _ := range system {
        if isUpper(node) {
            allowance[node] = math.MaxInt64
        } else if node != toTry {
            allowance[node] = 1
        }
    }
    return allowance
}

func findPathsWithAllowanceFrom(system map[string][]string,
        allowance map[string]int, src string) [][]string {
    if src == "end" {
        return [][]string{[]string{src}}
    }

    allowance[src]--
    paths := [][]string{}

    for _, dest := range system[src] {
        if allowance[dest] > 0 {
            childPaths := findPathsWithAllowanceFrom(system, allowance, dest)
            crossProduct := [][]string{}
            for _, path := range childPaths {
                path = append(path, src)
                crossProduct = append(crossProduct, path)
            }
            paths = append(paths, crossProduct...)
        }
    }

    allowance[src]++

    return paths
}


func CountPaths(system map[string][]string) int {
    unique := make(map[string]bool)

    for toTry, _ := range system {
        if isUpper(toTry) || toTry == "start" || toTry == "end" {
            continue
        }

        allowance := generateAllowance(system, toTry)
        paths := findPathsWithAllowanceFrom(system, allowance, "start")
        for _, path := range paths {
            joined := strings.Join(path, ",")
            unique[joined] = true
        }
    }

    return len(unique)
}

func main() {
    path := "./files/handout.txt"
    system := ReadCaveSystem(path)
    count := CountPaths(system)
    fmt.Println(count)
}
