package main

// breadth first search
func Uniform(state *State) float64 {
	return state.Dist
}

// heuristic depth first search
func Greedy(state *State) float64 {
	return state.H
}

// CCC-Combo Breaker!!
func Astar(state *State) float64 {
	return state.Dist + state.H
}
