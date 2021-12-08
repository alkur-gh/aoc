package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "sort"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func IteratePatternsAndOutputs(path string, ch chan []string) {
    f, err := os.Open(path)
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " | ")
        patterns := strings.Split(parts[0], " ")
        outputs := strings.Split(parts[1], " ")
        sort.Slice(patterns, func (i, j int) bool {
            return len(patterns[i]) < len(patterns[j])
        })
        ch <- patterns
        ch <- outputs
    }
    ch <- nil
}

func removeRunes(from, which string) string {
    res := from
    for _, r := range which {
        res = strings.ReplaceAll(res, string(r), "")
    }
    return res
}

var digits = map[string]int {
    "abcefg": 0, "cf": 1, "acdeg": 2, "acdfg": 3, "bcdf": 4,
    "abdfg": 5, "abdefg": 6, "acf": 7, "abcdefg": 8, "abcdfg": 9,
}

func determineDigit(str string) int {
    for k, v := range digits {
        if len(k) == len(str) && len(removeRunes(str, k)) == 0 {
            return v
        }
    }
    panic("can't determine the digit")
}

func determineOutput(outputs []string, mapping map[string]rune) int {
    mapFunc := func (r rune) rune { return mapping[string(r)] }
    result := 0
    for _, output := range outputs {
        digit := determineDigit(strings.Map(mapFunc, output))
        result = result * 10 + digit
    }
    return result
}

func determineMapping(patterns []string) map[string]rune {
    /* finding a, bd, cf */
    cf := patterns[0]
    a := removeRunes(patterns[1], cf)
    bd := removeRunes(patterns[2], cf + a)


    /* finding g, e */
    var g string
    var e string
    five1 := removeRunes(patterns[3], cf + a + bd)
    five2 := removeRunes(patterns[4], cf + a + bd)
    five3 := removeRunes(patterns[5], cf + a + bd)
    if len(five1) == 1 {
        g = five1
        if len(five2) == 2 {
            e = removeRunes(five2, g)
        } else {
            e = removeRunes(five3, g)
        }
    } else {
        g = five2
        e = removeRunes(five1, g)
    }

    /* finding b, d */
    var b string
    var d string

    six1 := removeRunes(patterns[6], cf + a + g + e)
    six2 := removeRunes(patterns[7], cf + a + g + e)
    six3 := removeRunes(patterns[8], cf + a + g + e)
    if len(six1) == 1 {
        b = six1
    } else if len(six2) == 1 {
        b = six2
    } else {
        b = six3
    }

    d = removeRunes(bd, b)

    /* finding c, f */
    var c string
    var f string
    six1 = removeRunes(patterns[6], bd + a + g + e)
    six2 = removeRunes(patterns[7], bd + a + g + e)
    six3 = removeRunes(patterns[8], bd + a + g + e)

    if len(six1) == 1 {
        f = six1
    } else if len(six2) == 1 {
        f = six2
    } else {
        f = six3
    }

    c = removeRunes(cf, f)

    return map[string]rune {
        a: 'a', b: 'b', c: 'c', d: 'd', e: 'e', f: 'f', g: 'g',
    }
}

func SumOfAllOutputs(ch chan []string) int {
    var patterns, outputs []string
    result := 0

    for patterns = <-ch; patterns != nil; patterns = <-ch {
        outputs = <-ch
        mapping := determineMapping(patterns)
        result += determineOutput(outputs, mapping)
    }

    return result
}

func main() {
    path := "./tests/files/input.txt"
    ch := make(chan []string, 2)
    go IteratePatternsAndOutputs(path, ch)
    result := SumOfAllOutputs(ch)
    fmt.Println(result)
}
