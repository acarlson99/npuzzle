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

	var endFile, startFile string
	flag.StringVar(&endFile, "end", "", "file containing goal state")
	flag.StringVar(&startFile, "start", "", "file containing start state")

	help := flag.Bool("h", false, "display help message")

	flag.Parse()

	if *help == true {
		usage(0)
	}

	args := flag.Args()
	fmt.Println(args)

	// setup
	SetHCalc(ManhattanDist)

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
	fmt.Println(start.ToStr())
	fmt.Println(end.ToStr())

	// solve
	info := SolvePuzzle(start, end)
	fmt.Println("RETURN: ", info.End.ToStr())

	if info.End != nil {
		info.End.PrintParents()
	} else {
		fmt.Println("Unsolvable")
	}
}
