package main

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	_, err := http.Get(a)
	if err != nil {
		return ""
	}
	aDuration := time.Since(startA)

	startB := time.Now()
	_, err = http.Get(b)
	if err != nil {
		return ""
	}
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}
