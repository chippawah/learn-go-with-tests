package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"https://google.com",
		"http://blog.gypsydave.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"https://google.com":        true,
		"http://blog.gypsydave.com": true,
		"waat://furhurterwe.geds":   false,
	}
	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v but got %v", want, got)
	}
}

func slowWebsiteCheckerStub(url string) bool {
	time.Sleep(50 * time.Millisecond)
	return true
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteCheckerStub, urls)
	}
}
