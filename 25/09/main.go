package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Tile struct {
	X, Y int
}

func (f *Tile) GetStepsTo(to *Tile) (int, int) {
	stepY := to.Y - f.Y
	if stepY != 0 {
		stepY /= int(math.Abs(float64(stepY)))
	}
	stepX := to.X - f.X
	if stepX != 0 {
		stepX /= int(math.Abs(float64(stepX)))
	}
	return stepX, stepY
}

type Square struct {
	From Tile
	To   Tile
	Area int
}

func (s *Square) SetArea() {
	width := int(math.Abs(float64(s.To.X-s.From.X)) + 1)
	height := int(math.Abs(float64(s.To.Y-s.From.Y)) + 1)
	s.Area = width * height
}

func (s *Square) GetSteps() (int, int) {
	return s.From.GetStepsTo(&s.To)
}

func (s *Square) Walk() {
	stepY := s.To.Y - s.From.Y
	if stepY != 0 {
		stepY /= int(math.Abs(float64(stepY)))
	}
	stepX := s.To.X - s.From.X
	if stepX != 0 {
		stepX /= int(math.Abs(float64(stepX)))
	}

}

func (t Tile) String() string {
	return fmt.Sprintf("(%d,%d)", t.X, t.Y)
}

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

func readTiles() ([]Tile, int, int) {
	input, cleanup := inputScanner()
	defer cleanup()
	maxX := 0
	maxY := 0
	reds := make([]Tile, 0)
	for input.Scan() {
		var tile Tile
		row := input.Text()
		fmt.Sscanf(row, "%d,%d", &tile.X, &tile.Y)
		if tile.X > maxX {
			maxX = tile.X
		}
		if tile.Y > maxY {
			maxY = tile.Y
		}
		reds = append(reds, tile)
	}
	return reds, maxX, maxY
}

func printShape(shape map[int]map[int]bool, maxX, maxY int) {
	// fmt.Printf("   ")
	// for x := 0; x <= maxX; x++ {
	// 	fmt.Printf("%3d", x)
	// }
	// fmt.Println()

	for y := 0; y <= maxY; y++ {
		// fmt.Printf("%d: ", y)
		for x := 0; x <= maxX; x++ {
			_, ok := shape[y][x]
			if ok {
				fmt.Printf("%3d", x)
			} else {
				fmt.Printf("  .")
			}
		}
		fmt.Println()
	}
}
func printGrid(grid [][]byte) {
	fmt.Println()
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			b := grid[y][x]
			if b == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", b)
			}
		}
		fmt.Println()
	}
}

func isInside(square Square, grid [][]byte) bool {
	stepX, stepY := square.GetSteps()

	for x := square.From.X; x != square.To.X+stepX; x += stepX {
		for y := square.From.Y; y != square.To.Y+stepY; y += stepY {
			if grid[y][x] != 'X' {
				return false
			}
		}

	}
	return true
}

func main() {
	fmt.Printf("Reading and parsing input...")
	reds, maxX, maxY := readTiles()

	fmt.Printf("[OK]\nCalculating areas...........")
	n := len(reds)
	squares := []Square{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			square := Square{From: reds[i], To: reds[j]}
			square.SetArea()
			squares = append(squares, square)
		}
	}

	fmt.Printf("[OK]\nCreating the grid...........")
	grid := make([][]byte, maxY+1)
	for y := 0; y <= maxY; y++ {
		grid[y] = make([]byte, maxX+1)
	}
	for r := 0; r < n; r++ {
		// ensure the last and first ones connects
		src := reds[r%n]
		dst := reds[(r+1)%n]
		stepX, stepY := src.GetStepsTo(&dst)
		for x, y := src.X, src.Y; x != dst.X+stepX || y != dst.Y+stepY; x, y = x+stepX, y+stepY {
			grid[y][x] = 'X'
		}
	}
	fmt.Printf("[OK]\nFilling the grid............")
	for y := 0; y <= maxY; y++ {
		borderLeft := maxX
		borderRight := 0
		for x := 0; x <= maxX; x++ {
			if grid[y][x] == 'X' {
				if x < borderLeft {
					borderLeft = x
				}
				if x > borderRight {
					borderRight = x
				}
			}
		}
		for x := borderLeft; x <= borderRight; x++ {
			grid[y][x] = 'X'
		}
	}
	fmt.Printf("[OK]\n")

	// sort descending
	sort.Slice(squares, func(i, j int) bool {
		return squares[i].Area > squares[j].Area
	})

	fmt.Println("Searching for largest valid square...")
	for _, square := range squares {
		if isInside(square, grid) {
			fmt.Println("🚀 Largest valid square: ", square)
			break
		} else {
			fmt.Println("❌ Invalid square: ", square)
		}
	}
}
