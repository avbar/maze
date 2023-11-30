package maze

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type walls [][]bool

type Maze struct {
	cols      int
	rows      int
	colWidth  float32
	rowHeight float32
	vWalls    walls
	hWalls    walls
}

func NewMaze(width, height int, cols, rows int) *Maze {
	m := &Maze{
		cols:      cols,
		rows:      rows,
		colWidth:  float32(width) / float32(cols),
		rowHeight: float32(height) / float32(rows),
	}

	m.Generate()

	return m
}

func (m *Maze) Draw(screen *ebiten.Image) {
	fillColor, lineColor := color.RGBA{0xe8, 0xe9, 0xeb, 0xff}, color.Black
	screenWidth, screenHeight := m.colWidth*float32(m.cols), m.rowHeight*float32(m.rows)

	// Draw background
	vector.DrawFilledRect(screen, 0, 0, screenWidth, screenHeight, fillColor, false)

	// Draw top and left borders
	vector.StrokeLine(screen, 0, 0, screenWidth, 0, 2, lineColor, false)
	vector.StrokeLine(screen, 0, 0, 0, screenHeight, 2, lineColor, false)

	// Draw cells
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			// Top left corner
			x0, y0 := float32(j)*m.colWidth, float32(i)*m.rowHeight
			// Bottom right corner
			x1, y1 := float32(j+1)*m.colWidth, float32(i+1)*m.rowHeight
			// Left wall
			if m.vWalls[i][j] {
				vector.StrokeLine(screen, x1, y0, x1, y1, 1, lineColor, false)
			}
			// Bottom wall
			if m.hWalls[i][j] {
				vector.StrokeLine(screen, x0, y1, x1, y1, 1, lineColor, false)
			}
		}
	}
}
