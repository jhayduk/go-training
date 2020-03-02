// Package utils is a collection of miscellaneous utilities used
// throughout the code.
package utils

import (
	"image"
	"os"

	// Import the image file formats that are to be supported
	_ "image/png"

	"github.com/faiface/pixel"
)

var (
	// OrigWindowSize is the original size on the NES platform that the
	// map was visible in.
	OrigWindowSize pixel.Vec = pixel.V(256.0, 176.0)

	// OrigToCurrentWindowScaling is the scaling factor to apply to
	// OrigWindowSize to get to WindowSize.
	OrigToCurrentWindowScaling float64 = 2.5

	// WindowSize is the actual window size that the game takes on the
	// screen.
	WindowSize pixel.Vec = OrigWindowSize.Scaled(OrigToCurrentWindowScaling)
)

//
// LoadPicture is a helper function to load pictures from files
// into a PictureData object.
//
func LoadPicture(path string) (pixel.Picture, error) {
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
