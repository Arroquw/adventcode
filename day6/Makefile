GOC = go build
GOV = go vet
GOL = golint

BENCH = ../bench.sh

IN = input.txt

BIN = $(addprefix aoc_,$(shell basename $(CURDIR)))
SRC = $(shell basename $(CURDIR)).go

build:
		$(GOC) -o $(BIN) $(SRC)

check:
		$(GOV) $(SRC)
		$(GOL) $(SRC)

clean:
		go clean
		rm -f $(BIN)

run:
		cat $(IN) | go run ./$(SRC)

.PHONY: bench build check clean run
