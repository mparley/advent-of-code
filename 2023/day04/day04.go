package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cards := strings.Split(strings.TrimSpace(string(f)), "\n")
	totalPoints := 0
	winCounts := make([]int, len(cards))
	cardCounts := make([]int, len(cards))
	cardSum := len(cards)

	for i, card := range cards {
		cardCounts[i]++
		tokens := strings.Fields(card)
		//fmt.Println(tokens)
		winning := make(map[int]struct{})
		var j int = 2
		for ; tokens[j] != "|"; j++ {
			val, err := strconv.Atoi(tokens[j])
			if err != nil {
				log.Fatal(err)
			}
			winning[val] = struct{}{}
		}

		j++

		winCount := 0

		for ; j < len(tokens); j++ {
			val, err := strconv.Atoi(tokens[j])
			if err != nil {
				log.Fatal(err)
			}

			_, ok := winning[val]
			if ok {
				winCount++
			}
		}

		winCounts[i] = winCount

		if winCount > 0 {
			totalPoints += (1 << (winCount - 1))
		}

		for k := 0; k < winCount; k++ {
			cardCounts[i+k+1] += (cardCounts[i])
			cardSum += (cardCounts[i])
		}
	}

	//	sum := 0
	//	for i, c := range cardCounts {
	//		fmt.Println(i, ":", c)
	//		sum += c
	//	}
	//	fmt.Println(sum)

	fmt.Println("Part 1:", totalPoints)
	fmt.Println("Part 2:", cardSum)
}
