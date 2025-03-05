package main

import (
	"bufio"
	"fmt"
	"os"
)

func lines(path string) ([]string, bool) {
	f, err := os.Open(path)
	if err != nil {
		return nil, false
	}
	defer f.Close()

	var cont []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		cont = append(cont, s.Text())
	}

	if s.Err() != nil {
		return nil, false
	}

	return cont, true
}

type FileDiff struct {
	from, to []string
	path     []Dir
}

func NewFileDiff(to, from string) (*FileDiff, bool) {
	ct, ok := lines(to)
	if !ok {
		return nil, false
	}
	cf, ok := lines(from)
	if !ok {
		return nil, false
	}
	return &FileDiff{from: cf, to: ct}, true
}

func (d *FileDiff) Dim() (int, int) {
	return len(d.to), len(d.from)
}

func (d *FileDiff) String() string {
	if d.path == nil {
		gs := NewGraphSolver(d)
		d.path = gs.Optimal()
	}

	s := ""
	i, j := 0, 0
	for _, dir := range d.path {
		switch dir {
		case DirS:
			s += fmt.Sprintln("+ ", d.to[i])
			i++
		case DirE:
			s += fmt.Sprintln("- ", d.from[j])
			j++
		case DirSE:
			if d.to[i] == d.from[j] {
				s += fmt.Sprintln("  ", d.to[i])
			} else {
				s += fmt.Sprintln("- ", d.from[j])
				s += fmt.Sprintln("+ ", d.to[i])
			}
			i++
			j++
		}
	}
	return s
}

func (d *FileDiff) S(i, j int) int {
	return 1
}

func (d *FileDiff) E(i, j int) int {
	return 1
}

func (d *FileDiff) SE(i, j int) int {
	if d.to[i] == d.from[j] {
		return 0
	} else {
		return 1
	}
}
