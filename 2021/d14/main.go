package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PolymerTemplate []rune

func ReadInput(path string) (PolymerTemplate, map[string]string) {
	f, err := os.Open(path)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	template := []rune(scanner.Text())

	pairs := make(map[string]string)
	scanner.Scan()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		pairs[parts[0]] = parts[1]
	}

	return template, pairs
}

func (template PolymerTemplate) Pairs() map[string]int {
	result := make(map[string]int)
	n := len(template)
	for i := 0; i < n-1; i++ {
		from := string(template[i : i+2])
		result[from]++
	}
	return result
}

func getMinMaxFromCounter(counter map[rune]int) (rune, rune) {
	var min, max rune
	for k, v := range counter {
		if max == 0 {
			min, max = k, k
		} else {
			if counter[min] > v {
				min = k
			}
			if counter[max] < v {
				max = k
			}
		}
	}
	return min, max
}

func applyInsertionPairs(
	tmpPairs map[string]int,
	insPairs map[string]string) (map[string]int, map[rune]int) {

	result := make(map[string]int)
	insCount := make(map[rune]int)
	for from, v := range tmpPairs {
		insCount[rune(insPairs[from][0])] += v
		insCount[rune(from[1])] += v
		left := string(from[0]) + insPairs[from]
		right := insPairs[from] + string(from[1])
		result[left] += v
		result[right] += v
	}
	return result, insCount
}

func getCounterFromTemplatePairs(pairs map[string]int) map[rune]int {
	result := make(map[rune]int)
	for k, v := range pairs {
		result[rune(k[0])] += v
		result[rune(k[1])] += v
	}
	return result
}

func Solve(
	template PolymerTemplate,
	insPairs map[string]string, steps int) int {

	tmpPairs := template.Pairs()
	var insCount map[rune]int
	for i := 0; i < steps; i++ {
		tmpPairs, insCount = applyInsertionPairs(tmpPairs, insPairs)
	}
	insCount[template[len(template)-1]]--
	counter := getCounterFromTemplatePairs(tmpPairs)
	min, max := getMinMaxFromCounter(counter)
	return (counter[max] - insCount[max]) - (counter[min] - insCount[min])
}

func main() {
	path := "./files/handout.txt"
	steps := 40
	template, insPairs := ReadInput(path)
	fmt.Println(Solve(template, insPairs, steps))
}
