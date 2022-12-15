package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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

func ParseCoord(xs, ys string) (c Coord) {
	var err error
	c.x, err = strconv.Atoi(xs)
	CheckErr(err)
	c.y, err = strconv.Atoi(ys)
	CheckErr(err)
	return
}

func IAbs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func ManDist(l, r Coord) int {
	return IAbs(l.x-r.x) + IAbs(l.y-r.y)
}

func FindOpen(sbm *map[Coord]Coord, bound int) Coord {
	for y := 0; y <= bound; y++ {
		if x := FindGapX(sbm, y, bound); x < bound {
			return Coord{x, y}
		}
	}
	return Coord{-1, -1}
}

func FindGapX(sbm *map[Coord]Coord, row int, bound int) int {
	bounds := []Coord{}
	for sens, beac := range *sbm {
		d := ManDist(sens, beac)
		if sens.y+d >= row && sens.y-d <= row {
			r := d - IAbs(row-sens.y)
			bounds = append(bounds, Coord{sens.x - r, sens.x + r})
		}
	}

	sort.Slice(bounds, func(l, r int) bool {
		return bounds[l].x < bounds[r].x
	})

	i := 0
	for _, v := range bounds {
		if i >= v.x && i <= v.y {
			i = v.y + 1
		}
	}
	return i

}

func RuleOutInRow(sbm *map[Coord]Coord, row int) int {
	nb := map[Coord]bool{}
	beacs := []Coord{}
	for sensor, beacon := range *sbm {
		d := ManDist(sensor, beacon)

		if beacon.y == row {
			beacs = append(beacs, beacon)
		}

		if sensor.y+d < row {
			continue
		}

		for x := sensor.x; ManDist(sensor, Coord{x, row}) <= d; x++ {
			nb[Coord{x, row}] = true
		}

		for x := sensor.x; ManDist(sensor, Coord{x, row}) <= d; x-- {
			nb[Coord{x, row}] = true
		}
	}

	for _, b := range beacs {
		if nb[b] {
			delete(nb, b)
		}
	}
	return len(nb)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	CheckErr(err)

	row, err := strconv.Atoi(os.Args[2])
	CheckErr(err)

	bound, err := strconv.Atoi(os.Args[3])
	CheckErr(err)

	closestBeacon := map[Coord]Coord{}

	lines := strings.Split(string(data), "\n")
	rx := regexp.MustCompile(`-?\d+`)

	for _, line := range lines {
		coords := rx.FindAllString(line, -1)
		if len(coords) == 4 {
			sen := ParseCoord(coords[0], coords[1])
			bea := ParseCoord(coords[2], coords[3])
			closestBeacon[sen] = bea
		}
	}

	part1 := RuleOutInRow(&closestBeacon, row)
	fmt.Println("Part 1:", part1)

	gap := FindOpen(&closestBeacon, bound)
	fmt.Println("Part 2:", (gap.x*4000000)+gap.y)

}
