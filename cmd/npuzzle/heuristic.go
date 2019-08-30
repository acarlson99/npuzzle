package main

import (
	"fmt"
	"math"
)

// heuristic functions compare two states and return H value of first param
func ManhattanDist(state *State, end *State) int {
	total := 0
	for ii, _ := range state.Board {
		if ii/state.Size == 0 || ii%state.Size == 0 {
			n, err := state.FindN(end.Board[ii])
			if err != nil {
				panic(fmt.Sprintf("%d found in end, but not start", end.Board[ii])) // TODO: address panic
			}
			total += int(math.Abs(float64(ii % state.Size - n % state.Size)) +
				math.Abs(float64(ii / state.Size - n / state.Size)))
		}
	}
	return total
}

func Max(state *State, end *State) int {
	total := 0
	for ii, _ := range state.Board {
		if ii / state.Size == 0 || ii % state.Size == 0 {
			total += int(math.Max(float64(Atomic(state, end)), float64(ManhattanDist(state, end))))
		}
	}
	return total
}

func Conflict(state *State, end *State) int {
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

func Atomic(state *State, end *State) int {
	total := 0
	for ii, _ := range state.Board {
		if state.Board[ii] != end.Board[ii] {
			total += 1
		}
	}
	return total
}
