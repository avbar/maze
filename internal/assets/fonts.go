package assets

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	dpi              = 72
	TitleFontSize    = 28
	SubtitleFontSize = 20
	MenuFontSize     = 20
)

var (
	TitleFont    = mustLoadFont(fonts.PressStart2P_ttf, TitleFontSize)
	SubtitleFont = mustLoadFont(fonts.PressStart2P_ttf, SubtitleFontSize)
	MenuFont     = mustLoadFont(fonts.MPlus1pRegular_ttf, MenuFontSize)
)

func mustLoadFont(src []byte, size int) font.Face {
	tt, err := opentype.Parse(src)
	if err != nil {
		panic(err)
	}

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return font
}
