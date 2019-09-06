package data

import (
	"math/rand"
	"testing"
)

func BenchmarkStorage_Get(b *testing.B) {

	s := New()
	for i := 0; i < 10000000; i++ {
		str := randomString(6)
		s.Add(str, str)
	}

	input := "abcd"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.Get(input)
	}
}

var chars = "abcdefghijklmnop"

func randomString(n int) string {
	l := len(chars) - 1
	res := ""
	m := rand.Intn(n)

	for i := 0; i <= m; i++ {
		res += string(chars[rand.Intn(l)])
	}
	return res
}
