package main

import (
	"log"

	"github.com/avbar/maze/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	Cols         = 10
	Rows         = 10
)

func main() {
	g := game.NewGame(ScreenWidth, ScreenHeight, Cols, Rows)

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Maze")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
