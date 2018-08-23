package util

import "math/rand"

const letterBytes = "abcdef0123456789"

func GetRandomColor() string {
	return "#" + randStringBytes(6)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
