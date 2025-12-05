package challenges

import (
	"aoc2025/shared"
	"fmt"
	"strings"
)

func hasForkliftAccess(g *shared.Grid, i int, j int) bool {
	if g.Elements[g.To1D(i, j)] != 1 {
		return false
	}

	n := g.NeighborsWithDiagValid(i, j)
	count := 0
	for k := range n {
		if n[k] == 1 {
			count++
		}
	}

	if count < 4 {
		return true
	} else {
		return false
	}
}

func iterateForklift(grid *shared.Grid) int {
	possiblePos := []int{}
	for i := range grid.Rows {
		for j := range grid.Columns {
			as1D := grid.To1D(i, j)
			if hasForkliftAccess(grid, i, j) {
				possiblePos = append(possiblePos, as1D)
			}
		}
	}

	for k := range possiblePos {
		grid.Elements[possiblePos[k]] = 0
	}
	return len(possiblePos)
}

func Day4(useExample bool, part int) {
	part2 := part == 2
	var grid shared.Grid
	var lines = strings.Split(shared.GetStringForDay(4, useExample), "\n")
	var valUint8 uint8
	grid.Initialize(len(lines), len(lines[0]))

	for i := range lines {
		for j := range lines[i] {
			val := string(lines[i][j])
			if val == "." {
				valUint8 = 0
			} else {
				valUint8 = 1
			}
			grid.Elements[grid.To1D(i, j)] = valUint8
		}
	}

	currentCount := iterateForklift(&grid)

	if !part2 {
		fmt.Printf("Part %d: %d\n", part, currentCount)
		return
	}

	totalCount := currentCount
	for currentCount > 0 {
		currentCount = iterateForklift(&grid)
		totalCount += currentCount
	}
	fmt.Printf("Part %d: %d\n", part, totalCount)
}
