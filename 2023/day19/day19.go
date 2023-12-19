package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type category int

const (
	xtreme category = iota
	musical
	aerodyn
	shiny
)

type machinePart []int

type instruction struct {
	cat    category
	op     rune
	val    int
	sendTo string
}

type vrange struct {
	lo, hi int
}

func getCat(c rune) category {
	return map[rune]category{'x': 0, 'm': 1, 'a': 2, 's': 3}[c]
}

func examinePart(p machinePart, workflows map[string][]instruction) bool {
	cur := "in"
	for cur != "A" && cur != "R" {
		for _, ins := range workflows[cur] {
			if ins.op == '.' {
				cur = ins.sendTo
				break
			} else if ins.op == '>' && p[ins.cat] > ins.val {
				cur = ins.sendTo
				break
			} else if ins.op == '<' && p[ins.cat] < ins.val {
				cur = ins.sendTo
				break
			}
		}
	}

	return cur == "A"
}

func splitRanges(splat int, vrs []vrange) ([]vrange, []vrange) {
	sort.Slice(vrs, func(i, j int) bool {
		return vrs[i].lo < vrs[j].lo
	})

	lo := make([]vrange, 0, len(vrs))
	hi := make([]vrange, 0, len(vrs))
	for i, vr := range vrs {
		if splat > vr.lo && splat <= vr.hi {
			lo = append(lo, vrs[:i]...)
			lo = append(lo, vrange{vr.lo, splat})
			hi = append(hi, vrange{splat, vr.hi})
			hi = append(hi, vrs[i+1:]...)
			break
		}
	}
	return lo, hi

}

func sumRanges(vrs []vrange) int {
	sum := 0
	for _, vr := range vrs {
		sum += (vr.hi - vr.lo)
	}
	return sum
}

func findRanges(wfs map[string][]instruction, w string, vrs [][]vrange) int {
	//	fmt.Println("Calling")
	//	fmt.Println(w, vrs)
	//	fmt.Println()
	out := 0
	for w != "A" && w != "R" {
		for _, ins := range wfs[w] {
			if ins.op == '.' {
				w = ins.sendTo
				continue
			}

			vr2 := make([][]vrange, 4)
			copy(vr2, vrs)
			//	fmt.Println("Splitting")
			//	fmt.Println(vrs)
			if ins.op == '>' {
				vrs[ins.cat], vr2[ins.cat] = splitRanges(ins.val, vrs[ins.cat])
			} else if ins.op == '<' {
				vr2[ins.cat], vrs[ins.cat] = splitRanges(ins.val-1, vrs[ins.cat])
			}
			//	fmt.Println(vrs)
			//	fmt.Println(vr2)
			//	fmt.Println()
			out += findRanges(wfs, ins.sendTo, vr2)
		}
	}

	if w == "A" {
		val := 1
		for _, vr := range vrs {
			val *= sumRanges(vr)
		}
		out += val
	}

	//	fmt.Println("Returning")
	//	fmt.Println(w, vrs)
	//	fmt.Println(out)
	//	fmt.Println()
	return out
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Fields(string(f))

	workflows := make(map[string][]instruction)
	parts := make([]machinePart, 0, len(input)/2)

	for _, line := range input {
		if line[0] != '{' {
			tokens := strings.FieldsFunc(line, func(r rune) bool {
				return r == '{' || r == '}' || r == ':' || r == ','
			})
			ins := make([]instruction, 0, 10)
			for i := 1; i < len(tokens)-1; i += 2 {
				cat := getCat(rune(tokens[i][0]))
				op := rune(tokens[i][1])
				val, _ := strconv.Atoi(tokens[i][2:])
				ins = append(ins, instruction{cat, op, val, tokens[i+1]})
			}
			ins = append(ins, instruction{0, '.', 0, tokens[len(tokens)-1]})
			workflows[tokens[0]] = ins
		} else {
			var x, m, a, s int
			fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
			parts = append(parts, machinePart{x, m, a, s})
		}
	}

	//	for w, f := range workflows {
	//		fmt.Println(w, f)
	//	}

	sum := 0
	for _, p := range parts {
		//fmt.Println(p)
		if examinePart(p, workflows) {
			sum += p[xtreme] + p[musical] + p[aerodyn] + p[shiny]
		}
	}

	fmt.Println("Part 1:", sum)

	vrs := make([][]vrange, 4)
	for i := range vrs {
		vrs[i] = append(vrs[i], vrange{0, 4000})
	}
	fmt.Println("Part 2:", findRanges(workflows, "in", vrs))

}
