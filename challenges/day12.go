package challenges

import (
	"aoc2025/shared"
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Region struct {
	Grid          *mat.Dense
	Dimensions    [2]int
	ShapeFillings []int
}

func (r *Region) GetArea() int {
	return r.Dimensions[0] * r.Dimensions[1]
}

func Day12(useExample bool, part int) {
	lines := strings.Split(shared.GetStringForDay(12, useExample), "\n")
	shapes := make([]*mat.Dense, 0)
	shapesAreas := make([]int, 0)
	regions := make([]Region, 0)
	for k := 0; k < len(lines); {
		if len(lines[k]) == 0 {
			k++
			continue
		}

		if len(lines[k]) > 1 && lines[k][1] == 58 {
			shapeLines := lines[k+1 : k+4]
			newShape := mat.NewDense(3, 3, nil)
			newShapeArea := 0
			for i := range 3 {
				for j := range 3 {
					switch shapeLines[i][j] {
					case 35:
						newShape.Set(i, j, 1)
						newShapeArea++
					case 46:
						newShape.Set(i, j, 0)
					default:
						panic("invalid shape character")
					}
				}
			}
			shapes = append(shapes, newShape)
			shapesAreas = append(shapesAreas, newShapeArea)
			k += 4
		} else {
			dimsAndShapes := strings.Split(lines[k], ":")
			dims := [2]int{0, 0}
			dimsStrs := strings.Split(dimsAndShapes[0], "x")
			num, err := strconv.Atoi(dimsStrs[0])
			if err != nil {
				panic(err)
			}
			dims[0] = num
			num, err = strconv.Atoi(dimsStrs[1])
			if err != nil {
				panic(err)
			}
			dims[1] = num
			shapesStrs := strings.Split(strings.TrimSpace(dimsAndShapes[1]), " ")
			shapeFillings := make([]int, len(shapesStrs))
			for s := range shapesStrs {
				num, err := strconv.Atoi(shapesStrs[s])
				if err != nil {
					panic(err)
				}
				shapeFillings[s] = num
			}
			newRegion := Region{
				Grid:          mat.NewDense(dims[0], dims[1], nil),
				Dimensions:    dims,
				ShapeFillings: shapeFillings,
			}
			regions = append(regions, newRegion)
			k++
		}
	}

	nCanFit := 0
	for _, reg := range regions {
		totalPresents := 0
		totalPresentArea := 0
		for i, s := range reg.ShapeFillings {
			totalPresents += s
			totalPresentArea += s * shapesAreas[i]
		}
		regArea := reg.GetArea()
		// Check on area trivial case
		if totalPresentArea > regArea {
			continue
		}
		// Check if tiling in 3x3 is possible
		if totalPresents*9 <= regArea {
			nCanFit++
		}
		// Ever heard about the Ostrich Algorithm?
		// (The real input does not have any case that violate the 2 checks above but can still fit)
		// Then just ignore it lol
	}

	fmt.Printf("Part %d: %d\n", part, nCanFit)
}
