package main

import (
	"log"

	"github.com/avbar/maze/internal/config"
	"github.com/avbar/maze/internal/gameengine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	cfg := config.MustLoad()
	log.Printf("config: %+v", *cfg)

	log.Println("starting the game")
	g := gameengine.NewGameEngine(cfg.ScreenWidth, cfg.ScreenHeight, cfg.Settings)

	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle("Maze")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
