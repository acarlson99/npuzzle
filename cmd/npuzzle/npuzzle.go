package main

import (
	"flag"
	"fmt"
	"os"
)

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

	newState := g.MoveUp()
	fmt.Println(g)
	fmt.Println(newState)
	sndState := newState.MoveUp()
	fmt.Println(sndState)
	c := newState.MoveLeft()
	d := newState.MoveRight()
	fmt.Println(c)
	fmt.Println(d)
}
