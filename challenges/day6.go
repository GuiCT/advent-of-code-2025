package challenges

import (
	"aoc2025/shared"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regexnumber = regexp.MustCompile(`\d+`)

func simplifyWhitespace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func isAlignedLeft(pos [][]int) bool {
	curr := pos[0][0]
	for k := range pos {
		if pos[k][0] != curr {
			return false
		}
	}

	return true
}

func parseOperandsColumn(idx *shared.Indexing, operandsStrs []string, operandsPosMatrix [][]int, j int) []int {
	var subIdx shared.Indexing
	var charMatrix []byte

	maxStringLen := 0
	for k := range operandsStrs {
		thisLen := len(operandsStrs[k])
		if maxStringLen < len(operandsStrs[k]) {
			maxStringLen = thisLen
		}
	}

	subIdx[0] = idx[0]
	subIdx[1] = maxStringLen
	charMatrix = make([]byte, subIdx.N())
	toRight := isAlignedLeft(shared.GetColumn(operandsPosMatrix, idx, j))

	for k := range subIdx[0] {
		var newStr string
		thisLen := len(operandsStrs[k])
		filling := strings.Repeat(" ", maxStringLen-thisLen)
		if toRight {
			newStr = operandsStrs[k] + filling
		} else {
			newStr = filling + operandsStrs[k]
		}
		copy(charMatrix[k*subIdx[1]:(k+1)*subIdx[1]], []byte(newStr))
	}

	operandsNumbers := make([]int, maxStringLen)
	for j := range maxStringLen {
		currNumber := make([]byte, 0, subIdx[0])
		for i := range subIdx[0] {
			currNumber = append(currNumber, charMatrix[subIdx.To1D(i, j)])
		}
		trimmed := strings.TrimSpace(string(currNumber))
		parsed, err := strconv.Atoi(trimmed)
		if err != nil {
			panic(err)
		}
		operandsNumbers[j] = parsed
	}
	return operandsNumbers
}

func parseOperandsRow(idx *shared.Indexing, operandsStrs []string) []int {
	operandsNumbers := make([]int, idx[0])
	for i := range idx[0] {
		val, err := strconv.Atoi(operandsStrs[i])
		if err != nil {
			panic(err)
		}
		operandsNumbers[i] = val
	}
	return operandsNumbers
}

func Day6(useExample bool, part int) {
	var idx shared.Indexing

	part2 := part == 2
	lines := strings.Split(shared.GetStringForDay(6, useExample), "\n")
	numLines := len(lines)
	numRows := numLines - 1
	operations := strings.Split(simplifyWhitespace(lines[numRows]), " ")
	numColumns := len(operations)
	idx[0] = numRows
	idx[1] = numColumns
	operandsMatrix := make([]string, idx.N())
	operandsPosMatrix := make([][]int, idx.N())

	for i := range numRows {
		pos := regexnumber.FindAllStringIndex(lines[i], -1)
		copy(operandsPosMatrix[i*idx[1]:(i+1)*idx[1]], [][]int(pos))
		simplifiedSplit := strings.Split(simplifyWhitespace(lines[i]), " ")
		for j := range numColumns {
			operandsMatrix[idx.To1D(i, j)] = simplifiedSplit[j]
		}
	}

	total := 0
	for j := range operations {
		var subTotal int
		var operandsNumbers []int

		operandsStrs := shared.GetColumn(operandsMatrix, &idx, j)

		if part2 {
			operandsNumbers = parseOperandsColumn(&idx, operandsStrs, operandsPosMatrix, j)
		} else {
			operandsNumbers = parseOperandsRow(&idx, operandsStrs)
		}

		switch operations[j] {
		case "+":
			subTotal = 0
			for k2 := range operandsNumbers {
				subTotal += operandsNumbers[k2]
			}
		case "*":
			subTotal = 1
			for k2 := range operandsNumbers {
				subTotal *= operandsNumbers[k2]
			}
		}
		total += subTotal
	}

	fmt.Printf("Part %d: %d\n", part, total)
}
