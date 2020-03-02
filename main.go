package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	// Import the image file formats that are to be supported
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/jhayduk/go-training/overworld"
	"github.com/jhayduk/go-training/utils"
	"golang.org/x/image/colornames"
)

//
// run is the main function of the game. It performs the initialization
// of the main game window and contains the update loop to continuously
// redraw the window.
//
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, utils.WindowSize.X, utils.WindowSize.Y),
		// Limit screen updates to the refresh rate of the monitor.
		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spritesheet, err := utils.LoadPicture("assets/trees.png")
	if err != nil {
		panic(err)
	}

	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	var (
		camPos       = pixel.ZV
		camSpeed     = 250.0
		camZoom      = 1.0
		camZoomSpeed = 1.2
		frames       = 0
		second       = time.Tick(time.Second)
	)

	//
	// Continuously update the window until the close button is hit
	//
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			tree := pixel.NewSprite(spritesheet, treesFrames[rand.Intn(len(treesFrames))])
			// Unproject transforms to game space
			mouse := cam.Unproject(win.MousePosition())
			tree.Draw(batch, pixel.IM.Scaled(pixel.ZV, 4).Moved(mouse))
		}

		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= camSpeed * dt
			if camPos.X <= pixel.ZV.X {
				camPos.X = pixel.ZV.X
			}
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += camSpeed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= camSpeed * dt
			if camPos.Y <= pixel.ZV.Y {
				camPos.Y = pixel.ZV.Y
			}
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += camSpeed * dt
		}
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)

		win.Clear(colornames.Blue)
		batch.Draw(win)
		overworld.Draw(win)
		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

//
// The main function only calls pxelgl.Run(run) to let the run() function,
// which is the main function for the game, execute in the main thread.
func main() {
	pixelgl.Run(run)
}
