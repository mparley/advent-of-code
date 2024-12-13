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

func (l pos) add(r pos) pos {
	return pos{l.x + r.x, l.y + r.y}
}

type direction int

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

func (d direction) step() pos {
	return []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}[d]
}

// gets next direction (90 degree turn)
func (d direction) next() direction {
	return (d + 1) % 4
}

type state struct {
	p pos
	d direction
}

// checks if p is outside the dimensions
func outside(dim, p pos) bool {
	return p.x < 0 || p.y < 0 || p.x >= dim.x || p.y >= dim.y
}

// runs the gaurd's path, breaks and returns true when it detects a loop
func runAndLoopDetect(dim pos, obs map[pos]struct{}, start state) (map[state]struct{}, bool) {
	path := make(map[state]struct{}, 0)
	path[start] = struct{}{}

	// while current position is inside get next position and check if its in the
	// obstruction so we can rotate. If not we move the gaurd and check for a loop
	for curr, dir := start.p, start.d; !outside(dim, curr); {
		next := curr.add(dir.step())
		if _, ok := obs[next]; ok {
			dir = dir.next()
		} else {
			curr = next

			// loop check
			if _, ok := path[state{curr, dir}]; ok {
				return path, true
			}

			path[state{curr, dir}] = struct{}{}
		}
	}

	return path, false
}

// solves for both parts depending on bool passed
// basically runs the gaurd func to get a path
// if we want part 1: it creates a set of positions so we can count them
// part 2: it runs the gaurd for every possible blockage on the path and counts
// the loops returned as true
func solve(dim pos, obs map[pos]struct{}, start state, part1 bool) int {
	path, _ := runAndLoopDetect(dim, obs, start)

	posSet := make(map[pos]struct{})
	for key := range path {
		posSet[key.p] = struct{}{}
	}

	if part1 {
		return len(posSet) - 1 // -1 because the function captures the step outside
	}

	// since we can't block the start pos we need to delete it
	delete(posSet, start.p)

	count := 0
	for p := range posSet {
		// we make a deep copy of the obstructions and add the pos to block to it
		obsCopy := make(map[pos]struct{})
		for key := range obs {
			obsCopy[key] = struct{}{}
		}
		obsCopy[p] = struct{}{}

		if _, loop := runAndLoopDetect(dim, obsCopy, start); loop {
			count++
		}
	}
	return count
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Fields(string(data))
	dim := pos{len(rows[0]), len(rows)}
	obstructions := make(map[pos]struct{}, 0)
	var start state

	// input parsing
	for y, row := range rows {
		for x, col := range row {
			if col == '#' {
				obstructions[pos{x, y}] = struct{}{}
			} else if col == '^' {
				start = state{pos{x, y}, NORTH}
			}
		}
	}

	fmt.Println("Part 1:", solve(dim, obstructions, start, true))
	fmt.Println("Part 2:", solve(dim, obstructions, start, false))
}
