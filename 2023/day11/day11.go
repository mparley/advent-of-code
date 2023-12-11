package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type pair struct {
	a, b point
}

func newPair(a, b point) pair {
	if a.x == b.x {
		if a.y < b.y {
			return pair{a, b}
		} else {
			return pair{b, a}
		}
	}
	if a.x < b.x {
		return pair{a, b}
	}
	return pair{b, a}
}

func absDiff(l, r int) int {
	if l < r {
		return r - l
	}
	return l - r
}

func (p pair) distance() int {
	return absDiff(p.b.x, p.a.x) + absDiff(p.b.y, p.a.y)
}

func inRange(lo, hi, target int) bool {
	if lo > hi {
		lo, hi = hi, lo
	}
	return lo < target && target < hi
}

func sumDists(gpairs map[pair]int, ex int) int {
	sum := 0
	for gpair, emptyCount := range gpairs {
		sum += gpair.distance() + (ex * emptyCount)
	}
	return sum
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")
	emptyRows := make(map[int]struct{})
	emptyCols := make(map[int]struct{})
	galaxies := make(map[point]struct{})

	for y, line := range input {
		empty := true
		for x, col := range line {
			if y == 0 && col == '.' {
				emptyCols[x] = struct{}{}
			}

			if col == '#' {
				galaxies[point{x, y}] = struct{}{}
				if _, ok := emptyCols[x]; ok {
					delete(emptyCols, x)
				}
				empty = false
			}
		}
		if empty {
			emptyRows[y] = struct{}{}
		}
	}

	// gpairs tracks the num of empty rows and cols between
	// the pair of points
	gpairs := make(map[pair]int)
	for g1 := range galaxies {
		for g2 := range galaxies {
			if g1 == g2 {
				continue
			}
			gp := newPair(g1, g2)
			if _, ok := gpairs[gp]; ok {
				continue
			} else {
				mts := 0
				for col := range emptyCols {
					if inRange(gp.a.x, gp.b.x, col) {
						mts++
					}
				}
				for row := range emptyRows {
					if inRange(gp.a.y, gp.b.y, row) {
						mts++
					}
				}
				gpairs[gp] = mts
			}
		}
	}

	fmt.Println("Part 1:", sumDists(gpairs, 2-1))
	fmt.Println("Part 2:", sumDists(gpairs, 1000000-1))
}
