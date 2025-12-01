package challenges

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day1(use_example bool) {
	var the_path string
	var next_pointer int
	var amount int

	if use_example {
		the_path = "examples/day1.txt"
	} else {
		the_path = "inputs/day1.txt"
	}

	data, err := os.ReadFile(the_path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	r := regexp.MustCompile(`(?P<Direction>[LR])(?P<Amount>\d+)`)
	current_pointer := 50
	amount_zeros := 0
	amount_zeros_p2 := 0
	for i := range lines {
		passed := false
		line := lines[i]
		if line == "" {
			break
		}
		m := r.FindStringSubmatch(line)
		amount_abs, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		if m[1] == "L" {
			amount = -amount_abs
		} else {
			amount = amount_abs
		}
		amount_zeros_p2 += amount_abs / 100
		next_pointer = current_pointer + (amount % 100)
		if next_pointer > 99 {
			next_pointer -= 100
			passed = true
		} else if next_pointer < 0 {
			next_pointer += 100
			passed = true
		}
		if passed && next_pointer != 0 && current_pointer != 0 {
			amount_zeros_p2++
		}
		if next_pointer == 0 {
			amount_zeros++
			amount_zeros_p2++
		}
		current_pointer = next_pointer
	}
	fmt.Printf("Part 1: %d\nPart2: %d\n", amount_zeros, amount_zeros_p2)
}
