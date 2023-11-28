package main

import (
	"bytes"
	"math/rand"
	"time"
)

type appContext struct {
	cfg           *appConfiguration
	stats         *appStats
	buffer        *bytes.Buffer
	figure        *figure
	busy_blocks   *blocks
	prevFrameData [][]int
	frameData     [][]int
}

func newAppContext(c *appConfiguration) *appContext {
	context := appContext{
		cfg:           c,
		stats:         newAppStats(),
		buffer:        new(bytes.Buffer),
		figure:        newFigure(3, 0, c.frameWidth/2, c.frameHeight/4),
		busy_blocks:   newBlocks(),
		prevFrameData: make([][]int, c.frameHeight),
		frameData:     make([][]int, c.frameHeight),
	}

	for lineNumber := 0; lineNumber < c.frameHeight; lineNumber++ {
		context.frameData[lineNumber] = make([]int, c.frameWidth)
		context.prevFrameData[lineNumber] = make([]int, c.frameWidth)
		for colNumber := 0; colNumber < c.frameWidth; colNumber++ {
			context.frameData[lineNumber][colNumber] = EMPTY_AREA_CHARACTER
			context.prevFrameData[lineNumber][colNumber] = EMPTY_AREA_CHARACTER
		}
	}

	return &context
}

func (ctx *appContext) updateFigure() bool {

	f := ctx.figure

	if 3 < time.Since(f.lastTurnover).Seconds() {
		f.turn()
	}

	if 0.5 < time.Since(f.lastMovement).Seconds() {
		if !f.moveDown(ctx) {
			ctx.busy_blocks.addFigure(f)
			ctx.figure = newFigure(rand.Intn(3)+1, 0, ctx.cfg.frameWidth/2, 0)
			if ctx.busy_blocks.areHere(ctx.figure.x, ctx.figure.getDown()) {
				return false
			}
		}
	}
	return true
}
