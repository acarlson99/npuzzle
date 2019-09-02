package main

import (
	"fmt"
	"math"
)

// heuristic functions compare two states and return H value of first param
func ManhattanDist(state, goal *State) float64 {
	var total float64
	total = 0
	for ii := range state.Board {
		if state.Board[ii] != 0 {
			n, err := state.FindN(goal.Board[ii])
			if err != nil {
				panic(fmt.Sprintf("%d found in goal, but not start", goal.Board[ii])) // TODO: address panic
			}
			total += (math.Abs(float64(ii%state.Size-n%state.Size)) +
				math.Abs(float64(ii/state.Size-n/state.Size)))
		}
	}
	return total
}

func HammingDist(state, goal *State) float64 {
	var total float64
	total = 0
	for ii := range state.Board {
		if state.Board[ii] != 0 && state.Board[ii] != goal.Board[ii] {
			total += 1
		}
	}
	return total
}

func MaxDist(state, goal *State) float64 {
	return math.Max(HammingDist(state, goal), ManhattanDist(state, goal))
}

func LinearConflict(state, goal *State) float64 {
	var total float64
	total = 0
	return total
}

func ConflictManhattan(state, goal *State) float64 {
	return ManhattanDist(state, goal) + (LinearConflict(state, goal) * 2)
}
