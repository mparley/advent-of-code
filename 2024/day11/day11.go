package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// helper to count digits of int
func digits(n int) (d int) {
	for i := n; i > 0; i /= 10 {
		d++
	}
	return
}

// Main part of solution
// Sort of like previous years problems - I forget the name
// but basically instead of trying to simulate it in a big ass
// array, you count the occurances of stones (since order doesn't matter)
func blink(stones map[int]int) map[int]int {
	blinked := make(map[int]int)

	for stone, count := range stones {

		if stone == 0 {
			blinked[1] += count

		} else if d := digits(stone); d%2 == 0 {
			mid := 1
			for i := 0; i < d/2; i++ {
				mid *= 10
			}
			blinked[stone/mid] += count
			blinked[stone%mid] += count

		} else {
			blinked[stone*2024] += count
		}
	}

	return blinked
}

func sum(stones map[int]int) (sum int) {
	for _, count := range stones {
		sum += count
	}
	return
}

func parseInput(input []string) map[int]int {
	out := make(map[int]int)
	for _, s := range input {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		out[val] += 1
	}
	return out
}

// Args[2] is for setting amount of blinks (and get part 2)
// when its set over 25 it will print out part 1 solution too
func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	stones := parseInput(strings.Fields(string(data)))

	numBlinks := 25
	if len(os.Args) > 2 {
		numBlinks, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	// fmt.Println(stones)
	for i := 0; i < numBlinks; i++ {
		if i == 25 {
			fmt.Println("After 25 blinks:", sum(stones))
		}
		stones = blink(stones)
		// fmt.Println(stones)
	}
	fmt.Println("After", numBlinks, "blinks:", sum(stones))

}
