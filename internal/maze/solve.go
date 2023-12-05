package maze

import (
	"github.com/avbar/maze/internal/common"
)

func (m *Maze) Solve(start, finish common.Pos) common.Path {
	steps := m.makeWave(start, finish)
	return m.makePath(steps, start)
}

func (m *Maze) makeWave(start, finish common.Pos) [][]int {
	steps := make([][]int, m.rows)
	for i := 0; i < m.rows; i++ {
		steps[i] = make([]int, m.cols)
	}

	steps[finish.Row][finish.Col] = 1
	cells := []common.Pos{finish}

	for len(cells) > 0 {
		c, r := cells[0].Col, cells[0].Row
		step := steps[r][c]

		if c > 0 && steps[r][c-1] == 0 && !m.vWalls[r][c-1] {
			steps[r][c-1] = step + 1
			if r == start.Row && c-1 == start.Col {
				break
			}
			cells = append(cells, common.Pos{Col: c - 1, Row: r})
		}

		if c < m.cols-1 && steps[r][c+1] == 0 && !m.vWalls[r][c] {
			steps[r][c+1] = step + 1
			if r == start.Row && c+1 == start.Col {
				break
			}
			cells = append(cells, common.Pos{Col: c + 1, Row: r})
		}

		if r > 0 && steps[r-1][c] == 0 && !m.hWalls[r-1][c] {
			steps[r-1][c] = step + 1
			if r-1 == start.Row && c == start.Col {
				break
			}
			cells = append(cells, common.Pos{Col: c, Row: r - 1})
		}

		if r < m.rows-1 && steps[r+1][c] == 0 && !m.hWalls[r][c] {
			steps[r+1][c] = step + 1
			if r+1 == start.Row && c == start.Col {
				break
			}
			cells = append(cells, common.Pos{Col: c, Row: r + 1})
		}

		cells = cells[1:]
	}

	return steps
}

func (m *Maze) makePath(steps [][]int, start common.Pos) common.Path {
	var p common.Path
	c, r := start.Col, start.Row

	for steps[r][c] > 1 {
		if c > 0 && steps[r][c-1] == steps[r][c]-1 && !m.vWalls[r][c-1] {
			c -= 1
		} else if c < m.cols-1 && steps[r][c+1] == steps[r][c]-1 && !m.vWalls[r][c] {
			c += 1
		} else if r > 0 && steps[r-1][c] == steps[r][c]-1 && !m.hWalls[r-1][c] {
			r -= 1
		} else if r < m.rows-1 && steps[r+1][c] == steps[r][c]-1 && !m.hWalls[r][c] {
			r += 1
		}
		p = append(p, common.Pos{Col: c, Row: r})
	}

	return p
}
