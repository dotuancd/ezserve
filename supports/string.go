package supports

import (
	"math/rand"
	"time"
)

func StringRand(n int) string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	max := len(chars)
	r := ""
	index := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		index = rand.Intn(max)
		r += chars[index:index+1]
	}
	return r
}