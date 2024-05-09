package testmodule

import "math/rand"

func PublicGetRand() int {
	return rand.Int()
}

func PublicRandFloat() float64 {
	return rand.Float64()
}
