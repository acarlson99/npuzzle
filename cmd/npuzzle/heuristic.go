package main

// heuristic functions compare two states and return H value of first param
func ManhattanDist(state *State, end *State) int {
	return (0)
}

func Conflict(state *State, end *State) int {
	return (0)
}

func RightPlace(state *State, end *State) int {
	total := 0
	for ii, _ := range state.Board {
		if state.Board[ii] != end.Board[ii] {
			total += 1
		}
	}
	return total
}
