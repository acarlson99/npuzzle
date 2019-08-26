package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func SolvePuzzle(start, goal *State) *Info {
	info := new(Info)
	if start == nil || goal == nil {
		fmt.Println("Start or Goal is nil")
		return nil
	} else if start.Size != goal.Size {
		fmt.Println("Start is a different size than goal")
		return nil
	}
	open_states := prque.New()
	closed_states := make(map[uint64]*State)

	open_states.Push(start, -float32(start.F))

	for !open_states.Empty() && info.End == nil {
		ii, _ := open_states.Pop()

		state := ii.(*State)
		// fmt.Println("CHECKING")
		// state.PrintBoard()
		// if isfinal
		if state.Hash == goal.Hash {
			// TODO: handle good case
			info.End = state
			// return info
		} else {
			closed_states[state.Hash] = state
			for _, newState := range []*State{state.MoveUp(), state.MoveDown(),
				state.MoveLeft(), state.MoveRight()} {
				if newState == nil {
					// fmt.Println("Newstate nil")
					continue
				} else {
					if closed_states[newState.Hash] != nil {
						// handle conflict here.  Update values if better
						var oldState *State
						if closed_states[newState.Hash] != nil {
							oldState = closed_states[newState.Hash]
						} else {
							oldState = nil
						}
						if newState.F < oldState.F {
							if closed_states[newState.Hash] != nil {
								closed_states[newState.Hash] = nil
							}
							open_states.Push(newState, -float32(newState.F))
						}
						// TODO: check if in open_states as well
						// If in closed_states remove from closed and insert into open_states
					} else {
						// fmt.Println(newState.ToStr())
						open_states.Push(newState, -float32(newState.F))
					}
				}
			}
			// fmt.Println(closed_states[state.Hash].ToStr())
		}

	}
	return info
}
