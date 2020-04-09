package conway

// Status ...
type Status string

const (
	// Alive ...
	Alive Status = "ALIVE"
	// Dead ...
	Dead Status = "DEAD"
)

// Cell ...
type Cell struct {
	Status   Status
	Position Position
}

// Position ...
type Position struct {
	X int
	Y int
}

// Next ...
func (cell Cell) Next(neighbors []Cell, rule Rule) Cell {
	return Cell{
		Status:   rule.NextStatus(cell, neighbors),
		Position: cell.Position,
	}
}
