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

func FindCommon(left string, right string) rune {
	left_map := map[rune]int{}

	for _, v := range left {
		left_map[v] += 1
	}

	for _, c := range right {
		_, exists := left_map[c]
		if exists {
			return c
		}
	}

	return ' '
}

func Priority(ch rune) int {
	if ch >= 'a' {
		return int(ch-'a') + 1
	} else {
		return int(ch-'A') + 27
	}
}

func GroupBadge(groups [3]string) rune {
	m1 := map[rune]int{}
	for _, c := range groups[0] {
		m1[c]++
	}

	m2 := map[rune]int{}
	for _, c := range groups[1] {
		_, exists := m1[c]
		if exists {
			m2[c]++
		}
	}

	for _, c := range groups[2] {
		_, exists := m2[c]
		if exists {
			return c
		}
	}

	return ' '
}

func main() {
	f, err := os.Open(os.Args[1])
	CheckErr(err)
	defer f.Close()

	priority_sum := 0

	groups := [3]string{"", "", ""}
	group_sum := 0

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		groups[i%3] = line

		half := len(line) / 2
		common := FindCommon(line[:half], line[half:])

		// fmt.Println(line[:half], line[half:])
		// fmt.Println("Common:", string(common), Priority(common))

		priority_sum += Priority(common)

		if i%3 == 2 {
			group_sum += Priority(GroupBadge(groups))
		}
	}

	fmt.Println("Priority sum:", priority_sum)
	fmt.Println("Group priority sum:", group_sum)
}
