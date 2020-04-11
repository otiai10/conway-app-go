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
	return Cell{
		Position: Position{X: x, Y: y},
		Status:   GetRandomStatus(),
	}
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
		Rule:   world.Rule,
	}
}

// GetNeighborsOf ...
func (world World) GetNeighborsOf(x, y int) []Cell {
	neighbors := []Cell{}
	for _, ny := range world.getRowIndicesFor(y) {
		for _, nx := range world.getColIndicesFor(x) {
			// fmt.Println(nx, ny)
			if ny != y || nx != x {
				neighbors = append(neighbors, world.Matrix[ny][nx])
				// fmt.Println(world.Matrix[ny][nx])
			}
		}
	}
	// fmt.Println("-----")
	return neighbors
}

func (world World) getRowIndicesFor(y int) []int {
	if y == 0 {
		return []int{y, y + 1}
	}
	if y == world.Height-1 {
		return []int{y - 1, y}
	}
	return []int{y - 1, y, y + 1}
}

func (world World) getColIndicesFor(x int) []int {
	if x == 0 {
		return []int{x, x + 1}
	}
	if x == world.Width-1 {
		return []int{x - 1, x}
	}
	return []int{x - 1, x, x + 1}
}
