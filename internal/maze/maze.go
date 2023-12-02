package maze

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	fillColor = color.RGBA{0xe8, 0xe9, 0xeb, 0xff}
	lineColor = color.Black
)

type walls [][]bool

type Maze struct {
	// Maze settings
	cols   int
	rows   int
	vWalls walls
	hWalls walls
	// Data for drawing
	colWidth  float64
	rowHeight float64
}

func NewMaze(cols, rows int, colWidth, rowHeight float64) *Maze {
	m := &Maze{
		cols:      cols,
		rows:      rows,
		colWidth:  colWidth,
		rowHeight: rowHeight,
	}

	m.Generate()

	return m
}

func (m *Maze) IsLeftWall(col, row int) bool {
	return col == 0 || m.vWalls[row][col-1]
}

func (m *Maze) IsRightWall(col, row int) bool {
	return m.vWalls[row][col]
}

func (m *Maze) IsTopWall(col, row int) bool {
	return row == 0 || m.hWalls[row-1][col]
}

func (m *Maze) IsBottomWall(col, row int) bool {
	return m.hWalls[row][col]
}

func (m *Maze) ColumnWidth() float64 {
	return m.colWidth
}

func (m *Maze) RowHeight() float64 {
	return m.rowHeight
}

func (m *Maze) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := float32(m.colWidth)*float32(m.cols), float32(m.rowHeight)*float32(m.rows)

	// Draw background
	vector.DrawFilledRect(screen, 0, 0, screenWidth, screenHeight, fillColor, false)

	// Draw top and left borders
	vector.StrokeLine(screen, 0, 0, screenWidth, 0, 2, lineColor, false)
	vector.StrokeLine(screen, 0, 0, 0, screenHeight, 2, lineColor, false)

	// Draw cells
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			// Top left corner
			x0, y0 := float32(j)*float32(m.colWidth), float32(i)*float32(m.rowHeight)
			// Bottom right corner
			x1, y1 := float32(j+1)*float32(m.colWidth), float32(i+1)*float32(m.rowHeight)
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
