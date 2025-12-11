package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CacheEntry struct {
	Start      string
	VisitedFFT bool
	VisitedDAC bool
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

func Walk(graph map[string][]string, start string, dac bool, fft bool, cache map[CacheEntry]int) int {
	switch start {
	case "out":
		if dac && fft {
			return 1
		}
		return 0
	case "fft":
		fft = true
	case "dac":
		dac = true
	}

	entry := CacheEntry{Start: start, VisitedFFT: fft, VisitedDAC: dac}
	if v, ok := cache[entry]; ok {
		return v
	}
	count := 0
	for _, output := range graph[start] {
		count += Walk(graph, output, dac, fft, cache)
	}
	cache[entry] = count
	return count
}

func main() {
	input, cleanup := inputScanner()
	defer cleanup()

	devices := map[string][]string{}

	for input.Scan() {
		row := input.Text()
		parts := strings.Split(row, ":")
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid input"))
		}
		device := parts[0]
		outputs := strings.Fields(parts[1])
		devices[device] = outputs
	}

	cache1 := map[CacheEntry]int{}
	fmt.Println("Part 2:", Walk(devices, "you", true, true, cache1))
	cache2 := map[CacheEntry]int{}
	fmt.Println("Part 2:", Walk(devices, "svr", false, false, cache2))
}
