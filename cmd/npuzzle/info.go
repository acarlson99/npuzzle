package main

type Info struct { // TODO: rename.  Stores info about search
	Num_opened     int    // total num of opened states
	Num_closed     int    // size of closed states
	Start          *State // beginning state
	Goal           *State // goal state
	End            *State // ptr to final state. nil if unsolvable
}
