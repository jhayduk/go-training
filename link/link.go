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

// linkSprite is a Link sprite facing forward
// TODO - this needs to be reworked, we won't have a separate pointer for
// each way Link is facing, and each step of him walking. This is just to get
// started.
var linkSprite *pixel.Sprite

func initializeLink() {
	var err error
	spritesheet, err = utils.LoadPicture("assets/NES-TheLegendOfZelda-Link.png")
	if err != nil {
		panic(err)
	}

	linkSprite = pixel.NewSprite(spritesheet, pixel.R(1.0, 283.0, 1.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
}

// Draw draws link to the target.
func Draw(target pixel.Target) {
	once.Do(initializeLink)
	linkSprite.Draw(target, pixel.IM.Moved(nes.StartingLocation).Scaled(pixel.ZV, utils.NESToCurrentScaling))
}
