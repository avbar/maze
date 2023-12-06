package cookie

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/common"

	"github.com/hajimehoshi/ebiten/v2"
)

type Cookie struct {
	// Maze position
	pos common.Pos

	// Data for drawing
	width     float64
	height    float64
	screenPos common.Vector
	sprite    *ebiten.Image
}

func NewCookie(pos common.Pos, width, height float64) *Cookie {
	screenPos := common.Vector{
		X: float64(pos.Col) * width,
		Y: float64(pos.Row) * height,
	}

	return &Cookie{
		pos:       pos,
		width:     width,
		height:    height,
		screenPos: screenPos,
		sprite:    assets.CookieSprite,
	}
}

func (c *Cookie) Pos() common.Pos {
	return c.pos
}

func (c *Cookie) Draw(screen *ebiten.Image) {
	bounds := c.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	scaleX := c.width / float64(bounds.Dx())
	scaleY := c.height / float64(bounds.Dy())
	scaleFactor := scaleX
	if scaleY < scaleFactor {
		scaleFactor = scaleY
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Scale(scaleFactor, scaleFactor)
	op.GeoM.Translate(c.width/2, c.height/2)

	op.GeoM.Translate(c.screenPos.X, c.screenPos.Y)

	screen.DrawImage(c.sprite, op)
}
