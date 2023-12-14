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

func (a *pos) add(b pos) {
	a.x += b.x
	a.y += b.y
}

type dir int

const (
	NORTH dir = iota
	WEST
	SOUTH
	EAST
)

func oppo(d dir) dir {
	return (d + 2) % 4
}

func getDir(d dir) pos {
	switch d {
	case NORTH:
		return pos{0, -1}
	case EAST:
		return pos{1, 0}
	case SOUTH:
		return pos{0, 1}
	case WEST:
		return pos{-1, 0}
	}
	return pos{0, 0}
}

func platformString(rocks map[pos]rune, w, h int) string {
	s := ""
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if _, ok := rocks[pos{x, y}]; ok {
				s += string(rocks[pos{x, y}])
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func inBounds(p pos, w, h int) bool {
	return 0 <= p.x && p.x < w && 0 <= p.y && p.y < h
}

func tilt(rocks map[pos]rune, d dir, w, h int) {
	r := rocks

	for rock, shape := range rocks {
		if shape != 'O' {
			continue
		}
		delete(r, rock)
		i := rock
		for ; inBounds(i, w, h); i.add(getDir(d)) {
			if _, ok := rocks[i]; ok && rocks[i] == '#' {
				break
			}
		}
		i.add(getDir(oppo(d)))
		for ; inBounds(i, w, h); i.add(getDir(oppo(d))) {
			if _, ok := r[i]; !ok {
				r[i] = 'O'
				break
			}
		}
	}
	rocks = r
}

func northLoad(rocks map[pos]rune, h int) int {
	load := 0
	for position, shape := range rocks {
		if shape == 'O' {
			load += h - position.y
		}
	}
	return load
}

func cycle(rocks map[pos]rune, w, h int) {
	for d := NORTH; d <= EAST; d++ {
		tilt(rocks, d, w, h)
	}
}

func part1(rocks map[pos]rune, w, h int) int {
	//fmt.Println(platformString(rocks,w,h))
	r := rocks
	tilt(r, NORTH, w, h)
	//fmt.Println()
	//fmt.Println(platformString(rocks,w,h))
	return northLoad(r, h)
}

func part2(rocks map[pos]rune, w, h, times int) int {
	cache := make(map[string]int)
	r := rocks
	loop := 0
	end := 0

	for i := 0; i < times; i++ {
		ps := platformString(r, w, h)
		if _, ok := cache[ps]; ok {
			loop = i - cache[ps]
			end = i
			break
		} else {
			cache[ps] = i
		}
		cycle(r, w, h)
	}
	remaining := (times - end) % loop
	for i := 0; i < remaining; i++ {
		cycle(r, w, h)
	}

	return northLoad(r, h)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	rocks := make(map[pos]rune)
	w, h := len(input[0]), len(input)

	for y, line := range input {
		for x, ch := range line {
			if ch == 'O' {
				rocks[pos{x, y}] = 'O'
			} else if ch == '#' {
				rocks[pos{x, y}] = '#'
			}
		}
	}

	fmt.Println("Part 1:", part1(rocks, w, h))
	fmt.Println("Part 2:", part2(rocks, w, h, 1000000000))
}
