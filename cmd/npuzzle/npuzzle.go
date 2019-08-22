package main

import (
	"flag"
	"fmt"
	"os"
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

	SetHCalc(ManhattanDist)

	fmt.Println(args)
	g := new(State)
	g.Init([]int{1, 2, 3, 0, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 1, 2)

	state := g.MoveUp()
	fmt.Println("BASE")
	g.PrintBoard()
	fmt.Println(solvePuzzle(g, state).ToStr())
}
