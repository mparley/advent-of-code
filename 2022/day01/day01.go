package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	args := os.Args[1:]

	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	elves := []int{}
	sum := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, sum)
			sum = 0
			continue
		}

		cal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += cal
	}

	elves = append(elves, sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("The highest total calorie count carried by and elf is %v\n", elves[0])
	fmt.Printf("The three top elves' calories summed is: %v\n", elves[0]+elves[1]+elves[2])
}
