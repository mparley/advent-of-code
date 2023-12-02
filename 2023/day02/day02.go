package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	var games [][]map[string]int

	//  i
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		var game []map[string]int
		hand := make(map[string]int)
		for j := 2; j < len(tokens); j += 2 {
			val, _ := strconv.Atoi(tokens[j])
			hand[string(tokens[j+1][0])] += val
			if tokens[j+1][len(tokens[j+1])-1] == ';' || j+1 == len(tokens)-1 {
				game = append(game, hand)
				hand = make(map[string]int)
			}
		}
		games = append(games, game)
		// fmt.Println(i + 1, game)
	}

	var bags []map[string]int
	for _, game := range games {
		bag := map[string]int{"r": 0, "g": 0, "b": 0}
		for _, hand := range game {
			if hand["r"] > bag["r"] {
				bag["r"] = hand["r"]
			}
			if hand["b"] > bag["b"] {
				bag["b"] = hand["b"]
			}
			if hand["g"] > bag["g"] {
				bag["g"] = hand["g"]
			}
		}
		bags = append(bags, bag)
	}

	idSum := 0
	powerSum := 0
	for i, bag := range bags {
		if bag["r"] <= 12 && bag["g"] <= 13 && bag["b"] <= 14 {
			idSum += i + 1
		}
		powerSum += bag["r"] * bag["g"] * bag["b"]
	}

	fmt.Println(idSum)
	fmt.Println(powerSum)
}
