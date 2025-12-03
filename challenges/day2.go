package challenges

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Dumb way for now, maybe think of a better approach?
func getInvalids(start int, end int) []int {
	invalids := []int{}
	for i := start; i <= end; i++ {
		i_str := strconv.Itoa(i)
		median := len(i_str) / 2
		before := i_str[:median]
		after := i_str[median:]
		if before == after {
			invalids = append(invalids, i)
		}
	}
	return invalids
}

func getInvalidsPart2(start int, end int) []int {
	invalids := []int{}
	for i := start; i <= end; i++ {
		i_str := strconv.Itoa(i)
		str_len := len(i_str)
		median := str_len / 2
		for j := 1; j <= median; j++ {
			if str_len%j > 0 {
				continue
			}
			seq := i_str[:j]
			target, err := strconv.Atoi(strings.Repeat(seq, str_len/j))
			if err != nil {
				panic(err)
			}
			if i == target {
				invalids = append(invalids, i)
				break
			}
		}
	}
	return invalids
}

func Day2(use_example bool, part int) {
	var the_path string
	var total_invalids []int
	part2 := part == 2

	if use_example {
		the_path = "examples/day2.txt"
	} else {
		the_path = "inputs/day2.txt"
	}

	data, err := os.ReadFile(the_path)
	if err != nil {
		panic(err)
	}
	intervals := strings.Split(string(data), ",")
	r := regexp.MustCompile(`(?P<Start>\d+)-(?P<End>\d+)`)
	total_sum := 0
	for i := range intervals {
		interval := intervals[i]
		m := r.FindStringSubmatch(interval)
		start, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		if part2 {
			total_invalids = getInvalidsPart2(start, end)
		} else {
			total_invalids = getInvalids(start, end)
		}
		for s := range total_invalids {
			total_sum += total_invalids[s]
		}
	}
	fmt.Printf("Part %d: %d\n", part, total_sum)
}
