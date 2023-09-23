package main

import "bytes"

type appContext struct {
	cfg       *appConfiguration
	stats     *appStats
	buffer    *bytes.Buffer
	figure    *figure
	pastFrame [][]int
}

func newAppContext(c *appConfiguration) *appContext {
	context := appContext{
		cfg:       c,
		stats:     newAppStats(),
		buffer:    new(bytes.Buffer),
		figure:    newFigure(3, 0, c.frameWidth/2, c.frameHeight/4),
		pastFrame: make([][]int, c.frameHeight),
	}

	for lineNumber := 0; lineNumber < c.frameHeight; lineNumber++ {
		context.pastFrame[lineNumber] = make([]int, c.frameWidth)
	}

	return &context
}
