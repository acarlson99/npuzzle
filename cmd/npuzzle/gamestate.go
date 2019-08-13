package main

type Gamestate struct {
	Score  int
	Board  []int
	EmptyX int
	EmptyY int
	Size   int
}

func (g *Gamestate) calcScore() int {
	return (0)
}

func (g *Gamestate) Init(board []int, emptyX, emptyY, size int) {
	g.Score = g.calcScore()
	g.Board = make([]int, len(board))
	copy(g.Board, board)
	g.EmptyX = emptyX
	g.EmptyY = emptyY
	g.Size = size
	g.Size = size
}

func (g *Gamestate) Copy() *Gamestate {
	newState := new(Gamestate)
	newState.Init(g.Board, g.EmptyX, g.EmptyY, g.Size)
	return (newState)
}

func (g *Gamestate) shiftTile(x, y int) *Gamestate {
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

func (g *Gamestate) MoveUp() *Gamestate {
	return (g.shiftTile(0, -1))
	// if g.EmptyY == 0 {
	// 	return (nil)
	// }

	// newState := g.Copy()
	// newState.EmptyY -= 1

	// newEmpty := (newState.EmptyY * newState.Size) + newState.EmptyX
	// gEmpty := (g.EmptyY * g.Size) + g.EmptyX

	// emptyVal := g.Board[gEmpty]
	// newState.Board[gEmpty] = g.Board[newEmpty]
	// newState.Board[newEmpty] = emptyVal

	// return (newState)
}

func (g *Gamestate) MoveDown() *Gamestate {
	return (g.shiftTile(0, 1))
}

func (g *Gamestate) MoveLeft() *Gamestate {
	return (g.shiftTile(-1, 0))
}

func (g *Gamestate) MoveRight() *Gamestate {
	return (g.shiftTile(1, 0))
}
