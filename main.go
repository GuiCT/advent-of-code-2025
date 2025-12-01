package main

import (
	"aoc2025/challenges"
	"flag"
	"fmt"
	"time"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 0, "Specify the day of the challenge to run")
	flag.Parse()
	switch day {
	case 1:
		timeStart := time.Now()
		challenges.Day1(false)
		fmt.Printf("Time elapsed: %.2fms", float64(time.Since(timeStart).Microseconds())/1000)
	}
}
