package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Jesus I should have just hardcoded my input, why is parsing harder than the
// puzzle???
func ParseInput(filename string) (stacks [][]rune, instructions [][]int) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// First loop getting the stacks
	for scanner.Scan() {
		line := scanner.Text()

		// Break when we see the stack numbers
		if unicode.IsDigit(rune(line[1])) {
			break
		}

		// 4 characters are used for every stack: two brackets, a space, and the crate
		// so we know the actual stack is /4 the index in string
		for i, v := range line {
			if unicode.IsLetter(v) {
				si := i / 4
				// Allocate slice for stack if it doesn't exist
				if len(stacks) < si+1 {
					stacks = append(stacks, make([][]rune, si+1-len(stacks))...)
				}
				stacks[si] = append(stacks[si], v)
			}
		}
	}

	// Second loop for parsing the instructions into groups of 3 ints
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")
		instruction := []int{}

		// Actual numbers on i 1, 3, and 5
		for i := 1; i < len(tokens); i += 2 {
			v, err := strconv.Atoi(tokens[i])
			if err != nil {
				log.Fatal(err)
			}
			instruction = append(instruction, v)
		}

		instructions = append(instructions, instruction)
	}

	return
}

// Push and Pop (to/from front of slice) - Why doesn't Go have stacks???
// These are reworked to take slices of size num so it can do part 2 as well
func PushCratesF(stack *[]rune, crates []rune) {
	*stack = append(crates, *stack...)
}

func PopCratesF(stack *[]rune, num int) (crates []rune, err error) {
	if len(*stack) < num {
		return nil, fmt.Errorf("Not enough crates in stack!")
	}
	crates = append(crates, (*stack)[:num]...)
	*stack = (*stack)[num:]
	return crates, nil
}

func main() {
	stacks1, instructions := ParseInput(os.Args[1])
	stacks2 := append([][]rune{}, stacks1...)

	for _, v := range instructions {

		// Part 1 we only push and pop one crate at a time so we loop
		for i := 0; i < v[0]; i++ {
			crates, _ := PopCratesF(&stacks1[v[1]-1], 1)
			PushCratesF(&stacks1[v[2]-1], crates)
		}

		// Part 2 grabs multiple crates at a time
		crates, _ := PopCratesF(&stacks2[v[1]-1], v[0])
		PushCratesF(&stacks2[v[2]-1], crates)
	}

	// Results
	fmt.Println("Part 1 Stacks (top first):")
	solution1, solution2 := "", ""
	for _, v := range stacks1 {
		fmt.Println(string(v))
		solution1 += string(v[0])
	}
	fmt.Println("Solution:", solution1)

	fmt.Println("\nPart 2 Stacks (top first):")
	for _, v := range stacks2 {
		fmt.Println(string(v))
		solution2 += string(v[0])
	}
	fmt.Println("Solution:", solution2)
}
