package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	ADD operation = iota
	MULTIPLY
	CONCAT
	LASTOP
)

func digits(n int) int {
	if n/10 == 0 {
		return 1
	}
	return 1 + digits(n/10)
}

// Parses the inputs into int slices, might have been easier for
// part 2 and the concatenation operator if we kept them as strings
// but oh well
func parseInput(data []byte) [][]int {
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	out := make([][]int, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Fields(line)
		equations := make([]int, 0, len(tokens))

		for i, token := range tokens {
			var val int
			var err error
			if i == 0 {
				val, err = strconv.Atoi(token[:len(token)-1])
			} else {
				val, err = strconv.Atoi(token)
			}
			if err != nil {
				log.Fatal(err)
			}
			equations = append(equations, val)
		}
		out = append(out, equations)
	}
	return out
}

// recursive solver
// takes the list of numbers and operation
// performs operation on first 2 entries and then passes rest
// of list into itself
func solver(target int, op operation, nums []int, part2 bool) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	var val int
	switch op {
	case ADD:
		val = nums[0] + nums[1]
	case MULTIPLY:
		val = nums[0] * nums[1]
	case CONCAT:
		val = nums[0]
		for i := 0; i < digits(nums[1]); i++ {
			val *= 10
		}
		val += nums[1]
	default:
		val = 0
	}

	copy := append([]int{val}, nums[2:]...)

	for i := ADD; i < LASTOP; i++ {
		if !part2 && i == CONCAT {
			continue
		}
		found := solver(target, i, copy, part2)
		if found {
			return true
		}
	}

	return false
}

// helper function, probably unnecessary but cleaner
func solvable(equation []int, part2 bool) bool {
	for i := ADD; i < LASTOP; i++ {
		if !part2 && i == CONCAT {
			continue
		}
		found := solver(equation[0], i, equation[1:], part2)
		if found {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	equations := parseInput(data)

	sum1, count1, sum2, count2 := 0, 0, 0, 0
	for _, equation := range equations {
		if solvable(equation, false) {
			// fmt.Println(equation[0], "is solvable!")
			count1++
			sum1 += equation[0]
		}
		if solvable(equation, true) {
			count2++
			sum2 += equation[0]
		}
	}

	fmt.Println("Part 1:", sum1, "(", count1, "equations )")
	fmt.Println("Part 2:", sum2, "(", count2, "equations )")
}
