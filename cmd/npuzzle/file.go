package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// get start and goal states
func InitStates(startFile, goalFile string) (*State, *State, error) {
	var scanner *bufio.Scanner
	if startFile == "" {
		fmt.Println("No start file.  Reading from stdin...")
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(startFile)
		if err != nil {
			return nil, nil, err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	}
	start, err := ReadStateFromScanner(scanner)
	if err != nil {
		return nil, nil, err
	}
	goal := new(State)
	if goalFile != "" {
		goal, err = ReadStateFromFile(goalFile)
		if err != nil {
			return nil, nil, err
		}
	} else {
		fmt.Println("Goal state unspecified.  Inferring")
		board := make([]int, len(start.Board))
		copy(board, start.Board)
		sort.Slice(board, func(ii, jj int) bool { return board[ii] < board[jj] })
		num := board[0]
		copy(board, board[1:])
		board[((start.Size-1)*start.Size)+start.Size-1] = num
		goal.SoftInit(board, (start.Size*start.Size)-1, start.Size)
	}
	if start.Size != goal.Size {
		return nil, nil,
			fmt.Errorf("start size %d != goal size %d", start.Size, goal.Size)
	}
	return start, goal, nil
}

// read state from file given filename
func ReadStateFromFile(filename string) (*State, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	return ReadStateFromScanner(scanner)
}

// read state from scanner
func ReadStateFromScanner(scanner *bufio.Scanner) (*State, error) {
	scanner.Split(bufio.ScanLines)
	var lines []string

	size := 0

	for size == 0 && scanner.Scan() {
		lines = append(lines, scanner.Text())
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			if len(word) <= 0 {
				continue
			} else if word[0] == '#' {
				break
			} else if size == 0 {
				num, err := strconv.ParseInt(word, 10, 32)
				if err != nil {
					return nil, err
				} else {
					size = int(num)
					if size < 2 {
						return nil,
							fmt.Errorf("Wait a size of %d is ridiculous. Use a size of at least 2", size)
					}
				}
			} else {
				return nil,
					fmt.Errorf("Unexpected token %s.  Expected newline", word)
			}
		}
	}
	if size < 2 {
		return nil,
			fmt.Errorf("Wait a size of %d is ridiculous. Use a size of at least 2", size)
	}

	check := make([]int, size*size)
	for ii, _ := range check {
		check[ii] = 1
	}
	board := make([]int, size*size)
	empty := 0
	posY := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		posX := 0
		words := strings.Split(scanner.Text(), " ")
		for _, word := range words {
			if len(word) <= 0 {
				continue
			} else if word[0] == '#' {
				break
			} else {
				num, err := strconv.ParseInt(word, 10, 32)
				if err != nil {
					return nil, err
				} else if posX >= size || posY >= size {
					return nil,
						fmt.Errorf("Too many arguments for size %d: \"%s\"",
							size, word)
				}
				if num >= int64(size*size) || num < 0 {
					return nil,
						fmt.Errorf("Number outsize of range %dx%d: %d", size, size, num)
				}
				if check[num] != 1 {
					return nil,
						fmt.Errorf("Duplicate detected: %d", num)
				}
				check[num] = 0
				board[(posY*size)+posX] = int(num)
				// if empty tile
				if num == 0 {
					empty = posX + (posY * size)
				}
				posX += 1
			}
		}
		// if it read a line of the map
		if posX == size {
			posY += 1
		} else if posX != 0 {
			return nil,
				fmt.Errorf("Not enough arguments for size %d: \"%s\"", size, scanner.Text())
		}
	}

	if empty < 0 {
		return nil,
			fmt.Errorf("No empty tile found.  An empty tile is marked with a '0'")
	} else if posY != size {
		return nil, fmt.Errorf("Not enough lines for argument %d.  Expected %d", size, posY)
	}

	// NOTE: this will probably never be called
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	state := new(State)
	state.SoftInit(board, empty, size)
	return state, nil
}
