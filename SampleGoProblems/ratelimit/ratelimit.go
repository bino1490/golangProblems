package main

import (
	"net/http"

	"golang.org/x/time/rate"
)

var ratelimit = rate.NewLimiter(2, 5)

func handlerFunc(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ratelimit.Allow() {
			w.Write([]byte("not allowed"))
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe(":8080", handlerFunc(handler))

}
