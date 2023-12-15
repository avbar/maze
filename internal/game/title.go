package game

import (
	"fmt"
	"image/color"

	"github.com/avbar/maze/internal/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	textColor = color.RGBA{0x80, 0x80, 0x80, 0xff}
)

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

	x := (g.screenWidth - len(titleText)*assets.TitleFontSize) / 2
	y := g.screenHeight / 4
	text.Draw(screen, titleText, assets.TitleFont, x, y, textColor)

	texts := []string{"PRESS SPACE BAR OR CLICK TO START", "", "USE ARROW KEYS TO CONTROL THE GOPHER",
		"", "", "", "PRESS F10 TO CHANGE THE SETTINGS"}

	for i, t := range texts {
		x := (g.screenWidth - len(t)*assets.SubtitleFontSize) / 2
		y := g.screenHeight*2/3 + i*assets.SubtitleFontSize
		text.Draw(screen, t, assets.SubtitleFont, x, y, textColor)
	}

	// Score
	if g.mode == ModeWin || g.mode == ModeLose {
		x := g.screenWidth / 10
		y := g.screenHeight / 10
		text.Draw(screen, fmt.Sprintf("%03d", g.playerPoints), assets.TitleFont, x, y, textColor)
		x = g.screenWidth*9/10 - 3*assets.TitleFontSize
		text.Draw(screen, fmt.Sprintf("%03d", g.rivalPoints), assets.TitleFont, x, y, textColor)
	}
}
