package main

type Part struct {
	Value   int
	Indices [][]int
}

func (p Part) getYOffset() int {
	return p.Indices[0][0]
}
