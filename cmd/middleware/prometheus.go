package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/meysam81/tarzan/cmd/metrics"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(wrapped.statusCode)

		metrics.RecordHTTPRequest(r.Method, r.URL.Path, statusCode)
		metrics.RecordHTTPDuration(r.Method, r.URL.Path, duration)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
