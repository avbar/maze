package game

import (
	"github.com/avbar/maze/internal/maze"
	"github.com/avbar/maze/internal/player"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	screenWidth  int
	screenHeight int
	maze         *maze.Maze
	player       *player.Player
}

func NewGame(screenWidth, screenHeight int, cols, rows int) *Game {
	colWidth := float64(screenWidth) / float64(cols)
	rowHeight := float64(screenHeight) / float64(rows)

	maze := maze.NewMaze(cols, rows, colWidth, rowHeight)
	player := player.NewPlayer(maze, 0, 0)

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		maze:         maze,
		player:       player,
	}
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.maze.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
