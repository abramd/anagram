package server

import (
	"encoding/json"
	"github.com/abramd/anagram/internal/data"
	"github.com/abramd/anagram/internal/search"
	"github.com/abramd/anagram/pkg/sort"
	"net/http"
)

const (
	wordParam = "word"
)

type AddRequest struct {
	Data []string `json:"data"`
}

type SearchResponse struct {
	Data []string `json:"data"`
}

type Server struct {
	searcher *search.Searcher
	data     data.Provider
	sorter   sort.Sorter
}

func NewServer(searcher *search.Searcher, data data.Provider, sorter sort.Sorter) *Server {
	return &Server{
		searcher: searcher,
		data:     data,
		sorter:   sorter,
	}
}

func (s *Server) Search(w http.ResponseWriter, r *http.Request) {
	word := r.Form.Get(wordParam)
	res := s.searcher.Search(word)
	b, err := json.Marshal(&SearchResponse{Data: res})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(b)
}

func (s *Server) Add(w http.ResponseWriter, r *http.Request) {
	req := new(AddRequest)
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "invalid data", http.StatusBadRequest)
		return
	}

	for _, v := range req.Data {
		k := s.sorter.Sort(v)
		s.data.Add(k, v)
	}
}

func Method(mthd string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if mthd != r.Method {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next(w, r)
	}
}
