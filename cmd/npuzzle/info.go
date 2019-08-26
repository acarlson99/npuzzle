package main

import "fmt"

type Info struct { // TODO: rename.  Stores info about search
	Num_opened int // total num of opened states
	Num_closed int // size of closed states
	Max_opened int
	Max_closed int
	Start      *State // beginning state
	Goal       *State // goal state
	End        *State // ptr to final state. nil if unsolvable
}

func (info *Info) PrintInfo() {
	fmt.Println("")
	if info == nil {
		fmt.Println("No information to print")
		return
	}
	if info.End != nil {
		fmt.Println("\nNumber of parents:", info.End.PrintParents()) // TODO: remove this print
		fmt.Println("Length of solution:", info.End.G)
	} else {
		fmt.Println("Unsolvable")
	}
	fmt.Println("Number of opened states:", info.Num_opened)
	fmt.Println("Number of closed states:", info.Num_closed)
	fmt.Println("Max number of opened states:", info.Max_opened)
	fmt.Println("Max number of closed states:", info.Max_closed)
}
