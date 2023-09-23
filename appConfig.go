package main

type appConfiguration struct {
	frameHeight int
	frameWidth  int
}

func newAppConfiguration(h, w int) *appConfiguration {
	return &appConfiguration{
		frameHeight: h,
		frameWidth:  w,
	}
}
