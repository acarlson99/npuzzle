package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

type Container struct { // TODO: rename.  Stores info about search
	Num_opened     int    // total num of opened states
	Num_closed     int    // size of closed states
	Max_num_states int    // max number of states ever stored in memory
	Num_moves      int    // initial to final
	Start          *State // beginning state
	Goal           *State // goal state
	End            *State // ptr to final state. nil if unsolvable
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ./npuzzle [inputfile]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	fmt.Println(args)
	g := new(State)
	g.Init([]int{1, 2, 3, 4, 0, 5, 6, 7, 8, 9, 10, 11}, 0, 1, 4)

	SetScoreCalc(ManhattanDist)

	state := g.MoveUp()
	// open_states.Push(state, state.Score)
	fmt.Println(g)
	fmt.Println(state)
	sndState := state.MoveUp()
	fmt.Println(sndState)
	c := state.MoveLeft()
	d := state.MoveRight()
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(d.ToStr())
	fmt.Printf("%+v\n", d)
}

func solvePuzzle(cont *Container) {
	open_states := prque.New()
	closed_states := make(map[float32]State)

	open_states.Push(cont.Start, cont.Start.Score)

	for !open_states.Empty() {
		newState, val := open_states.Pop()

	}
}
