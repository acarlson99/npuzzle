package main

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strconv"
)

type State struct {
	Score  float64 // F score
	Dist   float64 // nodes traversed from start to current node.  G
	H      float64 // heuristic distance from goal
	Board  []int
	Hash   uint64
	Empty  int
	Size   int
	Parent *State
}

var (
	calcH     func(*State) float64
	calcScore func(*State) float64
)

// takes function that compares to goal state
// e.g. SetHCalc(func(state *State) int { return RightPlace(state, goal) })
func SetHCalc(f func(*State) float64) {
	calcH = f
}

func SetScoreCalc(f func(*State) float64) {
	calcScore = f
}

func (state *State) CalcHash() uint64 {
	var buffer bytes.Buffer
	for _, num := range state.Board {
		buffer.WriteString(string(strconv.AppendUint([]byte(nil), uint64(num), 10)))
		buffer.WriteString(",")
	}
	f := fnv.New64a()
	f.Write([]byte(buffer.String()))
	return f.Sum64()
}

func (state *State) SoftInit(board []int, empty, size int) {
	state.Board = make([]int, size*size)
	copy(state.Board, board)
	state.Empty = empty
	state.Size = size
	state.Parent = nil
	state.Hash = 0
	state.Dist = 0
	state.H = 0
	state.Score = 0
}

func (state *State) Init(board []int, empty, size int) {
	state.Board = make([]int, size*size)
	copy(state.Board, board)
	state.Empty = empty
	state.Size = size
	state.Parent = nil
	state.Hash = state.CalcHash()
	state.Dist = 0
	state.H = calcH(state)
	state.Score = calcScore(state)
}

func (state *State) CalcValues() {
	state.Hash = state.CalcHash()
	state.H = calcH(state)
	state.Score = calcScore(state)
}

func (state *State) Solvable() bool {
	invCount := 0
	for ii := 0; ii < state.Size*state.Size; ii += 1 {
		for jj := ii + 1; jj < state.Size*state.Size; jj += 1 {
			if state.Board[ii] != 0 && state.Board[jj] != 0 && state.Board[ii] > state.Board[jj] {
				invCount += 1
			}
		}
	}
	if state.Size%2 == 1 {
		return invCount%2 == 0
	} else if (state.Empty/state.Size)%2 == 0 {
		return invCount%2 == 1
	} else {
		return invCount%2 == 0
	}
}

// return copy of state
func (state *State) copyState() *State {
	newState := new(State)

	// copy over all info
	newState.Score = state.Score
	newState.Dist = state.Dist
	newState.H = state.H
	newState.Board = make([]int, state.Size*state.Size)
	copy(newState.Board, state.Board)
	newState.Hash = state.Hash
	newState.Empty = state.Empty
	newState.Size = state.Size
	newState.Parent = state.Parent

	newState.Parent = state
	newState.Dist = state.Dist
	newState.H = state.H
	newState.Score = calcScore(newState)
	return (newState)
}

func (state *State) Idx(x, y int) int {
	return state.Board[(y*state.Size)+x]
}

func (state *State) FindN(n int) (int, error) {
	for ii := 0; ii < state.Size*state.Size; ii += 1 {
		if state.Board[ii] == n {
			return ii, nil
		}
	}
	return 0, fmt.Errorf("%d not in board state", n)
}

func (state *State) shiftTile(x, y int) *State {
	if state == nil || (state.Empty%state.Size)+x < 0 || (state.Empty%state.Size)+x >= state.Size ||
		(state.Empty/state.Size)+y < 0 || (state.Empty/state.Size)+y >= state.Size {
		return (nil)
	}
	newState := state.copyState()
	newState.Empty += x + (y * state.Size)

	emptyVal := state.Board[state.Empty]
	newState.Board[state.Empty] = state.Board[newState.Empty]
	newState.Board[newState.Empty] = emptyVal

	newState.Dist += 1
	newState.H = calcH(newState)
	newState.Score = calcScore(newState)
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
	return fmt.Sprintf("%+v\n%s", state, state.Parent.ToStr())
}

// BOARD PRINTING
func (state *State) PrintBoardWidth(width uint) uint {
	for ii := 0; ii < state.Size; ii += 1 {
		fmt.Printf("[")
		for jj := 0; jj < state.Size; jj += 1 {
			if jj != 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%*d", width, state.Board[(ii*state.Size)+jj])
		}
		fmt.Println("]")
	}
	if verbose {
		fmt.Printf("%+v\n", state)
	}
	return uint(2) + uint(state.Size-1) + (width * uint(state.Size))
}

// returns width of board printed
func (state *State) PrintBoard() uint {
	width := BoardNumWidth(state.Board)
	return state.PrintBoardWidth(width)
}

func (state *State) PrintParentsWidth(width uint) (uint, uint) {
	var depth uint
	if state.Parent != nil {
		offset, numParents := state.Parent.PrintParentsWidth(width)
		offset = offset/2 + 1
		depth = numParents + 1
		fmt.Printf("%*c\n%*c\n", offset, '|', offset, 'V')
	}
	if verbose {
		fmt.Println(depth)
	}
	return state.PrintBoardWidth(width), depth
}

// returns number of parents
func (state *State) PrintParents() uint {
	width := BoardNumWidth(state.Board)
	_, numParents := state.PrintParentsWidth(width)
	return numParents
}
