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
	open_states.Push(state, -float32(state.Score))
}

func insertClosed(state *State, closed_states map[uint64]*State, info *Info) {
	closed_states[state.Hash] = state
	info.Num_closed += 1
	if info.Max_closed < info.Num_closed {
		info.Max_closed = info.Num_closed
	}
}

func Solve(start, goal *State) (*Info, error) {
	inc := 100000
	info := new(Info)
	info.Start = start
	info.Goal = goal
	if start == nil || goal == nil {
		return nil, fmt.Errorf("Start or Goal is nil")
	} else if start.Size != goal.Size {

		return nil, fmt.Errorf("Start is a different size than goal")
	}
	open_states := prque.New()
	closed_states := make(map[uint64]*State)

	push(start, open_states, info)

	for !open_states.Empty() && info.End == nil {
		ii, _ := open_states.Pop()
		info.Num_opened -= 1

		state := ii.(*State)
		if verbose {
			inc -= 1
			if inc == 0 {
				inc = 100000
				fmt.Printf("%+v\n", state)
			}
		}
		// if final
		if state.Hash == goal.Hash {
			info.End = state
		} else {
			insertClosed(state, closed_states, info)
			for _, newState := range []*State{state.MoveUp(), state.MoveDown(),
				state.MoveLeft(), state.MoveRight()} {
				if newState == nil || closed_states[newState.Hash] != nil {
					continue
				} else {
					push(newState, open_states, info)
				}
			}
		}

	}
	return info, nil
}
