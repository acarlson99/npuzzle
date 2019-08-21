package main

import (
	"fmt"
	"log"
	"os"
)

type State struct {
	Score  float32
	Board  []int
	EmptyX int
	EmptyY int
	Size   int
	Parent *State
	G      int // nodes traversed from start to current node
	H      float32 // distance from goal.  Might change back to int
}

var (
	calcScore func(*State) float32
)

func SetScoreCalc(f func(*State) float32) {
	calcScore = f
}

func (state *State) CalcScore() float32 {
	if calcScore == nil {
		ll := log.New(os.Stderr, "", 0)
		ll.Println("calcScore has not been set\n")
		os.Exit(1)
	}
	return (calcScore(state))
}

func (state *State) Init(board []int, emptyX, emptyY, size int) {
	state.Board = make([]int, size*size)
	copy(state.Board, board)
	state.EmptyX = emptyX
	state.EmptyY = emptyY
	state.Size = size
	state.Parent = nil
	state.G = -1
	state.H = -1
}

// TODO: make hashmap to hold children.  Return child if map contains it
func (state *State) MakeChild() *State {
	newState := new(State)
	newState.Init(state.Board, state.EmptyX, state.EmptyY, state.Size)
	newState.Parent = state
	newState.G = state.G + 1
	newState.H = state.CalcScore()
	return (newState)
}

func (state *State) shiftTile(x, y int) *State {
	if state.EmptyX+x < 0 || state.EmptyX+x > state.Size || state.EmptyY+y < 0 || state.EmptyY+y > state.Size {
		return (nil)
	}
	newState := state.MakeChild()
	newState.EmptyX += x
	newState.EmptyY += y

	newEmpty := (newState.EmptyY * newState.Size) + newState.EmptyX
	gEmpty := (state.EmptyY * state.Size) + state.EmptyX

	emptyVal := state.Board[gEmpty]
	newState.Board[gEmpty] = state.Board[newEmpty]
	newState.Board[newEmpty] = emptyVal

	return (newState)
}

func (state *State) MoveUp() *State {
	return (state.shiftTile(0, -1))
}

func (state *State) MoveDown() *State {
	return (state.shiftTile(0, 1))
}

func (state *State) MoveLeft() *State {
	return (state.shiftTile(-1, 0))
}

func (state *State) MoveRight() *State {
	return (state.shiftTile(1, 0))
}

func (state *State) ToStr() string {
	str := fmt.Sprintf("&{Score:%d Board:%v EmptyX:%d EmptyY:%d Size:%d", state.Score, state.Board, state.EmptyX, state.EmptyY, state.Size)
	if state.Parent != nil {
		return (fmt.Sprintf("%s Parent:%s}", str, state.Parent.ToStr()))
	} else {
		return (fmt.Sprintf("%s Parent:%v}", str, state.Parent))
	}
}
