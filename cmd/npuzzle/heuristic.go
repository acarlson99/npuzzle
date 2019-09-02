package main

import (
	"fmt"
	"math"
)

// heuristic functions compare two states and return H value of first param
func ManhattanDist(state *State, goal *State) int {
	total := 0
	for ii := range state.Board {
		if ii/state.Size == 0 || ii%state.Size == 0 {
			n, err := state.FindN(goal.Board[ii])
			if err != nil {
				panic(fmt.Sprintf("%d found in goal, but not start", goal.Board[ii])) // TODO: address panic
			}
			total += int(math.Abs(float64(ii%state.Size-n%state.Size)) +
				math.Abs(float64(ii/state.Size-n/state.Size)))
		}
	}
	return total
}

func HammingDist(state *State, goal *State) int {
	total := 0
	for ii := range state.Board {
		if state.Board[ii] != goal.Board[ii] {
			total += 1
		}
	}
	return total
}

func MaxDist(state *State, goal *State) int {
	return int(math.Max(float64(HammingDist(state, goal)), float64(ManhattanDist(state, goal))))
}

func LinearConflict(state *State, goal *State) int {
	total := 0
	for ii := 0; ii < state.Size; ii += 1 {
		for jj := 0; jj < state.Size; jj += 1 {
			// ns := state.Idx(ii, jj)
			// ne := state.FindN(ns)
			total += 1
		}
	}
	return total
}
