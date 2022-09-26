package helpers

import "math/rand"

func GenerateRandomNum(min, max int) int {
	return rand.Intn(max-min) + min;
}