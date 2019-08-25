package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"strconv"
)

type State struct {
	F      int // f score
	G      int // nodes traversed from start to current node
	H      int // distance from goal
	Board  []uint
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
		buffer.WriteString(string(strconv.AppendUint([]byte(nil), uint64(num), 10)))
		// buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(",")
	}
	f := fnv.New64a()
	f.Write([]byte(buffer.String()))
	return f.Sum64()
}

func (state *State) Init(board []uint, emptyX, emptyY, size int) {
	state.Board = make([]uint, size*size)
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
	newState.Board = make([]uint, state.Size*state.Size)
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
	if state == nil || state.EmptyX+x < 0 || state.EmptyX+x >= state.Size ||
		state.EmptyY+y < 0 || state.EmptyY+y >= state.Size {
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
}

// PRINTING
func (state *State) PrintBoardWidth(width uint) uint {
	for ii := 0; ii < state.Size; ii += 1 {
		fmt.Printf("[")
		for jj := 0; jj < state.Size; jj += 1 {
			if jj != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%-*d", width, state.Board[(ii*state.Size)+jj])
		}
		fmt.Println("]")
	}
	return uint(2) + uint(state.Size-1) + (width * uint(state.Size))
}

// format to size (ceiling (log (max 1 num) 10)))
// returns width of board printed
func (state *State) PrintBoard() uint {
	width := BoardNumWidth(state.Board)
	return state.PrintBoardWidth(width)
}

func (state *State) PrintParentsWidth(width uint) uint {
	if state.Parent != nil {
		offset := state.Parent.PrintParentsWidth(width)/2 + 1
		fmt.Printf("%*c\n%*c\n", offset, '|', offset, 'V')
	}
	return state.PrintBoardWidth(width)
}

func (state *State) PrintParents() uint {
	width := BoardNumWidth(state.Board)
	state.PrintParentsWidth(width)
	return uint(2) + uint(state.Size-1) + (width * uint(state.Size))
}
