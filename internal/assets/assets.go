package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed png/*
var assets embed.FS

var (
	PlayerSprite    = mustLoadImage("png/gopher.png")
	PlayerWinSprite = mustLoadImage("png/gopher_win.png")
	RivalSprite     = mustLoadImage("png/brown_gopher.png")
	RivalWinSprite  = mustLoadImage("png/brown_gopher_win.png")
	CookieSprite    = mustLoadImage("png/cookie.png")
)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
