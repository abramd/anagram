package search

import (
	"github.com/abramd/anagram/internal/data"
	"github.com/abramd/anagram/pkg/sort"
)

type Searcher struct {
	data   data.Provider
	sorter sort.Sorter
}

func New(data data.Provider, sorter sort.Sorter) *Searcher {
	return &Searcher{
		sorter: sorter,
		data:   data,
	}
}

func (s *Searcher) Search(input string) []string {
	key := s.sorter.Sort(input)
	if res, ok := s.data.Get(key); ok {
		return res
	}
	return []string{}
}
