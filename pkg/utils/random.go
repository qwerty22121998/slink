package utils

import (
	"math/rand"
	"time"
)

const RandomPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = RandomPool[rand.Int63()%int64(len(RandomPool))]
	}
	return string(b)
}
