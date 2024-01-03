package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type graph struct {
	adj map[string][]string
	wts map[string]int
}

func newGraph(components map[string][]string) *graph {
	adj := map[string][]string{}
	wts := map[string]int{}
	for k, v := range components {
		adj[k] = append([]string{}, v...)
		wts[k] = 1
	}
	return &graph{adj, wts}
}

func (g *graph) nodes() []string {
	out := make([]string, 0, len(g.adj))
	for k := range g.adj {
		out = append(out, k)
	}
	return out
}

func (g *graph) length() int {
	return len(g.adj)
}

func (g *graph) contract(from, to, newNode string) {
	g.adj[newNode] = make([]string, 0, len(g.adj[from])+len(g.adj[to]))
	for _, s := range append(g.adj[from], g.adj[to]...) {
		if s != from && s != to {
			g.adj[newNode] = append(g.adj[newNode], s)
		}
	}

	for k, v := range g.adj {
		if k == from || k == to {
			continue
		}
		nadj := make([]string, 0, len(v))
		for _, s := range v {
			if s == from || s == to {
				s = newNode
			}
			nadj = append(nadj, s)
		}
		g.adj[k] = nadj
	}

	g.wts[newNode] = g.wts[from] + g.wts[to]

	delete(g.adj, from)
	delete(g.adj, to)
	delete(g.wts, from)
	delete(g.wts, to)
}

func (g *graph) randomEdge() (string, string) {
	keys := g.nodes()
	r1 := rand.Intn(len(keys))
	from := keys[r1]
	r2 := rand.Intn(len(g.adj[from]))
	to := g.adj[from][r2]
	return from, to
}

func part1(components map[string][]string) int {
	for {
		g := *newGraph(components)

		for g.length() > 2 {
			from, to := g.randomEdge()
			newNode := from + to
			g.contract(from, to, newNode)
		}

		keys := g.nodes()
		group1, group2 := keys[0], keys[1]
		//fmt.Println(len(g.adj[group1]))
		if len(g.adj[group1]) == 3 {
			return g.wts[group1] * g.wts[group2]
		}
	}
}

func main() {
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	components := map[string][]string{}

	for _, line := range input {
		tokens := strings.Fields(line)
		comp := tokens[0][:len(tokens[0])-1]
		components[comp] = append(components[comp], tokens[1:]...)
		for _, c := range tokens[1:] {
			components[c] = append(components[c], comp)
		}
	}

	fmt.Println("Part 1:", part1(components))
}
