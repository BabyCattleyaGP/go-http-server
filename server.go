package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore interface
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer struct
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// StubPlayerStore type to store score
type StubPlayerStore struct {
	scores map[string]int
}

// GetPlayerScore function, return player score
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
