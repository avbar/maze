package maze_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/avbar/maze/internal/game/coord"
	"github.com/avbar/maze/internal/game/maze"
)

type walls struct {
	vWalls maze.Walls
	hWalls maze.Walls
}

func isEqual(cols, rows int, w walls, m *maze.Maze) bool {
	if cols != m.Cols() || rows != m.Rows() {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if w.vWalls[i][j] != m.IsRightWall(j, i) ||
				w.hWalls[i][j] != m.IsBottomWall(j, i) {
				return false
			}
		}
	}

	return true
}

func TestGenerate(t *testing.T) {
	type args struct {
		cols  int
		rows  int
		walls walls
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "generate maze 10x10",
			args: args{
				cols: 10,
				rows: 10,
				walls: walls{
					vWalls: maze.Walls{
						{true, false, true, true, true, true, false, true, true, true},
						{false, true, true, true, false, false, false, false, false, true},
						{false, false, true, true, true, false, true, true, false, true},
						{true, true, false, false, true, true, true, false, false, true},
						{true, true, false, true, true, false, true, false, false, true},
						{true, false, false, true, true, true, true, true, false, true},
						{true, true, true, true, true, true, false, false, true, true},
						{false, false, true, false, false, false, false, true, true, true},
						{true, false, true, true, true, true, false, true, false, true},
						{false, true, true, false, false, false, true, true, false, true},
					},
					hWalls: maze.Walls{
						{false, true, false, false, false, false, false, true, false, false},
						{true, false, false, false, false, false, true, true, true, true},
						{false, false, true, false, false, false, false, false, false, true},
						{true, false, true, false, false, true, false, false, true, true},
						{false, false, true, true, false, false, false, false, true, true},
						{false, false, false, false, false, false, false, false, true, false},
						{false, false, true, false, true, true, false, true, false, false},
						{true, true, false, false, false, false, false, true, false, false},
						{false, false, false, true, true, true, false, false, true, false},
						{true, true, true, true, true, true, true, true, true, true},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(4)
			m := maze.NewMaze(tt.args.cols, tt.args.rows, 0, 0)
			m.Generate()

			if !isEqual(tt.args.cols, tt.args.rows, tt.args.walls, m) {
				t.Errorf("Generated maze is different")
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		cols   int
		rows   int
		start  coord.Pos
		finish coord.Pos
	}

	tests := []struct {
		name string
		args args
		want coord.Path
	}{
		{
			name: "solve maze 10x10",
			args: args{
				cols: 10,
				rows: 10,
				start: coord.Pos{
					Col: 9,
					Row: 0,
				},
				finish: coord.Pos{
					Col: 4,
					Row: 9,
				},
			},
			want: coord.Path{
				{Col: 9, Row: 1}, {Col: 8, Row: 1}, {Col: 7, Row: 1}, {Col: 6, Row: 1},
				{Col: 5, Row: 1}, {Col: 5, Row: 2}, {Col: 6, Row: 2}, {Col: 6, Row: 3},
				{Col: 6, Row: 4}, {Col: 6, Row: 5}, {Col: 6, Row: 6}, {Col: 6, Row: 7},
				{Col: 6, Row: 8}, {Col: 6, Row: 9}, {Col: 5, Row: 9}, {Col: 4, Row: 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(4)
			m := maze.NewMaze(tt.args.cols, tt.args.rows, 0, 0)
			m.Generate()
			path := m.Solve(tt.args.start, tt.args.finish)

			if !reflect.DeepEqual(tt.want, path) {
				t.Errorf("Wrong path in the maze")
			}
		})
	}
}
