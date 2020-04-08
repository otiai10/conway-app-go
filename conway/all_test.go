package conway

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestInitCell(t *testing.T) {
	cell := InitCell()
	Expect(t, cell).TypeOf("conway.Cell")
}