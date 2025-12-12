package challenges

import (
	"aoc2025/shared"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var d8regexnums = regexp.MustCompile(`(\d+),(\d+),(\d+)`)

func Day8(useExample bool, part int) {
	var posIdx shared.Indexing
	var distIdx shared.UTSIndexing
	var numConnections int

	part1 := part == 1
	part2 := part == 2
	if useExample {
		numConnections = 10
	} else {
		numConnections = 1000
	}

	lines := strings.Split(shared.GetStringForDay(8, useExample), "\n")
	numRows := len(lines)
	numColumns := 3
	posIdx[0] = numRows
	posIdx[1] = numColumns
	posMatrix := make([]int, posIdx.N())

	// Relation between every single one of the boxes
	distIdx[0] = numRows

	distMatrix := make([]float64, distIdx.N())

	for i := range numRows {
		groups := d8regexnums.FindStringSubmatch(lines[i])
		curr, err := strconv.Atoi(groups[1])
		if err != nil {
			panic(err)
		}
		posMatrix[posIdx.To1D(i, 0)] = curr
		curr, err = strconv.Atoi(groups[2])
		if err != nil {
			panic(err)
		}
		posMatrix[posIdx.To1D(i, 1)] = curr
		curr, err = strconv.Atoi(groups[3])
		if err != nil {
			panic(err)
		}
		posMatrix[posIdx.To1D(i, 2)] = curr
	}

	for k := range distIdx.N() {
		i, j := distIdx.To2D(k)
		xi := posMatrix[posIdx.To1D(i, 0)]
		xj := posMatrix[posIdx.To1D(j, 0)]
		yi := posMatrix[posIdx.To1D(i, 1)]
		yj := posMatrix[posIdx.To1D(j, 1)]
		zi := posMatrix[posIdx.To1D(i, 2)]
		zj := posMatrix[posIdx.To1D(j, 2)]
		distX := math.Pow(float64(xi-xj), 2.)
		distY := math.Pow(float64(yi-yj), 2.)
		distZ := math.Pow(float64(zi-zj), 2.)
		dist := math.Sqrt(distX + distY + distZ)
		distMatrix[distIdx.To1D(i, j)] = dist
	}

	// Matrix is a upper triangular, so distances of same pair are represented only once
	// We will ignore the 0s later.
	argSorted := shared.ArgsortSlice(distMatrix, &distIdx)
	parent := make([]int, numRows)
	size := make([]int, numRows)
	for i := range numRows {
		parent[i] = i
		size[i] = 1
	}

	find := func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}

	union := func(x, y int) bool {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return false
		}

		isYBigger := size[rootY] > size[rootX]

		if isYBigger {
			parent[rootX] = rootY
			size[rootY] += size[rootX]
			size[rootX] = 0
		} else {
			parent[rootY] = rootX
			size[rootX] += size[rootY]
			size[rootY] = 0
		}
		return true
	}

	connectionCount := 0
	circuitCount := numRows
	var result int
	// Ignore indexes from 0 to numRows - 1, they will always be 0.
	// From numRows beyond, connect the machines how many times it was asked to
	for k := numRows; circuitCount > 1; k++ {
		i, j := distIdx.To2D(argSorted[k])
		if i == j {
			continue
		}
		connectionCount++
		if union(i, j) {
			circuitCount--
			if part2 && circuitCount == 1 {
				iX := posMatrix[posIdx.To1D(i, 0)]
				jX := posMatrix[posIdx.To1D(j, 0)]
				result = iX * jX
				break
			}

			if part1 && connectionCount == numConnections {
				// Get sizes of all sets
				sizes := make([]int, 0, numRows)
				seen := make(map[int]bool)
				for n := range numRows {
					root := find(n)
					if !seen[root] {
						sizes = append(sizes, size[root])
						seen[root] = true
					}
				}
				sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
				result = 1
				for k := 0; k < 3 && k < len(sizes); k++ {
					result *= sizes[k]
				}
				break
			}
			if circuitCount == 1 {
				// All connected, can break early
				break
			}
		}
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
