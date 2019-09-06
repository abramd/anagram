package main

import (
	"github.com/abramd/anagram/internal/data"
	"github.com/abramd/anagram/internal/search"
	"github.com/abramd/anagram/internal/server"
	"github.com/abramd/anagram/pkg/sort"
	"log"
	"net/http"
)

func main() {
	storage := data.New()
	searcher := search.New(storage, sort.New())
	s := server.NewServer(searcher, storage, sort.New())

	http.HandleFunc("/add", s.Add)
	http.HandleFunc("/search", s.Search)
	log.Fatalln(http.ListenAndServe(":80", http.DefaultServeMux))
}
