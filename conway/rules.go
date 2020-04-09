package conway

// Rule ...
type Rule interface {
	NextStatus(self Cell, neighbors []Cell) Status
}

// DefaultRule ...
type DefaultRule struct{}

// NextStatus ...
func (rule DefaultRule) NextStatus(self Cell, neighbors []Cell) Status {
	return Dead
}
