package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Junction struct {
	X, Y, Z int
}

type Edge struct {
	A, B   int
	Weight float64
}

func (e Edge) String() string {
	return fmt.Sprintf("{%2d -> %2d}: %5.f", e.A, e.B, e.Weight)
}

func (j Junction) String() string {
	return fmt.Sprintf("(%3d, %3d, %3d)", j.X, j.Y, j.Z)
}

func main() {
	// Require at least one argument: the input file path
	if len(os.Args) < 2 {
		panic("Please provide an input file")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Optional: limit of connections to process
	var maxConnections int
	if len(os.Args) >= 3 {
		maxConnections, err = strconv.Atoi(os.Args[2])
		if err != nil {
			maxConnections = 10
		}
	}
	if maxConnections <= 0 {
		maxConnections = math.MaxInt32
	}

	content := strings.Split(string(data), "\n")
	junctions := readJunctions(content)
	distances := getDistances(junctions)

	circuitOf := make([]int, len(junctions))
	circuitSize := make([]int, len(junctions))
	lastCircuit := 0
	for d, c := 0, 0; d < len(distances) && c < maxConnections; d, c = d+1, c+1 {
		edge := distances[d]
		var currentCircuit int

		if circuitOf[edge.A] == 0 && circuitOf[edge.B] == 0 {
			// New circuit
			fmt.Printf("%2d. New circuit %v -> %v (%.f) | ", d, junctions[edge.A], junctions[edge.B], edge.Weight)
			lastCircuit++
			currentCircuit = lastCircuit
			circuitOf[edge.A] = currentCircuit
			circuitOf[edge.B] = currentCircuit
			circuitSize[currentCircuit] += 2

		} else if circuitOf[edge.A] == 0 || circuitOf[edge.B] == 0 {
			// Attach the free node to the existing circuit
			currentCircuit = max(circuitOf[edge.A], circuitOf[edge.B])
			circuitOf[edge.A] = currentCircuit
			circuitOf[edge.B] = currentCircuit
			circuitSize[currentCircuit] += 1
			fmt.Printf("%2d. Connected   %v to %v (%.f) | ", d, junctions[edge.A], junctions[edge.B], edge.Weight)

		} else if circuitOf[edge.A] != circuitOf[edge.B] {
			// Merge two different circuits; prefer merging smaller into larger
			var old, new int
			if circuitSize[circuitOf[edge.A]] > circuitSize[circuitOf[edge.B]] {
				old = circuitOf[edge.B]
				new = circuitOf[edge.A]
			} else {
				old = circuitOf[edge.A]
				new = circuitOf[edge.B]
			}
			for i := 0; i < len(circuitOf); i++ {
				if circuitOf[i] == old {
					circuitOf[i] = new
					circuitSize[old] -= 1
					circuitSize[new] += 1
				}
			}
			currentCircuit = new
			fmt.Printf("%2d. Merging    %v and %v (%.f) | ", d, junctions[edge.A], junctions[edge.B], edge.Weight)
		} else {
			// Already connected within the same circuit
			fmt.Printf("%2d. Skipping    %v -> %v (%.f) | ", d, junctions[edge.A], junctions[edge.B], edge.Weight)
			currentCircuit = circuitOf[edge.A]
		}

		// Early exit if all junctions ended up in the same circuit
		if circuitSize[currentCircuit] >= len(junctions) {
			fmt.Printf("ALL CONNECTED! Last connection between %v and %v (Ans: %d)\n", junctions[edge.A], junctions[edge.B], junctions[edge.A].X*junctions[edge.B].X)
			break
		}
		fmt.Println(countCircuits(circuitOf))
	}

	if result := top3(countCircuits(circuitOf)); result > 0 {
		fmt.Println("Multiplying top 3 circuits:", result)
	}
}

func top3(m map[int]int) int {
	if len(m) < 3 {
		return 0
	}
	counts := make([]int, 0, len(m))
	delete(m, 0) // drop the zero circuit
	for _, v := range m {
		counts = append(counts, v)
	}
	sort.Ints(counts)
	fmt.Println("Final circuits:", counts)
	return counts[len(counts)-1] * counts[len(counts)-2] * counts[len(counts)-3]
}

func countCircuits(nums []int) map[int]int {
	h := make(map[int]int, len(nums))
	for _, n := range nums {
		h[n]++
	}
	return h
}

func readJunctions(content []string) []Junction {
	junctions := make([]Junction, len(content))
	for row := range content {
		var j Junction
		fmt.Sscanf(content[row], "%d,%d,%d", &j.X, &j.Y, &j.Z)
		junctions[row] = j
	}
	return junctions
}

func getDistances(junctions []Junction) []Edge {
	edges := make([]Edge, 0)
	for i := 0; i < len(junctions); i++ {
		for j := i + 1; j < len(junctions); j++ {
			// Compute Euclidean distance for readability
			d := euclideanDistance(junctions[i], junctions[j])
			edges = append(edges, Edge{A: i, B: j, Weight: d})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	return edges
}

func euclideanDistance(a, b Junction) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	dz := float64(a.Z - b.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
