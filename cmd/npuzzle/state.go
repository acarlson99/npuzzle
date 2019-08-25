package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"hash/fnv"
)

type State struct {
	F      int // f score
	G      int // nodes traversed from start to current node
	H      int // distance from goal
	Board  []int
	Hash   uint64
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

func (state *State) CalcHash() uint64 {
	var buffer bytes.Buffer
	for _, num := range state.Board {
		buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(",")
	}
	f := fnv.New64a()
	f.Write([]byte(buffer.String()))
	return f.Sum64()
}

func (state *State) PrintBoard() {
	for ii := 0; ii < state.Size; ii += 1 {
		fmt.Println(state.Board[ii*state.Size : ii*state.Size+state.Size])
	}
}

func (state *State) PrintParents() {
	if state.Parent != nil {
		state.Parent.PrintParents()
		offset := state.Parent.Size + 1
		fmt.Printf("%*c\n%*c\n", offset, '|', offset, 'V')
	}
	state.PrintBoard()
}

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
func (state *State) copyState() *State {
	newState := new(State)

	// copy over all info
	newState.F = state.F
	newState.G = state.G
	newState.H = state.H
	newState.Board = make([]int, state.Size*state.Size)
	copy(newState.Board, state.Board)
	newState.Hash = state.Hash
	newState.EmptyX = state.EmptyX
	newState.EmptyY = state.EmptyY
	newState.Size = state.Size
	newState.Parent = state.Parent

	// newState.Init(state.Board, state.EmptyX, state.EmptyY, state.Size)
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
	newState := state.copyState()
	newState.EmptyX += x
	newState.EmptyY += y

	newEmptyIdx := (newState.EmptyY * newState.Size) + newState.EmptyX
	oldEmptyIdx := (state.EmptyY * state.Size) + state.EmptyX

	emptyVal := state.Board[oldEmptyIdx]
	newState.Board[oldEmptyIdx] = state.Board[newEmptyIdx]
	newState.Board[newEmptyIdx] = emptyVal

	newState.G += 1
	newState.H = newState.CalcH()
	newState.F = newState.G + newState.H
	newState.Hash = newState.CalcHash()
	newState.Parent = state

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
	if state == nil {
		return fmt.Sprintf("%+v", state)
	}
	return fmt.Sprintf("(%+v %s", state, state.Parent.ToStr())
	// str := fmt.Sprintf("&{Score:%d Board:%v EmptyX:%d EmptyY:%d Size:%d", state.Score, state.Board, state.EmptyX, state.EmptyY, state.Size)
	// if state.Parent != nil {
	// 	return (fmt.Sprintf("%s Parent:%s}", str, state.Parent.ToStr()))
	// } else {
	// 	return (fmt.Sprintf("%s Parent:%v}", str, state.Parent))
	// }
}
