package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func parseRaces(s string) []int {
	fields := strings.Fields(s)[1:]
	ret := make([]int, len(fields))

	for i, field := range fields {
		val, err := strconv.Atoi(string(field))
		if err != nil {
			log.Fatal(err)
		}
		ret[i] = val
	}
	return ret
}

func parseRace(s string) int {
	justNums := func(r rune) rune {
		if !unicode.IsDigit(r) {
			return -1
		}
		return r
	}

	numstr := strings.Map(justNums, s)

	val, err := strconv.Atoi(numstr)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func waysToWin(time, distance int) int {
	failures := 0
	for i := 0; i < distance; i++ {
		if i*(time-i) > distance {
			failures = (i - 1) * 2
			break
		}
	}
	return time - 1 - failures
}

func part1(times, distances []int) int {
	product := 1
	for i := 0; i < len(times); i++ {
		product *= waysToWin(times[i], distances[i])
	}
	return product
}

func part2(time, distance int) int {
	return waysToWin(time, distance)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	inputs := strings.Split(strings.TrimSpace(string(f)), "\n")

	//fmt.Println(inputs[0])
	//fmt.Println(inputs[1])

	if len(os.Args) == 2 || os.Args[2] == "1" {
		times := parseRaces(inputs[0])
		distances := parseRaces(inputs[1])
		fmt.Println("Part 1:", part1(times, distances))
	}

	if len(os.Args) == 2 || os.Args[2] == "2" {
		time := parseRace(inputs[0])
		distance := parseRace(inputs[1])
		fmt.Println("Part 2:", part2(time, distance))
	}
}
