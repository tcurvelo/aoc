package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	input, cleanup := inputScanner()
	defer cleanup()

	acc := 0
	wordLen := 12

	for input.Scan() {
		row := input.Text()
		word := make([]byte, wordLen)
		lastDigitPos := -1
		startIdx := len(row) - wordLen

		for i := startIdx; i < len(row); i++ {
			maxChar := row[i]
			maxPos := i

			for j := i; j > lastDigitPos; j-- {
				if row[j] >= maxChar {
					maxChar = row[j]
					maxPos = j
				}
			}

			word[i-startIdx] = maxChar
			lastDigitPos = maxPos
		}

		fmt.Println("Max pair:", row, string(word))
		total, _ := strconv.Atoi(string(word))
		acc += total
	}
	fmt.Println("Total:", acc)
}
