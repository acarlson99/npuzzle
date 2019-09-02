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

	var goalFile, startFile, heuristic, search string
	var visu bool
	var multiplier float64
	flag.StringVar(&goalFile, "goal", "", "file containing goal state")
	// flag.StringVar(&startFile, "start", "", "file containing start state")
	flag.StringVar(&heuristic, "heuristic", "max", "heuristic function (max, hamming, manhattan, conflict)")
	flag.StringVar(&search, "search", "astar", "type of search (astar, uniform, greedy)")
	flag.BoolVar(&verbose, "verbose", false, "verbose search output")
	flag.BoolVar(&visu, "visu", false, "enable visu")
	flag.Float64Var(&multiplier, "multiplier", 1.0, "heuristic multiplier")

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
		heuristicF = LinearConflict
	case "hamming":
		heuristicF = HammingDist
	case "max":
		heuristicF = MaxDist
	default:
		fmt.Println("Invalid heuristic:", heuristic)
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
		fmt.Println("Invalid search:", search)
		usage(1)
	}

	// read states
	start, goal, err := InitStates(startFile, goalFile)
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
	SetHCalc(func(state *State) int { return int(float64(heuristicF(state, goal)) * multiplier) })
	SetScoreCalc(searchF)

	start.CalcValues()
	goal.CalcValues()

	fmt.Println("")
	fmt.Println("START:")
	start.PrintBoard()
	fmt.Println("")
	fmt.Println("GOAL:")
	goal.PrintBoard()

	if start.Solvable() != goal.Solvable() {
		fmt.Println("WARNING: STATE NOT SOLVABLE")
	}

	// solve
	startT := time.Now()
	info, err := Solve(start, goal)
	endT := time.Now()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	info.Print()
	fmt.Println("Time:", endT.Sub(startT))
	if visu && info.End != nil {
		DisplayVisu(info)
	}
}
