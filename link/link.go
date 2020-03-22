// Package link contains the code for the player character, Link.
package link

import (
	"math"
	"sync"

	"github.com/faiface/pixel"
	"github.com/jhayduk/go-training/utils"
	"github.com/jhayduk/go-training/utils/nes"
)

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

// North, East, South and West are elements of the enumeration Direction
// and are used to indicate what direction Link is or should be facing.
const (
	North Direction = iota
	East
	South
	West
	numDirections
)

// speed is the speed, in world pixels per second, with which Link moves in
// each direction when he is moving.
var speed = [numDirections]pixel.Vec{
	pixel.V(0, 100.0),  // North
	pixel.V(100.0, 0),  // East
	pixel.V(0, -100.0), // South
	pixel.V(-100.0, 0), // West
}

// Step is an enumeration of the individual sprites each character has in
// in a certain direction before repeating. There are only two, so, for
// simplicity in naming, they are called left and right.
type Step int

// left and right are elements of the enumeration Step and are used to
// pick which version of a sprite facing in a certain direction is used.
const (
	right Step = iota
	left
	numSteps
)

// stepsPerSecond are the number of right/left steps Link takes each second
// he is moving
const stepPerSecond float64 = 4.0

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

// currentlyFacing indicates which way Link is currently facing.
// He starts off facing south (or forward) and then changes as keys
// are pressed and the various Turn*() functions are called.
var currentlyFacing Direction = South

// currentStep is the current version of the sprite (right or left) that
// should be used for the direction currently faced.
var currentStep Step = right

// currentRawStep is a float64 representation of the current step value
// that accumlates over time. A modulus of this value is used to determine
// the enumerated step to use.
var currentRawStep float64 = float64(right)

// currentXYLocation is Link's current XY location in world coordinates
var currentXYLocation pixel.Vec = nes.StartingLocation

// linkSprite is the link sprite that should be drawn next.
var linkSprite *pixel.Sprite

// basicSprites is a Link sprite facing in each direction when walking
// normally (meaning without any weapons)
// TODO - this needs to be reworked to include the multiple steps link
// can take.
var basicSprites [numDirections][numSteps]*pixel.Sprite

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

	// Load the linkSprite sprites for each direction and step.
	basicSprites[North][right] = pixel.NewSprite(spritesheet, pixel.R(69.0, 283.0, 69.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[North][left] = pixel.NewSprite(spritesheet, pixel.R(86.0, 283.0, 86.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[East][right] = pixel.NewSprite(spritesheet, pixel.R(35.0, 283.0, 35.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[East][left] = pixel.NewSprite(spritesheet, pixel.R(52.0, 283.0, 52.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[South][right] = pixel.NewSprite(spritesheet, pixel.R(1.0, 283.0, 1.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[South][left] = pixel.NewSprite(spritesheet, pixel.R(18.0, 283.0, 18.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[West][right] = pixel.NewSprite(flippedSpriteSheet, pixel.R(320.0, 283.0, 320.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))
	basicSprites[West][left] = pixel.NewSprite(flippedSpriteSheet, pixel.R(303.0, 283.0, 303.0+nes.SpriteSize.X, 283.0+nes.SpriteSize.Y))

	// Start off link facing south (forward)
	currentlyFacing = South
	currentStep = right
	currentRawStep = float64(currentStep)
	linkSprite = basicSprites[currentlyFacing][currentStep]
	currentXYLocation = nes.StartingLocation
}

// Draw draws link to the target.
func Draw(target pixel.Target) {
	once.Do(initializeLink)
	linkSprite.Draw(target, pixel.IM.Moved(currentXYLocation).Scaled(pixel.ZV, utils.NESToCurrentScaling))
}

// Move takes a direction and a deltaTime in seconds, and makes sure Link is
// turned in the direction given, turning him if he is not already facing
// that way, and the moving him the appropriate amount.
func Move(direction Direction, deltaTime float64) {
	currentlyFacing = direction
	currentRawStep = currentRawStep + stepPerSecond*deltaTime
	_, currentSubStep := math.Modf(currentRawStep)
	currentStep = Step(math.Round(currentSubStep))
	currentXYLocation = currentXYLocation.Add(speed[direction].Scaled(deltaTime))
	linkSprite = basicSprites[currentlyFacing][currentStep]
}
