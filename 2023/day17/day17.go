package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type coord struct {
	x, y int
}

type node struct {
	pos         coord
	dir         direction
	steps, dist int
}

type nodeQ []*node

func (nq nodeQ) Len() int { return len(nq) }

func (nq nodeQ) Less(i, j int) bool {
	return nq[i].dist < nq[j].dist
}

func (nq nodeQ) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
}

func (nq *nodeQ) Push(x any) {
	item := x.(*node)
	*nq = append(*nq, item)
}

func (nq *nodeQ) Pop() any {
	old := *nq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*nq = old[0 : n-1]
	return item
}

func inBounds(cur coord, w, h int) bool {
	return 0 <= cur.x && cur.x < w && 0 <= cur.y && cur.y < h
}

func nextPos(cur coord, dir direction) coord {
	add := []coord{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}[dir]
	return coord{cur.x + add.x, cur.y + add.y}
}

func neighbors(cur node, heatMap [][]int, lo, hi int) []node {
	w, h := len(heatMap[0]), len(heatMap)
	out := make([]node, 0, 3)
	for d := (4 + cur.dir - 1) % 4; d != (cur.dir+2)%4; d = (d + 1) % 4 {
		steps := 1
		if cur.steps > 0 {
			if d == cur.dir {
				if cur.steps >= hi {
					continue
				}
				steps += cur.steps
			} else if cur.steps < lo {
				continue
			}
		}
		nPos := nextPos(cur.pos, d)
		if inBounds(nPos, w, h) {
			nDist := cur.dist + heatMap[nPos.y][nPos.x]
			out = append(out, node{nPos, d, steps, nDist})
		}
	}
	return out
}

func heatLoss(heatMap [][]int, source, target coord, lo, hi int) int {
	srcNode := node{source, east, -1, 0}

	nQ := &nodeQ{&srcNode}
	heap.Init(nQ)
	seen := make(map[string]struct{})

	for len(*nQ) > 0 {
		u := heap.Pop(nQ).(*node)
		us := fmt.Sprint(u.pos, u.dir, u.steps)

		if _, ok := seen[us]; ok {
			continue
		}
		seen[us] = struct{}{}

		if u.pos.x == target.x && u.pos.y == target.y {
			return u.dist
		}

		nodes := neighbors(*u, heatMap, lo, hi)
		//fmt.Println(nodes)
		for i := range nodes {
			heap.Push(nQ, &nodes[i])
		}
	}

	return -1
}

func part1(heatMap [][]int) int {
	w, h := len(heatMap[0]), len(heatMap)
	return heatLoss(heatMap, coord{0, 0}, coord{w - 1, h - 1}, 0, 3)
}

func part2(heatMap [][]int) int {
	w, h := len(heatMap[0]), len(heatMap)
	return heatLoss(heatMap, coord{0, 0}, coord{w - 1, h - 1}, 4, 10)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Fields(string(f))
	heatMap := make([][]int, len(input))
	for y, row := range input {
		for _, col := range row {
			val, err := strconv.Atoi(string(col))
			if err != nil {
				log.Fatal(err)
			}
			heatMap[y] = append(heatMap[y], val)
		}
	}
	//fmt.Println(heatMap)

	fmt.Println("Part 1:", part1(heatMap))
	fmt.Println("Part 2:", part2(heatMap))
}
