package shared

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type HasNInfo interface {
	N() int
}

func ArgsortSlice[O constraints.Ordered](data []O, idx HasNInfo) []int {
	indices := make([]int, idx.N())
	for i := range indices {
		indices[i] = i
	}

	sort.SliceStable(indices, func(i, j int) bool {
		return data[indices[i]] < data[indices[j]]
	})

	return indices
}
