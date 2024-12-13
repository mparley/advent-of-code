package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func checkX(puzzle []string, x, y int) int {
	width, height := len(puzzle[0]), len(puzzle)
	count := 0

	//right
	if width-x > 3 && puzzle[y][x+1:x+4] == "MAS" {
		count++
	}

	//left
	if x >= 3 && puzzle[y][x-3:x] == "SAM" {
		count++
	}

	//up
	if y >= 3 {
		if puzzle[y-1][x] == 'M' && puzzle[y-2][x] == 'A' && puzzle[y-3][x] == 'S' {
			count++
		}

		//up-right
		if width-x > 3 {
			if puzzle[y-1][x+1] == 'M' && puzzle[y-2][x+2] == 'A' && puzzle[y-3][x+3] == 'S' {
				count++
			}
		}

		//up-left
		if x >= 3 {
			if puzzle[y-1][x-1] == 'M' && puzzle[y-2][x-2] == 'A' && puzzle[y-3][x-3] == 'S' {
				count++
			}
		}
	}

	//down
	if height-y > 3 {
		if puzzle[y+1][x] == 'M' && puzzle[y+2][x] == 'A' && puzzle[y+3][x] == 'S' {
			count++
		}

		//down-right
		if width-x > 3 {
			if puzzle[y+1][x+1] == 'M' && puzzle[y+2][x+2] == 'A' && puzzle[y+3][x+3] == 'S' {
				count++
			}
		}

		//down-left
		if x >= 3 {
			if puzzle[y+1][x-1] == 'M' && puzzle[y+2][x-2] == 'A' && puzzle[y+3][x-3] == 'S' {
				count++
			}
		}
	}

	return count
}

func checkA(puzzle []string, x, y int) bool {
	width, height := len(puzzle[0]), len(puzzle)

	if y == 0 || y == height-1 || x == 0 || x == width-1 {
		return false
	}

	// \-diagonal
	back := (puzzle[y-1][x-1] == 'M' && puzzle[y+1][x+1] == 'S') ||
		(puzzle[y-1][x-1] == 'S' && puzzle[y+1][x+1] == 'M')

	// /-diagonal
	forward := (puzzle[y+1][x-1] == 'M' && puzzle[y-1][x+1] == 'S') ||
		(puzzle[y+1][x-1] == 'S' && puzzle[y-1][x+1] == 'M')

	return back && forward
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	sum1, sum2 := 0, 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if lines[y][x] == 'X' {
				sum1 += checkX(lines, x, y)
			}
			if lines[y][x] == 'A' && checkA(lines, x, y) {
				// fmt.Println("x,y:", x, y)
				sum2++
			}
		}
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
