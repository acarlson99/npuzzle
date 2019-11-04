package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	verbose bool
)

type Info struct {
	Num_opened int // total num of opened states
	Num_closed int // size of closed states
	Max_opened int
	Max_closed int
	Start      *State // beginning state
	Goal       *State // goal state
	End        *State // ptr to final state. nil if unsolvable
}

func (info *Info) Print() {
	fmt.Println("")
	if info == nil {
		fmt.Println("No information to print")
		return
	}
	if info.End != nil {
		numParents := info.End.PrintParents()
		fmt.Println("Length of solution:", info.End.Dist)
		if numParents != uint(info.End.Dist) {
			defer log.Fatalln(fmt.Sprint("numParents = %d; info.End.Dist = %d",
				numParents, info.End.Dist))
		}
	} else {
		fmt.Println("Unsolvable")
	}
	fmt.Println("Number of opened states:", info.Num_opened)
	fmt.Println("Number of closed states:", info.Num_closed)
	fmt.Println("Max number of opened states:", info.Max_opened)
	fmt.Println("Max number of closed states:", info.Max_closed)
}

func usage(ret int) {
	fmt.Println("usage: ./npuzzle [options] [startFile]")
	flag.PrintDefaults()
	os.Exit(ret)
}

func main() {
	// handle args
	flag.Usage = func() { usage(1) }

	var goalFile, startFile, heuristic, search string
	var vis bool
	var multiplier float64
	flag.StringVar(&goalFile, "goal", "", "file containing goal state")
	flag.StringVar(&heuristic, "heuristic", "conflict-manhattan", "heuristic function (hamming, manhattan, max, euclidean, conflict, conflict-manhattan)")
	flag.StringVar(&search, "search", "astar", "type of search (uniform, greedy, astar)")
	flag.BoolVar(&verbose, "verbose", false, "verbose search output")
	flag.BoolVar(&vis, "vis", false, "enable visualizer")
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
	var heuristicF func(*State, *State) float64
	switch heuristic {
	case "manhattan":
		heuristicF = ManhattanDist
	case "conflict":
		heuristicF = LinearConflict
	case "hamming":
		heuristicF = HammingDist
	case "max":
		heuristicF = MaxDist
	case "euclidean":
		heuristicF = EuclideanDist
	case "conflict-manhattan":
		heuristicF = ConflictManhattan
	default:
		fmt.Println("Invalid heuristic:", heuristic)
		usage(1)
	}

	var searchF func(*State) float64
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
		fmt.Println("Error:", err)
		fmt.Println("NOTE: good file looks like this")
		fmt.Println(`# I am a comment
3 # size of puzzle
1 2 3 # contents of puzzle
4 5 6
7 8 0 # 0 is empty tile`)
		os.Exit(1)
	}

	fmt.Println("Start:")
	start.PrintBoard()
	fmt.Println("\nGoal:")
	goal.PrintBoard()

	// setup
	SetHCalc(func(state *State) float64 { return heuristicF(state, goal) * multiplier })
	SetScoreCalc(searchF)

	start.CalcValues()
	goal.CalcValues()

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
	if vis && info.End != nil {
		DisplayVis(info)
	}
}
