package challenges

import (
	"aoc2025/shared"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day5(useExample bool, part int) {
	part2 := part == 2
	portions := strings.Split(shared.GetStringForDay(5, useExample), "\n\n")
	intervalsStrs := strings.Split(portions[0], "\n")
	amountOfRanges := len(intervalsStrs)
	evalsStrs := strings.Split(portions[1], "\n")
	r := regexp.MustCompile(`(?P<Start>\d+)-(?P<End>\d+)`)

	intervals := make([][2]int, amountOfRanges)
	for i := range intervalsStrs {
		intervalStr := intervalsStrs[i]
		m := r.FindStringSubmatch(intervalStr)
		start, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		intervals[i][0] = start
		intervals[i][1] = end
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := range amountOfRanges - 1 {
		if intervals[i][0] == -1 {
			continue
		}

		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][0] < intervals[i+1][0] {
				intervals[i+1][0] = intervals[i][0]
			}
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
			intervals[i][0] = -1
		}
	}
	validIntervals := make([][2]int, 0, amountOfRanges)
	for i := range amountOfRanges {
		if intervals[i][0] != -1 {
			validIntervals = append(validIntervals, intervals[i])
		}
	}

	if !part2 {
		evals := make([]int, len(evalsStrs))
		for i := range evalsStrs {
			val, err := strconv.Atoi(evalsStrs[i])
			if err != nil {
				panic(err)
			}
			evals[i] = val
		}
		valid := 0
		for i := range evals {
			val := evals[i]
			within := false
			for j := range validIntervals {
				newval := (val >= validIntervals[j][0]) && (val <= validIntervals[j][1])
				within = newval
				if within {
					valid++
					break
				}
			}
		}
		fmt.Printf("Part %d: %d\n", part, valid)
		return
	}

	totalSum := 0
	for i := range validIntervals {
		totalSum += (validIntervals[i][1] - validIntervals[i][0]) + 1
	}
	fmt.Printf("Part %d: %d\n", part, totalSum)
}
