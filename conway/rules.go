package conway

// Rule ...
type Rule interface {
	NextStatus(self Cell, neighbors []Cell) Status
}

// DefaultRule ...
type DefaultRule struct{}

// NextStatus ...
func (rule DefaultRule) NextStatus(self Cell, neighbors []Cell) Status {
	aliveNeighbors := rule.countAliveNeighbors(neighbors)
	switch self.Status {
	case Alive:
		switch aliveNeighbors {
		case 0, 1:
			return Dead
		case 2, 3:
			return Alive
		default:
			return Dead
		}
	case Dead:
		switch aliveNeighbors {
		case 3:
			return Alive
		default:
			return Dead
		}
	}
	return Dead
}

func (rule DefaultRule) countAliveNeighbors(neighbors []Cell) int {
	count := 0
	for _, cell := range neighbors {
		if cell.Status == Alive {
			count++
		}
	}
	return count
}
