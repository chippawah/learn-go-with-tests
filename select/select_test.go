package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the faster responding website url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if neither server responds within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(2)

		defer serverA.Close()

		_, err := ConfirgurableRacer(serverA.URL, serverA.URL, 1)

		if err == nil {
			t.Errorf("expected and error but got none")
		}
	})
}

func makeDelayedServer(delay time.Duration) (delayedServer *httptest.Server) {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
