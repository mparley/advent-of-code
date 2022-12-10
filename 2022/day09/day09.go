package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func Abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func Unit(i int) int {
	if i >= 0 {
		return 1
	}
	return -1
}

func MoveHead(dir string, dist int, rope *[]Coord) []Coord {
	var tail []Coord
	for i := 0; i < dist; i++ {
		switch dir {
		case "R":
			(*rope)[0].x++
		case "L":
			(*rope)[0].x--
		case "D":
			(*rope)[0].y++
		case "U":
			(*rope)[0].y--
		}

		for j := 1; j < len(*rope); j++ {
			Update(&(*rope)[j-1], &(*rope)[j])
		}

		tail = append(tail, (*rope)[len(*rope)-1])
	}
	return tail
}

func Update(h *Coord, t *Coord) {
	xdif := h.x - t.x
	ydif := h.y - t.y

	if Abs(ydif) == 2 {
		if h.x != t.x {
			t.x += Unit(xdif)
		}
		t.y += Unit(ydif)
	} else if Abs(xdif) == 2 {
		if h.y != t.y {
			t.y += Unit(ydif)
		}
		t.x += Unit(xdif)
	}
}

func PrintRope(rope []Coord) {
	max := Coord{0, 0}
	min := Coord{0, 0}

	rmap := map[Coord]string{}
	for i := len(rope) - 1; i >= 0; i-- {
		rmap[rope[i]] = strconv.Itoa(i)
	}

	for k := range rmap {
		if k.x > max.x {
			max.x = k.x
		} else if k.x < min.x {
			min.x = k.x
		}

		if k.y > max.y {
			max.y = k.y
		} else if k.y < min.y {
			min.y = k.y
		}
	}

	fmt.Println()
	for y := min.y; y <= max.y; y++ {
		line := ""
		for x := min.x; x <= max.x; x++ {
			s, exists := rmap[Coord{x, y}]
			if x == 0 && y == 0 {
				line += "s"
			} else if exists {
				line += s
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

}

func PrintMap(tmap map[Coord]bool) {
	max := Coord{0, 0}
	min := Coord{0, 0}

	for k := range tmap {
		if k.x > max.x {
			max.x = k.x
		} else if k.x < min.x {
			min.x = k.x
		}

		if k.y > max.y {
			max.y = k.y
		} else if k.y < min.y {
			min.y = k.y
		}
	}

	for y := min.y; y <= max.y; y++ {
		line := ""
		for x := min.x; x <= max.x; x++ {
			if x == 0 && y == 0 {
				line += "s"
			} else if tmap[Coord{x, y}] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(data), "\n")
	rope_size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	rope := make([]Coord, rope_size)
	tailmap := map[Coord]bool{{0, 0}: true}

	for _, line := range inputs {
		if len(line) == 0 {
			continue
		}
		ins := strings.Split(line, " ")
		dist, _ := strconv.Atoi(ins[1])
		moves := MoveHead(ins[0], dist, &rope)
		for _, v := range moves {
			tailmap[v] = true
		}
		// PrintRope(rope)
	}

	// PrintMap(tailmap)
	fmt.Println(len(tailmap))
}
