package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Coord struct {
	x, y int
}

func Push(q *[]Coord, c Coord) {
	*q = append(*q, c)
}

func Pop(q *[]Coord) Coord {
	c := (*q)[0]
	*q = (*q)[1:]
	return c
}

func Neighbors(hmap *[][]int, c Coord) *[]Coord {
	width := len((*hmap)[0])
	height := len(*hmap)
	ns := []Coord{}

	if c.x-1 >= 0 {
		if (*hmap)[c.y][c.x-1] <= (*hmap)[c.y][c.x]+1 {
			ns = append(ns, Coord{c.x - 1, c.y})
		}
	}
	if c.x+1 < width {
		if (*hmap)[c.y][c.x+1] <= (*hmap)[c.y][c.x]+1 {
			ns = append(ns, Coord{c.x + 1, c.y})
		}
	}
	if c.y-1 >= 0 {
		if (*hmap)[c.y-1][c.x] <= (*hmap)[c.y][c.x]+1 {
			ns = append(ns, Coord{c.x, c.y - 1})
		}
	}
	if c.y+1 < height {
		if (*hmap)[c.y+1][c.x] <= (*hmap)[c.y][c.x]+1 {
			ns = append(ns, Coord{c.x, c.y + 1})
		}
	}

	return &ns
}

func BfsPath(hmap *[][]int, s Coord, e Coord) *[]Coord {
	prev := map[Coord]Coord{}

	q := []Coord{s}

	for len(q) != 0 {
		current := Pop(&q)
		if current == e {
			break
		}

		neighbors := Neighbors(hmap, current)
		for _, n := range *neighbors {
			_, exists := prev[n]
			if exists {
				continue
			}
			prev[n] = current
			Push(&q, n)
		}
	}

	_, exists := prev[e]
	if !exists {
		return nil
	}

	seq := []Coord{}
	cur := e
	for cur != s {
		seq = append(seq, cur)
		cur = prev[cur]
	}
	seq = append(seq, s)

	return &seq
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	CheckErr(err)

	starts := make([]Coord, 1)
	var target Coord

	heightmap := [][]int{}

	for i, line := range strings.Fields(string(data)) {
		heightmap = append(heightmap, make([]int, len(line)))
		for j, char := range line {
			if char == 'S' {
				starts[0] = Coord{j, i}
				char = 'a'
			} else if char == 'E' {
				target = Coord{j, i}
				char = 'z'
			} else if char == 'a' {
				starts = append(starts, Coord{j, i})
			}
			heightmap[i][j] = int(char - 'a')
		}
	}

	path := BfsPath(&heightmap, starts[0], target)
	if path != nil {
		fmt.Println(len(*path) - 1)
	}

	lengths := []int{len(*path)}
	for _, start := range starts[1:] {
		path = BfsPath(&heightmap, start, target)
		if path != nil {
			lengths = append(lengths, len(*path))
		}
	}
	sort.Ints(lengths)
	fmt.Println(lengths[0] - 1)
}
