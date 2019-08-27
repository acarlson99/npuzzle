package main

// breadth first search
func Uniform(state *State) int {
	return state.Dist
}

// heuristic depth first search
func Greedy(state *State) int {
	return state.H
}

// CCC-Combo Breaker!!
func Astar(state *State) int {
	return state.Dist + state.H
}
