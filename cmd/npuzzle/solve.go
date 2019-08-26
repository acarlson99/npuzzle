package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

func push(state *State, open_states *prque.Prque, info *Info) {
	info.Num_opened += 1
	if info.Max_opened < info.Num_opened {
		info.Max_opened = info.Num_opened
	}
	open_states.Push(state, -float32(state.F))
}

func insertClosed(state *State, closed_states map[uint64]*State, info *Info) {
	closed_states[state.Hash] = state
	info.Num_closed += 1
	if info.Max_closed < info.Num_closed {
		info.Max_closed = info.Num_closed
	}
}

// If in closed_states remove from closed and insert into open_states
func checkClosed(state *State, open_states *prque.Prque, closed_states map[uint64]*State, info *Info) {
	// handle conflict here.  Update values if better
	var oldState *State
	if closed_states[state.Hash] != nil {
		oldState = closed_states[state.Hash]
	} else {
		oldState = nil
	}
	if state.F < oldState.F {
		if closed_states[state.Hash] != nil {
			closed_states[state.Hash] = nil
			info.Num_closed -= 1
		}
		push(state, open_states, info)
		// info.Num_opened += 1
	}
	// TODO: check if in open_states as well
}

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

	push(start, open_states, info)

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
			insertClosed(state, closed_states, info)
			for _, newState := range []*State{state.MoveUp(), state.MoveDown(),
				state.MoveLeft(), state.MoveRight()} {
				if newState == nil {
					// fmt.Println("Newstate nil")
					continue
				} else {
					if closed_states[newState.Hash] != nil {
						checkClosed(newState, open_states, closed_states, info)
					} else {
						// fmt.Println(newState.ToStr())
						push(newState, open_states, info)
					}
				}
			}
			// fmt.Println(closed_states[state.Hash].ToStr())
		}

	}
	return info
}
