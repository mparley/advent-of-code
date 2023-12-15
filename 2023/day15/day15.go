package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	label    string
	focalLen int
}

func holdiayASH(s string) int {
	curr := 0
	for _, ch := range s {
		curr += int(ch)
		curr *= 17
		curr %= 256
	}
	return curr
}

func removeLens(box *[]lens, label string) {
	pos := -1
	for i, l := range *box {
		if l.label == label {
			pos = i
			break
		}
	}
	if pos >= 0 {
		*box = append((*box)[:pos], (*box)[pos+1:]...)
	}
}

func addLens(box *[]lens, l lens) {
	for i := range *box {
		if (*box)[i].label == l.label {
			(*box)[i].focalLen = l.focalLen
			return
		}
	}
	*box = append(*box, l)
}

func printBoxes(lightBoxes [][]lens) {
	for i, box := range lightBoxes {
		if len(box) > 0 {
			fmt.Println(i, ":", box)
		}
	}
}

func parseInstruction(input string) (string, bool, int) {
	label := ""

	for i, ch := range input {
		if ch == '-' {
			return label, false, 0
		}
		if ch == '=' {
			focalLen, err := strconv.Atoi(input[i+1:])
			if err != nil {
				log.Fatal(err)
			}
			return label, true, focalLen
		}
		label += string(ch)
	}

	log.Fatal("parseInstruction didn't find = or -")
	return label, false, 0
}

func part1(input []string) int {
	sum := 0
	for _, step := range input {
		sum += holdiayASH(step)
	}
	return sum
}

func part2(input []string) int {
	lightBoxes := make([][]lens, 256)
	for _, line := range input {
		label, insert, focal := parseInstruction(line)
		box := holdiayASH(label)
		if insert {
			addLens(&lightBoxes[box], lens{label, focal})
		} else {
			removeLens(&lightBoxes[box], label)
		}
	}

	sum := 0
	for i, box := range lightBoxes {
		for j, lens := range box {
			sum += (i + 1) * (j + 1) * lens.focalLen
		}
	}
	return sum
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), ",")

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))

}
