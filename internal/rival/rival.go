package rival

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/common"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	DirectionRight = 1.0
	DirectionLeft  = -1.0
)

type Rival struct {
	// Game data
	pos  common.Pos
	path common.Path
	// Data for drawing
	width     float64
	height    float64
	screenPos common.Vector
	direction float64
	sprite    *ebiten.Image
}

func NewRival(col, row int, width, height float64, path common.Path) *Rival {
	return &Rival{
		pos:       common.Pos{Col: col, Row: row},
		path:      path,
		width:     width,
		height:    height,
		screenPos: common.Vector{X: float64(col) * width, Y: float64(row) * height},
		direction: 1,
		sprite:    assets.RivalSprite,
	}
}

func (r *Rival) Pos() common.Pos {
	return r.pos
}

func (r *Rival) Update() {
	if len(r.path) == 0 {
		return
	}

	col, row := r.path[0].Col, r.path[0].Row
	r.path = r.path[1:]

	hStep := col - r.pos.Col

	r.pos.Col = col
	r.pos.Row = row
	r.screenPos.X = float64(r.pos.Col) * r.width
	r.screenPos.Y = float64(r.pos.Row) * r.height

	if hStep < 0 {
		r.direction = DirectionRight
	} else if hStep > 0 {
		r.direction = DirectionLeft
	}
}

func (r *Rival) Draw(screen *ebiten.Image) {
	bounds := r.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	scaleX := r.width / float64(bounds.Dx())
	scaleY := r.height / float64(bounds.Dy())
	scaleFactor := scaleX
	if scaleY < scaleFactor {
		scaleFactor = scaleY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Scale(scaleFactor*float64(r.direction), scaleFactor)
	op.GeoM.Translate(r.width/2, r.height/2)

	op.GeoM.Translate(r.screenPos.X, r.screenPos.Y)

	screen.DrawImage(r.sprite, op)
}
