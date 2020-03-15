// Package overworld provides the facility to load the overworld map.
//
// The original overworld had 128 maps arranged in a 16 x 8 grid.
// For this implementation, a border of 1 map all around is added so
// that having the map dynamically scroll while Link is walking to the edge
// of the original map will always show a drawn map. The border is added
// as deep water so that it appears that the whole adventure takes place
// on a rectangular island.
// The new overworld is, thus, arranged in an 18 x 10 grid with 180 maps.
//
// The overworld itself is kept in a pixel.Batch as an already scaled
// image that can be drawn on a pixel.Target with the Draw() function.
package overworld

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/jhayduk/go-training/utils"
	"github.com/jhayduk/go-training/utils/nes"
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

	//
	// First, load each of the maps from the spritesheet.
	//
	var mapFrames []pixel.Rect
	numRows := 0.0
	numCols := 0.0
	// Note that there is a 1 pixel width border between every submap, and an
	// additional footer of 181 pixel height at the bottom of the spritesheet.
	// So, to get to the origin of the first submap on the bottom left of the
	// spritesheet, we have to skip to the right 1 pixel and up 182 pixels.
	// Submaps are then loaded into the array one column at a time, bottom
	// up and then left to right. When all done, mapFrames contain the
	// frames for all of the actual submaps with the footer and the borders
	// stripped away/
	for x := spritesheet.Bounds().Min.X + 1; x < spritesheet.Bounds().Max.X; x += nes.MapSize.X + 1 {
		numRows = 0.0
		numCols += 1.0
		for y := spritesheet.Bounds().Min.Y + 182; y < spritesheet.Bounds().Max.Y; y += nes.MapSize.Y + 1 {
			numRows += 1.0
			mapFrames = append(mapFrames, pixel.R(x, y, x+nes.MapSize.X, y+nes.MapSize.Y))
		}
	}

	//
	// Now, draw each frame into the overworldMap where it belongs.
	// Note that the map is drawn in scaled world coordinates with the bottom
	// left edge anchored at pixel.ZV (0, 0), but that there is a one map
	// border around the maps added here.
	//
	frameIndex := 0
	for col := 1.0; col <= numCols; col += 1.0 {
		for row := 1.0; row <= numRows; row += 1.0 {
			submapFrame := pixel.NewSprite(spritesheet, mapFrames[frameIndex])
			submapFrame.Draw(overworldMap, pixel.IM.Moved(pixel.V(col*nes.MapSize.X, row*nes.MapSize.Y)).Scaled(pixel.ZV, utils.NESToCurrentScaling))
			frameIndex++
		}
	}
}

// Draw draws the overworld map to the target. The entire map is
// drawn as a single pixel.Batch each time this is called. This is fast
// because it is a single Batch draw.
func Draw(target pixel.Target) {
	once.Do(initializeOverworldMap)
	overworldMap.Draw(target)
}
