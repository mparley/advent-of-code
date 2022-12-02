package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	CheckErr(err)
	defer f.Close()

	total_score := 0
	total_score2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 3 {
			continue
		}

		var p1 int = int(line[0] - 'A')
		var p2 int = int(line[2] - 'X')
		var p3 int

		// Part 1
		switch p2 - p1 {
		case 1:
			total_score += 6 + p2 + 1
		case -2:
			total_score += 6 + p2 + 1
		case 0:
			total_score += 3 + p2 + 1
		default:
			total_score += p2 + 1
		}

		// Part 2
		switch p2 {
		case 0:
			p3 = (3 + p1 - 1) % 3
		case 1:
			p3 = p1
		default:
			p3 = (p1 + 1) % 3
		}
		total_score2 += (p2 * 3) + p3 + 1
	}

	fmt.Println("Total Score Part 1:", total_score)
	fmt.Println("Total Score Part 2:", total_score2)

	CheckErr(scanner.Err())
}
