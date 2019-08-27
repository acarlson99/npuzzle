package main

// heuristic functions compare two states and return H value of first param
func ManhattanDist(state *State, end *State) int {
	// // (math.Abs(Target->getX() - Source->getX()) + math.Abs(Target->getY() - Source->getY()))
	// // Heuristic::DistanceManhattan(new Point(col - row->begin(), row - Values.begin()), Heuristic::GetCoordFromValue(*col));
	// distance := 0
	// for ii = 0; ii < state.Size; ii += 1 {
	// 	for jj = 0; jj < state.Size; jj += 1 {

	// 	}
	// }
	return 0
}

func Conflict(state *State, end *State) int {
	return 0
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
