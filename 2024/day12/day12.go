package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) inside(w, h int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < w && p.y < h
}

func (p pos) neighbors() []pos {
	return []pos{{p.x, p.y - 1}, {p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}}
}

func maxint(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func minint(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// A dfs search to identify and return a region
// visited is tracked so we don't look at regions we already identified
func search(garden []string, start pos, visited map[pos]bool) map[pos]bool {
	curr := garden[start.y][start.x]
	group := make(map[pos]bool)

	st := []pos{start}
	w, h := len(garden[0]), len(garden)

	for len(st) > 0 {
		p := st[0]
		st = append([]pos{}, st[1:]...)
		visited[p] = true
		group[p] = true

		for _, n := range p.neighbors() {
			if n.inside(w, h) && !visited[n] && garden[n.y][n.x] == curr {
				st = append([]pos{n}, st...)
			}
		}
	}

	return group
}

// Counts perimeter sides
// first finds the limits of the area
// then runs through that region vertically then horizontally
// kinda jank to save writing everything twice
func perimSides(region map[pos]bool) int {
	min := []int{math.MaxInt, math.MaxInt}
	max := []int{0, 0}

	for p := range region {
		min[0] = minint(min[0], p.x)
		max[0] = maxint(max[0], p.x)
		min[1] = minint(min[1], p.y)
		max[1] = maxint(max[1], p.y)
	}

	count := 0

	// first will be a vert sweep then second will be horizontal
	for i := range min {
		for j := min[i]; j <= max[i]; j++ {
			edge := []bool{false, false}
			for k := min[(i+1)%2]; k <= max[(i+1)%2]+1; k++ {
				x, y := j, k
				next := []pos{{x + 1, y}, {x - 1, y}}

				if i == 1 {
					x, y = k, j
					next = []pos{{x, y + 1}, {x, y - 1}}
				}

				// if we find empty spots above/below or left/right of current
				// pos then we set a flag. When the side ends we reset the flag
				// and increment count.
				for l := range edge {
					if region[pos{x, y}] && !region[next[l]] {
						edge[l] = true
					} else if edge[l] {
						edge[l] = false
						count++
					}
				}
			}
		}
	}

	return count
}

// price function
func price(region map[pos]bool, part2 bool) int {
	area := len(region)
	perim := 0

	if part2 {
		return area * perimSides(region)
	}

	// for part one we just increment for every neighbor not in the region
	for p := range region {
		for _, n := range p.neighbors() {
			if !region[n] {
				perim++
			}
		}
	}

	return area * perim
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	garden := strings.Fields(string(data))

	regions := make([]map[pos]bool, 0)
	visited := make(map[pos]bool)

	for y, row := range garden {
		for x := range row {
			if !visited[pos{x, y}] {
				regions = append(regions, search(garden, pos{x, y}, visited))
			}
		}
	}

	sum1, sum2 := 0, 0
	for _, region := range regions {
		sum1 += price(region, false)
		sum2 += price(region, true)
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
