package vents

import (
    "fmt"
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

type Segment struct {
    x1, y1, x2, y2 int
    dx, dy int
}

func ParseSegment(line string) Segment {
    parts := strings.Split(line, " -> ")
    src := strings.Split(parts[0], ",")
    dest := strings.Split(parts[1], ",")

    x1, err := strconv.Atoi(src[0])
    check(err)
    y1, err := strconv.Atoi(src[1])
    check(err)

    x2, err := strconv.Atoi(dest[0])
    check(err)
    y2, err := strconv.Atoi(dest[1])
    check(err)

    dx := 0
    if x2 > x1 {
        dx = 1
    } else if x1 > x2 {
        dx = -1
    }

    dy := 0
    if y2 > y1 {
        dy = 1
    } else if y1 > y2 {
        dy = -1
    }

    return Segment{x1, y1, x2, y2, dx, dy}
}

func ReadSegments(path string) (int, int, []Segment) {
    f, err := os.Open(path)
    check(err)
    scanner := bufio.NewScanner(f)

    var segments []Segment
    rows, cols := 0, 0
    for scanner.Scan() {
        segment := ParseSegment(scanner.Text())
        segments = append(segments, segment)
        if segment.y1 >= rows {
            rows = segment.y1 + 1
        }
        if segment.y2 >= rows {
            rows = segment.y2 + 1
        }
        if segment.x1 >= cols {
            cols = segment.x1 + 1
        }
        if segment.x2 >= cols {
            cols = segment.x2 + 1
        }
    }

    return rows, cols, segments
}

type VentsMap struct {
    rows, cols int
    array []int
}

func NewVentsMap(rows, cols int) VentsMap {
    return VentsMap{rows, cols, make([]int, rows * cols)}
}

func (vents *VentsMap) AddSegment(segment Segment) {
    x1, y1, x2, y2 := segment.x1, segment.y1, segment.x2, segment.y2
    dx, dy := segment.dx, segment.dy
    for x1 != x2 || y1 != y2 {
        vents.array[y1 * vents.cols + x1]++
        x1 += dx
        y1 += dy
    }
    vents.array[y1 * vents.cols + x1]++
}

func (vents *VentsMap) Print() {
    rows, cols := vents.rows, vents.cols
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Print(vents.array[i * cols + j], " ")
        }
        fmt.Println()
    }
}

func (vents *VentsMap) CountIntersections() int {
    count := 0
    for _, v := range vents.array {
        if v > 1 {
            count++
        }
    }
    return count
}
