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

func (p pos) neighbors() []pos {
	return []pos{{p.x, p.y - 1}, {p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}}
}

func (p pos) inside(w, h int) bool {
	return p.x >= 0 && p.y >= 0 && p.x < w && p.y < h
}

func (p pos) val(topo [][]int) int {
	return topo[p.y][p.x]
}

func parseInput(lines []string) ([][]int, []pos) {
	topo := make([][]int, 0, len(lines))
	starts := make([]pos, 0)
	for y, line := range lines {
		row := make([]int, 0, len(lines[0]))
		for x, c := range line {
			switch c {
			case '.':
				row = append(row, -1)
			case '0':
				row = append(row, 0)
				starts = append(starts, pos{x, y})
			default:
				row = append(row, int(c-'0'))
			}
		}
		topo = append(topo, row)
	}
	return topo, starts
}

// A basic dfs
// Don't need to keep track of visited positions because the trail always
// increments by 1. Keeps track of steps to track the longest path in the
// ends map. The count is for every path that hits a 9
func search(topo [][]int, p pos, steps, count int, ends map[pos]int) int {
	steps++

	if p.val(topo) == 9 {
		if ends[p] < steps {
			ends[p] = steps
		}
		return count + 1
	}

	w, h := len(topo[0]), len(topo)
	currCount := count

	for _, n := range p.neighbors() {
		if n.inside(w, h) && n.val(topo)-p.val(topo) == 1 {
			currCount += search(topo, n, steps, count, ends)
		}
	}

	return currCount
}

func score(topo [][]int, trailhead pos) (int, int) {
	trailends := make(map[pos]int)
	rating := search(topo, trailhead, 0, 0, trailends)
	return len(trailends), rating
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	topo, trailheads := parseInput(strings.Fields(string(data)))
	sum1, sum2 := 0, 0

	for _, trailhead := range trailheads {
		score, rating := score(topo, trailhead)
		sum1 += score
		sum2 += rating
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)

}
