package conway

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestWorld(t *testing.T) {
	world := InitWorld(20, 20)
	Expect(t, world).TypeOf("conway.World")
}

func TestWorld_Next(t *testing.T) {
	world := InitWorld(20, 20)
	next := world.Next()
	Expect(t, next).TypeOf("conway.World")
}
