package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		rd *responseData
	}
)

func (lgw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lgw.ResponseWriter.Write(b)
	lgw.rd.size += size
	return size, err
}

func (lgw *loggingResponseWriter) WriteHeader(statusCode int) {
	lgw.ResponseWriter.WriteHeader(statusCode)
	lgw.rd.status = statusCode
}

func MiddlewareHandlerLog(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			rd := &responseData{
				status: 0,
				size:   0,
			}
			lgw := loggingResponseWriter{
				ResponseWriter: w,
				rd:             rd,
			}
			h.ServeHTTP(&lgw, r)

			duration := time.Since(start)
			fileds := []zap.Field{
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.Int("status", rd.status),
				zap.Duration("duration", duration),
				zap.Int("size", rd.size),
			}
			log.Info("Received request", fileds...)
		})
	}
}
