package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type signal struct {
	pulse               bool
	sender, destination string
}

type signalQ []signal

type module struct {
	name    string
	mtype   rune
	state   bool
	memory  map[string]bool
	outputs []string
}

func (sq *signalQ) push(signals ...signal) {
	*sq = append(*sq, signals...)
}

func (sq *signalQ) pop() signal {
	s := (*sq)[0]
	*sq = (*sq)[1:]
	return s
}

func (m *module) process(in signal) []signal {
	outPulse := false

	switch m.mtype {
	case 'b':
		outPulse = in.pulse
	case '%':
		if !in.pulse {
			m.state = !m.state
			outPulse = m.state
		} else {
			return []signal{}
		}
	case '&':
		allHi := true
		m.memory[in.sender] = in.pulse
		for _, v := range m.memory {
			if !v {
				allHi = false
				break
			}
		}
		outPulse = !allHi
	default:
		return []signal{}
	}

	ret := make([]signal, 0, len(m.outputs))
	for _, output := range m.outputs {
		ret = append(ret, signal{outPulse, m.name, output})
	}
	return ret
}

func parseModules(input []string) map[string]*module {
	modules := make(map[string]*module)
	modOutputs := make(map[string][]string)

	for _, line := range input {
		tokens := strings.Fields(line)

		name := tokens[0]
		if name[0] != 'b' {
			name = name[1:]
		}
		mtype := rune(tokens[0][0])

		modules[name] = &module{name, mtype, false, map[string]bool{}, make([]string, 0, len(tokens)-2)}

		for i := 2; i < len(tokens); i++ {
			if i != len(tokens)-1 {
				n := len(tokens[i])
				tokens[i] = tokens[i][:n-1]
			}
			modOutputs[name] = append(modOutputs[name], tokens[i])
		}
	}

	for mod, outputs := range modOutputs {
		for _, output := range outputs {
			modules[mod].outputs = append(modules[mod].outputs, output)
			if _, ok := modules[output]; ok {
				modules[output].memory[mod] = false
			} else {
				m := &module{output, '.', false, map[string]bool{mod: false}, []string{}}
				modules[output] = m
			}
		}
	}

	return modules
}

func pressButton(modules map[string]*module, observe map[string]struct{}) (int, int, map[string]struct{}) {
	found := make(map[string]struct{})

	sQ := signalQ{signal{false, "button", "broadcaster"}}
	hiCount, loCount := 0, 0

	for len(sQ) > 0 {
		cur := sQ.pop()
		if _, ok := observe[cur.sender]; ok {
			if cur.pulse {
				//fmt.Println("found", cur.sender)
				found[cur.sender] = struct{}{}
			}
		}

		if cur.pulse {
			hiCount++
		} else {
			loCount++
		}

		next := modules[cur.destination].process(cur)
		sQ.push(next...)
	}
	return loCount, hiCount, found
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(vals ...int) int {
	lcm := 1
	for i := 0; i < len(vals); i++ {
		lcm = (vals[i] * lcm) / gcd(vals[i], lcm)
	}

	return lcm
}

func part1(modules map[string]*module) int {
	loCount, hiCount := 0, 0

	for i := 0; i < 1000; i++ {
		l, h, _ := pressButton(modules, map[string]struct{}{})
		loCount += l
		hiCount += h
	}

	return loCount * hiCount
}

func part2(modules map[string]*module, target string) int {
	watch := map[string]struct{}{}
	vals := make([]int, 0, 10)

	var observe *module
	for name := range modules[target].memory {
		if modules[name].mtype == '&' {
			observe = modules[name]
			break
		}
	}

	for name := range observe.memory {
		watch[name] = struct{}{}
	}

	for i := 1; len(watch) > 0; i++ {
		_, _, found := pressButton(modules, watch)

		for name := range found {
			if _, ok := watch[name]; ok {
				vals = append(vals, i)
				delete(watch, name)
			}
		}
	}

	return lcm(vals...)
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	modules := parseModules(input)
	modules2 := map[string]*module{}
	for k, v := range modules {
		m := *v
		modules2[k] = &m
	}

	fmt.Println("Part 1:", part1(modules))
	fmt.Println("Part 2:", part2(modules2, "rx"))
}
