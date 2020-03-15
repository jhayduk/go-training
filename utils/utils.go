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

// NESToCurrentScaling is the scaling factor to apply to
// the NES objects to get to the size displayed in the window.
const NESToCurrentScaling float64 = 2.5

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
