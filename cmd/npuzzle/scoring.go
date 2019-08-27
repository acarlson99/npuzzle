package main

func Astar(state *State) int {
	return state.Dist + state.H
}

func Uniform(state *State) int {
	return state.Dist
}

func Greedy(state *State) int {
	return state.H
}
