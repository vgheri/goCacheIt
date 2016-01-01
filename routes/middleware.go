package routes

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func middleware(requestHandler http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := newResponseWriterWrapper(w)
		start := time.Now()
		requestHandler.ServeHTTP(res, r)
		duration := time.Since(start)
		vars := mux.Vars(r)
		key := vars["key"]
		log.Printf(
			"\t%s\t%s\t%s\t%s\t%d\t%s",
			r.Method,
			r.RequestURI,
			key,
			routeName,
			res.Status(),
			duration,
		)
	})
}
