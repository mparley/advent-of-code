package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func findBounds(input string, dos, donts [][]int) [][]int {
	l := len(input)

	var bounds [][]int

	on := true
	for i := 0; i < l; {
		if on {
			if len(donts) == 0 {
				bounds = append(bounds, []int{i, l})
				break
			}

			for len(donts) > 0 {
				val := donts[0][0]
				donts = donts[1:]
				if val > i {
					on = false
					bounds = append(bounds, []int{i, val})
					i = val + 1
					break
				}
			}
		} else {
			if len(dos) == 0 {
				break
			}

			for len(dos) > 0 {
				val := dos[0][0]
				dos = dos[1:]
				if val > i {
					on = true
					i = val
					break
				}
			}
		}
	}

	return bounds
}

func evaluate(matches [][]string) int {
	sum := 0
	for _, match := range matches {
		l, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}

		r, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(l, "*", r, "=", l*r)
		sum += l * r
	}

	return sum
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	reMul := regexp.MustCompile(`mul\(([0-9]*),([0-9]*)\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	results := reMul.FindAllStringSubmatch(input, -1)
	// fmt.Println(results)
	fmt.Println("Part 1:", evaluate(results))

	dos := reDo.FindAllStringIndex(input, -1)
	donts := reDont.FindAllStringIndex(input, -1)
	// fmt.Println(dos)
	// fmt.Println(donts)

	bounds := findBounds(input, dos, donts)
	// fmt.Println(bounds)

	sum2 := 0
	for _, bound := range bounds {
		results = reMul.FindAllStringSubmatch(input[bound[0]:bound[1]], -1)
		sum2 += evaluate(results)
	}
	fmt.Println("Part 2:", sum2)
}
