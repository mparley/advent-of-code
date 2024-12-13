package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) outside(w, h int) bool {
	return p.x < 0 || p.y < 0 || p.x >= w || p.y >= h
}

// takes a list of positions and gets all pairings
// probably could have been smarter and dont do pairs twice
// but this way we can go each direction separately and the pairAnts func is simpler
func pairs(locations []pos) [][]pos {
	pairList := make([][]pos, 0)
	for _, loc := range locations {
		for _, l := range locations {
			if l == loc {
				continue
			}
			pairList = append(pairList, []pos{loc, l})
		}
	}
	return pairList
}

// solution
// basically finds the distance between the first and second
// and then adds that onto one end once for part1 and until out of bounds for part 2
func pairAnts(pair []pos, width, height int, single bool) []pos {
	if len(pair) != 2 {
		log.Fatal("Pairs need to be 2 elements")
	}

	dx := pair[0].x - pair[1].x
	dy := pair[0].y - pair[1].y

	if single {
		np := pos{pair[0].x + dx, pair[0].y + dy}
		if np.outside(width, height) {
			return []pos{}
		}
		return []pos{{pair[0].x + dx, pair[0].y + dy}}
	}

	out := make([]pos, 0)

	// tricky i = 0 because for part 2 the antenna location is included as well
	for i := 0; true; i++ {
		np := pos{pair[0].x + (dx * i), pair[0].y + (dy * i)}
		if np.outside(width, height) {
			break
		}
		out = append(out, np)
	}

	return out
}

// useful to visualize the antinodes
func printAnts(width, height int, antinodes map[pos]struct{}) {
	out := ""
	for y := 0; y < height; y++ {
		s := ""
		for x := 0; x < width; x++ {
			if _, ok := antinodes[pos{x, y}]; ok {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
		out += s
	}
	fmt.Print(out)
}

// gets the list of antinodes from the map of antennas
func antinodes(w, h int, antennas map[rune][]pos, part1 bool) map[pos]struct{} {
	antinodes := make(map[pos]struct{})

	for _, locations := range antennas {
		pairs := pairs(locations)
		for _, pair := range pairs {
			a := pairAnts(pair, w, h, part1)
			for i := range a {
				antinodes[a[i]] = struct{}{}
			}
		}
	}

	return antinodes
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	city := strings.Fields(string(data))
	width, height := len(city[0]), len(city)
	antennas := make(map[rune][]pos)

	for y, row := range city {
		for x, col := range row {
			if col != '.' {
				antennas[col] = append([]pos{{x, y}}, antennas[col]...)
			}
		}
	}

	a1 := antinodes(width, height, antennas, true)
	a2 := antinodes(width, height, antennas, false)

	// printAnts(width, height, a1)
	printAnts(width, height, a2)

	fmt.Println("Part 1:", len(a1))
	fmt.Println("Part 2:", len(a2))
}
