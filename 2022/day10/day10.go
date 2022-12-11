package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cycle := 0
	x := 1
	xlog := []int{1}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		input := strings.Split(scanner.Text(), " ")

		switch input[0] {
		case "noop":
			xlog = append(xlog, x)
			cycle += 1
		case "addx":
			xlog = append(xlog, x, x)
			cycle += 2
			value, _ := strconv.Atoi(input[1])
			x += value
		}
	}

	sum := 0
	for i := 20; i < len(xlog); i += 40 {
		// fmt.Println(i, "*", xlog[i], "=", xlog[i]*i)
		sum += (xlog[i] * i)
	}

	fmt.Println("Sum of signal strengths:", sum)

	scanline := ""
	for i := 1; i <= 240; i++ {
		pixel := (i - 1) % 40
		if pixel >= xlog[i]-1 && pixel <= xlog[i]+1 {
			scanline += "#"
		} else {
			scanline += "."
		}

		if pixel == 39 {
			fmt.Println(scanline)
			scanline = ""
		}
	}
}
