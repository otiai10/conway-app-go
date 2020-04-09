package conway

// World ...
type World struct {
	Width  int
	Height int
	Matrix [][]Cell
	Rule   Rule
}

// InitWorld ...
func InitWorld(w, h int, rule ...Rule) World {
	matrix := [][]Cell{}
	for y := 0; y < h; y++ {
		row := []Cell{}
		for x := 0; x < w; x++ {
			row = append(row, CreateCell(x, y, Alive))
		}
		matrix = append(matrix, row)
	}
	rule = append(rule, DefaultRule{})
	return World{
		Width:  w,
		Height: h,
		Matrix: matrix,
		Rule:   rule[0],
	}
}

// CreateCell ...
func CreateCell(x, y int, status Status) Cell {
	return Cell{}
}

// Next provides next world.
func (world World) Next() World {
	matrix := [][]Cell{}
	for y := 0; y < world.Height; y++ {
		row := []Cell{}
		for x := 0; x < world.Width; x++ {
			cell := world.Matrix[y][x]
			neighbors := world.GetNeighborsOf(x, y)
			row = append(row, cell.Next(neighbors, world.Rule))
		}
		matrix = append(matrix, row)
	}
	return World{
		Width:  world.Width,
		Height: world.Height,
		Matrix: matrix,
	}
}

// GetNeighborsOf ...
func (world World) GetNeighborsOf(x, y int) []Cell {
	neighbors := []Cell{}

	return neighbors
}
