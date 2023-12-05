package game

import (
	"math/rand"
	"time"

	"github.com/avbar/maze/internal/common"
	"github.com/avbar/maze/internal/cookie"
	"github.com/avbar/maze/internal/maze"
	"github.com/avbar/maze/internal/player"
	"github.com/avbar/maze/internal/rival"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	screenWidth  int
	screenHeight int
	maze         *maze.Maze
	player       *player.Player
	rival        *rival.Rival
	rivalTimer   *common.Timer
	cookie       *cookie.Cookie
}

func NewGame(screenWidth, screenHeight int, cols, rows int) *Game {
	colWidth := float64(screenWidth) / float64(cols)
	rowHeight := float64(screenHeight) / float64(rows)

	maze := maze.NewMaze(cols, rows, colWidth, rowHeight)
	cookieCol := rand.Intn(cols)
	path := maze.Solve(common.Pos{Col: cols - 1, Row: 0}, common.Pos{Col: cookieCol, Row: rows - 1})

	player := player.NewPlayer(0, 0, colWidth, rowHeight)
	rival := rival.NewRival(cols-1, 0, colWidth, rowHeight, path)
	rivalTimer := common.NewTimer(500 * time.Millisecond)
	cookie := cookie.NewCookie(cookieCol, rows-1, colWidth, rowHeight)

	return &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		maze:         maze,
		player:       player,
		rival:        rival,
		rivalTimer:   rivalTimer,
		cookie:       cookie,
	}
}

func (g *Game) Reset() {
	cols := g.maze.Cols()
	rows := g.maze.Rows()
	colWidth := float64(g.screenWidth) / float64(cols)
	rowHeight := float64(g.screenHeight) / float64(rows)

	g.maze = maze.NewMaze(cols, rows, colWidth, rowHeight)
	cookieCol := rand.Intn(cols)
	path := g.maze.Solve(common.Pos{Col: cols - 1, Row: 0}, common.Pos{Col: cookieCol, Row: rows - 1})

	g.player = player.NewPlayer(0, 0, colWidth, rowHeight)
	g.rival = rival.NewRival(cols-1, 0, colWidth, rowHeight, path)
	g.cookie = cookie.NewCookie(cookieCol, rows-1, colWidth, rowHeight)
}

func (g *Game) playerNextStep() common.Pos {
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

	return common.Pos{
		Col: col,
		Row: row,
	}
}

func (g *Game) Update() error {
	if g.player.Pos() == g.cookie.Pos() || g.rival.Pos() == g.cookie.Pos() {
		// Game is over, start new game
		g.Reset()
	}

	if g.player.Pos() == g.rival.Pos() {
		// Scatter players randomly
		cols := g.maze.Cols()
		rows := g.maze.Rows()
		g.player.Update(rand.Intn(cols), rand.Intn(rows))

		rivalCol, rivalRow := rand.Intn(cols), rand.Intn(rows)
		path := g.maze.Solve(common.Pos{Col: rivalCol, Row: rivalRow}, g.cookie.Pos())
		g.rival.SetPos(rivalCol, rivalRow, path)
	} else {
		// Next step
		pos := g.playerNextStep()
		g.player.Update(pos.Col, pos.Row)

		g.rivalTimer.Update()
		if g.rivalTimer.IsReady() {
			g.rivalTimer.Reset()
			g.rival.Update()
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.maze.Draw(screen)
	g.cookie.Draw(screen)
	g.player.Draw(screen)
	g.rival.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
