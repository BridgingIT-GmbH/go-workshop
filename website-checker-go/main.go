package main

import (
	"fmt"
	"net/http"
	"time"
)

var urlsToCheck = []string{
	"https://gobyexample.com",
	"https://github.com",
	"https://go.dev",
	"https://go.dev/tour/welcome/1",
	"https://go.dev/learn",
	"https://www.bridging-it.de",
	"https://www.bridging-it.de/content/karriere",
	"https://bridgingit.github.io/go-workshop",
	"https://www.destroyallsoftware.com/talks/wat",
	"https://pkg.go.dev",
	"https://github.com/golang/go",
	"https://kernel.org",
}

func main() {
	channel := make(chan string, len(urlsToCheck))
	start := time.Now()

	for _, url := range urlsToCheck {
		go check(url, channel)
	}

	for index := 0; index < len(urlsToCheck); index++ {
		fmt.Println(<-channel)
	}

	fmt.Printf("Checking websites took %d ms\n", time.Since(start).Milliseconds())
}

func check(url string, channel chan string) {
	result, err := http.Get(url)
	if err != nil {
		channel <- url + ": " + err.Error()
	} else {
		channel <- url + ": " + result.Status
	}
}
