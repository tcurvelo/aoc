package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fileOrStdin() *bufio.Scanner {
	file := os.Stdin
	if len(os.Args) > 1 {
		var err error
		file, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
	return bufio.NewScanner(file)
}

func main() {
	input := fileOrStdin()
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
