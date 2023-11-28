package main

import "time"

type figure struct {
	figureType        int
	figureOrientation int
	x                 int
	y                 int
	lastMovement      time.Time
	lastTurnover      time.Time
}

func newFigure(ft, o, x, y int) *figure {
	return &figure{
		figureType:        ft,
		figureOrientation: o,
		x:                 x,
		y:                 y,
		lastMovement:      time.Now(),
		lastTurnover:      time.Now(),
	}
}

func (f *figure) turn(ctx *appContext) bool {

	f.figureOrientation++

	// check if turn intersects any of byusy blocks
	b := ctx.busy_blocks
	fb := f.blocks()
	if b.areHere(fb[0][0], fb[0][1]) ||
		b.areHere(fb[1][0], fb[1][1]) ||
		b.areHere(fb[2][0], fb[2][1]) ||
		b.areHere(fb[3][0], fb[3][1]) {
		f.figureOrientation--
		return false
	}

	// TODO: check if turn intersects any of frame borders

	f.lastTurnover = time.Now()
	return true
}

func (f *figure) moveDown(ctx *appContext) bool {

	b := ctx.busy_blocks
	fb := f.blocks()

	if !b.areHere(fb[0][0], fb[0][1]+1) &&
		!b.areHere(fb[1][0], fb[1][1]+1) &&
		!b.areHere(fb[2][0], fb[2][1]+1) &&
		!b.areHere(fb[3][0], fb[3][1]+1) &&
		fb[0][1]+1 != ctx.cfg.frameHeight &&
		fb[1][1]+1 != ctx.cfg.frameHeight &&
		fb[2][1]+1 != ctx.cfg.frameHeight &&
		fb[3][1]+1 != ctx.cfg.frameHeight {
		f.lastMovement = time.Now()
		f.y++
		return true
	}
	return false
}

func (f *figure) moveRight() {

	f.lastMovement = time.Now()
	f.x++

}

func (f *figure) moveLeft() {

	f.lastMovement = time.Now()
	f.x--

}

func (f *figure) getRight() int {
	b := f.blocks()

	return max(b[0][0], b[1][0], b[2][0], b[3][0])
}

func (f *figure) getLeft() int {
	b := f.blocks()

	return min(b[0][0], b[1][0], b[2][0], b[3][0])
}

func (f *figure) getDown() int {
	b := f.blocks()

	return max(b[0][1], b[1][1], b[2][1], b[3][1])
}

func (f *figure) getUp() int {
	b := f.blocks()

	return min(b[0][1], b[1][1], b[2][1], b[3][1])
}

func (f *figure) isHere(h, v int) bool {

	b := f.blocks()

	if h == b[0][0] && v == b[0][1] {
		return true
	}

	if h == b[1][0] && v == b[1][1] {
		return true
	}

	if h == b[2][0] && v == b[2][1] {
		return true
	}

	if h == b[3][0] && v == b[3][1] {
		return true
	}

	return false
}

func (f *figure) block(h, v int) int {
	b := f.blocks()

	if h == b[0][0] && v == b[0][1] {
		return '█'
	}

	// if h == b[1][0] && v == b[1][1] {
	// 	return true
	// }

	// if h == b[2][0] && v == b[2][1] {
	// 	return true
	// }

	// if h == b[3][0] && v == b[3][1] {
	// 	return true
	// }
	return '█'
}

func (f *figure) blocks() [4][2]int {
	blocks := [4][2]int{}

	switch f.figureType {
	case 1:
		blocks = f.blocksSquare()
	case 2:
		blocks = f.blocksLine()
	case 3:
		blocks = f.blocksL()
	case 4:
		blocks = f.blocksT()
	}

	return blocks
}

// XX
// XX
func (f *figure) blocksSquare() [4][2]int {
	return [4][2]int{
		{f.x - 1, f.y - 1}, {f.x, f.y - 1},
		{f.x - 1, f.y}, {f.x, f.y},
	}

}

// X
// X
// X
// X
func (f *figure) blocksLine() [4][2]int {

	var positions = [2][4][2]int{
		{
			{f.x, f.y - 3},
			{f.x, f.y - 2},
			{f.x, f.y - 1},
			{f.x, f.y},
		},
		{
			{f.x - 3, f.y}, {f.x - 2, f.y}, {f.x - 1, f.y}, {f.x, f.y},
		},
	}

	return positions[f.figureOrientation%2]
}

// X
// X
// XX
func (f *figure) blocksL() [4][2]int {
	var positions = [8][4][2]int{
		{
			{f.x - 1, f.y - 2},
			{f.x - 1, f.y - 1},
			{f.x - 1, f.y}, {f.x, f.y},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x + 1, f.y - 1}, {f.x + 2, f.y - 1},
		},
		{
			{f.x, f.y},
			{f.x + 1, f.y},
			{f.x + 1, f.y + 1},
			{f.x + 1, f.y + 2},
		},
		{
			{f.x, f.y}, {f.x, f.y + 1}, {f.x - 1, f.y + 1}, {f.x - 2, f.y + 1},
		},
		{
			{f.x, f.y},
			{f.x - 1, f.y},
			{f.x - 1, f.y + 1},
			{f.x - 1, f.y + 2},
		},
		{
			{f.x, f.y}, {f.x, f.y + 1}, {f.x + 1, f.y + 1}, {f.x + 2, f.y + 1},
		},
		{
			{f.x, f.y},
			{f.x + 1, f.y},
			{f.x + 1, f.y - 1},
			{f.x + 1, f.y - 2},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x - 1, f.y - 1}, {f.x - 2, f.y - 1},
		},
	}

	return positions[f.figureOrientation%8]
}

// .X
// XXX
func (f *figure) blocksT() [4][2]int {
	var positions = [4][4][2]int{
		{
			{f.x, f.y}, {f.x + 1, f.y}, {f.x - 1, f.y}, {f.x, f.y + 1},
		},
		{
			{f.x, f.y}, {f.x + 1, f.y}, {f.x - 1, f.y}, {f.x, f.y - 1},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x, f.y + 1}, {f.x - 1, f.y},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x, f.y + 1}, {f.x + 1, f.y},
		},
	}

	return positions[f.figureOrientation%4]
}
