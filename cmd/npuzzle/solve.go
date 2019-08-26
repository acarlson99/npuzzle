package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func SolvePuzzle(start, goal *State) *Info {
	info := new(Info)
	info.Start = start
	info.Goal = goal
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
	info.Num_opened += 1

	for !open_states.Empty() && info.End == nil {
		ii, _ := open_states.Pop()
		info.Num_opened -= 1

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
			info.Num_closed += 1
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
								info.Num_closed -= 1
							}
							open_states.Push(newState, -float32(newState.F))
							info.Num_opened += 1
						}
						// TODO: check if in open_states as well
						// If in closed_states remove from closed and insert into open_states
					} else {
						// fmt.Println(newState.ToStr())
						open_states.Push(newState, -float32(newState.F))
						info.Num_opened += 1
					}
				}
			}
			// fmt.Println(closed_states[state.Hash].ToStr())
		}

	}
	return info
}
