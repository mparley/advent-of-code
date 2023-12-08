package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type pair struct {
	l string
	r string
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func steps(curr string, ins string, nmap map[string]pair,
	tgt func(s string) bool) int {

	steps := 0
	for ; !tgt(curr); steps++ {
		if ins[steps%len(ins)] == 'L' {
			curr = nmap[curr].l
		} else {
			curr = nmap[curr].r
		}
	}
	return steps
}

func part1(curr string, ins string, nmap map[string]pair) int {
	return steps(curr, ins, nmap, func(s string) bool {
		return s == "ZZZ"
	})
}

func part2(starts []string, ins string, nmap map[string]pair) int {
	zsteps := make([]int, 0, len(starts))
	checkZ := func(s string) bool {
		return s[2] == 'Z'
	}

	for _, start := range starts {
		zsteps = append(zsteps, steps(start, ins, nmap, checkZ))
	}

	lcm := zsteps[0]
	for i := 1; i < len(zsteps); i++ {
		lcm = (zsteps[i] * lcm) / gcd(zsteps[i], lcm)
	}

	return lcm
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	checkErr(err)
	inputs := strings.Split(strings.TrimSpace(string(f)), "\n")

	instructions := inputs[0]
	nodeMap := make(map[string]pair)

	re := regexp.MustCompile("([A-Z0-9]+)")
	var starts []string

	for _, input := range inputs[1:] {
		nodes := re.FindAllString(input, -1)
		if len(nodes) == 3 {
			nodeMap[nodes[0]] = pair{nodes[1], nodes[2]}
			if nodes[0][2] == 'A' {
				starts = append(starts, nodes[0])
			}
		}
	}

	if len(os.Args) == 2 || os.Args[2] == "1" {
		fmt.Println("Part 1:", part1("AAA", instructions, nodeMap))
	}
	if len(os.Args) == 2 || os.Args[2] == "2" {
		fmt.Println("Part 2:", part2(starts, instructions, nodeMap))
	}
}
