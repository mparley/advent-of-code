package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type v3 struct {
	x, y, z float64
}

type hail struct {
	pos, vel v3
}

type indexPair struct {
	i1, i2 int
}

func newIndexPair(i1, i2 int) indexPair {
	if i1 > i2 {
		return indexPair{i2, i1}
	}
	return indexPair{i1, i2}
}

func pathCross(h1, h2 hail, from, to float64) (bool, v3) {
	//m1, m2 := a.vel.y/a.vel.x, b.vel.y/b.vel.x
	a := (h1.pos.y + h1.vel.y - h1.pos.y) / (h1.pos.x + h1.vel.x - h1.pos.x)
	b := (h2.pos.y + h2.vel.y - h2.pos.y) / (h2.pos.x + h2.vel.x - h2.pos.x)

	if a == b {
		return false, v3{0, 0, 0}
	}

	c := h1.pos.y - a*h1.pos.x
	d := h2.pos.y - b*h2.pos.x

	ix := (d - c) / (a - b)
	iy := a*(d-c)/(a-b) + c

	if ((ix-h1.pos.x < 0) != (h1.vel.x < 0)) || //
		((iy-h1.pos.y < 0) != (h1.vel.y < 0)) || //
		((ix-h2.pos.x < 0) != (h2.vel.x < 0)) || //
		((iy-h2.pos.y < 0) != (h2.vel.y < 0)) {

		return false, v3{0, 0, 0}
	}

	return (from <= ix && ix <= to && from <= iy && iy <= to), v3{ix, iy, 0}
}

func countIntersections(hailstones []hail, from, to float64, print bool) int {
	checked := map[indexPair]struct{}{}
	count := 0
	for i := range hailstones {
		for j := range hailstones {
			ipair := newIndexPair(i, j)
			if _, ok := checked[ipair]; ok {
				continue
			}

			if crossed, at := pathCross(hailstones[i], hailstones[j], from, to); crossed {
				count++
				if print {
					fmt.Println(hailstones[i], hailstones[j], "crossed @", at)
				}
			}

			checked[ipair] = struct{}{}
		}
	}
	return count
}

func z3SolveAndFuckMath(hailstones []hail) {
	fmt.Println("This is where I tried to import the z3 go package and get it working")
	fmt.Println("This is where I failed to do that")
	fmt.Println("So just take these equations and throw them into some math tool to be solved")
	r := rand.Intn(len(hailstones) - 3)
	for i := 0; i < 3; i++ {
		h := hailstones[r+i]
		fmt.Println("x +  t", i, " * vx = ", int(h.pos.x), " +  t", i, " * ", int(h.vel.x))
		fmt.Println("y +  t", i, " * vy = ", int(h.pos.y), " +  t", i, " * ", int(h.vel.y))
		fmt.Println("z +  t", i, " * vz = ", int(h.pos.z), " +  t", i, " * ", int(h.vel.z))
		fmt.Println()
	}
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	hailstones := []hail{}

	for _, line := range input {
		var x, y, z, vx, vy, vz float64 = 0, 0, 0, 0, 0, 0
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f", &x, &y, &z, &vx, &vy, &vz)
		hailstones = append(hailstones, hail{v3{x, y, z}, v3{vx, vy, vz}})
	}

	fmt.Println("Part 1:", countIntersections(hailstones, 200000000000000, 400000000000000, false))
	fmt.Println("Part 2:")
	z3SolveAndFuckMath(hailstones)

}
