package maze

import (
	"math/rand"
)

// Creates random maze using Eller's Algorithm
func (m *Maze) Generate() {
	m.initWalls()

	sets := make([]int, m.cols)

	currSet := 0
	for i := 0; i < m.rows; i++ {
		// Assign sets for cells in the row
		for j := 0; j < m.cols; j++ {
			if sets[j] == 0 {
				currSet++
				sets[j] = currSet
			}
		}

		// Vertical walls
		for j := 0; j < m.cols-1; j++ {
			if randomBool() || sets[j] == sets[j+1] {
				m.vWalls[i][j] = true
			} else {
				for k := m.cols - 1; k > j; k-- {
					if sets[k] == sets[j+1] {
						sets[k] = sets[j]
					}
				}
			}
		}
		m.vWalls[i][m.cols-1] = true

		if i < m.rows-1 {
			// Horizontal walls
			for j := 0; j < m.cols; j++ {
				if randomBool() {
					emptyBottom := 0
					for k := 0; k < m.cols; k++ {
						if sets[k] == sets[j] && !m.hWalls[i][k] {
							emptyBottom++
							if emptyBottom > 1 {
								m.hWalls[i][j] = true
								break
							}
						}
					}
				}
			}

			// Copy current row to next
			for j := 0; j < m.cols; j++ {
				if m.hWalls[i][j] {
					sets[j] = 0
				}
			}
		} else {
			// Last row
			for j := 0; j < m.cols; j++ {
				m.hWalls[i][j] = true
			}

			for j := 0; j < m.cols-1; j++ {
				if sets[j] != sets[j+1] {
					m.vWalls[i][j] = false

					for k := m.cols - 1; k > j; k-- {
						if sets[k] == sets[j+1] {
							sets[k] = sets[j]
						}
					}
				}
			}
		}
	}
}

func allocateWalls(cols, rows int) walls {
	w := make(walls, rows)
	for i := 0; i < rows; i++ {
		w[i] = make([]bool, cols)
	}

	return w
}

func (m *Maze) initWalls() {
	m.vWalls = allocateWalls(m.cols, m.rows)
	m.hWalls = allocateWalls(m.cols, m.rows)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}
