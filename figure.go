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

	// check if turn intersects any of busy blocks
	b := ctx.busy_blocks
	fb := f.blocks()
	if b.areHere(fb[0][0], fb[0][1]) ||
		b.areHere(fb[1][0], fb[1][1]) ||
		b.areHere(fb[2][0], fb[2][1]) ||
		b.areHere(fb[3][0], fb[3][1]) {
		f.figureOrientation--
		return false
	}

	// check if turn intersects any of frame borders
	if f.getLeft() <= 0 ||
		f.getRight() >= ctx.cfg.frameWidth-1 {
		f.figureOrientation--
		return false
	}

	f.lastTurnover = time.Now()
	return true
}

func (f *figure) moveDown(ctx *appContext) bool {

	f.y++

	b := ctx.busy_blocks
	fb := f.blocks()

	if !b.areHere(fb[0][0], fb[0][1]) &&
		!b.areHere(fb[1][0], fb[1][1]) &&
		!b.areHere(fb[2][0], fb[2][1]) &&
		!b.areHere(fb[3][0], fb[3][1]) &&
		f.getDown() != ctx.cfg.frameHeight-1 {
		f.lastMovement = time.Now()
		return true
	}

	f.y--

	return false
}

func (f *figure) moveRight(ctx *appContext) bool {

	f.x++

	b := ctx.busy_blocks
	fb := f.blocks()

	if !b.areHere(fb[0][0], fb[0][1]) &&
		!b.areHere(fb[1][0], fb[1][1]) &&
		!b.areHere(fb[2][0], fb[2][1]) &&
		!b.areHere(fb[3][0], fb[3][1]) &&
		f.getRight() < ctx.cfg.frameWidth-1 {
		f.lastMovement = time.Now()
		return true
	}

	f.x--
	return false

}

func (f *figure) moveLeft(ctx *appContext) bool {

	f.x--

	b := ctx.busy_blocks
	fb := f.blocks()

	if !b.areHere(fb[0][0], fb[0][1]) &&
		!b.areHere(fb[1][0], fb[1][1]) &&
		!b.areHere(fb[2][0], fb[2][1]) &&
		!b.areHere(fb[3][0], fb[3][1]) &&
		f.getLeft() > 0 {
		f.lastMovement = time.Now()
		return true
	}

	f.x++
	return false

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

/*
	 func (f *figure) getUp() int {
		b := f.blocks()

		return min(b[0][1], b[1][1], b[2][1], b[3][1])
	}
*/
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

func (f *figure) blockData(h, v int) int {
	b := f.blocks()

	if h == b[0][0] && v == b[0][1] {
		return '█'
	}

	if h == b[1][0] && v == b[1][1] {
		return '█' //'░'
	}

	if h == b[2][0] && v == b[2][1] {
		return '█' //'▒'
	}

	if h == b[3][0] && v == b[3][1] {
		return '█' //'▓'
	}
	return 'S'
}

func (f *figure) blocks() [4][2]int {
	blocks := [4][2]int{}

	switch f.figureType % 7 {
	case 0:
		blocks = f.blocksSquare()
	case 1:
		blocks = f.blocksLine()
	case 2:
		blocks = f.blocksL()
	case 3:
		blocks = f.blocksT()
	case 4:
		blocks = f.blocksS()
	case 5:
		blocks = f.blocksX()
	case 6:
		blocks = f.blocks6()
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

/*
X X   X

	.  X . X

X X   X
*/
func (f *figure) blocksX() [4][2]int {

	var positions = [2][4][2]int{
		{
			{f.x - 1, f.y - 1}, {f.x + 1, f.y - 1},
			{f.x - 1, f.y + 1}, {f.x + 1, f.y + 1},
		},
		{
			{f.x - 1, f.y}, {f.x + 1, f.y}, {f.x, f.y - 1}, {f.x, f.y + 1},
		},
	}

	return positions[f.figureOrientation%2]

}

// .........X.X
// .XX.XX..XX.XX
// XX...XX.X...X
func (f *figure) blocksS() [4][2]int {
	var positions = [4][4][2]int{
		{
			{f.x, f.y}, {f.x + 1, f.y}, {f.x + 1, f.y - 1}, {f.x + 2, f.y - 1},
		},
		{
			{f.x + 1, f.y}, {f.x + 2, f.y}, {f.x, f.y - 1}, {f.x + 1, f.y - 1},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x + 1, f.y - 1}, {f.x + 1, f.y - 2},
		},
		{
			{f.x + 1, f.y}, {f.x + 1, f.y - 1}, {f.x, f.y - 1}, {f.x, f.y - 2},
		},
	}

	return positions[f.figureOrientation%4]
}

// ........X.......X
// .........X.....X.....XX.
// .X..X....X.....X....X..X
// ..XX....X.......X.......
func (f *figure) blocks6() [4][2]int {
	var positions = [4][4][2]int{
		{
			{f.x, f.y}, {f.x + 1, f.y}, {f.x - 1, f.y - 1}, {f.x + 2, f.y - 1},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x - 1, f.y - 2}, {f.x - 1, f.y + 1},
		},
		{
			{f.x, f.y}, {f.x, f.y - 1}, {f.x + 1, f.y - 2}, {f.x + 1, f.y + 1},
		},
		{
			{f.x, f.y}, {f.x + 1, f.y}, {f.x - 1, f.y + 1}, {f.x + 2, f.y + 1},
		},
	}

	return positions[f.figureOrientation%4]
}
