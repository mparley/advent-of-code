package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type point struct {
	x, y, z int
}

type brick struct {
	a, b point
}

func absint(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func (b brick) lows() (int, int, int) {
	x, y, z := b.a.x, b.a.y, b.a.z
	if b.b.x < x {
		x = b.b.x
	}
	if b.b.y < y {
		y = b.b.y
	}
	if b.b.z < z {
		z = b.b.z
	}
	return x, y, z
}

func (b brick) dims() (int, int, int) {
	width, height, depth := b.a.x-b.b.x, b.a.y-b.b.y, b.a.z-b.b.z
	return absint(width), absint(height), absint(depth)
}

func (b *brick) fallTo(z int) {
	dif := b.a.z - z
	b.a.z -= dif
	b.b.z -= dif
}

func (b brick) collides(other brick) bool {
	x1, y1, _ := b.lows()
	x2, y2, _ := other.lows()
	w1, h1, _ := b.dims()
	w2, h2, _ := other.dims()

	return x1 <= x2+w2 && x2 <= x1+w1 && y1 <= y2+h2 && y2 <= y1+h1
}

func part2(supported, supports map[int][]int, cannot map[int]struct{}) int {
	sum := 0
	for b := range cannot {
		q := []int{b}
		gone := map[int]struct{}{}

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			gone[cur] = struct{}{}

			for _, sed := range supports[cur] {
				all := true
				for _, ss := range supported[sed] {
					if _, ok := gone[ss]; !ok {
						all = false
						break
					}
				}
				if all {
					sum++
					q = append(q, sed)
				}
			}
		}
	}

	return sum
}

func solve(bricks []brick) (int, int) {
	layers := [][]int{}
	layer := []int{}
	for i, cur := 0, bricks[0].a.z; i < len(bricks); i++ {
		if bricks[i].a.z > cur || i == len(bricks)-1 {
			layers = append(layers, layer)
			layer = []int{}
			cur = bricks[i].a.z
		}

		layer = append(layer, i)
	}
	layers = append(layers, layer)

	supported := map[int][]int{}
	supports := map[int][]int{}

	for i, l := range layers {
		for _, b := range l {
			bottomed := false

			if i == 0 {
				bricks[b].fallTo(1)
				supported[b] = []int{-1}
				continue
			}

			maxZ := 0
			s := []int{}
			for j := i - 1; j >= 0; j-- {
				for _, k := range layers[j] {
					if bricks[b].collides(bricks[k]) {
						bottomed = true
						if bricks[k].b.z > maxZ {
							maxZ = bricks[k].b.z
							s = []int{}
						}
						if bricks[k].b.z == maxZ {
							s = append(s, k)
						}
					}
				}
			}

			if bottomed {
				bricks[b].fallTo(maxZ + 1)
				supported[b] = s
				for _, t := range s {
					supports[t] = append(supports[t], b)
				}
			} else {
				bricks[b].fallTo(1)
				supported[b] = []int{-1}
			}
		}
	}

	//	for _, l := range layers {
	//		for _, b := range l {
	//			fmt.Println(bricks[b])
	//		}
	//	}

	cannot := map[int]struct{}{}
	for _, support := range supported {
		if len(support) == 1 && support[0] != -1 {
			cannot[support[0]] = struct{}{}
		}
	}

	p2 := part2(supported, supports, cannot)

	return (len(bricks) - len(cannot)), p2
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Fields(string(f))
	bricks := make([]brick, 0, len(input))
	for _, line := range input {
		x1, x2, y1, y2, z1, z2 := 0, 0, 0, 0, 0, 0
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &x1, &y1, &z1, &x2, &y2, &z2)
		p1, p2 := point{x1, y1, z1}, point{x2, y2, z2}
		if p1.z > p2.z {
			p1, p2 = p2, p1
		}
		bricks = append(bricks, brick{p1, p2})
	}
	slices.SortFunc(bricks, func(a, b brick) int {
		return cmp.Compare(a.a.z, b.a.z)
	})

	part1, part2 := solve(bricks)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
