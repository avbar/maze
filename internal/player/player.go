package player

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DirectionRight = 1.0
	DirectionLeft  = -1.0
)

type Player struct {
	// Game data
	pos common.Pos
	// Data for drawing
	width     float64
	height    float64
	screenPos common.Vector
	direction float64
	sprite    *ebiten.Image
}

func NewPlayer(col, row int, width, height float64) *Player {
	return &Player{
		pos:       common.Pos{Col: col, Row: row},
		width:     width,
		height:    height,
		screenPos: common.Vector{},
		direction: 1,
		sprite:    assets.PlayerSprite,
	}
}

func (p *Player) Pos() common.Pos {
	return p.pos
}

func (p *Player) Update(col, row int) {
	hStep := col - p.pos.Col

	p.pos.Col = col
	p.pos.Row = row
	p.screenPos.X = float64(p.pos.Col) * p.width
	p.screenPos.Y = float64(p.pos.Row) * p.height

	if hStep < 0 {
		p.direction = DirectionLeft
	} else if hStep > 0 {
		p.direction = DirectionRight
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	scaleX := p.width / float64(bounds.Dx())
	scaleY := p.height / float64(bounds.Dy())
	scaleFactor := scaleX
	if scaleY < scaleFactor {
		scaleFactor = scaleY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Scale(scaleFactor*float64(p.direction), scaleFactor)
	op.GeoM.Translate(p.width/2, p.height/2)

	op.GeoM.Translate(p.screenPos.X, p.screenPos.Y)

	screen.DrawImage(p.sprite, op)
}
