package gameengine

import (
	"github.com/avbar/maze/internal/common"
	"github.com/avbar/maze/internal/game"
	"github.com/avbar/maze/internal/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type GameEngine struct {
	screenWidth  int
	screenHeight int
	settings     common.Settings
	game         *game.Game
	menu         *menu.Menu
	scene        Scene
}

func NewGameEngine(screenWidth, screenHeight int, cols, rows int) *GameEngine {
	settings := common.Settings{
		Cols: cols,
		Rows: rows,
	}
	g := &GameEngine{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		settings:     settings,
	}

	g.game = game.NewGame(screenWidth, screenHeight, cols, rows, g.switchToMenu)
	g.menu = menu.NewMenu(settings, g.switchToGame)
	g.scene = g.game

	return g
}

func (g *GameEngine) switchToMenu() {
	g.scene = g.menu
}

func (g *GameEngine) switchToGame() {
	if g.settings != g.menu.Settings() {
		g.settings = g.menu.Settings()
		g.game = game.NewGame(g.screenWidth, g.screenHeight, g.settings.Cols, g.settings.Rows, g.switchToMenu)
	}
	g.scene = g.game
}

func (g *GameEngine) Update() error {
	g.scene.Update()
	return nil
}

func (g *GameEngine) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *GameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}
