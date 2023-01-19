package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// while passing channels to func we can restrict it to either write-ONLY or read-ONLY end
// here by saying chan <- result means the get func can only write-ONLY into the channel.
func get(url string, ch chan<- result) { //write end of chan
	now := time.Now()
	var t time.Duration
	if respone, err := http.Get(url); err != nil {
		t = time.Since(now).Round(time.Millisecond)
		ch <- result{url, err, t}
	} else {
		t = time.Since(now).Round(time.Millisecond)
		ch <- result{url, err, t}
		respone.Body.Close()
	}
}

func main() {
	results := make(chan result) // creating a chan which can store result struct ONLY.
	urls := []string{
		"www.google.com",
		"www.google.uk",
		"www.amazon.com",
		"www.nytimes.com",
		"www.localhost.com",
	}
	for _, url := range urls {
		go get(url, results)
	}

	// ranging on urls because 4 goroutines would have written into the channels so i have to only read those number of responses.
	//NOTE: channels can only be closed onces.
	// Also if you try to read from a empty channel it will get blocked becuasee the reader would wait
	stopTimer := time.After(1 * time.Second) //ONE TIME TIMER
	for range urls {
		select {
		case res := <-results:
			log.Println(res.url, res.latency)
		case <-stopTimer:
			log.Fatal("Timeout")
		}
	}
}
