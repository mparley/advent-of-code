package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readin(file string) [][]string {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(f), "\n")

	var ret [][]string
	last := 0
	for i, line := range input {
		if line == "" {
			ret = append(ret, input[last:i])
			last = i + 1
		}
	}
	return ret
}

func rotate(block []string) []string {
	newBlock := make([]string, len(block[0]))
	for _, line := range block {
		for j, col := range line {
			newBlock[j] += string(col)
		}
	}
	return newBlock
}

func compare(a, b string) int {
	diffs := 0
	for i, c := range a {
		if c != rune(b[i]) {
			diffs++
		}
	}
	return diffs
}

func summarize(block []string, limit int) int {
	possible := make(map[int]int)

	for i, line := range block {
		if i == 0 {
			continue
		}
		diffs := compare(line, block[i-1])
		if diffs <= limit {
			possible[i] = 0
		}
	}

	for k := range possible {
		loDist := k - 1
		if loDist > len(block)-k-1 {
			loDist = len(block) - k - 1
		}
		found := true
		for i := 0; i <= loDist; i++ {
			possible[k] += compare(block[k-1-i], block[k+i])
			if possible[k] > limit {
				found = false
				break
			}
		}
		if found && possible[k] == limit {
			return k
		}
	}
	return 0
}

func run(blocks [][]string, limit int) int {
	sum := 0
	for _, block := range blocks {
		hor := summarize(block, limit)
		if hor > 0 {
			sum += hor * 100
		} else {
			sum += summarize(rotate(block), limit)
		}
	}
	return sum
}

func main() {
	blocks := readin(os.Args[1])

	fmt.Println("Part 1:", run(blocks, 0))
	fmt.Println("Part 2:", run(blocks, 1))
}
