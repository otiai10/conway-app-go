package conway

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// GetRandomStatus ...
func GetRandomStatus() Status {
	switch rand.Intn(2) {
	case 0:
		return Dead
	default:
		return Alive
	}
}
