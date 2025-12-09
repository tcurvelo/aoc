package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputScanner() (input *bufio.Scanner, cleanup func()) {
	file := os.Stdin
	if len(os.Args) > 1 {
		var err error
		file, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		return bufio.NewScanner(file), func() { file.Close() }
	}
	return bufio.NewScanner(file), func() {}
}

func checkDirection(grid [][]bool, x, y, dx, dy int) int {
	if x+dx < 0 || x+dx == len(grid) || y+dy < 0 || y+dy == len(grid[x]) {
		return 0
	}
	if grid[x+dx][y+dy] {
		return 1
	}
	return 0
}

type Point struct {
	X int
	Y int
}

func main() {
	input, cleanup := inputScanner()
	defer cleanup()

	currGrid := [][]bool{}
	mapping := map[string]bool{
		".": false,
		"@": true,
	}

	for input.Scan() {
		row := input.Text()
		rowBool := make([]bool, len(row))
		for i, ch := range row {
			rowBool[i] = mapping[string(ch)]
		}
		currGrid = append(currGrid, rowBool)
	}

	removed := 0
	for {
		available := 0
		toRemove := []Point{}
		for i := 0; i < len(currGrid); i++ {
			for j := 0; j < len(currGrid[0]); j++ {
				if currGrid[i][j] {
					adjacents := (checkDirection(currGrid, i, j, -1, -1) + // check up-left
						checkDirection(currGrid, i, j, -1, 0) + // check up
						checkDirection(currGrid, i, j, -1, 1) + // check up-right
						checkDirection(currGrid, i, j, 0, 1) + // check right
						checkDirection(currGrid, i, j, 1, 1) + // check down-right
						checkDirection(currGrid, i, j, 1, 0) + // check down
						checkDirection(currGrid, i, j, 1, -1) + // check down-left
						checkDirection(currGrid, i, j, 0, -1)) // check left
					if adjacents < 4 {
						available++
						toRemove = append(toRemove, Point{X: i, Y: j})
					}
				}
			}
		}

		fmt.Println("Available:", available)

		for _, p := range toRemove {
			currGrid[p.X][p.Y] = false
		}
		removed += available
		if available == 0 {
			break
		}
	}
	fmt.Println("Removed:", removed)
}
