package game

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	directionRight = 1.0
	directionLeft  = -1.0
)

type Player struct {
	// Maze position
	pos common.Pos

	// Data for drawing
	width     float64
	height    float64
	screenPos common.Vector
	direction float64
	sprite    *ebiten.Image
	spriteWin *ebiten.Image
}

func NewPlayer(pos common.Pos, width, height float64) *Player {
	screenPos := common.Vector{
		X: float64(pos.Col) * width,
		Y: float64(pos.Row) * height,
	}

	return &Player{
		pos:       pos,
		width:     width,
		height:    height,
		screenPos: screenPos,
		direction: directionRight,
		sprite:    assets.PlayerSprite,
		spriteWin: assets.PlayerWinSprite,
	}
}

func (p *Player) Pos() common.Pos {
	return p.pos
}

func (p *Player) Win() {
	p.sprite = p.spriteWin
}

func (p *Player) Update(pos common.Pos) {
	hStep := pos.Col - p.pos.Col

	p.pos = pos
	p.screenPos.X = float64(p.pos.Col) * p.width
	p.screenPos.Y = float64(p.pos.Row) * p.height

	if hStep < 0 {
		p.direction = directionLeft
	} else if hStep > 0 {
		p.direction = directionRight
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
