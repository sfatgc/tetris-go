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
		ctx.updateFigure()
		renderFrame(ctx)
		showFrame(ctx)
		ctx.stats.statsUpdate()
		dt := time.Since(t)
		time.Sleep((time.Second / 60) - dt)
	}

}

func renderFrame(ctx *appContext) {

	for h := 0; h < ctx.cfg.frameWidth; h++ {
		for v := 0; v < ctx.cfg.frameHeight; v++ {

			var current_char int = EMPTY_AREA_CHARACTER

			if ctx.figure.isHere(h, v) {
				current_char = ctx.figure.block(h, v)
			}

			if ctx.busy_blocks.areHere(h, v) {
				current_char = 'X'
			}

			if h == 0 || h == ctx.cfg.frameWidth-1 {
				current_char = '║'
			}

			if v == 0 {
				if h == 0 {
					current_char = '╔'
				} else if h == ctx.cfg.frameWidth-1 {
					current_char = '╗'
				} else {
					current_char = '═'
				}
			}

			if v == ctx.cfg.frameHeight-1 {
				if h == 0 {
					current_char = '╚'
				} else if h == ctx.cfg.frameWidth-1 {
					current_char = '╝'
				} else {
					current_char = '═'
				}
			}

			ctx.frameData[v][h] = current_char

		}
	}

	ctx.buffer.Reset()

	for v := 0; v < ctx.cfg.frameHeight; v++ {
		ctx.buffer.WriteString("\n\t\t\t")
		for h := 0; h < ctx.cfg.frameWidth; h++ {
			ctx.buffer.WriteString(string(ctx.frameData[v][h]))
		}
	}

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
