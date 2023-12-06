package testhandlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/storage/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorageGofermart(ctrl)
	cookies := cookie.NewCookie("test")
	hndlr := allhandlers.NewHandlers(mockStorage, cookies)

	gomock.InOrder(mockStorage.EXPECT().CheckUserLoginData(&handlersmodels.RequestLogin{
		Login:    "123",
		Password: "123",
		Ctx:      context.Background(),
	}).Return(nil, handlersmodels.ErrMissingDataInTable),
		mockStorage.EXPECT().CheckUserLoginData(&handlersmodels.RequestLogin{
			Login:    "test",
			Password: "1234",
			Ctx:      context.Background(),
		}).Return(&handlersmodels.ResultLogin{
			UserID:   1,
			Password: "$2a$10$w2Ksuvu8yOKxYB6rt9hh6u9axDbLQ0fRYWzepHeB2p.P1dHdl9Ea2",
		}, nil),
		mockStorage.EXPECT().CheckUserLoginData(&handlersmodels.RequestLogin{
			Login:    "test",
			Password: "123",
			Ctx:      context.Background(),
		}).Return(&handlersmodels.ResultLogin{
			UserID:   1,
			Password: "$2a$10$w2Ksuvu8yOKxYB6rt9hh6u9axDbLQ0fRYWzepHeB2p.P1dHdl9Ea2",
		}, nil))

	methodLogin := http.MethodPost
	urlLogin := "http://localhost:8080/api/user/login"
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
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with an empty request " +
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
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with the request body " +
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
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with incomplete request body, " +
				"missing password, should return the Status Code 400",
			req: req{
				body:        `{"login": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with a request body in which " +
				"the parameter does not match the models structure type handlers_models.RequestLogin, should return " +
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
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with non-existent data in the " +
				"request body, should return the Status Code 401",
			req: req{
				body:        `{"login": "123", "password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusUnauthorized,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with an invalid password in the " +
				"request body, should return the Status Code 401",
			req: req{
				body:        `{"login": "test", "password": "1234"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusUnauthorized,
				cookie:     nil,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/login\" with the request body with the " +
				"correct data, should return the Status Code 200",
			req: req{
				body:        `{"login": "test", "password": "123"}`,
				contentType: "application/json",
			},
			resp: resp{
				statusCode: http.StatusOK,
				cookie:     nil,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(methodLogin, urlLogin, strings.NewReader(test.body))
			request.Header.Set("Content-Type", test.contentType)
			w := httptest.NewRecorder()
			hndlr.Login(w, request)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, test.statusCode, resp.StatusCode)
			if test.cookie != nil {
				assert.NotNil(t, test.cookie)
			}
		})
	}
}
