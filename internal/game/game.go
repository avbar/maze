package game

import (
	"math/rand"
	"time"

	"github.com/avbar/maze/internal/common"
	"github.com/avbar/maze/internal/cookie"
	"github.com/avbar/maze/internal/maze"
	"github.com/avbar/maze/internal/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeWin
	ModeLose
)

type Game struct {
	screenWidth  int
	screenHeight int
	mode         Mode
	maze         *maze.Maze
	player       *player.Player
	rival        *player.Player
	rivalPath    common.Path
	rivalTimer   *common.Timer
	cookie       *cookie.Cookie
	playerPoints int
	rivalPoints  int
	speed        int
	showMenu     func()
}

func NewGame(screenWidth, screenHeight int, settings common.Settings, showMenu func()) *Game {
	colWidth := float64(screenWidth) / float64(settings.Cols)
	rowHeight := float64(screenHeight) / float64(settings.Rows)

	maze := maze.NewMaze(settings.Cols, settings.Rows, colWidth, rowHeight)

	game := &Game{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		mode:         ModeTitle,
		maze:         maze,
		speed:        settings.Speed,
		showMenu:     showMenu,
	}

	game.Reset()

	return game
}

func (g *Game) Reset() {
	cols := g.maze.Cols()
	rows := g.maze.Rows()
	colWidth := float64(g.screenWidth) / float64(cols)
	rowHeight := float64(g.screenHeight) / float64(rows)

	g.maze.Generate()

	playerPos := common.Pos{
		Col: 0,
		Row: 0,
	}
	rivalPos := common.Pos{
		Col: cols - 1,
		Row: 0,
	}
	cookiePos := common.Pos{
		Col: rand.Intn(cols),
		Row: rows - 1,
	}

	g.player = player.NewPlayer(playerPos, colWidth, rowHeight)
	g.rival = player.NewRival(rivalPos, colWidth, rowHeight)
	g.rivalPath = g.maze.Solve(rivalPos, cookiePos)
	g.rivalTimer = common.NewTimer(g.timerDuration())
	g.cookie = cookie.NewCookie(cookiePos, colWidth, rowHeight)
}

func (g *Game) timerDuration() time.Duration {
	return time.Duration(100*(11-g.speed)) * time.Millisecond
}

func (g *Game) playerStep() common.Pos {
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

func (g *Game) rivalStep() common.Pos {
	if len(g.rivalPath) == 0 {
		return g.rival.Pos()
	}

	pos := g.rivalPath[0]
	g.rivalPath = g.rivalPath[1:]
	return pos
}

func (g *Game) isStartKeyPressed() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}

func (g *Game) isMenuKeyPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyF10)
}

func (g *Game) Update() {
	if g.isMenuKeyPressed() {
		g.showMenu()
		return
	}

	switch g.mode {
	case ModeTitle:
		if g.isStartKeyPressed() {
			g.mode = ModeGame
		}

	case ModeGame:
		if g.player.Pos() == g.cookie.Pos() {
			g.mode = ModeWin
			g.player.Win()
			g.playerPoints++
		} else if g.rival.Pos() == g.cookie.Pos() {
			g.mode = ModeLose
			g.rival.Win()
			g.rivalPoints++
		} else if g.player.Pos() == g.rival.Pos() {
			// Scatter players randomly
			cols := g.maze.Cols()
			rows := g.maze.Rows()

			playerPos := common.Pos{
				Col: rand.Intn(cols),
				Row: rand.Intn(rows),
			}
			rivalPos := common.Pos{
				Col: rand.Intn(cols),
				Row: rand.Intn(rows),
			}

			g.player.Update(playerPos)
			g.rival.Update(rivalPos)
			g.rivalPath = g.maze.Solve(rivalPos, g.cookie.Pos())
		} else {
			// Next step
			g.player.Update(g.playerStep())

			g.rivalTimer.Update()
			if g.rivalTimer.IsReady() {
				g.rivalTimer.Reset()
				g.rival.Update(g.rivalStep())
			}
		}

	case ModeWin, ModeLose:
		if g.isStartKeyPressed() {
			g.mode = ModeGame
			g.Reset()
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.maze.Draw(screen)
	if g.mode == ModeTitle || g.mode == ModeGame {
		g.cookie.Draw(screen)
	}
	g.player.Draw(screen)
	g.rival.Draw(screen)

	if g.mode == ModeTitle || g.mode == ModeWin || g.mode == ModeLose {
		g.drawTitle(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
