package main

type block struct {
	is_busy bool
	data    int
}
type blocks struct {
	deletable_lines bool
	height          int
	width           int
	busy            [][]block
}

func newBlocks(width, height int) *blocks {
	b := blocks{
		height: height,
		width:  width,
		busy:   make([][]block, height),
	}
	for lineNumber := 0; lineNumber < height; lineNumber++ {
		b.busy[lineNumber] = make([]block, width)
	}
	return &b
}

func (b *blocks) areHere(h, v int) bool {
	return 0 <= h && h < b.width && 0 <= v && v < b.height && b.busy[v][h].is_busy
}

func (b *blocks) addFigure(f *figure) {
	for _, fb := range f.blocks() {
		h := fb[0]
		v := fb[1]
		if 0 <= h && h < b.width && 0 <= v && v < b.height && !b.areHere(h, v) {
			b.busy[v][h].is_busy = true
			b.busy[v][h].data = '▒'
			if b.line_full(v) {
				for h := 0; h < b.width; h++ {
					b.busy[v][h].data = '◘'
				}
				b.deletable_lines = true
			}
		}
	}
}

func (b *blocks) line_full(v int) bool {
	strike := true
	for h := 1; h < b.width-1; h++ {
		strike = strike && b.busy[v][h].is_busy
	}
	return strike
}

func (b *blocks) delete_deletable_lines() int {
	b.deletable_lines = false

	lines_deleted := 0

	for lineNumber := 0; lineNumber < b.height; lineNumber++ {
		if b.line_full(lineNumber) {
			b.delete_line(lineNumber)
			lines_deleted++
		}
	}

	return lines_deleted

}

func (b *blocks) delete_line(v int) {
	for lineNumber := v; lineNumber > 0; lineNumber-- {
		b.busy[lineNumber] = b.busy[lineNumber-1]
	}
	for columnNumber := 0; columnNumber < b.width; columnNumber++ {
		b.busy[0][columnNumber].is_busy = false
		b.busy[0][columnNumber].data = EMPTY_AREA_CHARACTER
	}
}
