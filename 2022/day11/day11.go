package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items       []uint64
	t, f        int
	inspections uint64
	operands    []string
	Operation   func(uint64, uint64) uint64
	test        uint64
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Round(barrel *[]*Monkey, d uint64) {
	for _, monkey := range *barrel {
		for len(monkey.items) > 0 {
			to, item := monkey.Inspect(d)
			(*barrel)[to].Catch(item)
		}
	}
}

func PrintMonkeys(barrel *[]*Monkey) {
	for i, m := range *barrel {
		fmt.Println("Monkey", i, ":", m.items)
	}
}

func (m *Monkey) Inspect(d uint64) (to int, item uint64) {
	item = m.items[0]
	m.items = m.items[1:]
	var r uint64
	if m.operands[1] == "old" {
		r = item
	} else {
		val, err := strconv.Atoi(m.operands[1])
		CheckErr(err)
		r = uint64(val)
	}
	item = m.Operation(item, r)

	if d == 0 {
		item = item / 3
	} else {
		item = item % d
	}

	if (item % m.test) == 0 {
		to = m.t
	} else {
		to = m.f
	}

	m.inspections++
	return to, item
}

func (m *Monkey) Catch(item uint64) {
	m.items = append(m.items, item)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	CheckErr(err)
	lines := strings.Split(string(data), "\n")
	var d uint64 = 0
	if os.Args[2] == "2" {
		d = 1
	}
	barrel := []*Monkey{}

	for i := 0; i < len(lines); i += 7 {
		m := Monkey{}
		tokens := strings.Fields(strings.ReplaceAll(lines[i+1], ",", ""))
		m.items = []uint64{}
		for j := 2; j < len(tokens); j++ {
			val, err := strconv.Atoi(tokens[j])
			CheckErr(err)
			m.items = append(m.items, uint64(val))
		}

		tokens = strings.Fields(lines[i+2])
		if tokens[4] == "*" {
			m.Operation = func(l uint64, r uint64) uint64 {
				return l * r
			}
		} else {
			m.Operation = func(l uint64, r uint64) uint64 {
				return l + r
			}
		}

		m.operands = []string{tokens[3], tokens[5]}

		tokens = strings.Fields(lines[i+3])
		val, err := strconv.Atoi(tokens[len(tokens)-1])
		CheckErr(err)
		m.test = uint64(val)
		d *= m.test

		tokens = strings.Fields(lines[i+4])
		val, err = strconv.Atoi(tokens[len(tokens)-1])
		CheckErr(err)
		m.t = val

		tokens = strings.Fields(lines[i+5])
		val, err = strconv.Atoi(tokens[len(tokens)-1])
		CheckErr(err)
		m.f = val

		m.inspections = 0

		barrel = append(barrel, &m)
	}

	for _, m := range barrel {
		fmt.Println(m)
	}
	fmt.Println()

	rounds := 20
	if d != 0 {
		rounds = 10000
	}

	// PrintMonkeys(&barrel)
	for i := 0; i < rounds; i++ {
		Round(&barrel, d)
		// PrintMonkeys(&barrel)
		// fmt.Println("...")
	}

	counts := []uint64{}
	for i, m := range barrel {
		fmt.Println("Monkey", i, "inspected items", m.inspections)
		counts = append(counts, m.inspections)
	}
	sort.Slice(counts, func(l, r int) bool { return counts[l] < counts[r] })
	sol := counts[len(counts)-1] * counts[len(counts)-2]
	fmt.Println(sol)
}
