package game

import (
	"math/rand"

	"github.com/avbar/maze/internal/cookie"
	"github.com/avbar/maze/internal/maze"
	"github.com/avbar/maze/internal/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	screenWidth  int
	screenHeight int
	maze         *maze.Maze
	player       *player.Player
	cookie       *cookie.Cookie
}

func NewGame(screenWidth, screenHeight int, cols, rows int) *Game {
	colWidth := float64(screenWidth) / float64(cols)
	rowHeight := float64(screenHeight) / float64(rows)

	maze := maze.NewMaze(cols, rows, colWidth, rowHeight)
	player := player.NewPlayer(0, 0, colWidth, rowHeight)
	cookie := cookie.NewCookie(rand.Intn(cols), rows-1, colWidth, rowHeight)

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		maze:         maze,
		player:       player,
		cookie:       cookie,
	}
}

func (g *Game) Reset() {
	cols := g.maze.Cols()
	rows := g.maze.Rows()
	colWidth := float64(g.screenWidth) / float64(cols)
	rowHeight := float64(g.screenHeight) / float64(rows)

	g.maze = maze.NewMaze(cols, rows, colWidth, rowHeight)
	g.player = player.NewPlayer(0, 0, colWidth, rowHeight)
	g.cookie = cookie.NewCookie(rand.Intn(cols), rows-1, colWidth, rowHeight)
}

func (g *Game) Update() error {
	if g.player.Pos() == g.cookie.Pos() {
		g.Reset()
	}

	col, row := g.player.Pos().Col, g.player.Pos().Row
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && !g.maze.IsBottomWall(col, row) {
		row++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && !g.maze.IsTopWall(col, row) {
		row--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && !g.maze.IsLeftWall(col, row) {
		col--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && !g.maze.IsRightWall(col, row) {
		col++
	}

	g.player.Update(col, row)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.maze.Draw(screen)
	g.cookie.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
