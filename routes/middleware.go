package routes

import (
	"log"
	"net/http"
	"time"
)

func middleware(requestHandler http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestHandler.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			routeName,
			time.Since(start),
		)
	})
}
