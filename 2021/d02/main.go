package main

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

const (
    DFORWARD string = "forward"
    DDOWN           = "down"
    DUP             = "up"
)

type Player struct {
    x, y, aim int
}

type Command struct {
    args []string
}

func parseCommand(line string) Command {
    return Command{strings.Split(line, " ")}
}

func (p *Player) execute(c Command) {
    x, err := strconv.Atoi(c.args[1])
    check(err)

    switch c.args[0] {
    case DFORWARD:
        p.x += x
        p.y += p.aim * x
    case DDOWN:
        p.aim += x
    case DUP:
        p.aim -= x
    default:
        panic("illegal command")
    }
}

func runMovements(path string) int {
    f, err := os.Open(path)
    check(err)
    defer f.Close()

    player := Player{0, 0, 0}
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        command := parseCommand(scanner.Text())
        player.execute(command)
    }

    return player.x * player.y
}

func main() {
    path := "./tests/input.txt"
    fmt.Println("result", runMovements(path))
}
