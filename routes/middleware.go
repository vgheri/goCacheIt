package routes

import (
	"github.com/vgheri/goCacheIt/Godeps/_workspace/src/github.com/gorilla/mux"
	//"log"
	"fmt"
	"net/http"
	"time"
)

func middleware(requestHandler http.Handler, routeName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := newResponseWriterWrapper(w)
		start := time.Now()
		requestHandler.ServeHTTP(res, r)
		duration := time.Since(start)
		durationMs := duration.Seconds() * float64(time.Second/time.Millisecond)
		vars := mux.Vars(r)
		key := vars["key"]
		fmt.Printf(
			//"\t%s\t%s\t%s\t%s\t%d\t%f",
			"{time:'%s','method':'%s','path':'%s','key':'%s','route':'%s','statusCode':%d,'duration':%f}\n",
			time.Now(),
			r.Method,
			r.RequestURI,
			key,
			routeName,
			res.Status(),
			durationMs,
		)
	})
}
