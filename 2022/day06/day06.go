package main

import (
	"fmt"
	"log"
	"os"
)

func Unique(window []rune) bool {
	window_map := map[rune]bool{}
	for _, v := range window {
		if window_map[v] {
			return false
		}
		window_map[v] = true
	}
	return true
}

func FindMarker(buffer []byte, window_size int) int {
	window := []rune{}
	marker := -1

	for i, v := range buffer {
		window = append(window, rune(v))
		if len(window) < window_size {
			continue
		}

		if Unique(window) {
			marker = i + 1
			break
		}

		window = window[1:]
	}

	return marker
}

func main() {
	buffer, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First start-of-packet marker:", FindMarker(buffer, 4))
	fmt.Println("First start-of-message marker:", FindMarker(buffer, 14))
}
