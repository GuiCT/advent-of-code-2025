package shared

type Grid struct {
	Rows     int
	Columns  int
	Elements []uint8
}

func (g *Grid) Initialize(rows int, columns int) {
	g.Rows = rows
	g.Columns = columns
	numberElems := rows * columns
	g.Elements = make([]uint8, numberElems)
}

func (g *Grid) To1D(i int, j int) int {
	return i*g.Rows + j
}

func (g *Grid) To2D(k int) (int, int) {
	return k / g.Rows, k % g.Columns
}

func (g *Grid) Truncate(i int, j int) (int, int) {
	if i < 0 {
		i = 0
	}
	if i > g.Rows-1 {
		i = g.Rows - 1
	}
	if j < 0 {
		j = 0
	}
	if j > g.Columns-1 {
		j = g.Columns - 1
	}
	return i, j
}

func (g *Grid) IsValidPos(i int, j int) bool {
	iInRange := i >= 0 && i < g.Rows
	jInRange := j >= 0 && j < g.Columns
	return iInRange && jInRange
}

func (g *Grid) NeighborsWithDiag(i int, j int) [8]uint8 {
	var vals [8]uint8

	vals[0] = g.Elements[g.To1D(g.Truncate(i-1, j-1))]
	vals[1] = g.Elements[g.To1D(g.Truncate(i-1, j))]
	vals[2] = g.Elements[g.To1D(g.Truncate(i-1, j+1))]
	vals[3] = g.Elements[g.To1D(g.Truncate(i, j-1))]
	vals[4] = g.Elements[g.To1D(g.Truncate(i, j+1))]
	vals[5] = g.Elements[g.To1D(g.Truncate(i+1, j-1))]
	vals[6] = g.Elements[g.To1D(g.Truncate(i+1, j))]
	vals[7] = g.Elements[g.To1D(g.Truncate(i+1, j+1))]

	return vals
}

func (g *Grid) NeighborsWithDiagValid(i int, j int) []uint8 {
	vals := make([]uint8, 0, 8)

	if g.IsValidPos(i-1, j-1) {
		vals = append(vals, g.Elements[g.To1D(i-1, j-1)])
	}
	if g.IsValidPos(i-1, j) {
		vals = append(vals, g.Elements[g.To1D(i-1, j)])
	}
	if g.IsValidPos(i-1, j+1) {
		vals = append(vals, g.Elements[g.To1D(i-1, j+1)])
	}
	if g.IsValidPos(i, j-1) {
		vals = append(vals, g.Elements[g.To1D(i, j-1)])
	}
	if g.IsValidPos(i, j+1) {
		vals = append(vals, g.Elements[g.To1D(i, j+1)])
	}
	if g.IsValidPos(i+1, j-1) {
		vals = append(vals, g.Elements[g.To1D(i+1, j-1)])
	}
	if g.IsValidPos(i+1, j) {
		vals = append(vals, g.Elements[g.To1D(i+1, j)])
	}
	if g.IsValidPos(i+1, j+1) {
		vals = append(vals, g.Elements[g.To1D(i+1, j+1)])
	}

	return vals
}

func (g *Grid) Neighbors(i int, j int) [4]uint8 {
	var vals [4]uint8

	vals[0] = g.Elements[g.To1D(i-1, j)]
	vals[1] = g.Elements[g.To1D(i, j-1)]
	vals[2] = g.Elements[g.To1D(i, j+1)]
	vals[3] = g.Elements[g.To1D(i+1, j)]

	return vals
}
