package main

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.HandlerFunc

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		next.ServeHTTP(w, r)
		t1 := time.Since(t0).Milliseconds()
		log.Printf("%s %s %dms", r.Method, r.URL.Path, t1)
	}
}

func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}
