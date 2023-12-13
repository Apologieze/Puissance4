package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Constantes définissant les paramètres généraux du programme.
const (
	globalWidth         = globalNumTilesX * globalTileSize
	globalHeight        = (globalNumTilesY + 1) * globalTileSize
	globalTileSize      = 100
	globalNumTilesX     = 7
	globalNumTilesY     = 6
	globalCircleMargin  = 5
	globalBlinkDuration = 60
	globalNumColorLine  = 3
	globalNumColorCol   = 3
	globalNumColor      = globalNumColorLine * globalNumColorCol
)

// Variables définissant les paramètres généraux du programme.
var (
	globalBackgroundColor     color.Color = color.NRGBA{R: 176, G: 196, B: 222, A: 255}
	globalGridColor           color.Color = color.NRGBA{R: 119, G: 136, B: 153, A: 255}
	globalTextColor           color.Color = color.NRGBA{R: 25, G: 25, B: 5, A: 255}
	globalSelectColor         color.Color = color.NRGBA{R: 25, G: 25, B: 5, A: 255}
	globalOpponentSelectColor color.Color = color.NRGBA{R: 120, G: 120, B: 120, A: 255}
	smallFont                 font.Face
	largeFont                 font.Face
	globalTokenColors         [globalNumColor]color.Color = [globalNumColor]color.Color{
		color.NRGBA{R: 255, G: 201, B: 54, A: 255},
		color.NRGBA{R: 0, G: 196, B: 46, A: 255},
		color.NRGBA{R: 25, G: 252, B: 215, A: 255},
		color.NRGBA{R: 232, G: 85, B: 0, A: 255},
		color.NRGBA{R: 226, G: 255, B: 59, A: 255},
		color.NRGBA{R: 255, G: 0, B: 195, A: 255},
		color.NRGBA{R: 84, G: 255, B: 101, A: 255},
		color.NRGBA{R: 168, G: 8, B: 255, A: 255},
		color.NRGBA{R: 25, G: 40, B: 252, A: 255},
	}
	offScreenImage *ebiten.Image
)
