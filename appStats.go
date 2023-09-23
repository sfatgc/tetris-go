package main

import "time"

type appStats struct {
	start       time.Time
	fps         float64
	framesShown uint
}

func newAppStats() *appStats {
	return &appStats{
		start:       time.Now(),
		fps:         0,
		framesShown: 0,
	}
}

func (this *appStats) statsUpdate() {
	t := time.Now()
	this.framesShown++
	if this.framesShown == 30 {

		this.fps = float64(this.framesShown) / time.Since(this.start).Seconds()
		this.start = t
		this.framesShown = 0
	}

}
