package game

import (
	"github.com/avbar/maze/internal/maze"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

type Game struct {
	maze *maze.Maze
}

func NewGame(cols, rows int) *Game {
	return &Game{
		maze: maze.NewMaze(ScreenWidth, ScreenHeight, cols, rows),
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.maze.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
