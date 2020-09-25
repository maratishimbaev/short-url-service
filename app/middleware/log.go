package middleware

import (
	"context"
	"github.com/kataras/golog"
	"math/rand"
	"net/http"
)

const requestNumberSize = 4

var numbers = []rune("0123456789")

type StatusResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewStatusResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{ResponseWriter: w}
}

func (w *StatusResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func generateRequestNumber(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(str)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestNumber := generateRequestNumber(requestNumberSize)
		r = r.WithContext(context.WithValue(r.Context(), "id", requestNumber))
		
		golog.Infof("#%s: %s %s", requestNumber, r.Method, r.URL)

		sw := NewStatusResponseWriter(w)
		next.ServeHTTP(sw, r)
		
		golog.Infof("#%s: code %d", requestNumber, sw.StatusCode)
	})
}
