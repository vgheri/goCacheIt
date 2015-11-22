package routes

import (
	"github.com/vgheri/goCacheIt/metrics"
	"log"
	"net/http"
	"time"
)

func middleware(requestHandler http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go metrics.LogHit(routeName)
		start := time.Now()
		requestHandler.ServeHTTP(w, r)
		end := time.Since(start)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			routeName,
			end,
		)
		go metrics.LogDuration(routeName, end)
	})
}
