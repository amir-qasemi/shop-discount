package util

import "math/rand"

func init() {
	rand.Seed(42)
}

// Ptr Returns a pointer to the passed variable.
// Can be used to get the address of return value of function.
func Ptr[T any](a T) *T {
	return &a
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
