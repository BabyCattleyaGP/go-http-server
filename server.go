package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer function, get player name, call getplayerscore
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayerScore(player))
}

// GetPlayerScore function, return player score
func GetPlayerScore(name string) string {
	if name == "Baby" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
