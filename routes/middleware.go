package routes

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/vgheri/goCacheIt/metrics"
	"log"
	"net/http"
	"time"
)

func middleware(requestHandler http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestHandler.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			routeName,
			duration,
		)
		vars := mux.Vars(r)
		key := vars["key"]
		go metrics.LogMetrics(routeName, key, duration)
	})
}
