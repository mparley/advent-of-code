package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func IsSpelledOut(s string) int {
	if len(s) < 3 {
		return 0
	}

	if s[:3] == "one" {
		return 1
	}

	if s[:3] == "two" {
		return 2
	}

	if s[:3] == "six" {
		return 6
	}

	if len(s) < 4 {
		return 0
	}

	if s[:4] == "four" {
		return 4
	}

	if s[:4] == "nine" {
		return 9
	} 

	if s[:4] == "five" {
	    return 5
	}

	if len(s) < 5 {
		return 0
	}

	if s[:5] == "three" {
		return 3
	}

	if s[:5] == "seven" {
	    return 7
	}

	if s[:5] == "eight" {
	    return 8
	}

	return 0
}


func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	var sum1 int = 0
	var sum2 int = 0

	for _, line := range lines {
		val1 := 0
		//fmt.Println(line)
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				//fmt.Println("First:",int(line[i]-'0'))
				val1 += int(line[i]-'0') * 10
				break
			}
		}

		for i := len(line)-1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				//fmt.Println("Last:",int(line[i]-'0'))
				val1 += int(line[i]-'0')
				break
			}
		}
		//fmt.Println(val1)
		sum1 += val1
	}

	for _, line := range lines {
		val2 := 0
		fmt.Println(line)
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				fmt.Println("First:",int(line[i]-'0'))
				val2 += int(line[i]-'0') * 10
				break
			} else if sp := IsSpelledOut(line[i:]); sp > 0 {
				fmt.Println("First:",sp)
				val2 += sp * 10
				break
			}
		}

		for i := len(line)-1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				fmt.Println("Last:",int(line[i]-'0'))
				val2 += int(line[i]-'0')
				break
			} else if sp := IsSpelledOut(line[i:]); sp > 0 {
				fmt.Println("Last:",sp)
				val2 += sp
				break
			}
		}
		fmt.Println(val2)
		sum2 += val2
	}

	fmt.Println("Part 1:",sum1)
	fmt.Println("Part 2:",sum2)
}
