package conway

// World ...
type World struct {
	Width  int
	Height int
	Cells  [][]Cell
}

// InitWorld ...
func InitWorld(w, h int) *World {
	return &World{}
}

// InitCell ...
func InitCell() Cell {
	return Cell{}
}
