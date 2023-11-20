package testhandlers

import (
	"github.com/stretchr/testify/assert"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	flags := config.ParseFlags()
	log, _ := logger.NewZapLogger("", "local")
	strg, _ := storage.NewStorage("host=localhost password=1234 dbname=gofermart user=postgres "+
		"sslmode=disable", log.Log)
	hndlr := allhandlers.NewHandlers(log.Log, strg, flags.SecretKeyJWTToken)

	method := http.MethodPost
	url := "http://localhost:8080/api/user/login"
	type req struct {
		body        string
		contentType string
	}
	type resp struct {
		statusCode int
		cookie     *http.Cookie
	}
	tests := []struct {
		name string
		req
		resp
	}{
		{},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, strings.NewReader(test.body))
			request.Header.Set("Content-Type", test.contentType)
			w := httptest.NewRecorder()
			hndlr.Login(w, request)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, resp.StatusCode, test.statusCode)
			if test.cookie != nil {
				assert.NotNil(t, test.cookie)
			}
		})
	}
}
