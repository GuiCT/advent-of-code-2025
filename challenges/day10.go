package challenges

import (
	"aoc2025/shared"
	"bytes"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type JoltageMachine struct {
	DesiredState string
	Buttons      [][]int
	Joltage      []int
}

type JoltageStackState struct {
	CurrentState []byte
	StateSize    int
	MachineRef   *JoltageMachine
}

func d10ParseInput(line string) (JoltageMachine, error) {
	var machine JoltageMachine
	regexState := regexp.MustCompile(`\[([\.#]+)\]`)
	regexButton := regexp.MustCompile(`\((.+?)\)`)
	regexJoltage := regexp.MustCompile(`{(.+)}`)

	if !regexState.MatchString(line) {
		return machine, fmt.Errorf("could not parse machine state")
	}
	machine.DesiredState = regexState.FindStringSubmatch(line)[1]

	if !regexButton.MatchString(line) {
		return machine, fmt.Errorf("could not parse machine buttons")
	}
	buttonsStrs := regexButton.FindAllStringSubmatch(line, -1)
	machine.Buttons = make([][]int, len(buttonsStrs))
	for i := range buttonsStrs {
		thisButton := buttonsStrs[i][1]
		intsStrs := strings.Split(thisButton, ",")
		machine.Buttons[i] = make([]int, len(intsStrs))
		for j := range intsStrs {
			val, err := strconv.Atoi(intsStrs[j])
			if err != nil {
				return machine, err
			}
			machine.Buttons[i][j] = val
		}
	}

	if !regexJoltage.MatchString(line) {
		return machine, fmt.Errorf("could not parse machine joltage")
	}
	joltageStr := regexJoltage.FindStringSubmatch(line)[1]
	joltIntStrs := strings.Split(joltageStr, ",")
	machine.Joltage = make([]int, len(joltIntStrs))
	for i := range joltIntStrs {
		val, err := strconv.Atoi(joltIntStrs[i])
		if err != nil {
			return machine, err
		}
		machine.Joltage[i] = val
	}
	return machine, nil
}

func d10Step(s JoltageStackState) []JoltageStackState {
	numPossibilities := len(s.MachineRef.Buttons)
	nextStates := make([]JoltageStackState, numPossibilities)
	for p := range numPossibilities {
		newState := make([]byte, len(s.CurrentState))
		copy(newState, s.CurrentState)
		nextStates[p].StateSize = s.StateSize
		nextStates[p].MachineRef = s.MachineRef
		pressedButtonToggles := nextStates[p].MachineRef.Buttons[p]
		for b := range pressedButtonToggles {
			pos := pressedButtonToggles[b]
			switch newState[pos] {
			case 35:
				newState[pos] = 46
			case 46:
				newState[pos] = 35
			}
		}
		nextStates[p].CurrentState = newState
	}
	return nextStates
}

func Day10(useExample bool) {
	var wg sync.WaitGroup

	lines := strings.Split(shared.GetStringForDay(10, useExample), "\n")
	machines := make([]JoltageMachine, len(lines))
	for i := range lines {
		machine, err := d10ParseInput(lines[i])
		if err != nil {
			panic(err)
		}
		machines[i] = machine
	}

	minPresses := make([]int, len(machines))
	for k := range minPresses {
		wg.Go(func() {
			var initialState JoltageStackState
			var byteCondition []byte
			initialState.MachineRef = &machines[k]
			initialState.StateSize = len(machines[k].DesiredState)
			initialState.CurrentState = bytes.Repeat([]byte{46}, initialState.StateSize)
			byteCondition = []byte(machines[k].DesiredState)
			queue := []JoltageStackState{initialState}
			amountPresses := 0
		iterateQueue:
			for {
				amountPresses++
				newQueueSize := len(queue) * len(initialState.MachineRef.Buttons)
				newQueue := make([]JoltageStackState, 0, newQueueSize)
				for q := range queue {
					toAppend := d10Step(queue[q])
					newQueue = append(newQueue, toAppend...)
				}
				for nq := range newQueue {
					if slices.Equal(newQueue[nq].CurrentState, byteCondition) {
						minPresses[k] = amountPresses
						break iterateQueue
					}
				}
				queue = newQueue
			}
		})
	}
	wg.Wait()

	totalPresses := 0
	for k := range minPresses {
		totalPresses += minPresses[k]
	}

	fmt.Printf("Part %d: %d\n", 1, totalPresses)
}
