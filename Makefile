all: tetris-go
	./tetris-go

tetris-go: *.go
	go build