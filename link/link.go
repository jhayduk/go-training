// Package link contains the code for the player character, Link.
package link

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/jhayduk/go-training/utils"
	"github.com/jhayduk/go-training/utils/nes"
)

// once is used to call initializeLink() only once from the Draw()
// function no matter how many times the Draw() function gets called.
var once sync.Once

// spritesheet is the full spritesheet for the Link character. It is loaded
// the first time Draw is called.
var spritesheet pixel.Picture

// flippedSpritesheet is identical to spritesheet, except it is flipped
// horizontally (about its y axis). This is done because the link spritesheet
// does not contain any west (left) facing images, so the standard sheet
// was flipped and saved so the east (right) facing characters face west.
// This is also loaded the first time Draw is called.
var flippedSpriteSheet pixel.Picture

// Direction is an enumerated type of the different directions Link can be
// facing. South is considered as facing front towards the camera
//
//         North (Back/Backward/Up)
//                    ^
//                    |
// West (Left) <--- Link ---> East (Right)
//                    |
//                    v
//        South (Front/Forward/Down)
//
type Direction int

// The following is an enumerated list of directions that Link can face.
// Facing south means facing forward.
const (
	North Direction = iota
	East
	South
	West
	maxDirection
)

// facing indicates which way Link is currently facing.
// This starts off facing south (or forward) and then changes as keys
// are pressed and the various Turn*() functions are called.
var facing = South

// linkSprite is the link sprite that should be drawn next.
var linkSprite *pixel.Sprite

// basicSprites is a Link sprite facing in each direction when walking
// normally (meaning without any weapons)
// TODO - this needs to be reworked to include the multiple steps link
// can take.
var basicSprites [maxDirection]*pixel.Sprite

func initializeLink() {
	var err error
	spritesheet, err = utils.LoadPicture("assets/NES-TheLegendOfZelda-Link.png")
	if err != nil {
		panic(err)
	}

	flippedSpriteSheet, err = utils.LoadPicture("assets/NES-TheLegendOfZelda-Link-FlippedHorizontally.png")
	if err != nil {
		panic(err)
	}

	// Load the linkSprite sprites for each direction.
	basicSprites[North] = pixel.NewSprite(spritesheet, pixel.R(69.0, 283.0, 69.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[East] = pixel.NewSprite(spritesheet, pixel.R(35.0, 283.0, 35.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[South] = pixel.NewSprite(spritesheet, pixel.R(1.0, 283.0, 1.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[West] = pixel.NewSprite(flippedSpriteSheet, pixel.R(320.0, 283.0, 320.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))

	// Start off link facing south (forward)
	facing = South
	linkSprite = basicSprites[facing]
}

// Draw draws link to the target.
func Draw(target pixel.Target) {
	once.Do(initializeLink)
	linkSprite.Draw(target, pixel.IM.Moved(nes.StartingLocation).Scaled(pixel.ZV, utils.NESToCurrentScaling))
}

// Turn makes sure LInk is turned in the Direction given turning him if
// he is not already facing that way.
func Turn(d Direction) {
	if facing != d {
		facing = d
		linkSprite = basicSprites[facing]
	}
}
