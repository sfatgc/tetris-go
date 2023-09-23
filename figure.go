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

func (this *figure) update() {

	if 1.5 < time.Since(this.lastTurnover).Seconds() {
		this.lastTurnover = time.Now()
		this.figureOrientation++
	}

	if 2 < time.Since(this.lastMovement).Seconds() {
		this.lastMovement = time.Now()
		this.y++
	}

}

func (this *figure) here(h, v int) bool {

	b := this.blocks()

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

func (this *figure) block(h, v int) int {
	b := this.blocks()

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

func (this *figure) blocks() [4][2]int {
	blocks := [4][2]int{}

	switch this.figureType {
	case 1:
		blocks = this.blocksSquare()
	case 2:
		blocks = this.blocksLine()
	case 3:
		blocks = this.blocksL()
	case 4:
		blocks = this.blocksT()
	}

	return blocks
}

// XX
// XX
func (this *figure) blocksSquare() [4][2]int {
	return [4][2]int{
		{this.x - 1, this.y - 1}, {this.x, this.y - 1},
		{this.x - 1, this.y}, {this.x, this.y},
	}

}

// X
// X
// X
// X
func (this *figure) blocksLine() [4][2]int {

	var positions = [2][4][2]int{
		{
			{this.x, this.y - 3},
			{this.x, this.y - 2},
			{this.x, this.y - 1},
			{this.x, this.y},
		},
		{
			{this.x - 3, this.y}, {this.x - 2, this.y}, {this.x - 1, this.y}, {this.x, this.y},
		},
	}

	return positions[this.figureOrientation%2]
}

// X
// X
// XX
func (this *figure) blocksL() [4][2]int {
	var positions = [8][4][2]int{
		{
			{this.x - 1, this.y - 2},
			{this.x - 1, this.y - 1},
			{this.x - 1, this.y}, {this.x, this.y},
		},
		{
			{this.x, this.y}, {this.x, this.y - 1}, {this.x + 1, this.y - 1}, {this.x + 2, this.y - 1},
		},
		{
			{this.x, this.y},
			{this.x + 1, this.y},
			{this.x + 1, this.y + 1},
			{this.x + 1, this.y + 2},
		},
		{
			{this.x, this.y}, {this.x, this.y + 1}, {this.x - 1, this.y + 1}, {this.x - 2, this.y + 1},
		},
		{
			{this.x, this.y},
			{this.x - 1, this.y},
			{this.x - 1, this.y + 1},
			{this.x - 1, this.y + 2},
		},
		{
			{this.x, this.y}, {this.x, this.y + 1}, {this.x + 1, this.y + 1}, {this.x + 2, this.y + 1},
		},
		{
			{this.x, this.y},
			{this.x + 1, this.y},
			{this.x + 1, this.y - 1},
			{this.x + 1, this.y - 2},
		},
		{
			{this.x, this.y}, {this.x, this.y - 1}, {this.x - 1, this.y - 1}, {this.x - 2, this.y - 1},
		},
	}

	return positions[this.figureOrientation%8]
}

// ..XX
// XXXXXX
func (this *figure) blocksT() [4][2]int {
	var positions = [4][4][2]int{
		{
			{this.x, this.y}, {this.x + 1, this.y}, {this.x - 1, this.y}, {this.x, this.y + 1},
		},
		{
			{this.x, this.y}, {this.x + 1, this.y}, {this.x - 1, this.y}, {this.x, this.y - 1},
		},
		{
			{this.x, this.y}, {this.x, this.y - 1}, {this.x, this.y + 1}, {this.x - 1, this.y},
		},
		{
			{this.x, this.y}, {this.x, this.y - 1}, {this.x, this.y + 1}, {this.x + 1, this.y},
		},
	}

	return positions[this.figureOrientation%4]
}
