package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type dir int

const (
	UP dir = iota
	RIGHT
	DOWN
	LEFT
)

type pos struct {
	x, y int
}

func connects(p rune, indir dir) bool {
	switch indir {
	case 0:
		return p == '|' || p == '7' || p == 'F'
	case 1:
		return p == '-' || p == '7' || p == 'J'
	case 2:
		return p == '|' || p == 'J' || p == 'L'
	case 3:
		return p == '-' || p == 'L' || p == 'F'
	}
	return false
}

func connected(a, b rune, by dir) bool {
	opo := (by + 2) % 4
	return (a == 'S' || connects(a, opo)) && connects(b, by)
}

func inBounds(p pos, w, h int) bool {
	return p.x >= 0 && p.x < w && p.y >= 0 && p.y < h
}

func neighbors(p pos) []pos {
	return []pos{{p.x, p.y - 1}, {p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y}}
}

func replaceS(a, b dir) rune {
	if (a+2)%4 == b {
		if a == 1 || a == 3 {
			return '-'
		} else {
			return '|'
		}
	}

	if (a+1)%4 != b {
		a, b = b, a

	}

	switch a {
	case 0:
		return 'F'
	case 1:
		return '7'
	case 2:
		return 'J'
	case 3:
		return 'L'
	}

	return 'S'
}

func nextDir(cur rune, indir dir) dir {
	switch cur {
	case '|':
		if indir == UP {
			return UP
		} else {
			return DOWN
		}
	case '-':
		if indir == RIGHT {
			return RIGHT
		} else {
			return LEFT
		}
	case 'L':
		if indir == LEFT {
			return UP
		} else {
			return RIGHT
		}
	case 'F':
		if indir == UP {
			return RIGHT
		} else {
			return DOWN
		}
	case 'J':
		if indir == RIGHT {
			return UP
		} else {
			return LEFT
		}
	case '7':
		if indir == RIGHT {
			return DOWN
		} else {
			return LEFT
		}
	default:
		return UP
	}
}

func trace(start pos, pipes [][]rune) map[pos]struct{} {
	cur := start
	ch := pipes[start.y][start.x]
	path := make(map[pos]struct{})
	path[cur] = struct{}{}
	var next pos
	var d, dstart dir

	for i, n := range neighbors(start) {
		if !inBounds(n, len(pipes[0]), len(pipes)) {
			continue
		}
		ch2 := pipes[n.y][n.x]
		if connected(ch, ch2, dir(i)) {
			next = n
			d = dir(i)
			dstart = d
			break
		}
	}

	for next != start {
		cur = next
		path[cur] = struct{}{}
		ch = pipes[cur.y][cur.x]
		d = nextDir(ch, d)
		switch d {
		case UP:
			next = pos{cur.x, cur.y - 1}
		case RIGHT:
			next = pos{cur.x + 1, cur.y}
		case DOWN:
			next = pos{cur.x, cur.y + 1}
		case LEFT:
			next = pos{cur.x - 1, cur.y}
		}
	}

	pipes[start.y][start.x] = replaceS(dstart, d)
	return path
}

func printLoop(input []string, loop map[pos]struct{}) {
	for y, line := range input {
		l := ""
		for x, ch := range line {
			if _, ok := loop[pos{x, y}]; ok {
				l += string(ch)
			} else {
				l += "."
			}
		}
		fmt.Println(y, l)
	}
}

func inner(loop map[pos]struct{}, input [][]rune) int {
	inner := 0
	for y, in := range input {
		for x := range in {
			if _, ok := loop[pos{x, y}]; ok {
				continue
			}

			intersects := 0
			opening := '.'
			for i := x + 1; i < len(in); i++ {
				if _, ok := loop[pos{i, y}]; ok {
					ch := input[y][i]
					switch ch {
					case '|':
						intersects++
					case 'F', 'L':
						if opening == '.' {
							opening = ch
						}
					case 'J':
						if opening == 'F' {
							intersects++
						}
						opening = '.'
					case '7':
						if opening == 'L' {
							intersects++
						}
						opening = '.'
					}
				}
			}

			if intersects%2 != 0 {
				inner++
			}
		}
	}

	return inner
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	start := pos{-1, -1}
	var runes [][]rune
	for y, in := range input {
		r := []rune(in)
		if start.x == -1 {
			for x, i := range in {
				if i == 'S' {
					start = pos{x, y}
				}
			}
		}
		runes = append(runes, r)
	}

	loop := trace(start, runes)

	if len(os.Args) == 2 || os.Args[2] == "1" {
		fmt.Println(len(loop) / 2)
	}
	if len(os.Args) == 2 || os.Args[2] == "2" {
		fmt.Println(inner(loop, runes))
	}

}
