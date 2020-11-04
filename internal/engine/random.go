package engine

import "math/rand"

// RandomRange returns a pseudo-random number in [min,max]
func RandomRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
