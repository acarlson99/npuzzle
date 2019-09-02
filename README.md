# N-Puzzle Solver

## Dependencies

* go v 1.12
* sdl2
* sdl2 ttf

## Build

```sh
go build ./cmd/npuzzle
```

## Usage

```sh
./npuzzle -h # help output
./npuzzle start.txt
```

## Algorithm

### Search

* Greedy

	* Not optimal

	* Fast

	* Path based only on H (heuristic) score

	* Heuristic DFS

* Uniform

	* Optimal

	* Slow

	* Entirely ignores H score

	* BFS

* Astar

	* Optimal

	* Fast

	* Path based on distance from beginning and H score

### Heuristic

* Atomic

	* Atomic distance

	* Number of tiles out of place

* Manhattan

	* Manhattan Distance from goal

* Max

	* Max of Manhattan and Atomic

* Conflict

	* Linear Conflict

## Visu

* e - end of solution
* s - start of solution
* right - step forward
* left - step backward
* q esc - quit

## TROUBLESHOOTING

nfs file locks don't work which may lead to problems building. Try setting GOPATH to /tmp/go

May need to install sdl libs

* linux - apt install libsdl2-dev libsdl2-ttf-dev
* osx - brew install sdl2 sdl2_ttf

If gui breaks build too much try removing gui.go and the bit that calls gui stuff (end of main)
