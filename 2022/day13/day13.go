package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Compare(l, r []interface{}) (in_order int) {
	// fmt.Println("Compare", l, "vs", r)
	minl := len(l)
	if minl > len(r) {
		minl = len(r)
	}

	in_order = 0
	for i := 0; i < minl && in_order == 0; i++ {
		lf, lisf := l[i].(float64)
		rf, risf := r[i].(float64)

		// fmt.Println(l[i], r[i], lf, rf, lisf, risf)
		// fmt.Println(l[i], "vs", r[i])

		if lisf && risf && lf != rf {
			if lf < rf {
				return 1
			} else {
				return -1
			}
		} else if lisf && !risf {
			in_order := Compare([]interface{}{lf}, r[i].([]interface{}))
			// fmt.Println(lf, "vs", r[i], in_order)
			if in_order != 0 {
				return in_order
			}
		} else if !lisf && risf {
			in_order := Compare(l[i].([]interface{}), []interface{}{rf})
			// fmt.Println(l[i], "vs", rf, in_order)
			if in_order != 0 {
				return in_order
			}
		} else if !lisf && !risf {
			in_order := Compare(l[i].([]interface{}), r[i].([]interface{}))
			// fmt.Println(l[i], "vs", r[i], in_order)
			if in_order != 0 {
				return in_order
			}
		}

	}

	dif := len(r) - len(l)
	if dif == 0 {
		return 0
	} else if dif > 0 {
		return 1
	} else {
		return -1
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	CheckErr(err)
	input := strings.Fields(string(data))
	jlist := [][]interface{}{}

	for _, v := range input {
		var d []interface{}
		var b = []byte(v)
		err := json.Unmarshal(b, &d)
		CheckErr(err)
		jlist = append(jlist, d)
	}

	sum := 0
	for i := 1; i < len(jlist); i += 2 {
		// fmt.Println(jlist[i-1], "vs", jlist[i])
		in_order := Compare(jlist[i-1], jlist[i])
		// fmt.Println(in_order, (i+1)/2)
		if in_order == 1 {
			sum += ((i + 1) / 2)
		}
	}

	fmt.Println(sum)

	dividers := [][]byte{[]byte(`[[2]]`), []byte(`[[6]]`)}
	divider_unmarsh := [][]interface{}{}
	for _, divider := range dividers {
		var d []interface{}
		err = json.Unmarshal(divider, &d)
		CheckErr(err)
		jlist = append(jlist, d)
		divider_unmarsh = append(divider_unmarsh, d)
	}

	sort.Slice(jlist, func(l, r int) bool {
		return Compare(jlist[l], jlist[r]) == 1
	})

	decoder_key := 1
	for i, v := range jlist {
		// fmt.Println(v)
		if Compare(v, divider_unmarsh[0]) == 0 {
			decoder_key *= i + 1
		} else if Compare(v, divider_unmarsh[1]) == 0 {
			decoder_key *= i + 1
		}
	}

	fmt.Println(decoder_key)
}
