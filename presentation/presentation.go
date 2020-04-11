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
	r.clear()
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

func (r *Renderer) clear() {
	fmt.Fprint(r.Out, "\033c")
}

func (r *Renderer) newline() {
	fmt.Fprint(r.Out, "\n")
}
