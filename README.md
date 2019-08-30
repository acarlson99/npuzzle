# N-Puzzle Solver

## Dependencies

* go v 1.12

## Build

```sh
go build ./cmd/npuzzle
```

## Usage

```sh
./npuzzle -h # help output
./npuzzle start.txt
```

## NOTES

If you are on nfs file locks don't work which may lead to problems building. Try setting GOPATH to /tmp/go
