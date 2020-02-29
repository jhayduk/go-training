package main

import (
	"image"
	"os"
	"time"

	// Import the image file formats that are to be supported
	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//
// loadPicture is a helper function to load pictures from files
// into a PictureData object.
//
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

//
// run is the main function of the game. It performs the initialization
// of the main game window and contains the update loop to continuously
// redraw the window.
//
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
	win.SetSmooth(true)

	//
	// Load the "hiking gopher" picture and create a sprite from it.
	// For now, this is the player character (I presume)
	//
	pic, err := loadPicture("hiking.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())

	win.Clear(colornames.Firebrick)

	angle := 0.0

	//
	// Continuously update the window until the close button is hit
	//
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		angle += 3 * dt

		//
		// Note that sprites are anchored by their centers, the x,y axis
		// starts at the lower left and the axes increase "naturally" (x
		// increases to the right and y increases upward).
		mat := pixel.IM
		mat = mat.Rotated(pixel.ZV, angle)
		//
		// Moved(win.Bounds().Center()) does not, technically, mean "move to the
		// center of the window", it means "move by the amount specified by a
		// vector starting at the origin and ending at the center of the window".
		//
		mat = mat.Moved(win.Bounds().Center())

		win.Clear(colornames.Firebrick)
		sprite.Draw(win, mat)
		win.Update()
	}

}

//
// The main function only calls pxelgl.Run(run) to let the run() function,
// which is the main function for the game, execute in the main thread.
func main() {
	pixelgl.Run(run)
}
