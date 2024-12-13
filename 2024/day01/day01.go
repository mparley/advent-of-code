package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Fields(string(data))
	leftList := make([]int, 0, len(inputs)/2)
	rightList := make([]int, 0, len(inputs)/2)
	counts := make(map[int]int)

	for i, input := range inputs {
		val, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if i%2 == 0 {
			leftList = append(leftList, val)
		} else {
			rightList = append(rightList, val)
			counts[val]++
		}
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	sum1 := 0
	sum2 := 0

	for i, l := range leftList {
		diff := l - rightList[i]
		if diff < 0 {
			diff = -diff
		}
		sum1 += diff
		sum2 += l * counts[l]
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
