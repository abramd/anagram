package sort

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort([]byte(tt.input)); string(got) != tt.want {
				t.Errorf("MergeSort() = %v, want %v", string(got), tt.want)
			}
		})
	}
}
