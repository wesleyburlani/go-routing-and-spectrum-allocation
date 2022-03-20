package utils

import (
	"math/rand"
	"time"
)

func RandomNumberBetween(lower, upper int) int {
	rand.Seed(time.Now().Unix())
	return lower + rand.Intn(upper-lower+1)
}
