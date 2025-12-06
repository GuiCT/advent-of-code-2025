package shared

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
