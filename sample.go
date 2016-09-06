package sample

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var ch chan []byte

// Sum returns the sum of a and b
// Use GoSampleSum to call it from objective-c
func Sum(a int, b int) (result int) {
	return a + b
}

// Init initializes this example
func Init() {
	ch = make(chan []byte, 5)
	go startRequesting()
}

func startRequesting() {
	for url := range ch {
		makeHTTPRequest(url)
	}
}

// Crash attempts to reproduce the crash described in golang/go#16644
func Crash() {
	for i := 0; i < 10; i++ {
		ch <- []byte("http://www.authy.com")
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func makeHTTPRequest(url []byte) {
	res, err := http.Get(string(url))
	if err != nil {
		log.Fatal(err)
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
