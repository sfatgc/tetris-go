package main

import (
	"bytes"
	"math/rand"
	"time"
)

type appContext struct {
	cfg         *appConfiguration
	stats       *appStats
	rounds      int
	score       int
	buffer      *bytes.Buffer
	figure      *figure
	busy_blocks *blocks
	frameData   [][]int
}

func newAppContext(c *appConfiguration) *appContext {
	context := appContext{
		cfg:         c,
		stats:       newAppStats(),
		rounds:      0,
		score:       0,
		buffer:      new(bytes.Buffer),
		figure:      newFigure(rand.Intn(6), rand.Intn(8), c.frameWidth/2, c.frameHeight/4),
		busy_blocks: newBlocks(c.frameWidth, c.frameHeight),
		frameData:   make([][]int, c.frameHeight),
	}

	for lineNumber := 0; lineNumber < c.frameHeight; lineNumber++ {
		context.frameData[lineNumber] = make([]int, c.frameWidth)
	}

	return &context
}

func (ctx *appContext) update() bool {

	f := ctx.figure

	if int64(500-(ctx.score*20)) < time.Since(f.lastMovement).Milliseconds() {
		if ctx.busy_blocks.deletable_lines {
			ctx.score += ctx.busy_blocks.delete_deletable_lines()
		} else {

			if !f.moveDown(ctx) {
				ctx.busy_blocks.addFigure(f)
				ctx.figure = newFigure(rand.Intn(6), rand.Intn(8), ctx.cfg.frameWidth/2, 0)
				ctx.rounds++

				// TODO: check-up for edge case failures below
				if /* !ctx.busy_blocks.marked_lines && !ctx.busy_blocks.deletable_lines &&  */ ctx.busy_blocks.areHere(ctx.figure.x, ctx.figure.getDown()) {
					return false
				}

			}
		}
	}
	return true
}
