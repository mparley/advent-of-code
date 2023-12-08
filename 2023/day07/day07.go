package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handType int

const (
	HighCard handType = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func cardVals(part2 bool) map[rune]int {
	ret := map[rune]int{
		'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
		'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13,
	}
	if part2 {
		ret['J'] = 0
	}
	return ret
}

type hand struct {
	cards string
	bid   int
}

func evaluateHand(s string, part2 bool) int {
	handValue := 0
	cVals := cardVals(part2)

	cardCount := make(map[rune]int)
	maxCount := 0

	for i, ch := range s {
		cardCount[ch]++
		handValue += cVals[ch] << ((len(s) - 1 - i) * 4)
		if cardCount[ch] > maxCount {
			maxCount = cardCount[ch]
		}
	}

	//fmt.Println(s, handValue)

	if part2 && len(cardCount) > 1 {
		keys := make([]rune, 0, len(cardCount))
		for k := range cardCount {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(l, r int) bool {
			return cardCount[keys[l]] >= cardCount[keys[r]]
		})

		if keys[0] == 'J' {
			cardCount[keys[1]] += cardCount[keys[0]]
			maxCount = cardCount[keys[1]]
			delete(cardCount, 'J')
		} else if _, ok := cardCount['J']; ok {
			cardCount[keys[0]] += cardCount['J']
			maxCount = cardCount[keys[0]]
			delete(cardCount, 'J')
		}
	}

	var ht handType
	switch len(cardCount) {
	case 1:
		ht = FiveKind
	case 2:
		if maxCount == 4 {
			ht = FourKind
		} else {
			ht = FullHouse
		}
	case 3:
		if maxCount == 3 {
			ht = ThreeKind
		} else {
			ht = TwoPair
		}
	case 4:
		ht = OnePair
	default:
		ht = HighCard
	}

	return handValue + ((int(ht)) << (len(s) * 4))
}

func solve(hands []hand, part2 bool) int {
	handValues := make(map[string]int)
	for _, h := range hands {
		handValues[h.cards] = evaluateHand(h.cards, part2)
	}

	sort.Slice(hands, func(l, r int) bool {
		return handValues[hands[l].cards] < handValues[hands[r].cards]
	})
	//	fmt.Println(hands)

	totalWinnings := 0
	for i, h := range hands {
		totalWinnings += (i + 1) * h.bid
	}
	return totalWinnings
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	checkErr(err)

	inputs := strings.Fields(string(f))
	hands := make([]hand, 0, len(inputs)/2)

	for i := 0; i < len(inputs); i += 2 {
		bid, err := strconv.Atoi(inputs[i+1])
		checkErr(err)
		h := hand{inputs[i], bid}
		hands = append(hands, h)
	}

	fmt.Println("Part 1:", solve(hands, false))
	fmt.Println("Part 2:", solve(hands, true))
}
