package main

import (
	"fmt"
	"log"
	"os"
)

type State struct {
	F      int // f score
	G      int // nodes traversed from start to current node
	H      int // distance from goal
	Board  []int
	Hash   int
	EmptyX int
	EmptyY int
	Size   int
	Parent *State
}

var (
	calcH func(*State) int
)

func SetHCalc(f func(*State) int) {
	calcH = f
}

func (state *State) CalcH() int {
	if calcH == nil {
		ll := log.New(os.Stderr, "", 0)
		ll.Println("calcH has not been set")
		os.Exit(1)
	}
	return (calcH(state))
}

func (state *State) CalcHash() int {
	// TODO: implement hash
	var hash int
	for ii, val := range state.Board {
		hash += ii * val
	}
	return hash
}

func (state *State) PrintBoard() {
	for ii := 0; ii < state.Size; ii += 1 {
		fmt.Println(state.Board[ii*state.Size : ii*state.Size+state.Size])
	}
}

// TODO: make this less expensive
func (state *State) Init(board []int, emptyX, emptyY, size int) {
	state.Board = make([]int, size*size)
	copy(state.Board, board)
	state.EmptyX = emptyX
	state.EmptyY = emptyY
	state.Size = size
	state.Parent = nil
	state.Hash = state.CalcHash()
	state.G = 0
	state.H = state.CalcH()
	state.F = state.G + state.H
}

// TODO: maybe make hashmap to hold children.  Return child if map contains it
func (state *State) MakeChild() *State {
	newState := new(State)
	newState.Init(state.Board, state.EmptyX, state.EmptyY, state.Size)
	newState.Parent = state
	newState.G = state.G + 1
	newState.H = state.H
	newState.F = newState.G + newState.H
	return (newState)
}

func (state *State) shiftTile(x, y int) *State {
	if state == nil || state.EmptyX+x < 0 || state.EmptyX+x >= state.Size || state.EmptyY+y < 0 || state.EmptyY+y >= state.Size {
		return (nil)
	}
	newState := state.MakeChild()
	newState.EmptyX += x
	newState.EmptyY += y

	newEmptyIdx := (newState.EmptyY * newState.Size) + newState.EmptyX
	oldEmptyIdx := (state.EmptyY * state.Size) + state.EmptyX

	emptyVal := state.Board[oldEmptyIdx]
	fmt.Printf("MAKING CHILD\n%+v\n", newState)
	fmt.Printf("%+v\n", state)
	fmt.Println(newEmptyIdx)
	fmt.Println(oldEmptyIdx)
	newState.Board[oldEmptyIdx] = state.Board[newEmptyIdx]
	newState.Board[newEmptyIdx] = emptyVal

	newState.H = newState.CalcH()
	newState.F = newState.G + newState.H
	newState.Hash = newState.CalcHash()

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
	str := fmt.Sprintf("%+v", state)
	return str
	// str := fmt.Sprintf("&{Score:%d Board:%v EmptyX:%d EmptyY:%d Size:%d", state.Score, state.Board, state.EmptyX, state.EmptyY, state.Size)
	// if state.Parent != nil {
	// 	return (fmt.Sprintf("%s Parent:%s}", str, state.Parent.ToStr()))
	// } else {
	// 	return (fmt.Sprintf("%s Parent:%v}", str, state.Parent))
	// }
}
