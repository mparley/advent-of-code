package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type intSet struct {
	m map[int]struct{}
}

func (s intSet) has(val int) bool {
	_, ok := s.m[val]
	return ok
}

func (s intSet) add(val int) {
	s.m[val] = struct{}{}
}

func parseInput(data []byte) (map[int]intSet, [][]int) {
	parts := strings.Split(string(data), "\n\n")
	if len(parts) != 2 {
		log.Fatal("Couldn't Parse")
	}

	ruleMap := make(map[int]intSet)

	for _, rule := range strings.Fields(parts[0]) {
		stringvals := strings.Split(rule, "|")
		before, err := strconv.Atoi(stringvals[0])
		if err != nil {
			log.Fatal(err)
		}
		after, err := strconv.Atoi(stringvals[1])
		if err != nil {
			log.Fatal(err)
		}
		_, ok := ruleMap[after]
		if !ok {
			ruleMap[after] = intSet{map[int]struct{}{}}
		}
		ruleMap[after].add(before)
	}

	pageLists := make([][]int, 0)

	for _, pages := range strings.Fields(parts[1]) {
		pagenums := strings.Split(pages, ",")
		pageList := make([]int, 0, len(pagenums))
		for _, page := range pagenums {
			val, err := strconv.Atoi(page)
			if err != nil {
				log.Fatal(err)
			}
			pageList = append(pageList, val)
		}
		pageLists = append(pageLists, pageList)
	}

	return ruleMap, pageLists
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	ruleMap, pageLists := parseInput(data)

	// fmt.Println("rules:", ruleMap, "\npages:", pageLists)

	count := 0
	middleSum := 0
	bads := make([]int, 0)

	for update, pageList := range pageLists {
		failed := false
		for i := len(pageList) - 1; i >= 0 && !failed; i-- {
			curr := pageList[i]
			for j := i + 1; j < len(pageList); j++ {
				if ruleMap[curr].has(pageList[j]) {
					failed = true
					bads = append(bads, update)
					break
				}
			}
		}
		if !failed {
			count++
			middleSum += pageList[len(pageList)/2]
		}
	}

	fmt.Println("Part 1:", middleSum, "(", count, "correct )")

	middleSum = 0
	for _, bad := range bads {
		pageList := pageLists[bad]
		for i := len(pageList) - 1; i >= 0; i-- {
			curr := pageList[i]
			last := -1
			for j := i + 1; j < len(pageList); j++ {
				if ruleMap[curr].has(pageList[j]) {
					last = j
				}
			}

			if last > 0 {
				pageList = slices.Insert(pageList, last+1, curr)
				pageList = slices.Delete(pageList, i, i+1)
			}
		}
		middleSum += pageList[len(pageList)/2]
	}

	fmt.Println("Part 2:", middleSum, "(", len(bads), "incorrect )")
}
