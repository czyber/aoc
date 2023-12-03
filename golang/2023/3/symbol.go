package main

type Symbol struct {
	Value   string
	Indices []int
}

func (s Symbol) getYOffset() int {
	return s.Indices[0]
}

func (s Symbol) getXOffset() int {
	return s.Indices[1]
}

func (s Symbol) isNextTo(p Part) bool {
	coords := p.Indices
	for _, coord := range coords {
		if coord[0] == s.getYOffset() && (coord[1] == s.getXOffset()+1 || coord[1] == s.getXOffset()-1) {
			return true
		} else if (coord[0] == s.getYOffset()+1 || coord[0] == s.getYOffset()-1) && (coord[1] == s.getXOffset()+1 || coord[1] == s.getXOffset()-1) {
			return true
		} else if (coord[0] == s.getYOffset()+1 || coord[0] == s.getYOffset()-1) && coord[1] == s.getXOffset() {
			return true
		}
	}
	return false
}
