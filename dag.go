package main

import (
	"slices"
)

type DAGraph interface {
	S(i, j int) int
	E(i, j int) int
	SE(i, j int) int
	Dim() (int, int) // inclusive!
}

func NewGraphSolver(g DAGraph) *GraphSolver {
	return &GraphSolver{g, nil}
}

type GraphSolver struct {
	g       DAGraph
	optimal [][]int
}

func (g *GraphSolver) solve() {
	g.optimal = nil
	rows, cols := g.g.Dim()
	for row := range rows + 1 {
		g.optimal = append(g.optimal, []int{})
		for col := range cols + 1 {
			g.optimal[row] = append(g.optimal[row], g.best(row, col))
		}
	}
}

type Dir int

const (
	DirS = iota
	DirE
	DirSE
)

func (g *GraphSolver) Optimal() (path []Dir) {
	g.solve()

	i, j := g.g.Dim() // optimal[i][j] is valid!

	for i > 0 || j > 0 {
		switch {
		case i == 0:
			path = append(path, DirE)
			j--
		case j == 0:
			path = append(path, DirS)
			i--
		default:
			s := g.optimal[i-1][j] + g.g.S(i-1, j)
			e := g.optimal[i][j-1] + g.g.E(i, j-1)
			se := g.optimal[i-1][j-1] + g.g.SE(i-1, j-1)
			switch {
			case se < s && se < e:
				path = append(path, DirSE)
				i--
				j--
			case s < e:
				path = append(path, DirS)
				i--
			default:
				path = append(path, DirE)
				j--
			}
		}
	}
	slices.Reverse(path)
	return path
}

func (g *GraphSolver) best(i, j int) int {
	o := g.optimal
	switch {
	case i == 0 && j == 0:
		return 0
	case i == 0:
		return o[i][j-1] + g.g.E(i, j-1)
	case j == 0:
		return o[i-1][j] + g.g.S(i-1, j)
	default:
		a := o[i-1][j] + g.g.S(i-1, j)
		b := o[i][j-1] + g.g.E(i, j-1)
		c := o[i-1][j-1] + g.g.SE(i-1, j-1)
		return min(a, b, c)
	}
}
