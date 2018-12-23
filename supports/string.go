package supports

import (
	"github.com/spf13/cast"
	"math/rand"
	"strings"
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

// Put value to placeholder
// template = "User :user_id has been deleted."
// replacements = map[string]interface = {"user_id": 12}
// will be = "User 12 has been deleted."
func Replace(template string, replacements map[string]interface{}) string {
	var noLimit = -1
	for key, value := range replacements {
		template = strings.Replace(template, ":" + key, cast.ToString(value), noLimit)
	}

	return template
}