package challenges

import (
	"aoc2025/shared"
	"fmt"
	"slices"
	"strings"
	"sync"
)

func d7traverseMemo() func(charMatrix []byte, idx *shared.Indexing, i int, j int) int {
	cache := make(map[[2]int]int)
	var mu sync.RWMutex

	var fn func(charMatrix []byte, idx *shared.Indexing, i int, j int) int
	fn = func(charMatrix []byte, idx *shared.Indexing, i int, j int) int {
		if i == idx[0]-1 {
			return 1
		}

		mu.RLock()
		v, ok := cache[[2]int{i, j}]
		mu.RUnlock()
		if ok {
			return v
		}

		var retVal int
		if charMatrix[idx.To1D(i+1, j)] == 94 {
			retVal = fn(charMatrix, idx, i+1, j-1) + fn(charMatrix, idx, i+1, j+1)
		} else {
			retVal = fn(charMatrix, idx, i+1, j)
		}

		mu.Lock()
		cache[[2]int{i + 1, j}] = retVal
		mu.Unlock()

		return retVal
	}

	return fn
}

func d7SolvePart2(charMatrix []byte, idx *shared.Indexing, initPos [2]int) {
	fn := d7traverseMemo()
	res := fn(charMatrix, idx, initPos[0], initPos[1])
	fmt.Printf("Part %d: %d\n", 2, res)
}

func Day7(useExample bool, part int) {
	var idx shared.Indexing

	part2 := part == 2
	lines := strings.Split(shared.GetStringForDay(7, useExample), "\n")
	numRows := len(lines)
	numColumns := len(lines[0])
	idx[0] = numRows
	idx[1] = numColumns
	charMatrix := make([]byte, idx.N())

	for i := range numRows {
		copy(charMatrix[i*numColumns:(i+1)*numColumns], lines[i])
	}

	sPos := strings.Index(lines[0], "S")
	lastBeamsPos := [][2]int{{1, sPos}}
	// Usually i try to adapt part 1 and 2 to be all on the same function
	// I won't try this today
	if part2 {
		d7SolvePart2(charMatrix, &idx, [2]int{1, sPos})
		return
	}

	allLastBeamPosEnded := false

	totalSplits := 0
	for !allLastBeamPosEnded {
		newBeamPos := make([][2]int, 0, numColumns)
		for b := range lastBeamsPos {
			beam := lastBeamsPos[b]
			if charMatrix[idx.To1D(beam[0]+1, beam[1])] == 94 {
				totalSplits++
				if slices.Index(newBeamPos, [2]int{beam[0] + 1, beam[1] - 1}) == -1 {
					newBeamPos = append(newBeamPos, [2]int{beam[0] + 1, beam[1] - 1})
				}
				if slices.Index(newBeamPos, [2]int{beam[0] + 1, beam[1] + 1}) == -1 {
					newBeamPos = append(newBeamPos, [2]int{beam[0] + 1, beam[1] + 1})
				}
			} else {
				if slices.Index(newBeamPos, [2]int{beam[0] + 1, beam[1]}) == -1 {
					newBeamPos = append(newBeamPos, [2]int{beam[0] + 1, beam[1]})
				}
			}
		}

		for b := range newBeamPos {
			beam := newBeamPos[b]
			if beam[0] < numRows-1 {
				allLastBeamPosEnded = false
				break
			} else {
				allLastBeamPosEnded = true
			}
		}
		lastBeamsPos = newBeamPos
	}

	fmt.Printf("Part %d: %d\n", part, totalSplits)
}
