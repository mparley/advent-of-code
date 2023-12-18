package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

type instruction struct {
	dir    direction
	meters int
}

type coord struct {
	x, y int
}

func runeDir(r rune) direction {
	return map[rune]direction{'U': up, 'R': right, 'D': down, 'L': left}[r]
}

func delta(d direction) coord {
	return []coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}[d]
}

func addCoord(l, r coord) coord {
	return coord{l.x + r.x, l.y + r.y}
}

func magnitude(pos coord, m int) coord {
	return coord{pos.x * m, pos.y * m}
}

func updateLimits(pos coord, lo, hi *coord) {
	if pos.x > hi.x {
		hi.x = pos.x
	}
	if pos.y > hi.y {
		hi.y = pos.y
	}
	if pos.x < lo.x {
		lo.x = pos.x
	}
	if pos.y < lo.y {
		lo.y = pos.y
	}
}

func printDig(digMap map[coord]struct{}) {
	lo, hi := coord{math.MaxInt, math.MaxInt}, coord{0, 0}
	for pos := range digMap {
		updateLimits(pos, &lo, &hi)
	}

	for y := lo.y; y <= hi.y; y++ {
		for x := lo.x; x <= hi.x; x++ {
			if _, ok := digMap[coord{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func qPush(q *[]coord, pos coord) {
	*q = append(*q, pos)
}

func qPop(q *[]coord) coord {
	pos := (*q)[0]
	*q = (*q)[1:]
	return pos
}

func fillMap(digMap *map[coord]struct{}, start coord) {
	digQ := make([]coord, 0, len(*digMap))
	qPush(&digQ, start)

	for len(digQ) > 0 {
		pos := qPop(&digQ)
		if _, ok := (*digMap)[pos]; ok {
			continue
		}
		(*digMap)[pos] = struct{}{}
		for d := up; d <= left; d++ {
			qPush(&digQ, addCoord(pos, delta(d)))
		}
	}
}

func dig(ins []instruction) map[coord]struct{} {
	pos := coord{0, 0}
	digMap := make(map[coord]struct{})
	digMap[pos] = struct{}{}

	lo, hi := coord{math.MaxInt, math.MaxInt}, coord{0, 0}
	for _, in := range ins {
		for i := 0; i < in.meters; i++ {
			pos = addCoord(pos, delta(in.dir))
			digMap[pos] = struct{}{}
			updateLimits(pos, &lo, &hi)
		}
	}

	start := coord{0, 0}
	for x := lo.x; x <= hi.x; x++ {
		if _, ok := digMap[coord{x, lo.y}]; ok {
			if _, ok := digMap[coord{x, lo.y + 1}]; !ok {
				start = coord{x, lo.y + 1}
				break
			}
		}
	}

	fillMap(&digMap, start)
	return digMap
}

func printVertices(vertices []coord) {
	digMap := make(map[coord]struct{})
	for i := 0; i < len(vertices)-1; i++ {
		cur := vertices[i]
		next := vertices[i+1]
		if cur.x == next.x {
			if cur.y > next.y {
				cur, next = next, cur
			}
			for j := cur.y; j <= next.y; j++ {
				digMap[coord{cur.x, j}] = struct{}{}
			}
		} else if cur.y == next.y {
			if cur.x > next.x {
				cur, next = next, cur
			}
			for j := cur.x; j <= next.x; j++ {
				digMap[coord{j, cur.y}] = struct{}{}
			}
		}
	}
	printDig(digMap)
}

func bigDig(ins []instruction, print bool) int {
	vertices := make([]coord, 0, len(ins))
	vertices = append(vertices, coord{100, 100})

	edgeSum := 0
	for _, in := range ins {
		cur := vertices[len(vertices)-1]
		edgeSum += in.meters
		d := magnitude(delta(in.dir), in.meters)
		vertices = append(vertices, addCoord(cur, d))
	}

	if print {
		printVertices(vertices)
	}

	sum1, sum2 := 0, 0
	for i := 0; i < len(vertices)-1; i++ {
		sum1 += (vertices[i].x * vertices[i+1].y)
		sum2 += (vertices[i].y * vertices[i+1].x)
	}
	sum1 += (vertices[len(vertices)-1].x * vertices[0].y)
	sum2 += (vertices[0].x * vertices[len(vertices)-1].y)

	area := sum1 - sum2
	if area < 0 {
		area *= -1
	}

	return (area / 2) + (edgeSum / 2) + 1
}

func part1(input []string, dumb bool) int {
	instructions := make([]instruction, 0, len(input))
	for _, line := range input {
		var ch rune
		var m, r, g, b int
		fmt.Sscanf(line, "%c %d (#%02x%02x%02x)", &ch, &m, &r, &g, &b)
		ins := instruction{runeDir(ch), m}
		instructions = append(instructions, ins)
	}

	if !dumb {
		return bigDig(instructions, false)
	}

	digMap := dig(instructions)
	printDig(digMap)
	return len(digMap)
}

func part2(input []string) int {
	instructions := make([]instruction, 0, len(input))
	for _, line := range input {
		var m, d int
		fmt.Sscanf(strings.Fields(line)[2], "(#%05x%01x)", &m, &d)
		ins := instruction{direction((d + 1) % 4), m}
		instructions = append(instructions, ins)
	}

	return bigDig(instructions, false)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	fmt.Println("Part 1:", part1(input, false))
	fmt.Println("Part 2:", part2(input))
}
