package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type label struct {
	val int
	x   int
	y   int
}

func probe(m *[]string, x int, y int) label {
	if x < 0 || x >= len((*m)[0]) || y < 0 || y >= len((*m)) {
		return label{val: -1}
	}

	if !unicode.IsDigit(rune((*m)[y][x])) {
		return label{val: -1}
	}

	s := ""
	sx := x

	for i := x; i >= 0 && unicode.IsDigit(rune((*m)[y][i])); i-- {
		s = string((*m)[y][i]) + s
		sx = i
	}

	for i := x + 1; i < len((*m)[y]) && unicode.IsDigit(rune((*m)[y][i])); i++ {
		s += string((*m)[y][i])
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return label{v, sx, y}
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	schematic := strings.Split(string(f), "\n")
	parts := make(map[label]string)
	ratioSum := 0

	for i, line := range schematic {
		// cur := ""
		for j, ch := range line {
			if !unicode.IsDigit(ch) && ch != '.' {
				res := make([]label, 8)
				res[0] = probe(&schematic, j-1, i-1)
				res[1] = probe(&schematic, j, i-1)
				res[2] = probe(&schematic, j+1, i-1)
				res[3] = probe(&schematic, j-1, i)
				res[4] = probe(&schematic, j+1, i)
				res[5] = probe(&schematic, j-1, i+1)
				res[6] = probe(&schematic, j, i+1)
				res[7] = probe(&schematic, j+1, i+1)

				uniq := make(map[label]string)
				prod := 1

				for _, v := range res {
					if v.val > 0 {
						uniq[v] = string(ch)
					}
				}

				for k, v := range uniq {
					prod *= k.val
					parts[k] = v
				}

				if ch == '*' && len(uniq) == 2 {
					ratioSum += prod
				}
			}
		}
	}

	// fmt.Println(parts)

	sum := 0
	for part := range parts {
		sum += part.val
	}
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", ratioSum)

}
