package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	titleFontSize = 28
	fontSize      = 20
)

var (
	textColor       = color.RGBA{0x80, 0x80, 0x80, 0xff}
	titleArcadeFont font.Face
	arcadeFont      font.Face
)

func init() {
	const dpi = 72

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}

	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) drawTitle(screen *ebiten.Image) {
	var titleText string
	switch g.mode {
	case ModeTitle:
		titleText = "MAZE"
	case ModeWin:
		titleText = "YOU WIN"
	case ModeLose:
		titleText = "YOU LOSE"
	}

	texts := []string{"PRESS SPACE BAR OR CLICK TO START", "", "USE ARROW KEYS TO CONTROL THE GOPHER"}

	x := (g.screenWidth - len(titleText)*titleFontSize) / 2
	y := g.screenHeight / 4
	text.Draw(screen, titleText, titleArcadeFont, x, y, textColor)

	for i, t := range texts {
		x := (g.screenWidth - len(t)*fontSize) / 2
		y := g.screenHeight*2/3 + i*fontSize
		text.Draw(screen, t, arcadeFont, x, y, textColor)
	}
}
