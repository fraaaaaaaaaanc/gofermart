package testhandlers

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gofermart/internal/config"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	"gofermart/internal/storage"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	flags := config.ParseFlags()
	log, _ := logger.NewZapLogger("", "local")
	strg, _ := storage.NewStorage("host=localhost password=1234 dbname=gofermart user=postgres "+
		"sslmode=disable", log.Log)
	hndlr := allhandlers.NewHandlers(log.Log, strg, flags.SecretKeyJWTToken)

	method := http.MethodPost
	url := "http://localhost:8080/api/user/register"
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
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\" with an empty request " +
				"body, it should return the Status Code 400",
			req: req{
				body:        ``,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\" with the request body " +
				"in the appropriate json format, should return the Status Code 400",
			req: req{
				body:        `tests`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\" with a request body in which " +
				"the parameter does not match the models structure type models.RequestRegister, should return " +
				"the Status Code 400",
			req: req{
				body:        `{"login": 123, "password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\" with an incomplete request " +
				"body should return the Status Code 400",
			req: req{
				body:        `{"password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\" with the correct request " +
				"body should return the Status Code 200",
			req: req{
				body:        `{"login": "123", "password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusOK,
				cookie:     &http.Cookie{},
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/register\"with the same request " +
				"body should return the Status Code 409",
			req: req{
				body:        `{"login": "123", "password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusConflict,
				cookie:     &http.Cookie{},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, strings.NewReader(test.body))
			request.Header.Set("Content-Type", test.contentType)
			w := httptest.NewRecorder()
			hndlr.Register(w, request)
			resp := w.Result()
			defer func() {
				if err := resp.Body.Close(); err != nil {
					log.Error("request body closing error", zap.Error(err))
				}
			}()

			assert.Equal(t, test.statusCode, resp.StatusCode)
			if test.cookie != nil {
				assert.NotNil(t, resp.Cookies())
			}

		})
	}
}
