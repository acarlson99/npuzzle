package main

import (
	"flag"
	"fmt"
	"os"
)

func usage(ret int) {
	flag.PrintDefaults()
	os.Exit(ret)
}

func main() {
	// handle args
	flag.Usage = func() { usage(1) }

	var endFile, startFile, heuristic, search string
	flag.StringVar(&endFile, "end", "", "file containing goal state")
	flag.StringVar(&startFile, "start", "", "file containing start state")
	flag.StringVar(&heuristic, "heuristic", "manhattan", "heuristic function (manhattan, conflict, right-place)")
	flag.StringVar(&search, "search", "astar", "type of search (astar, uniform, greedy)")

	flag.Parse()

	args := flag.Args()
	fmt.Println(args)

	// setup
	var heuristicF func (*State, *State) int
	switch heuristic {
	case "manhattan":
		heuristicF = ManhattanDist
	case "conflict":
		heuristicF = Conflict
	case "right-place":
		heuristicF = RightPlace
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
