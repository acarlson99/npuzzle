package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	verbose bool
)

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
	flag.StringVar(&heuristic, "heuristic", "manhattan", "heuristic function (manhattan, conflict, atomic, max)")
	flag.StringVar(&search, "search", "astar", "type of search (astar, uniform, greedy)")
	flag.BoolVar(&verbose, "verbose", false, "verbose search output")

	flag.Parse()

	args := flag.Args()

	switch len(args) {
	case 0:
		startFile = ""
	case 1:
		startFile = args[0]
	default:
		usage(1)
	}

	// set functions
	var heuristicF func(*State, *State) int
	switch heuristic {
	case "manhattan":
		heuristicF = ManhattanDist
	case "conflict":
		heuristicF = Conflict
	case "atomic":
		heuristicF = Atomic
	case "max":
		heuristicF = Max
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
7 8 0 # 0 is empty tile`)
		os.Exit(1)
	}

	// setup
	SetHCalc(func(state *State) int { return heuristicF(state, end) })
	SetScoreCalc(searchF)

	start.CalcValues()
	end.CalcValues()

	fmt.Println("")
	fmt.Println("START:")
	start.PrintBoard()
	fmt.Println("")
	fmt.Println("END:")
	end.PrintBoard()

	if start.Solvable() != end.Solvable() {
		fmt.Println("WARNING: STATE NOT SOLVABLE")
	}

	// solve
	startT := time.Now()
	info, err := Solve(start, end)
	endT := time.Now()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	info.Print()
	fmt.Println("Time:", endT.Sub(startT))
}
