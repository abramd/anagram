package sort

import (
	"math/rand"
	"testing"
)

func TestMergeSort(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"1", "dadada", "aaaddd"},
		{"2", "qkeur", "ekqru"},
		{"3", "qa", "aq"},
		{"4", "a", "a"},
		{"5", "", ""},
		{"6", "zyxwvutsrqcba", "abcqrstuvwxyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort([]byte(tt.input)); string(got) != tt.want {
				t.Errorf("MergeSort() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

// 10 		- 747 ns
// 100 		- 6466 ns
// 1000	 	- 64765 ns
// 10000 	- 694299 ns
// 100000 	- 7515312 ns
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		str := randomString(100000)
		b.StartTimer()
		Merge(str)
	}
}

var chars = "abcdefghijklmnop"

func randomString(n int) string {
	l := len(chars) - 1
	res := ""

	for i := 0; i <= n; i++ {
		res += string(chars[rand.Intn(l)])
	}
	return res
}
