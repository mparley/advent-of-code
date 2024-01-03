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

func neighbors(p point, ch rune) []point {
	ns := []point{{p.x, p.y - 1}, {p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}}
	switch ch {
	case '^':
		return []point{ns[0]}
	case '>':
		return []point{ns[1]}
	case 'v':
		return []point{ns[2]}
	case '<':
		return []point{ns[3]}
	default:
		return ns
	}
}

func inBounds(p point, w, h int) bool {
	return 0 <= p.x && p.x < w && 0 <= p.y && p.y < h
}

func dfs(input []string, cur, end point, count int, //
	visited [][]bool, part2 bool) int {

	if cur == end {
		return count
	}

	visited[cur.y][cur.x] = true

	ch := rune(input[cur.y][cur.x])
	if part2 {
		ch = '.'
	}

	maxCount := 0
	for _, n := range neighbors(cur, ch) {
		if !inBounds(n, len(input[0]), len(input)) {
			continue
		}
		if input[n.y][n.x] == '#' {
			continue
		}
		if visited[n.y][n.x] == true {
			continue
		}

		ret := dfs(input, n, end, count+1, visited, part2)
		if ret > maxCount {
			maxCount = ret
		}
	}
	visited[cur.y][cur.x] = false

	return maxCount
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Fields(string(f))
	start := point{strings.Index(input[0], "."), 0}
	end := point{strings.Index(input[len(input)-1], "."), len(input) - 1}

	visited := make([][]bool, len(input))
	for i := range input {
		visited[i] = make([]bool, len(input[0]))
	}

	fmt.Println("Part 1:", dfs(input, start, end, 0, visited, false))
	fmt.Println("Part 2:", dfs(input, start, end, 0, visited, true))
}
