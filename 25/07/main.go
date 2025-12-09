package main

import (
	"fmt"
	"os"
	"strings"
)

var total int

func downHill(diagram [][]rune, row int, col int, cache map[[2]int]int) int {
	if row >= len(diagram) || col < 0 || col >= len(diagram[row]) {
		return 1
	}

	var paths int
	if v, ok := cache[[2]int{row, col}]; ok {
		return v
	}

	switch diagram[row][col] {
	case '.':
		diagram[row][col] = '|'
		paths = downHill(diagram, row+1, col, cache)
	case '|':
		paths = downHill(diagram, row+1, col, cache)
	case '^':
		total++
		left := downHill(diagram, row, col-1, cache)
		right := downHill(diagram, row, col+1, cache)
		paths = left + right
	default:
		paths = 0
	}

	cache[[2]int{row, col}] = paths
	return paths
}

func main() {
	if len(os.Args) < 1 {
		panic("Please provide an input file")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	diagram := make([][]rune, len(lines))
	for i, line := range lines {
		diagram[i] = []rune(line)
	}

	var p int
	for p = range lines[0] {
		if lines[0][p] == 'S' {
			break
		}
	}

	cache := make(map[[2]int]int)
	paths := downHill(diagram, 1, p, cache)
	for i := 0; i < len(diagram); i++ {
		for j := 0; j < len(diagram[i]); j++ {
			fmt.Printf("%s", string(diagram[i][j]))
		}
		fmt.Println()
	}

	fmt.Println("Total:", total)
	fmt.Println("Total paths:", paths)
}
