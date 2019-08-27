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
	var f func (*State, *State) int
	switch heuristic {
	case "manhattan":
		f = ManhattanDist
	case "conflict":
		f = Conflict
	case "right-place":
		f = RightPlace
	case "greedy":
		f = Greedy
	case "uniform":
		f = Uniform
	default:
		fmt.Println("Invalid heuristic")
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
	SetHCalc(func(state *State) int { return f(state, end) })

	// solve
	info := AStar(start, end)

	if info == nil {
		os.Exit(1)
	}

	info.PrintInfo()
	// fmt.Printf("%+v\n", info)
	// fmt.Println(info.End.ToStr())
}
