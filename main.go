package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// run is the main function of the game. It performs the initialization
// of the main game window and contains the update loop to continuously
// redraw the window.
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		// Limit screen updates to the refresh rate of the monitor.
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Clear the game screen by painting the whole screen Skyblue.
	win.Clear(colornames.Skyblue)

	for !win.Closed() {
		win.Update()
	}

}

//
// The main function only calls pxelgl.Run(run) to let the run() function,
// which is the main function for the game, execute in the main thread.
func main() {
	pixelgl.Run(run)
}
