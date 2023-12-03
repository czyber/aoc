package main

type EngineSchematic struct {
	parts   []Part
	symbols []Symbol
}

func (e EngineSchematic) getParts() []Part {
	return e.parts
}

func (e EngineSchematic) getSymbols() []Symbol {
	return e.symbols
}

func (e *EngineSchematic) addParts(ps []Part) {
	e.parts = append(e.parts, ps...)
}

func (e *EngineSchematic) addSymbols(ss []Symbol) {
	e.symbols = append(e.symbols, ss...)
}

func (e EngineSchematic) getAdjacentSymbols(p Part) bool {
	hasAdjacent := false
	for _, symbol := range e.symbols {
		if symbol.isNextTo(p) {
			hasAdjacent = true
		}
	}
	return hasAdjacent
}

func (e EngineSchematic) validatePart(p Part) bool {
	return e.getAdjacentSymbols(p)
}

func (e EngineSchematic) buildPartSum() int {
	isValid := false
	sum := 0
	for _, part := range e.parts {
		isValid = e.validatePart(part)
		if isValid {
			sum += part.Value
		} else {
		}

	}
	return sum
}
