package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Overlap(left []int, right []int) (over, cont bool) {
	if left[0] <= right[0] && left[1] >= right[1] {
		return true, true
	} else if left[0] >= right[0] && left[1] <= right[1] {
		return true, true
	} else if left[0] <= right[0] && right[0] <= left[1] {
		return true, false
	} else if left[0] <= right[1] && right[1] <= left[1] {
		return true, false
	} else {
		return false, false
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	CheckErr(err)
	defer f.Close()

	contain_count := 0
	overlap_count := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// Getting bounds and converting them to int
		rx := regexp.MustCompile("[0-9]+")
		bounds_str := rx.FindAllString(line, 4)
		var bounds [4]int
		for i, v := range bounds_str {
			bounds[i], _ = strconv.Atoi(v)
		}

		// fmt.Println(bounds)

		// Checking overlaps and contains
		overlap, contains := Overlap(bounds[:2], bounds[2:])
		if contains {
			contain_count++
		}
		if overlap {
			overlap_count++
		}
	}

	fmt.Println("Fully contain count:", contain_count)
	fmt.Println("overlap count:", overlap_count)
}
