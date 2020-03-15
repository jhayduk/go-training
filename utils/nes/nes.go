// Package nes contains contants and utilities specifically
// geared towards the nes system, or The Legend of Zelda game on the nes
// system. The main distiction is that sizes and coordinates are the
// originals for the NES system regardless of any scaling that might be
// done externally when the game is displayed now on a screen.
package nes

import (
	"github.com/faiface/pixel"
)

// ScreenSize is a vector representing the total screen display size of an
// NES on a NTSC TV in pixels.
// Technically, the total screen size for an NES system was 256 x 240, but
// only 256 x 224 was normally shown.
// Reference: http://forums.nesdev.com/viewtopic.php?f=2&t=10141
var ScreenSize pixel.Vec = pixel.V(256.0, 224.0)

// StatsWindowSize is a vector representing the 256 x 48 display area at
// the top of the screen in The Legend of Zelda game where the stats
// information would be dislayed.
var StatsWindowSize pixel.Vec = pixel.V(256.0, 224.0)

// MapSize is a vector representing the size of the map in The Legend of Zelda
// game displayed underneath the stats window. This region fills the rest of
// the screen.
var MapSize pixel.Vec = pixel.V(256.0, 176.0)

// StartingLocation is a vector representing the starting location in the
// overworld where the adventure starts. Link should be placed centered at
// this location, and the camera should start centered in that location as
// well. On the spritesheet, this is in the center of the 8th map on the
// bottom 1st row of maps. This is essentially at the center of the bottom
// of the overworld.
var StartingLocation pixel.Vec = MapSize.ScaledXY(pixel.V(8.0, 1.0))

// SpriteSize is a vector representing the size of an individual sprite in
// pixels.
// Most characters (except bosses) are one sprite in size.
var SpriteSize pixel.Vec = pixel.V(16.0, 16.0)
