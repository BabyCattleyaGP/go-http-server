package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"syreclabs.com/go/faker"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Baby's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Baby", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := faker.Number().Number(2)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
