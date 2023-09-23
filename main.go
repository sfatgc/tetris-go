package main

import (
	"fmt"
	"os"
	"time"
)

const DEFAULT_FRAME_HEIGHT = 30
const DEFAULT_FRAME_WIDTH = 50
const EMPTY_AREA_CHARACTER = ' '

func main() {

	ctx := newAppContext(newAppConfiguration(DEFAULT_FRAME_HEIGHT, DEFAULT_FRAME_WIDTH))

	for {
		t := time.Now()
		ctx.figure.update()
		renderFrame(ctx)
		showFrame(ctx)
		ctx.stats.statsUpdate()
		dt := time.Since(t)
		time.Sleep((time.Second / 60) - dt)
	}

}

func renderFrame(ctx *appContext) {

	frameData := ctx.pastFrame

	for h := 0; h < ctx.cfg.frameWidth; h++ {
		for v := 0; v < ctx.cfg.frameHeight; v++ {

			frameData[v][h] = EMPTY_AREA_CHARACTER

			if ctx.figure.here(h, v) {
				frameData[v][h] = ctx.figure.block(h, v)
			}

			if h == 0 || h == ctx.cfg.frameWidth-1 {
				frameData[v][h] = '║'
			}

			if v == 0 {
				if h == 0 {
					frameData[v][h] = '╔'
				} else if h == ctx.cfg.frameWidth-1 {
					frameData[v][h] = '╗'
				} else {
					frameData[v][h] = '═'
				}
			}

			if v == ctx.cfg.frameHeight-1 {
				if h == 0 {
					frameData[v][h] = '╚'
				} else if h == ctx.cfg.frameWidth-1 {
					frameData[v][h] = '╝'
				} else {
					frameData[v][h] = '═'
				}
			}

		}
	}

	ctx.buffer.Reset()

	for v := 0; v < ctx.cfg.frameHeight; v++ {
		ctx.buffer.WriteString("\n\t\t\t")
		for h := 0; h < ctx.cfg.frameWidth; h++ {
			ctx.buffer.WriteString(string(frameData[v][h]))
		}
	}

	ctx.pastFrame = frameData

}

func showFrame(ctx *appContext) {
	fmt.Fprintf(os.Stdout,
		"\033[2J\033[1;1H"+
			"Initialized frame height: %d, width: %d"+
			"%s"+
			"\nFrame rate: %f\n",
		ctx.cfg.frameHeight,
		ctx.cfg.frameWidth,
		ctx.buffer.String(),
		ctx.stats.fps)
}
