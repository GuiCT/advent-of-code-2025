package main

import (
	"aoc2025/challenges"
	"flag"
	"fmt"
	"time"
)

func main() {
	var day int
	var part int
	var useExample bool

	flag.BoolVar(&useExample, "example", false, "Specify if using example or not")
	flag.IntVar(&day, "day", 0, "Specify the day of the challenge to run")
	flag.IntVar(&part, "part", 0, "Specify the part of the challenge to run")
	flag.Parse()
	timeStart := time.Now()
	switch day {
	case 1:
		challenges.Day1(useExample)
	case 2:
		challenges.Day2(useExample, part)
	case 3:
		challenges.Day3(useExample, part)
	case 4:
		challenges.Day4(useExample, part)
	}
	fmt.Printf("Time elapsed: %.2fms", float64(time.Since(timeStart).Microseconds())/1000)
}
