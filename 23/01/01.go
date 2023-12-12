package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lookup map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	accPt1 := 0
	accPt2 := 0
	re1 := regexp.MustCompile(`(\d)`)
	re2 := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		row := input.Text()

		// Part 1
		digitsPt1 := re1.FindAllString(row, -1)
		first := digitsPt1[0]
		last := digitsPt1[len(digitsPt1)-1]

		value, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		accPt1 += value

		// Part 2
		digitsPt2 := []string{}
		loc := []int{}

		// find one match at time
		for i := 0; i < len(row); {
			loc = re2.FindIndex([]byte(row[i:]))
			if loc == nil {
				break
			}
			digitsPt2 = append(digitsPt2, row[i+loc[0]:i+loc[1]])
			i += loc[0] + 1
		}

		first = digitsPt2[0]
		last = digitsPt2[len(digitsPt2)-1]

		if _, err := strconv.Atoi(first); err != nil {
			first = lookup[first]
		}
		if _, err := strconv.Atoi(last); err != nil {
			last = lookup[last]
		}

		value, err = strconv.Atoi(first + last)
		if err != nil {
			panic(err)

		}
		accPt2 += value

	}
	fmt.Println("Part 1:", accPt1)
	fmt.Println("Part 2:", accPt2)
}
