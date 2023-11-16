package compress

import (
	"compress/gzip"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type compressWriter struct {
	w  http.ResponseWriter
	gz *gzip.Writer
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		w:  w,
		gz: gzip.NewWriter(w),
	}
}

func (cw *compressWriter) Header() http.Header {
	return cw.w.Header()
}

func (cw *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		cw.w.Header().Set("Content-Encoding", "gzip")
	}
	cw.w.WriteHeader(statusCode)
}

func (cw *compressWriter) Write(b []byte) (int, error) {
	if len(b) < 10 {
		return cw.w.Write(b)
	}
	return cw.gz.Write(b)
}

func (cw *compressWriter) Close() error {
	return cw.gz.Close()
}

type compressReader struct {
	r  io.ReadCloser
	rz *gzip.Reader
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	rz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &compressReader{
		r:  r,
		rz: rz,
	}, nil
}

func (cr *compressReader) Read(b []byte) (int, error) {
	return cr.rz.Read(b)
}

func (cr *compressReader) Close() error {
	return cr.rz.Close()
}

func MiddlewareCompress(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var cw *compressWriter
			var cr *compressReader

			ow := w
			zipFormAccept := r.Header.Values("Accept-Encoding")
			for _, elem := range zipFormAccept {
				if elem == "gzip" {
					cw = newCompressWriter(w)
					ow = cw
				}
				defer func() {
					if cw != nil {
						if err := cw.Close(); err != nil {
							w.WriteHeader(http.StatusInternalServerError)
							log.Error("Error", zap.Error(err))
							return
						}
					}
				}()
			}

			zipFormContent := r.Header.Values("Content-Encoding")
			for _, elem := range zipFormContent {
				if elem == "gzip" {
					cr, err := newCompressReader(r.Body)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						log.Error("Error:", zap.Error(err))
						return
					}
					r.Body = cr
				}
			}
			defer func() {
				if cr != nil {
					if err := cr.Close(); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						log.Error("Error:", zap.Error(err))
						return
					}
				}
			}()
			h.ServeHTTP(ow, r)
		})
	}
}
