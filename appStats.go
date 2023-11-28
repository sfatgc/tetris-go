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

func (stats *appStats) statsUpdate() {
	t := time.Now()
	stats.framesShown++
	// if stats.framesShown == 10 {
	if time.Since(stats.start).Milliseconds() >= 1000 {

		stats.fps = float64(stats.framesShown) / time.Since(stats.start).Seconds()
		stats.start = t
		stats.framesShown = 0
	}

}
