package compress

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareCompress(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	requestBody := `{"login":"test","password":"123"}`

	t.Run("send_gzip", func(t *testing.T) {
		// Test case for sending GZIP-encoded data.
		// Compress the request body and make an HTTP POST request with GZIP-encoded data.
		buf := bytes.NewBuffer(nil)
		zb := gzip.NewWriter(buf)
		_, err := zb.Write([]byte(requestBody))
		assert.NoError(t, err)
		err = zb.Close()
		assert.NoError(t, err)
		fmt.Println(buf)

		request := httptest.NewRequest(http.MethodPost, "/", buf)
		assert.NoError(t, err)
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		request.Header.Set("Content-Encoding", "gzip")
		request.RequestURI = ""
		recorder := httptest.NewRecorder()

		handler := MiddlewareCompress()(mockHandler)
		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)

		b, err := io.ReadAll(recorder.Body)
		assert.NoError(t, err)
		assert.NotNil(t, string(b))
	})

	t.Run("accept_gzip", func(t *testing.T) {
		// Test case for accepting GZIP-encoded data.
		// Send an HTTP POST request with an "Accept-Encoding: gzip" header.
		buf := bytes.NewBufferString(requestBody)
		request := httptest.NewRequest(http.MethodPost, "/", buf)
		request.RequestURI = ""
		request.Header.Set("Accept-Encoding", "gzip")
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		recorder := httptest.NewRecorder()

		handler := MiddlewareCompress()(mockHandler)
		handler.ServeHTTP(recorder, request)

		zr, err := gzip.NewReader(recorder.Body)
		assert.NoError(t, err)

		b, err := io.ReadAll(zr)
		assert.NoError(t, err)
		assert.NotNil(t, string(b))
	})
}
