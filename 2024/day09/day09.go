package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type block struct {
	id   int
	size int
}

// Part 1 solution move each id at end to first open space
// because of how I'm storing the blocks I have to make a new free block
// when the file doesn't take up the whole space
func defrag1(disk []block) []block {
	blocks := append([]block{}, disk...)

	// kinda janky but i is for free space and j if for files
	for i, j := 0, len(blocks)-1; i < j; {
		for blocks[j].id == -1 {
			j--
		}
		for blocks[i].id != -1 {
			i++
		}
		if i >= j { // need to check if free space isn't found
			continue
		}

		sizediff := blocks[j].size - blocks[i].size
		if sizediff == 0 {
			blocks[i], blocks[j] = blocks[j], blocks[i]

		} else if sizediff > 0 {
			blocks[j].size = sizediff
			blocks[i].id = blocks[j].id

			// we need to add another block if the file doesn't fill the free block
		} else {
			sizediff = -sizediff
			blocks[i] = blocks[j]
			blocks[j].id = -1
			end := append([]block{{-1, sizediff}}, blocks[i+1:]...)
			blocks = append(blocks[:i+1], end...)
		}
	}

	return blocks
}

// Part 2 solution similar to part one but we need to reset i for each id
func defrag2(disk []block) []block {
	blocks := append([]block{}, disk...)

	for i, j := 0, len(blocks)-1; j >= 0; j-- {
		for j >= 0 && (blocks[j].id == -1) {
			j--
		}
		if j < 0 {
			continue
		}

		// looking for first free block *that is large enough*
		for i = 0; i < j && (blocks[i].id != -1 || blocks[i].size < blocks[j].size); i++ {
		}
		if i >= j {
			continue
		}

		id := blocks[j].id
		sizediff := blocks[i].size - blocks[j].size

		if sizediff == 0 {
			blocks[i], blocks[j] = blocks[j], blocks[i]

			// Here we check if the next block is a free block, if so we can just add to that
		} else if blocks[i+1].id == -1 {
			blocks[i+1].size += sizediff
			blocks[i].id = id
			blocks[j].id = -1

			// Again we have to add another block for when there is extra free space
		} else {
			blocks[j].id = -1
			blocks[i] = block{id, blocks[j].size}
			end := append([]block{{-1, sizediff}}, blocks[i+1:]...)
			blocks = append(blocks[:i+1], end...)
			j++ // since we add to the list make sure to increment j
		}
	}

	return blocks
}

func checksum(disk []block) int {
	sum := 0
	i := 0
	for _, b := range disk {
		for j := 0; j < b.size; j++ {
			if b.id != -1 {
				sum += (i * b.id)
			}
			i++
		}
	}
	return sum
}

func printDisk(disk []block) {
	for _, d := range disk {
		for i := 0; i < d.size; i++ {
			if d.id == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(strconv.Itoa(d.id))
			}
		}
	}
	fmt.Print("\n")
}

func parseDisk(diskStr string) []block {
	blocks := make([]block, 0)

	for i, id := 0, 0; i < len(diskStr); i++ {
		val := int(diskStr[i] - '0')
		if i%2 == 0 {
			blocks = append(blocks, block{id, val})
		} else {
			blocks = append(blocks, block{-1, val})
			id++
		}
	}

	return blocks
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	blocks := parseDisk(strings.TrimSpace(string(data)))

	p1 := defrag1(blocks)
	p2 := defrag2(blocks)

	if len(os.Args) > 2 {
		printDisk(p1)
		printDisk(p2)
	}

	fmt.Println("Part 1:", checksum(p1))
	fmt.Println("Part 2:", checksum(p2))
}
