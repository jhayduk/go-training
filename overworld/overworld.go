// Package overworld provides the facility to load the overworld map.
package overworld

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/jhayduk/go-training/utils"
)

// once is used to call initializeOverworldMap() only once from the Draw()
// function no matter how many times the Draw() function gets called.
var once sync.Once

// overworldMap is a Pixel batch of the entire overworld image. It starts out
// as nil, and is populated the first time it is asked to be drawn.
// After that, it is always redrawn from the batch.
var overworldMap *pixel.Batch

// initializeOverworldMap loads the entire map fro m the assets file and
// loads it into a pixel.Batch object that can be drawn each update
// cycle.
// It is expected that this will be called once either during initialization,
// or the first time the overworld map needs to be drawn.
func initializeOverworldMap() {
	spritesheet, err := utils.LoadPicture("assets/NES-TheLegendOfZelda-Overworld.png")
	if err != nil {
		panic(err)
	}
	overworldMap = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)

	const (
		submapWidth  = 256.0
		submapHeight = 176.0
	)
	var submapFrames []pixel.Rect
	numRows := 0.0
	numCols := 0.0
	// Note that there is a 1 pixel width border between every submap, and an
	// additional footer of 181 pixel height at the bottom of the spritesheet.
	// So, to get to the origin of the first submap on the bottom left of the
	// spritesheet, we have to skip to the right 1 pixel and up 182 pixels.
	// Submaps are then loaded into the array one column at a time, bottom
	// up and then left to right. When all done, submapFrames contain the
	// frames for all of the actual submaps with the footer and the borders
	// stripped away/
	for x := spritesheet.Bounds().Min.X + 1; x < spritesheet.Bounds().Max.X; x += submapWidth + 1 {
		numRows = 0.0
		numCols += 1.0
		for y := spritesheet.Bounds().Min.Y + 182; y < spritesheet.Bounds().Max.Y; y += submapHeight + 1 {
			numRows += 1.0
			submapFrames = append(submapFrames, pixel.R(x, y, x+submapWidth, y+submapHeight))
		}
	}
	// Now, draw each frame into the overworldMap where it belongs.
	frameIndex := 0
	for col := 0.0; col < numCols; col += 1.0 {
		for row := 0.0; row < numRows; row += 1.0 {
			submapFrame := pixel.NewSprite(spritesheet, submapFrames[frameIndex])
			submapFrame.Draw(overworldMap, pixel.IM.Moved(pixel.V(col*submapWidth, row*submapHeight)).Scaled(pixel.ZV, utils.OrigToCurrentWindowScaling))
			frameIndex++
		}
	}
}

// Draw draws the overworld map to the target. The entire map is
// drawn as a single pixel.Batch each time this is called.
func Draw(target pixel.Target) {
	once.Do(initializeOverworldMap)
	overworldMap.Draw(target)
}
