package main

import (
	"log"

	"github.com/avbar/maze/internal/common"
	"github.com/avbar/maze/internal/gameengine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var Settings = common.Settings{
	Cols:  10,
	Rows:  10,
	Speed: 6,
}

func main() {
	g := gameengine.NewGameEngine(ScreenWidth, ScreenHeight, Settings)

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Maze")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
