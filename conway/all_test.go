package conway

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestWorld(t *testing.T) {
	world := InitWorld(20, 20)
	Expect(t, world).TypeOf("conway.World")
}

var (
	A = Cell{Status: Alive}
	D = Cell{Status: Dead}
)

type testcase struct {
	Previous [][]Cell
	Next     [][]Cell
}

var testcases = []testcase{
	{
		Previous: [][]Cell{
			{D, A, A},
			{A, A, A},
			{A, A, A},
		},
		Next: [][]Cell{
			{A, D, A},
			{D, D, D},
			{A, D, A},
		},
	},
}

func TestWorld_GetNeighborsOf(t *testing.T) {
	world := InitWorld(3, 3)
	world.Matrix = [][]Cell{
		{D, A, A},
		{A, A, A},
		{A, A, A},
	}
	neighbors := world.GetNeighborsOf(0, 0)
	Expect(t, len(neighbors)).ToBe(3)
	next := world.Next()
	Expect(t, next.Matrix[0][0].Status).ToBe(Alive)
	neighbors = world.GetNeighborsOf(2, 2)
	Expect(t, len(neighbors)).ToBe(3)
}

func TestWorld_Next(t *testing.T) {
	for _, tc := range testcases {
		world := InitWorld(3, 3)
		world.Matrix = tc.Previous
		next := world.Next()
		Expect(t, next.Matrix).Deeply().ToBe(tc.Next)
	}
}
