package challenges

import (
	"aoc2025/shared"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var d9regexnums = regexp.MustCompile(`(\d+),(\d+)`)

// I will not pretend i got the raycast algorithm right
// I eventually just gave up and kinda used this guy code as my reference
// https://github.com/amirsina-mandegari/adventofcode2025/blob/main/day09/solution.py
func d9IsPointInside(posMatrix []int, posIdx *shared.Indexing, x, y int) bool {
	hits := 0
	for k := range posIdx[0] {
		v1x := posMatrix[posIdx.To1D(k, 0)]
		v1y := posMatrix[posIdx.To1D(k, 1)]
		v2x := posMatrix[posIdx.To1D((k+1)%posIdx[0], 0)]
		v2y := posMatrix[posIdx.To1D((k+1)%posIdx[0], 1)]

		if v1x == v2x && v2x == x {
			if min(v1y, v2y) <= y && y <= max(v1y, v2y) {
				return true
			}
		} else if v1y == v2y && v2y == y {
			if min(v1x, v2x) <= x && x <= max(v1x, v2x) {
				return true
			}
		}

		if v1x == v2x {
			edgeX := v1x
			edgeMinY := min(v1y, v2y)
			edgeMaxY := max(v1y, v2y)

			if edgeX < x && edgeMinY < y && y <= edgeMaxY {
				hits++
			}
		}
	}
	return hits%2 == 1
}

func d9RectangleHitEdge(posMatrix []int, posIdx *shared.Indexing, k1, k2, minX, maxX, minY, maxY int) bool {
	v1x := posMatrix[posIdx.To1D(k1, 0)]
	v1y := posMatrix[posIdx.To1D(k1, 1)]
	v2x := posMatrix[posIdx.To1D(k2, 0)]
	v2y := posMatrix[posIdx.To1D(k2, 1)]

	if v1x == v2x {
		edgeX := v2x
		edgeMinY := min(v1y, v2y)
		edgeMaxY := max(v1y, v2y)
		withinX := minX < edgeX && edgeX < maxX
		validY := edgeMinY < maxY && edgeMaxY > minY
		if withinX && validY {
			return true
		}
	} else {
		edgeY := v2y
		edgeMinX := min(v1x, v2x)
		edgeMaxX := max(v1x, v2x)
		withinY := minY < edgeY && edgeY < maxY
		validX := edgeMinX < maxX && edgeMaxX > minX
		if withinY && validX {
			return true
		}
	}

	return false
}

func d9RectIsValid(posMatrix []int, posIdx *shared.Indexing, minX, maxX, minY, maxY int) bool {
	validFirst := d9IsPointInside(posMatrix, posIdx, minX, minY)
	validSecond := d9IsPointInside(posMatrix, posIdx, minX, maxY)
	validThird := d9IsPointInside(posMatrix, posIdx, maxX, minY)
	validFourth := d9IsPointInside(posMatrix, posIdx, maxX, maxY)
	cornersValid := validFirst && validSecond && validThird && validFourth
	if !cornersValid {
		return false
	}

	for kEdge := range posIdx[0] {
		k1 := kEdge
		k2 := (kEdge + 1) % posIdx[0]
		if d9RectangleHitEdge(posMatrix, posIdx, k1, k2, minX, maxX, minY, maxY) {
			return false
		}
	}

	return true
}

func d9SimpleArea(x1, y1, x2, y2 int) int {
	xDist := x1 - x2
	if xDist < 0 {
		xDist = -xDist
	}
	xDist++
	yDist := y1 - y2
	if yDist < 0 {
		yDist = -yDist
	}
	yDist++
	return xDist * yDist
}

func d9AreaP2(posMatrix []int, posIdx *shared.Indexing, p1, p2 int) int {
	p1x := posMatrix[posIdx.To1D(p1, 0)]
	p1y := posMatrix[posIdx.To1D(p1, 1)]
	p2x := posMatrix[posIdx.To1D(p2, 0)]
	p2y := posMatrix[posIdx.To1D(p2, 1)]
	minX := min(p1x, p2x)
	maxX := max(p1x, p2x)
	minY := min(p1y, p2y)
	maxY := max(p1y, p2y)

	if d9RectIsValid(posMatrix, posIdx, minX, maxX, minY, maxY) {
		return d9SimpleArea(p1x, p1y, p2x, p2y)
	} else {
		return 0
	}
}

func Day9(useExample bool, part int) {
	part2 := part == 2
	lines := strings.Split(shared.GetStringForDay(9, useExample), "\n")
	var posIdx shared.Indexing
	posIdx[0] = len(lines)
	posIdx[1] = 2
	posMatrix := make([]int, posIdx.N())
	currentMaxArea := 0

	for i := range lines {
		groups := d9regexnums.FindStringSubmatch(lines[i])
		if groups == nil {
			panic(fmt.Errorf("day9: it was not possible to parse the positions at line %d", i))
		}
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
	}

	d9IsPointInside(posMatrix, &posIdx, 9, 2)

	// It is not needed to build the matrix to calculate the area.
	for i := range posIdx[0] {
		x1 := posMatrix[posIdx.To1D(i, 0)]
		y1 := posMatrix[posIdx.To1D(i, 1)]
		for j := i + 1; j < posIdx[0]; j++ {
			x2 := posMatrix[posIdx.To1D(j, 0)]
			y2 := posMatrix[posIdx.To1D(j, 1)]
			var area int
			if part2 {
				area = d9AreaP2(posMatrix, &posIdx, i, j)
			} else {
				area = d9SimpleArea(x1, y1, x2, y2)
			}
			if area > currentMaxArea {
				currentMaxArea = area
			}
		}
	}

	fmt.Printf("Part %d: %d\n", part, currentMaxArea)
}
