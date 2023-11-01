package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("return Adam's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Adam", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q \n", got, want)
		}

	})

}
