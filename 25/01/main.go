package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	re := regexp.MustCompile("([A-Z])([0-9]+)")
	dial := 50
	passwordP1 := 0
	passwordP2 := 0

	for input.Scan() {
		direction := 1
		row := input.Text()
		matches := re.FindStringSubmatch(row)

		if matches[1] == "L" {
			direction = -1
		}

		distance, _ := strconv.Atoi(matches[2])

		zeros := distance / 100
		distance = distance % 100

		dialBefore := dial
		dial += direction * distance
		if dialBefore < 0 && dial >= 0 {
			zeros++
		}
		if dialBefore > 0 && dial <= 0 || dial >= 100 {
			zeros++
		}

		dial = (dial + 100) % 100
		fmt.Printf("The dial is rotated %s to point at %d", matches[0], dial)
		if zeros > 0 {
			fmt.Printf("; during this rotation, it points at 0 a few times (%d).", zeros)
			passwordP2 += zeros
		}
		fmt.Println()

		if dial == 0 {
			passwordP1++
		}
	}

	fmt.Println(passwordP1)
	fmt.Println(passwordP2)
}
