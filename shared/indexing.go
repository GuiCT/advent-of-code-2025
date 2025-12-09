package shared

import "math"

type Indexing [2]int

func (idx *Indexing) N() int {
	return idx[0] * idx[1]
}

func (idx *Indexing) To1D(i int, j int) int {
	if i < 0 {
		i = i % idx[0]
		i = idx[0] + i
	}

	if j < 0 {
		j = j % idx[1]
		j = idx[1] + j
	}

	return idx[1]*i + j
}

func (idx *Indexing) To2D(k int) (int, int) {
	return k % idx[0], k / idx[1]
}

func GetColumn[V any](matrix []V, idx *Indexing, j int) []V {
	ret := make([]V, idx[0])
	for k := range idx[0] {
		ret[k] = matrix[idx.To1D(k, j)]
	}
	return ret
}

// Upper triangular, square matrix indexing
type UTSIndexing [1]int

func (idx *UTSIndexing) N() int {
	return (idx[0] * (idx[0] + 1)) / 2
}

func (idx *UTSIndexing) To1D(i int, j int) int {
	if i < 0 {
		i = i % idx[0]
		i = idx[0] + i
	}

	if j < 0 {
		j = j % idx[0]
		j = idx[0] + j
	}

	if i > j {
		i, j = j, i
	}

	return (j*(j+1))/2 + i
}

func IntSqrt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func (idx *UTSIndexing) To2D(k int) (int, int) {
	j := int(((-1 + IntSqrt(1+8*k)) / 2))
	i := k - (j*(j+1))/2
	return i, j
}
