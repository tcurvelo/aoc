package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	total := 0

	for input.Scan() {
		row := input.Text()
		intervals := strings.Split(row, ",")

		for _, interval := range intervals {
			matches := re.FindStringSubmatch(interval)
			min, _ := strconv.Atoi(matches[1])
			max, _ := strconv.Atoi(matches[2])

			for n := min; n <= max; n++ {
				nStr := strconv.Itoa(n)
				if isRepeating(nStr) {
					total += n
					fmt.Printf("Found repeated number (%s) in interval %d to %d \n", nStr, min, max)
				}
			}

		}
	}
	fmt.Printf("Total: %d\n", total)
}

func isRepeating(s string) bool {
	n := len(s)
	for patternLen := 1; patternLen <= n/2; patternLen++ {
		if n%patternLen != 0 {
			continue
		}
		pattern := s[:patternLen]
		matches := true
		for i := patternLen; i < n; i += patternLen {
			if s[i:i+patternLen] != pattern {
				matches = false
				break
			}
		}
		if matches {
			return true
		}
	}
	return false
}
