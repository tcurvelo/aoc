package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Range struct {
	min int
	max int
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	merged := []Range{ranges[0]}
	for _, curr := range ranges[1:] {
		last := &merged[len(merged)-1]
		if curr.min <= last.max+1 {
			if curr.max > last.max {
				last.max = curr.max
			}
		} else {
			merged = append(merged, curr)
		}
	}
	return merged
}

func main() {
	input, cleanup := inputScanner()
	defer cleanup()
	db := []Range{}
	count := 0

	for input.Scan() {
		row := input.Text()
		if row == "" {
			break
		}
		var r Range
		fmt.Sscanf(row, "%d-%d", &r.min, &r.max)
		db = append(db, r)
	}

	db = mergeRanges(db)

	for input.Scan() {
		row := input.Text()
		var id int
		fmt.Sscanf(row, "%d", &id)

		fresh := false
		for _, r := range db {
			if id >= r.min && id <= r.max {
				fmt.Println("ID", id, "is valid in range", r)
				fresh = true
			}
		}
		if fresh {
			fmt.Println("ID", id, "is fresh")
			count++
		}
	}
	total := 0
	for _, r := range db {
		total += r.max - r.min + 1
	}
	fmt.Println("Total fresh IDs:", count)
	fmt.Println("Total IDs in ranges:", total)
}
