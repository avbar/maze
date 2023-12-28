package maze

type Walls [][]bool

type Maze struct {
	// Size
	cols int
	rows int
	// Vertical walls
	vWalls Walls
	// Horizontal walls
	hWalls Walls

	// Data for drawing
	colWidth  float64
	rowHeight float64
}

func NewMaze(cols, rows int, colWidth, rowHeight float64) *Maze {
	return &Maze{
		cols:      cols,
		rows:      rows,
		colWidth:  colWidth,
		rowHeight: rowHeight,
	}
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

func (m *Maze) Cols() int {
	return m.cols
}

func (m *Maze) Rows() int {
	return m.rows
}
