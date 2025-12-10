package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func reduce(nums []int, op func(int, int) int) int {
	// fmt.Println("Reducing:", nums)
	acc := nums[0]
	for i := 1; i < len(nums); i++ {
		acc = op(acc, nums[i])
	}
	return acc
}

func main() {
	if len(os.Args) < 1 {
		panic("Please provide an input file")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	content := strings.Split(string(data), "\n")
	opsRaw := strings.Fields(content[len(content)-1])
	opMap := map[string]func(int, int) int{
		"+": add,
		"*": mul,
	}

	// get ops
	ops := make([]func(int, int) int, len(opsRaw))
	for i, op := range opsRaw {
		ops[i] = opMap[op]
	}
	partOne(content, ops)
	partTwo(content, ops)

}

func partOne(content []string, ops []func(int, int) int) {
	grid := make([][]int, len(content)-1)
	for row := range grid {
		cols := strings.Fields(content[row])
		for col := range cols {
			num, err := strconv.Atoi(cols[col])
			if err != nil {
				panic(err)
			}
			grid[row] = append(grid[row], num)
		}
	}
	acc := 0
	for col := 0; col < len(grid[0]); col++ {
		parts := []int{}
		for row := 0; row < len(grid); row++ {
			parts = append(parts, grid[row][col])
		}
		acc += reduce(parts, ops[col])
	}
	fmt.Println("Result p1:", acc)
}

func partTwo(content []string, ops []func(int, int) int) {
	matrix := make([][]byte, len(content)-1)
	for i := range matrix {
		matrix[i] = []byte(content[i])
	}

	totalCols := len(matrix[0])
	totalRows := len(matrix)

	acc := 0
	opIdx := len(ops) - 1
	parts := []int{}
	for col := totalCols - 1; col >= 0; col-- {
		numRaw := ""
		for row := 0; row < totalRows; row++ {
			numRaw += string(matrix[row][col])
		}
		numRaw = strings.TrimSpace(numRaw)
		num, err := strconv.Atoi(strings.TrimSpace(numRaw))
		if err != nil {
			acc += reduce(parts, ops[opIdx])
			opIdx--
			parts = []int{}
			continue
		}
		parts = append(parts, num)
		// fmt.Printf("Parsed number: %d\n", num)
	}
	acc += reduce(parts, ops[opIdx])
	fmt.Println("Result p2:", acc)
}
