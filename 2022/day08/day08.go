package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Coord struct {
	x, y int
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	treemap := [][]int{}

	for scanner.Scan() {
		treeline := []int{}
		for _, v := range scanner.Text() {
			w, _ := strconv.Atoi(string(v))
			treeline = append(treeline, w)
		}
		treemap = append(treemap, treeline)
	}

	// for _, v := range treemap {
	// 	fmt.Println(v)
	// }

	visible := map[Coord][]int{}
	width := len(treemap[0])
	height := len(treemap)

	for y, treeline := range treemap {
		highest := -1
		for x, tree := range treeline {
			if tree > highest {
				highest = tree
				visible[Coord{x, y}] = []int{x, -1, -1, -1}
			}
		}

		highest = -1
		for x := width - 1; x >= 0; x-- {
			if treeline[x] > highest {
				highest = treeline[x]
				_, exists := visible[Coord{x, y}]
				if !exists {
					visible[Coord{x, y}] = []int{-1, -1, width - x - 1, -1}
				} else {
					visible[Coord{x, y}][2] = width - x - 1
				}
			}
		}
	}

	for x := 0; x < width; x++ {
		highest := -1
		for y := 0; y < height; y++ {
			if treemap[y][x] > highest {
				highest = treemap[y][x]
				_, exists := visible[Coord{x, y}]
				if !exists {
					visible[Coord{x, y}] = []int{-1, y, -1, -1}
				} else {
					visible[Coord{x, y}][1] = y
				}
			}
		}

		highest = -1
		for y := height - 1; y >= 0; y-- {
			if treemap[y][x] > highest {
				highest = treemap[y][x]
				_, exists := visible[Coord{x, y}]
				if !exists {
					visible[Coord{x, y}] = []int{-1, -1, -1, height - y - 1}
				} else {
					visible[Coord{x, y}][3] = height - y - 1
				}
			}
		}
	}

	fmt.Println("Number of visible trees:", len(visible))

	high_score := 0
	for k, v := range visible {
		if v[0] == -1 {
			v[0] = k.x
			for x := k.x - 1; x >= 0; x-- {
				if treemap[k.y][x] >= treemap[k.y][k.x] {
					v[0] = k.x - x
					break
				}
			}
		}

		if v[1] == -1 {
			v[1] = k.y
			for y := k.y - 1; y >= 0; y-- {
				if treemap[y][k.x] >= treemap[k.y][k.x] {
					v[1] = k.y - y
					break
				}
			}
		}

		if v[2] == -1 {
			v[2] = width - 1 - k.x
			for x := k.x + 1; x < width; x++ {
				if treemap[k.y][x] >= treemap[k.y][k.x] {
					v[2] = x - k.x
					break
				}
			}
		}

		if v[3] == -1 {
			v[3] = height - 1 - k.y
			for y := k.y + 1; y < height; y++ {
				if treemap[y][k.x] >= treemap[k.y][k.x] {
					v[3] = y - k.y
					break
				}
			}
		}

		score := v[0] * v[1] * v[2] * v[3]
		// fmt.Println(k, "scores:", v, "score:", score)
		if score > high_score {
			high_score = score
		}
	}

	fmt.Println("Highest scenic score:", high_score)
}
