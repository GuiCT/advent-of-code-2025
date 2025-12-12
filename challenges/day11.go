package challenges

import (
	"aoc2025/shared"
	"fmt"
	"slices"
	"strings"
	"sync"
)

func Day11(useExample bool, part int) {
	var DFS func([]string) int
	nodeMap := make(map[string][]string)

	part2 := part == 2

	var lines []string
	if part2 && useExample {
		// For some reason the example is different across the parts?
		lines = strings.Split(`svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`, "\n")
	} else {
		lines = strings.Split(shared.GetStringForDay(11, useExample), "\n")
	}
	for k := range lines {
		split := strings.Split(lines[k], ":")
		source := split[0]
		endings := strings.Split(strings.TrimSpace(split[1]), " ")
		nodeMap[source] = endings
	}

	cache := make(map[string]int)
	var mu sync.RWMutex
	DFS = func(history []string) int {
		last := history[len(history)-1]
		out := last == "out"
		var condition bool = out
		var hasFFT, hasDAC bool
		hasFFT = slices.Contains(history, "fft")
		hasDAC = slices.Contains(history, "dac")
		condition = condition && hasFFT && hasDAC

		if out {
			if condition {
				return 1
			} else {
				return 0
			}
		}

		var cacheKey string
		if !part2 {
			cacheKey = last
		} else {
			// This should be obvious but i really took a long time to think about it
			// Composite caching yey
			cacheKey = fmt.Sprintf("%s|%t|%t", last, hasFFT, hasDAC)
		}

		mu.RLock()
		v, ok := cache[cacheKey]
		mu.RUnlock()
		if ok {
			return v
		}

		possibilities := nodeMap[last]
		totalSum := 0
		for p := range possibilities {
			totalSum += DFS(append(history, possibilities[p]))
		}

		mu.Lock()
		cache[cacheKey] = totalSum
		mu.Unlock()

		return totalSum
	}

	origin := "you"
	if part2 {
		origin = "svr"
	}
	amount := DFS([]string{origin})
	fmt.Printf("Part %d: %d\n", part, amount)
}
