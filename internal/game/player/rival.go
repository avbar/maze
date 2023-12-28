package player

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/game/coord"
)

func NewRival(pos coord.Pos, width, height float64) *Player {
	rival := NewPlayer(pos, width, height)
	rival.direction = directionLeft
	rival.sprite = assets.RivalSprite
	rival.spriteWin = assets.RivalWinSprite
	return rival
}
