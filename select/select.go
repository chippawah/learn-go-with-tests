package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(urlA, urlB string) (fasterResponse string, err error) {
	return ConfirgurableRacer(urlA, urlB, 10*time.Second)
}

func ConfirgurableRacer(urlA, urlB string, timeout time.Duration) (fasterRespondent string, err error) {
	select {
	case <-ping(urlA):
		return urlA, nil
	case <-ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout while waiting for response from %s and %s", urlA, urlB)
	}
}

func ping(url string) (ch chan struct{}) {
	ch = make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return
}
