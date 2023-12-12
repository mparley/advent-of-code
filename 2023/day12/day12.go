package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func contains(runes []rune, r rune) bool {
	for _, ch := range runes {
		if ch == r {
			return true
		}
	}
	return false
}

func requiredLength(sizes []int) int {
	sum := 0
	for _, s := range sizes {
		sum += s
	}
	return sum + len(sizes) - 1
}

func eval(springs []rune, sizes []int, cache *map[string]int) int {
	var ret int = 0
	s := fmt.Sprint(springs, sizes)
	if _, ok := (*cache)[s]; ok {
		return (*cache)[s]
	}

	// Only undamaged or '?' left and out of sizes, made it to end
	if !contains(springs, '#') && len(sizes) == 0 {
		ret = 1
		// At end, springs groupable into last size
	} else if !contains(springs, '.') && len(sizes) == 1 && len(springs) == sizes[0] {
		ret = 1
		// Not enough length in springs to accomodate number of groups, or no sizes
	} else if len(springs) < requiredLength(sizes) || len(sizes) == 0 {
		ret = 0
	} else {
		// if substring of first size can be grouped and terminated, eval rest of string
		if !contains(springs[:sizes[0]], '.') && springs[sizes[0]] != '#' {
			ret += eval(springs[sizes[0]+1:], sizes[1:], cache)
		}
		// first char can be '.' so eval rest of string as if it was '.'
		if springs[0] != '#' {
			ret += eval(springs[1:], sizes, cache)
		}
	}

	(*cache)[s] = ret
	return ret
}

func part1(springRecords [][]rune, groupSizes [][]int) int {
	var sum int = 0
	cache := make(map[string]int)
	for i, springs := range springRecords {
		count := eval(springs, groupSizes[i], &cache)
		sum += count
	}
	return sum
}

func part2(springRecords [][]rune, groupSizes [][]int) int {
	for i := range springRecords {
		springs := springRecords[i]
		sizes := groupSizes[i]
		for j := 0; j < 4; j++ {
			springRecords[i] = append(springRecords[i], '?')
			springRecords[i] = append(springRecords[i], springs...)
			groupSizes[i] = append(groupSizes[i], sizes...)
		}
	}
	return part1(springRecords, groupSizes)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")
	var springRecords [][]rune
	var damagedGroupSizes [][]int
	for _, line := range input {
		fields := strings.Fields(line)
		springRecords = append(springRecords, []rune(fields[0]))
		sizes := strings.Split(fields[1], ",")
		groupSize := make([]int, 0, len(sizes))
		for _, size := range sizes {
			s, err := strconv.Atoi(size)
			if err != nil {
				log.Fatal(err)
			}
			groupSize = append(groupSize, s)
		}
		damagedGroupSizes = append(damagedGroupSizes, groupSize)
	}

	fmt.Println("Part 1:", part1(springRecords, damagedGroupSizes))
	fmt.Println("Part 2:", part2(springRecords, damagedGroupSizes))
}
