package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type srange struct {
	start int
	end   int
}

func newSrange(start int, length int) *srange {
	return &srange{start, start + length}
}

func parseMappings(s string) (*srange, *srange) {
	ss := strings.Fields(s)
	dstart, _ := strconv.Atoi(ss[0])
	sstart, _ := strconv.Atoi(ss[1])
	length, _ := strconv.Atoi(ss[2])
	return &srange{sstart, sstart + length}, &srange{dstart, dstart + length}
}

func splitSranges(src *srange, val *srange) ([]*srange, int) {
	var r []*srange

	//  src
	// v a l
	if val.start < src.start && src.end < val.end {
		r = make([]*srange, 3)
		r[0] = &srange{val.start, src.start - 1}
		r[1] = &srange{src.start, src.end}
		r[2] = &srange{src.end + 1, val.end}
		return r, 1
	}

	// s r c
	//  val
	if src.start <= val.start && val.end <= src.end {
		return []*srange{val}, 0
	}

	//  s r c
	// v a l
	if val.start < src.start && src.start <= val.end && val.end <= src.end {
		r = make([]*srange, 2)
		r[0] = &srange{val.start, src.start - 1}
		r[1] = &srange{src.start, val.end}
		return r, 1
	}

	// s r c
	//  v a l
	if src.start <= val.start && val.start <= src.end && src.end < val.end {
		r = make([]*srange, 2)
		r[0] = &srange{val.start, src.end}
		r[1] = &srange{src.end + 1, val.end}
		return r, 0
	}

	return r, -1
}

func part1(seeds []int, input []string) int {
	lowest := math.MaxInt
	vals := make(map[int]struct{})
	changedVals := make(map[int]struct{})

	for _, seed := range seeds {
		vals[seed] = struct{}{}
	}

	for i, line := range input {
		if len(line) == 0 {
			continue
		}
		//fmt.Println(line)

		if unicode.IsDigit(rune(line[0])) {
			if len(vals) == 0 {
				continue
			}

			src, dst := parseMappings(line)

			for key := range vals {
				dif := dst.start - src.start
				if key >= src.start && key <= src.end {
					changedVals[key+dif] = struct{}{}
					delete(vals, key)
				}
			}
		}

		if !unicode.IsDigit(rune(line[0])) || i == len(input)-1 {
			for k := range changedVals {
				vals[k] = struct{}{}
				delete(changedVals, k)
			}
		}
	}

	for k := range vals {
		if k < lowest {
			lowest = k
		}
	}

	return lowest
}

func part2(seeds []int, input []string) int {
	lowest := math.MaxInt
	vals := make(map[srange]struct{})
	changedVals := make(map[srange]struct{})

	for i := 0; i < len(seeds); i += 2 {
		vals[*newSrange(seeds[i], seeds[i+1])] = struct{}{}
	}

	for i, line := range input {
		if len(line) == 0 {
			continue
		}
		//fmt.Println(line)

		if unicode.IsDigit(rune(line[0])) {
			if len(vals) == 0 {
				continue
			}

			src, dst := parseMappings(line)
			var addBack []srange

			for key := range vals {
				ret, m := splitSranges(src, &key)
				if m != -1 {
					for i, r := range ret {
						if i == m {
							dif := dst.start - src.start
							nr := srange{r.start + dif, r.end + dif}
							changedVals[nr] = struct{}{}
						} else {
							addBack = append(addBack, *r)
						}
					}
					delete(vals, key)
				}
			}

			for _, key := range addBack {
				vals[key] = struct{}{}
			}
		}

		if !unicode.IsDigit(rune(line[0])) || i == len(input)-1 {
			for k := range changedVals {
				vals[k] = struct{}{}
				delete(changedVals, k)
			}
		}
	}

	for k := range vals {
		if k.start < lowest {
			lowest = k.start
		}
	}

	return lowest
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	seedLine := strings.Fields(lines[0])
	lines = lines[1:]
	var seeds []int

	for i := 1; i < len(seedLine); i++ {
		val, err := strconv.Atoi(seedLine[i])
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, val)
	}

	if len(os.Args) == 2 || os.Args[2] == "1" {
		fmt.Println("Part 1:", part1(seeds, lines))
	}

	if len(os.Args) == 2 || os.Args[2] == "2" {
		fmt.Println("Part 2:", part2(seeds, lines))
	}
}
