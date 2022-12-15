package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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

func ParseCoord(s string) (c Coord) {
	var err error
	in := strings.Split(s, ",")
	c.x, err = strconv.Atoi(in[0])
	CheckErr(err)
	c.y, err = strconv.Atoi(in[1])
	CheckErr(err)
	return
}

func SimSand(cmap *map[Coord]string, spawn Coord, floor int) bool {
	grain := spawn
	for grain.y < floor {
		n := grain
		n.y++
		_, exists := (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		n.x--
		_, exists = (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		n.x += 2
		_, exists = (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		(*cmap)[grain] = "o"
		return true
	}

	return false
}

func SimSand2(cmap *map[Coord]string, spawn Coord, floor int) bool {
	grain := spawn
	for grain.y < floor-1 {
		n := grain
		n.y++
		_, exists := (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		n.x--
		_, exists = (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		n.x += 2
		_, exists = (*cmap)[n]
		if !exists {
			grain = n
			continue
		}

		(*cmap)[grain] = "o"
		if grain != spawn {
			return true
		} else {
			return false
		}
	}

	(*cmap)[grain] = "o"
	return true
}

func PrintMap(cmap *map[Coord]string, floor int) {
	lo := Coord{math.MaxInt, 0}
	hi := Coord{0, floor}
	for k := range *cmap {
		if lo.x > k.x {
			lo.x = k.x
		}
		if hi.x < k.x {
			hi.x = k.x
		}
	}

	for y := lo.y; y <= hi.y; y++ {
		line := ""
		for x := lo.x; x <= hi.x; x++ {
			if y == floor {
				line += "#"
				continue
			}
			v, exists := (*cmap)[Coord{x, y}]
			if !exists {
				line += "."
			} else {
				line += v
			}
		}
		fmt.Println(line)
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	CheckErr(err)
	rocklines := strings.Fields(string(data))

	sand_spawn := Coord{500, 0}
	collisions := map[Coord]string{sand_spawn: "+"}
	low := 0

	for i, v := range rocklines {
		if v == "->" {
			l := ParseCoord(rocklines[i-1])
			r := ParseCoord(rocklines[i+1])

			if l.y > low {
				low = l.y
			}
			if r.y > low {
				low = r.y
			}

			if l.x > r.x || l.y > r.y {
				l, r = r, l
			}

			for y := l.y; y <= r.y; y++ {
				for x := l.x; x <= r.x; x++ {
					collisions[Coord{x, y}] = "#"
				}
			}
		}
	}

	// PrintMap(&collisions, low)

	r := 0
	for SimSand(&collisions, sand_spawn, low) {
		r++
	}

	// PrintMap(&collisions, low)
	fmt.Println("Part 1:", r)

	r++
	for SimSand2(&collisions, sand_spawn, low+2) {
		r++
	}

	// PrintMap(&collisions, low+2)
	fmt.Println("Part 2:", r)
}
