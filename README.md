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

* Hamming

	* Hamming distance

	* Number of tiles out of place

* Manhattan

	* Manhattan Distance from goal

* Max

	* Max of Manhattan and Hamming

* Conflict

	* Tiles conflict if in same row or column, goal in row/column, blocked by other tile

* Conflict-Manhattan

	* Linear Conflict * 2 + Manhattan

	* Able to solve 4-puzzles relatively quickly

## Visu

* e - end of solution
* s - start of solution
* right - step forward
* left - step backward
* q esc - quit

## Resources

* [astar](https://en.wikipedia.org/wiki/A*_search_algorithm)

* [heuristics](https://algorithmsinsight.wordpress.com/graph-theory-2/a-star-in-general/implementing-a-star-to-solve-n-puzzle/)

## TROUBLESHOOTING

nfs file locks don't work which may lead to problems building. Try setting GOPATH to /tmp/go

May need to install sdl libs

* linux - apt install libsdl2-dev libsdl2-ttf-dev
* osx - brew install sdl2 sdl2_ttf

If visualizer breaks build too much try removing visu.go and the bit that calls visu stuff (end of main)
