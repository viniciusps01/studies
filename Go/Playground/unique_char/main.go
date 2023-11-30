package main

import (
	"fmt"
)

const (
	size = 26
)

func main() {
	data := []string{
		"abbbbb",
		"aaaaaab",
		"aaaaaacb",
		"aaaafasdafaadasjalcb",
		"caaaafasdafaadasjjallcbb",
	}

	for _, s := range data {
		fmt.Println(firstUniqueChar(s))
	}
}

func firstUniqueChar(s string) string {
	m := [size]int{}
	indexesInOrder := []int{}

	for _, c := range s {
		if m[c] == nil {
			q := 1
			m[c] = &q

			indexesInOrder = append(indexesInOrder, c)

			continue
		}

		*m[c]++
	}

	for _, k := range indexesInOrder {
		if *m[k] == 1 {
			return string(k)
		}
	}

	return "-"
}
