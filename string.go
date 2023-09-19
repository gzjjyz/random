package random

import "math/rand"

const (
	keyBase = 48
	keyRand = 122 - 48 + 1
)

func GenerateKey(n int) string {
	bytes := make([]byte, n)
	// key:ascII 48- 122
	for i := 0; i < n; i++ {
		bytes[i] = byte(rand.Intn(keyRand) + keyBase)
	}
	return string(bytes)
}
