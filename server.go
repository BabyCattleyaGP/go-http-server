package main

import (
	"fmt"
	"net/http"
)

// PlayerServer function, return player score
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
