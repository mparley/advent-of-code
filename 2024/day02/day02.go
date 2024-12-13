package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func safe(report []int) bool {
	decreasing := false

	if report[0]-report[1] == 0 {
		return false
	} else if report[0]-report[1] > 0 {
		decreasing = true
	}

	for i := 1; i < len(report); i++ {
		val := report[i-1] - report[i]

		if val == 0 {
			return false
		}

		if val < 0 {
			if decreasing {
				return false
			}
			val = -val
		} else if !decreasing {
			return false
		}

		if val > 3 {
			return false
		}
	}

	return true
}

func evaluate(reports [][]int, part2 bool) int {
	safeCount := 0
	for _, report := range reports {
		// fmt.Println(report)
		if safe(report) {
			// fmt.Println("SAFE")
			safeCount++
		} else if part2 {
			for i := range report {
				r := make([]int, 0, len(report))
				r = append(r, report[:]...)
				r = append(r[:i], r[i+1:]...)
				if safe(r) {
					// fmt.Println("SAFE")
					safeCount++
					break
				}
			}
			// fmt.Println("UNSAFE")
		} else {
			// fmt.Println("UNSAFE")
		}
	}

	return safeCount
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	reports := make([][]int, 0, len(input))

	for _, line := range input {
		report := make([]int, 0)
		vals := strings.Fields(line)
		for _, val := range vals {
			levels, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			report = append(report, levels)
		}
		reports = append(reports, report)
	}

	fmt.Println("Part 1:", evaluate(reports, false))
	fmt.Println("Part 2:", evaluate(reports, true))
}
