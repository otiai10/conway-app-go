package presentation

import (
	"bytes"
	"testing"

	"github.com/otiai10/conway-app-go/conway"

	. "github.com/otiai10/mint"
)

var (
	A = conway.Cell{Status: conway.Alive}
	D = conway.Cell{Status: conway.Dead}
)

func TestRenderer_Update(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	r := Renderer{Out: buf, Alive: "#", Dead: "!"}
	r.Update(conway.World{
		Matrix: [][]conway.Cell{
			{A, A, D},
			{D, D, A},
			{A, D, D},
		},
	})
	Expect(t, buf.String()).ToBe("\033[0;0H##!\n!!#\n#!!\n")
}
