package player

import (
	"github.com/avbar/maze/internal/assets"
	"github.com/avbar/maze/internal/maze"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Cell struct {
	col int
	row int
}

type Vector struct {
	X float64
	Y float64
}

const (
	DirectionRight = 1.0
	DirectionLeft  = -1.0
)

type Player struct {
	// Game data
	cell Cell
	maze *maze.Maze
	// Data for drawing
	width     float64
	height    float64
	screenPos Vector
	direction float64
	sprite    *ebiten.Image
}

func NewPlayer(maze *maze.Maze, col, row int) *Player {
	return &Player{
		cell:      Cell{col: col, row: row},
		maze:      maze,
		width:     maze.ColumnWidth(),
		height:    maze.RowHeight(),
		screenPos: Vector{},
		direction: 1,
		sprite:    assets.PlayerSprite,
	}
}

func (p *Player) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && !p.maze.IsBottomWall(p.cell.col, p.cell.row) {
		p.screenPos.Y += p.height
		p.cell.row++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && !p.maze.IsTopWall(p.cell.col, p.cell.row) {
		p.screenPos.Y -= p.height
		p.cell.row--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && !p.maze.IsLeftWall(p.cell.col, p.cell.row) {
		p.screenPos.X -= p.width
		p.cell.col--
		p.direction = DirectionLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && !p.maze.IsRightWall(p.cell.col, p.cell.row) {
		p.screenPos.X += p.width
		p.cell.col++
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
