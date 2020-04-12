package presentation

import (
	"fmt"
	"io"

	"github.com/otiai10/conway-app-go/conway"
)

// Renderer ...
type Renderer struct {
	Out   io.Writer
	Alive string
	Dead  string
}

// Update ...
func (r *Renderer) Update(world conway.World) error {
	r.moveToTop()
	for _, row := range world.Matrix {
		for _, cell := range row {
			switch cell.Status {
			case conway.Alive:
				fmt.Fprint(r.Out, r.Alive)
			case conway.Dead:
				fmt.Fprint(r.Out, r.Dead)
			}
		}
		r.newline()
	}
	return nil
}

// Clear ...
func (r *Renderer) Clear() {
	fmt.Fprint(r.Out, "\033c")
}

func (r *Renderer) moveToTop() {
	fmt.Fprint(r.Out, "\033[0;0H")
}

func (r *Renderer) newline() {
	fmt.Fprint(r.Out, "\n")
}
