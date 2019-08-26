package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// read state from file if specified.  Otherwise reads from stdin
// func ReadState(file string) (*State, error) {
// 	var scanner *bufio.Scanner
// 	if file == "" {
// 		scanner = bufio.NewScanner(os.Stdin)
// 	} else {
// 		file, err := os.Open(file)
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer file.Close()
// 		scanner = bufio.NewScanner(file)
// 	}
// 	state, err := ReadStateFromScanner(scanner)
// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return state, nil
// 	}
// }

// get start and end states
func InitStates(startFile, endFile string) (*State, *State, error) {
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
	end := new(State)
	if endFile != "" {
		end, err = ReadStateFromFile(endFile)
		if err != nil {
			return nil, nil, err
		}
	} else {
		fmt.Println("End state unspecified.  Inferring")
		board := make([]uint, len(start.Board))
		copy(board, start.Board)
		sort.SliceStable(board, func(ii, jj int) bool { return board[ii] < board[jj] })
		num := board[0]
		copy(board, board[1:])
		board[((start.Size-1)*start.Size)+start.Size-1] = num
		end.SoftInit(board, int(start.Size-1), int(start.Size-1), start.Size)
	}
	return start, end, nil
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

	board := make([]uint, size*size)
	emptyX := -1
	emptyY := -1
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
				num, err := strconv.ParseUint(word, 10, 32)
				if err != nil {
					return nil, err
				} else if posX >= size || posY >= size {
					return nil,
						fmt.Errorf("Too many arguments for size %d: \"%s\"",
							size, word)
				}
				board[(posY*size)+posX] = uint(num)
				// if empty tile
				if num == 0 {
					emptyX = posX
					emptyY = posY
				}
				posX += 1
			}
		}
		// if it read a line of the map
		if posX == size {
			posY += 1
		} else {
			return nil,
				fmt.Errorf("Not enough arguments for size %d: \"%s\"", size, scanner.Text())
		}
	}

	if emptyX < 0 || emptyY < 0 {
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
	state.SoftInit(board, emptyX, emptyY, size)
	fmt.Println("Input file:")
	for _, line := range lines {
		fmt.Println(line)
	}
	return state, nil
}
