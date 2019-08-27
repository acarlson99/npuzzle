package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	verbose bool
)

// func (state *State) Solvable() bool {
// 	invCount := 0
// 	for ii := 0; ii < state.Size * state.Size; ii += 1 {
// 		for jj := ii + 1; jj < state.Size * state.Size; jj += 1 {
// 			if state.Board[ii] > state.Board[jj] {
// 				invCount += 1
// 			}
// 		}
// 	}
// 	// return invCount
// 	if len(state.Board) % 2 == 1 {
// 		return invCount % 2 == 0
// 	} else if (((state.EmptyY * state.Size) + state.EmptyX) / state.Size) % 2 == 0 {
// 		return invCount % 2 == 1
// 	} else {
// 		return invCount % 2 == 0
// 	}
// }

func usage(ret int) {
	fmt.Println("usage: ./npuzzle [options] [startFile]")
	flag.PrintDefaults()
	os.Exit(ret)
}

func main() {
	// handle args
	flag.Usage = func() { usage(1) }

	var endFile, startFile, heuristic, search string
	flag.StringVar(&endFile, "end", "", "file containing goal state")
	// flag.StringVar(&startFile, "start", "", "file containing start state")
	flag.StringVar(&heuristic, "heuristic", "atomic", "heuristic function (manhattan, conflict, atomic)")
	flag.StringVar(&search, "search", "astar", "type of search (astar, uniform, greedy)")
	flag.BoolVar(&verbose, "verbose", false, "verbose search output")

	flag.Parse()

	args := flag.Args()
	// fmt.Println(args)

	switch len(args) {
	case 0:
		startFile = ""
	case 1:
		startFile = args[0]
	default:
		usage(1)
	}

	// setup
	var heuristicF func(*State, *State) int
	switch heuristic {
	case "manhattan":
		heuristicF = ManhattanDist
	case "conflict":
		heuristicF = Conflict
	case "atomic":
		heuristicF = Atomic
	default:
		fmt.Println("Invalid heuristic")
		usage(1)
	}

	var searchF func(*State) int
	switch search {
	case "astar":
		searchF = Astar
	case "greedy":
		searchF = Greedy
	case "uniform":
		searchF = Uniform
	default:
		fmt.Println("Invalid search")
		usage(1)
	}

	// read states
	start, end, err := InitStates(startFile, endFile)
	if err != nil {
		fmt.Println(err)
		fmt.Println("NOTE: good file looks like this")
		fmt.Println(`# I am a comment
3 # size of puzzle
1 2 3 # contents of puzzle
4 5 6
7 8 0`)
		os.Exit(1)
	}

	// fmt.Println(start.Solvable())
	// fmt.Println(end.Solvable())

	// if start.Solvable() != end.Solvable() {
	// 	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA NOT SOLVABLE")
	// }

	fmt.Println("")
	fmt.Println("START:")
	start.PrintBoard()
	fmt.Println("")
	fmt.Println("END:")
	end.PrintBoard()

	// setup
	SetHCalc(func(state *State) int { return heuristicF(state, end) })
	SetScoreCalc(searchF)

	// solve
	info := Solve(start, end)

	if info == nil {
		os.Exit(1)
	}

	info.PrintInfo()
	// fmt.Printf("%+v\n", info)
	// fmt.Println(info.End.ToStr())
}
