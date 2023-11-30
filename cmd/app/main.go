package main

import (
	"log"

	"github.com/avbar/maze/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	mazeCols = 10
	mazeRows = 10
)

func main() {
	g := game.NewGame(mazeCols, mazeRows)

	ebiten.SetWindowTitle("Maze")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
