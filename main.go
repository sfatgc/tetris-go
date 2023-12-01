package main

import (
	"fmt"
	"os"
	"time"

	term "github.com/nsf/termbox-go"
)

const DEFAULT_FRAME_HEIGHT = 30
const DEFAULT_FRAME_WIDTH = 15
const EMPTY_AREA_CHARACTER = ' '

func main() {

	ctx := newAppContext(newAppConfiguration(DEFAULT_FRAME_HEIGHT, DEFAULT_FRAME_WIDTH))

	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()

	inputChannel := make(chan term.Key, 100)
	timerChannel := time.Tick(10 * time.Millisecond)

	go getInput(ctx, inputChannel)

	for {
		if processEvents(ctx, inputChannel, timerChannel) && ctx.update() {
			renderFrame(ctx)
			showFrame(ctx)
			ctx.stats.statsUpdate()
		} else {
			break
		}
	}
	renderFrame(ctx)
	showFrame(ctx)
	timerChannel = time.After(5 * time.Second)
	processEvents(ctx, inputChannel, timerChannel)
}

func getInput(ctx *appContext, ic chan term.Key) {
	for {
		ev := term.PollEvent()

		switch ev.Type {
		case term.EventKey:
			term.Sync()
			ic <- ev.Key
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func processEvents(ctx *appContext, ic chan term.Key, tc <-chan time.Time) bool {

	select {
	case <-tc:
		return true
	case inputKey := <-ic:
		switch inputKey {
		case term.KeyEsc:
			return false
		case term.KeyArrowUp:
			ctx.figure.turn(ctx)
		case term.KeyArrowLeft:
			ctx.figure.moveLeft(ctx)
		case term.KeyArrowRight:
			ctx.figure.moveRight(ctx)
		case term.KeyArrowDown:
			ctx.figure.moveDown(ctx)
			ctx.figure.moveDown(ctx)
		}
	}

	return true
}

func renderFrameBorders(ctx *appContext, current_char, h, v int) int {

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

	return current_char
}

func renderFrame(ctx *appContext) {

	for h := 0; h < ctx.cfg.frameWidth; h++ {
		for v := 0; v < ctx.cfg.frameHeight; v++ {

			var current_char int = EMPTY_AREA_CHARACTER

			if ctx.figure.isHere(h, v) {
				current_char = ctx.figure.blockData(h, v)
			}

			if ctx.busy_blocks.areHere(h, v) {
				current_char = ctx.busy_blocks.busy[v][h].data
			}

			current_char = renderFrameBorders(ctx, current_char, h, v)

			ctx.frameData[v][h] = current_char

		}
	}

	ctx.buffer.Reset()

	for v := 0; v < ctx.cfg.frameHeight; v++ {
		ctx.buffer.WriteString("\n\t\t\t")
		for h := 0; h < ctx.cfg.frameWidth; h++ {
			ctx.buffer.WriteRune(rune(ctx.frameData[v][h]))
			if h > 0 && h < ctx.cfg.frameWidth-1 {
				ctx.buffer.WriteRune(rune(ctx.frameData[v][h]))
			}
		}
	}

}

func showFrame(ctx *appContext) {
	fmt.Fprintf(os.Stdout,
		"\033[2J\033[1;1H"+
			"Initialized frame height: %d, width: %d"+
			"%s"+
			"\nFrame rate: %f\n"+
			"Round: %d\tSCORE: %d",
		ctx.cfg.frameHeight,
		ctx.cfg.frameWidth,
		ctx.buffer.String(),
		ctx.stats.fps,
		ctx.rounds,
		ctx.score)
}
