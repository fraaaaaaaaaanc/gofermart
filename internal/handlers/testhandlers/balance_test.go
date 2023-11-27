package testhandlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofermart/internal/handlers/allhandlers"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/storage/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorageMock(ctrl)
	hndlrs := allhandlers.NewHandlers(mockStorage, "test")

	gomock.InOrder(
		mockStorage.EXPECT().GetUserBalance(gomock.Any()).Return(
			&handlers_models.RespUserBalance{
				UserBalance:      100,
				WithdrawnBalance: 0,
			}, nil),
		mockStorage.EXPECT().GetUserBalance(gomock.Any()).Return(
			&handlers_models.RespUserBalance{
				UserBalance:      100.8,
				WithdrawnBalance: 456.2,
			}, nil),
	)

	method := http.MethodGet
	url := "http://localhost:8080/api/user/balance"

	type req struct {
		userID int
	}
	type resp struct {
		statusCode int
		body       string
	}
	tests := []struct {
		name string
		req
		resp
	}{
		{
			name: "GET request was sent to \"http://localhost:8080/api/user/balance\", with a positive response, " +
				"should return th status code 200",
			req: req{
				userID: 1,
			},
			resp: resp{
				statusCode: http.StatusOK,
				body:       `{"current": 100,"withdrawn": 0}`,
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/api/user/balance\", with a positive response, " +
				"should return th status code 200",
			req: req{
				userID: 2,
			},
			resp: resp{
				statusCode: http.StatusOK,
				body:       `{"current": 100.8,"withdrawn": 456.2}`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, nil)
			ctx := context.WithValue(request.Context(), cookiemodels.UserID, test.userID)
			request = request.WithContext(ctx)
			w := httptest.NewRecorder()
			hndlrs.Balance(w, request)

			resp := w.Result()
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, test.statusCode, resp.StatusCode)
			assert.JSONEq(t, test.body, string(body))
		})
	}
}
