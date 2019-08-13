package main

type State struct {
	Score  int
	Board  []int
	EmptyX int
	EmptyY int
	Size   int
}

func (g *State) calcScore() int {
	return (0)
}

func (g *State) Init(board []int, emptyX, emptyY, size int) {
	g.Score = g.calcScore()
	g.Board = make([]int, len(board))
	copy(g.Board, board)
	g.EmptyX = emptyX
	g.EmptyY = emptyY
	g.Size = size
	g.Size = size
}

func (g *State) Copy() *State {
	newState := new(State)
	newState.Init(g.Board, g.EmptyX, g.EmptyY, g.Size)
	return (newState)
}

func (g *State) shiftTile(x, y int) *State {
	if g.EmptyX+x < 0 || g.EmptyX+x > g.Size || g.EmptyY+y < 0 || g.EmptyY+y > g.Size {
		return (nil)
	}
	newState := g.Copy()
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
