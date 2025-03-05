package utils

import (
	"math/rand"
)

func GetFakeUUID() uint64 {
	return rand.Uint64()
}
