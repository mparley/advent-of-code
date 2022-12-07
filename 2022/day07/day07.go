package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Goes through and recursively sets the size of directories
func CalcFolderSizes(dmap *map[string]map[string]uint64, start string) uint64 {
	if (*dmap)[start] == nil {
		return 0
	}

	var size uint64 = 0
	for k, v := range (*dmap)[start] {
		if v == 0 {
			(*dmap)[start][k] = CalcFolderSizes(dmap, k)
			size += (*dmap)[start][k]
		} else {
			size += v
		}
	}

	return size
}

// Sums up each directories total size - probably should have just built a map
// for this in CalcFolderSizes and return that but w/e
func FolderSize(directory *map[string]uint64) uint64 {
	var size uint64 = 0
	for _, v := range *directory {
		size += v
	}
	return size
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	directory_map := map[string]map[string]uint64{} // A map of maps
	directory_stack := []string{"/"}

	// Build the directory and file structure - note we need to store whole paths
	// because in the input directories aren't uniquely named (though they are in
	// the example :( )
	for _, v := range lines {
		current := strings.Join(directory_stack, "/")
		tokens := strings.Split(v, " ")

		// Executing line
		if tokens[0] == "$" {
			// on cd manage current directory stack
			if tokens[1] == "cd" {
				if tokens[2] == "/" {
					directory_stack = []string{"/"}
				} else if tokens[2] == ".." {
					directory_stack = directory_stack[:len(directory_stack)-1]
				} else {
					directory_stack = append(directory_stack, tokens[2])
				}

				// on ls build empty map for the current directory
			} else if tokens[1] == "ls" {
				directory_map[current] = map[string]uint64{}
			}

			// Adding directories and files
		} else if tokens[0] == "dir" {
			directory_map[current][current+"/"+tokens[1]] = 0
		} else {
			directory_map[current][current+"/"+tokens[1]], err = strconv.ParseUint(tokens[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	total := CalcFolderSizes(&directory_map, "/")
	// fmt.Println("total size:", total)
	// for k, v := range directory_map {
	// 	fmt.Println(k, ":", v)
	// }

	var sum1 uint64 = 0
	smallest := total
	const disk_size uint64 = 70000000
	const space_needed uint64 = 30000000

	// Loop through map and get sizes for part 1 and 2
	for _, v := range directory_map {
		size := FolderSize(&v)
		if size <= 100000 {
			sum1 += size
		}
		if (disk_size-total)+size >= space_needed && size < smallest {
			smallest = size
		}
	}

	// results
	fmt.Println("Sum of directories under 100000 size:", sum1)
	fmt.Println("Size of smallest directory to delete:", smallest)
}
