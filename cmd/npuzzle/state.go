package main

import "fmt"

type State struct {
	Score  int
	Board  []int
	EmptyX int
	EmptyY int
	Size   int
	Parent *State
	Dist   int
}

func (g *State) calcScore() int {
	return (0)
}

func (g *State) Init(board []int, emptyX, emptyY, size int) {
	g.Score = g.calcScore()
	g.Board = make([]int, size*size)
	copy(g.Board, board)
	g.EmptyX = emptyX
	g.EmptyY = emptyY
	g.Size = size
	g.Parent = nil
	g.Dist = 0
}

// TODO: make hashmap to hold children.  Return child if map contains it
func (g *State) MakeChild() *State {
	newState := new(State)
	newState.Init(g.Board, g.EmptyX, g.EmptyY, g.Size)
	newState.Parent = g
	newState.Dist = g.Dist + 1
	return (newState)
}

func (g *State) shiftTile(x, y int) *State {
	if g.EmptyX+x < 0 || g.EmptyX+x > g.Size || g.EmptyY+y < 0 || g.EmptyY+y > g.Size {
		return (nil)
	}
	newState := g.MakeChild()
	newState.EmptyX += x
	newState.EmptyY += y

	newEmpty := (newState.EmptyY * newState.Size) + newState.EmptyX
	gEmpty := (g.EmptyY * g.Size) + g.EmptyX

	emptyVal := g.Board[gEmpty]
	newState.Board[gEmpty] = g.Board[newEmpty]
	newState.Board[newEmpty] = emptyVal

	return (newState)
}

func (g *State) MoveUp() *State {
	return (g.shiftTile(0, -1))
}

func (g *State) MoveDown() *State {
	return (g.shiftTile(0, 1))
}

func (g *State) MoveLeft() *State {
	return (g.shiftTile(-1, 0))
}

func (g *State) MoveRight() *State {
	return (g.shiftTile(1, 0))
}

func (g *State) ToStr() string {
	str := fmt.Sprintf("&{Score:%d Board:%v EmptyX:%d EmptyY:%d Size:%d", g.Score, g.Board, g.EmptyX, g.EmptyY, g.Size)
	if g.Parent != nil {
		return (fmt.Sprintf("%s Parent:%s}", str, g.Parent.ToStr()))
	} else {
		return (fmt.Sprintf("%s Parent:%v}", str, g.Parent))
	}
}
