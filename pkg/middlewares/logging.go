package middlewares

import (
	"log"
	"net/http"
	"time"
)

type wrapperWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrapperWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &wrapperWriter{w, 200}
		next.ServeHTTP(wrapper, r)
		log.Println("request", r.Method, r.URL.Path, wrapper.status, time.Since(start), "ms")
	})
}
