package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type dir int

const (
	north dir = iota
	east
	south
	west
)

type pos struct {
	x, y int
}

func (l *pos) add(r pos) {
	(*l).x += r.x
	(*l).y += r.y
}

type beam struct {
	p pos
	d dir
}

type beamQ []beam

func (q *beamQ) push(b beam) {
	*q = append(*q, b)
}

func (q *beamQ) pop() beam {
	b := (*q)[0]
	*q = (*q)[1:]
	return b
}

func inBounds(p pos, w, h int) bool {
	return 0 <= p.x && p.x < w && 0 <= p.y && p.y < h
}

func getDir(d dir) pos {
	return []pos{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}[d]
}

func backslash(d dir) dir {
	return []dir{west, south, east, north}[d]
}

func forwardslash(d dir) dir {
	return []dir{east, north, west, south}[d]
}

func getNextBeams(curr beam, ch rune) []beam {
	if (ch == '-' && curr.d%2 == 0) || (ch == '|' && curr.d%2 == 1) {
		n2 := curr
		curr.d = (curr.d + 1) % 4
		n2.d = (4 + n2.d - 1) % 4
		curr.p.add(getDir(curr.d))
		n2.p.add(getDir(n2.d))
		return []beam{curr, n2}
	}
	if ch == '\\' {
		curr.d = backslash(curr.d)
	} else if ch == '/' {
		curr.d = forwardslash(curr.d)
	}
	curr.p.add(getDir(curr.d))
	return []beam{curr}
}

func energize(room [][]rune, start pos, sdir dir) map[beam]struct{} {
	lightMap := make(map[beam]struct{})
	q := make(beamQ, 0, 10)
	q.push(beam{start, sdir})

	for len(q) > 0 {
		b := q.pop()

		if !inBounds(b.p, len(room[0]), len(room)) {
			continue
		}
		if _, ok := lightMap[b]; ok {
			continue
		}

		lightMap[b] = struct{}{}
		for _, nextBeam := range getNextBeams(b, room[b.p.y][b.p.x]) {
			q.push(nextBeam)
		}
	}

	return lightMap
}

func dirChar(d dir) rune {
	return []rune{'^', '>', 'v', '<'}[d]
}

func collapseIntersects(lightMap map[beam]struct{}) map[pos]dir {
	ret := make(map[pos]dir)
	for b := range lightMap {
		ret[b.p] = b.d
	}
	return ret
}

func printLights(room [][]rune, lightMap map[pos]dir) {
	for y, row := range room {
		for x, ch := range row {
			if _, ok := lightMap[pos{x, y}]; ok && ch == '.' {
				fmt.Print(string(dirChar(lightMap[pos{x, y}])))
			} else {
				fmt.Print(string(ch))
			}
		}
		fmt.Println()
	}
}

func run(room [][]rune, start pos, sdir dir) int {
	lights := collapseIntersects(energize(room, start, sdir))
	return len(lights)
}

func part1(room [][]rune) int {
	return run(room, pos{0, 0}, east)
}

func part2(room [][]rune) int {
	maxEnergy := 0
	for x := range room[0] {
		energy := run(room, pos{x, 0}, south)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	for x := range room[len(room)-1] {
		energy := run(room, pos{x, len(room) - 1}, north)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	for y := range room {
		energy := run(room, pos{0, y}, east)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	for y := range room {
		energy := run(room, pos{len(room[0]) - 1, y}, west)
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}

	return maxEnergy
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	room := make([][]rune, len(input))
	for y, line := range input {
		room[y] = []rune(line)
	}

	fmt.Println(part1(room))
	fmt.Println(part2(room))
}
