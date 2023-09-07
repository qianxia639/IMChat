package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	if n == 0 {
		return ""
	}
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomStringWithString(str string, n int) string {
	if n == 0 {
		return ""
	}

	if len(str) == 0 {
		str = alphabet
	}

	var sb strings.Builder
	k := len(str)
	for i := 0; i < n; i++ {
		c := str[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
