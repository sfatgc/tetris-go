package main

type block struct {
	x, y int
}

type blocks struct {
	busy []block
}

func newBlocks() *blocks {
	return &blocks{}
}

func (b *blocks) areHere(h, v int) bool {
	for _, bs := range b.busy {
		if bs.x == h && bs.y == v {
			return true
		}
	}
	return false
}

func (b *blocks) addFigure(f *figure) {
	for _, fb := range f.blocks() {
		if !b.areHere(fb[0], fb[1]) {
			b.busy = append(b.busy, block{fb[0], fb[1]})
		}
	}
}
