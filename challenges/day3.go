package challenges

import (
	"aoc2025/shared"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getVoltage(input string, numDigits int) int {
	var voltage int
	var currIdx int
	maxs := make([]int, numDigits)
	maxRightmostIdx := -1

	for digit := 0; digit < numDigits; digit++ {
		currIdx = maxRightmostIdx + 1
		for ; currIdx < len(input)-(numDigits-digit)+1; currIdx++ {
			val, err := strconv.Atoi(string(input[currIdx]))
			if err != nil {
				panic(err)
			}
			if maxs[digit] < val {
				maxs[digit] = val
				maxRightmostIdx = currIdx
			}
			if val == 9 {
				break
			}
		}
	}

	voltage = 0
	for i := 0; i < numDigits; i++ {
		power := int(math.Pow(float64(10), float64(numDigits-i-1)))
		voltage += power * maxs[i]
	}

	return voltage
}

func Day3(useExample bool, part int) {
	numDigits := 2
	part2 := part == 2

	if part2 {
		numDigits = 12
	}

	var lines = strings.Split(shared.GetStringForDay(3, useExample), "\n")
	totalSum := 0
	for i := range lines {
		v := getVoltage(lines[i], numDigits)
		totalSum += v
	}
	fmt.Printf("Part %d: %d\n", part, totalSum)
}
