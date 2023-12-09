package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func predict(vals []int) (int, int) {
	all0 := true
	difs := make([]int, 0, len(vals)-1)

	for i := 1; i < len(vals); i++ {
		dif := vals[i] - vals[i-1]
		difs = append(difs, dif)
		if dif != 0 {
			all0 = false
		}
	}

	if all0 {
		return vals[len(vals)-1], vals[0]
	}

	hi, lo := predict(difs)
	return vals[len(vals)-1] + hi, vals[0] - lo
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(f)), "\n")
	var sandVals [][]int

	for _, input := range inputs {
		ins := strings.Fields(input)
		vals := make([]int, 0, len(ins))
		for _, i := range ins {
			v, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, v)
		}
		sandVals = append(sandVals, vals)
	}

	sum1 := 0
	sum2 := 0
	for _, sandv := range sandVals {
		hi, lo := predict(sandv)
		sum1 += hi
		sum2 += lo
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
