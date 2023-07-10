package common

import (
	"math/rand"
)

func RandValuesInt64(min,max int64) int64{
	if min == max {
		return max
	}

	if (max - min) <= 0 {
		return 0
	}

	randValue := min + rand.Int63n(max - min)

	return randValue
}

