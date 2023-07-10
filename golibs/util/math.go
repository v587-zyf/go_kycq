package util

import "math/rand"

//RouletteSelect
//轮盘选择，返回选中的index,未选中返回-1
func RouletteSelect(weights []int) int {
	sum := 0
	for _, weight := range weights {
		sum += weight
	}
	if sum == 0 {
		return -1
	}
	// get a random value
	value := rand.Intn(sum)
	// locate the random value based on the weights
	for index, weight := range weights {
		if weight <= 0 {
			continue
		}
		value -= weight
		if value < 0 {
			return index
		}
	}
	// only when rounding errors occur
	return -1
}
