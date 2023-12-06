package player

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/common"
)

func NewRival(pos common.Pos, width, height float64) *Player {
	rival := NewPlayer(pos, width, height)
	rival.direction = directionLeft
	rival.sprite = assets.RivalSprite
	return rival
}
