package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

type walker struct {
	pos   position
	steps int
}

type walkQ []walker

func addPos(poss ...position) position {
	out := position{0, 0}
	for _, p := range poss {
		out = position{out.x + p.x, out.y + p.y}
	}
	return out
}

func (wq *walkQ) pop() walker {
	out := (*wq)[0]
	*wq = (*wq)[1:]
	return out
}

func (wq *walkQ) push(w walker) {
	*wq = append(*wq, w)
}

func (w walker) walk() []walker {
	out := make([]walker, 0, 4)
	for _, d := range []position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		w := walker{addPos(w.pos, d), w.steps - 1}
		out = append(out, w)
	}
	return out
}

func inBounds(p position, w, h int) bool {
	return 0 <= p.x && p.x < w && 0 <= p.y && p.y < h
}

func printGarden(input []string, landings map[position]struct{}) {
	for y, row := range input {
		for x, col := range row {
			s := string(col)
			if _, ok := landings[position{x, y}]; ok {
				s = "O"
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
}

func rock(input []string, pos position) bool {
	y := (len(input) + pos.y) % len(input)
	x := (len(input[0]) + pos.x) % len(input[0])
	return input[y][x] == '#'
}

func bfs(input []string, start position, steps int) int {
	w, h := len(input[0]), len(input)
	count := 0

	wq := make(walkQ, 0, 100)
	wq.push(walker{start, steps})
	visited := make(map[position]struct{})

	for len(wq) > 0 {
		cur := wq.pop()
		if _, ok := visited[cur.pos]; ok {
			continue
		}

		visited[cur.pos] = struct{}{}
		if cur.steps%2 == 0 {
			count++
		}

		if cur.steps == 0 {
			continue
		}

		for _, wk := range cur.walk() {
			if !inBounds(wk.pos, w, h) || rock(input, wk.pos) {
				continue
			}
			wq.push(wk)
		}
	}

	return count
}

func part2(input []string, start position) int {
	steps := 26501365
	length := len(input)

	// on for same even/odd-ness as starting fold, off for opposite
	// started out doing even/odd variables but it was confusing
	on, off := 1, 0

	onCount := bfs(input, start, steps)
	offCount := bfs(input, start, steps+1)

	folds := steps / length

	for i := 1; i <= folds-1; i++ {
		if i%2 == 0 {
			on += 4 * i
		} else {
			off += 4 * i
		}
	}

	//fmt.Println(on, off)
	total := on*onCount + off*offCount

	// striaght ends
	// s = steps - (folds * length) + length - 1
	//   = steps - (steps / length * length) + length - 1
	//   = steps - (steps) + length - 1
	// length is added to bring from center to end
	// subtract 1 to get into last fold from end of prev
	s := length - 1
	total += bfs(input, position{length / 2, 0}, s)
	total += bfs(input, position{0, length / 2}, s)
	total += bfs(input, position{length / 2, length - 1}, s)
	total += bfs(input, position{length - 1, length / 2}, s)

	// big corners
	// there are big and little corners
	// s = length - 1 + length - length / 2
	// (+ length) - going down a block so add steps back in to bfs
	// (- length/2) - going left/right by half a block so subtract steps
	// bfs * by (folds - 1) because this corner is on the point ring
	s = (length*3)/2 - 1
	total += (folds - 1) * bfs(input, position{0, 0}, s)
	total += (folds - 1) * bfs(input, position{length - 1, 0}, s)
	total += (folds - 1) * bfs(input, position{0, length - 1}, s)
	total += (folds - 1) * bfs(input, position{length - 1, length - 1}, s)

	// small corners
	//s = ((length*3)/2 - 1) - length
	// (- length) going left/right by one block so subtract steps
	// bfs * by folds because this corner slightly cuts into border ring of diamond
	// one step beyond the point ring
	s = length/2 - 1
	total += folds * bfs(input, position{0, 0}, s)
	total += folds * bfs(input, position{length - 1, 0}, s)
	total += folds * bfs(input, position{0, length - 1}, s)
	total += folds * bfs(input, position{length - 1, length - 1}, s)

	return total
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Fields(string(f))

	steps := 64
	if len(os.Args) > 2 {
		val, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		steps = val
	}

	var start position
	rocks := make(map[position]struct{})
	for y, line := range input {
		for x, col := range line {
			if col == 'S' {
				start = position{x, y}
			} else if col == '#' {
				rocks[position{x, y}] = struct{}{}
			}
		}
	}

	fmt.Println("Part 1:", bfs(input, start, steps))
	fmt.Println("Part 2:", part2(input, start))
}
