# N-Puzzle Solver

## Dependencies

* go v 1.12
* sdl2

## Build

```sh
go build ./cmd/npuzzle
```

## Usage

```sh
./npuzzle -h # help output
./npuzzle start.txt
```

## TROUBLESHOOTING

nfs file locks don't work which may lead to problems building. Try setting GOPATH to /tmp/go

May need to install sdl libs

* linux - libsdl2-dev and libsdl2-ttf-dev

If gui breaks build too much try removing gui.go and the bit that calls gui stuff (end of main)
