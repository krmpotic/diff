package main

type StringDiff struct {
	from, to string
	path     []Dir
}

func NewStringDiff(to, from string) *StringDiff {
	return &StringDiff{from: from, to: to}
}

func (d *StringDiff) Dim() (int, int) {
	return len(d.to), len(d.from)
}

func (d *StringDiff) String() string {
	if d.path == nil {
		gs := NewGraphSolver(d)
		d.path = gs.Optimal()
	}

	s := ""
	n := 0
	for _, dir := range d.path {
		switch dir {
		case DirS:
			s += "-"
		case DirE:
			s += string(d.from[n])
			n++
		case DirSE:
			s += string(d.from[n])
			n++
		}
	}

	s += "\n"
	n = 0
	for _, dir := range d.path {
		switch dir {
		case DirS:
			s += string(d.to[n])
			n++
		case DirE:
			s += "-"
		case DirSE:
			s += string(d.to[n])
			n++
		}
	}
	s += "\n"
	return s
}

func (d *StringDiff) S(i, j int) int {
	return 1
}

func (d *StringDiff) E(i, j int) int {
	return 1
}

func (d *StringDiff) SE(i, j int) int {
	if d.to[i] == d.from[j] {
		return 0
	} else {
		return 1
	}
}
