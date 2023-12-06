package testhandlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofermart/internal/cookie"
	"gofermart/internal/handlers/allhandlers"
	"gofermart/internal/logger"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/storage/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrders(t *testing.T) {
	_ = logger.NewZapLogger("", "local")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorageGofermart(ctrl)
	cookies := cookie.NewCookie("test")
	hndlr := allhandlers.NewHandlers(mockStorage, cookies)

	gomock.InOrder(
		mockStorage.EXPECT().GetAllUserOrders(gomock.Any(), 1).Return(nil, handlersmodels.ErrTheAreNoOrders),
		mockStorage.EXPECT().GetAllUserOrders(gomock.Any(), 2).Return(
			[]handlersmodels.RespGetOrders{
				{
					OrderNumber: "545454545454",
					Status:      "PROCESSED",
					Accrual:     900,
					UploadedAt:  "2020-12-10T15:15:45+03:00",
				},
				{
					OrderNumber: "6709504120728607",
					Status:      "NEW",
					UploadedAt:  "2020-12-10T15:12:01+03:00",
				},
			}, nil),
	)

	method := http.MethodGet
	url := "http://localhost:8080/api/user/orders"

	type req struct {
		userID int
	}
	type resp struct {
		statusCode  int
		body        string
		contentType string
	}
	tests := []struct {
		name string
		req
		resp
	}{
		{
			name: "GET request was sent to \"http://localhost:8080/api/user/orders\", with userid that has no " +
				"registered orders, should return status code 204",
			req: req{
				userID: 1,
			},
			resp: resp{
				statusCode:  http.StatusNoContent,
				body:        "",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name: "GET request was sent to \"http://localhost:8080/api/user/orders\", with userid that has " +
				"registered orders, should return status code 200",
			req: req{
				userID: 2,
			},
			resp: resp{
				statusCode: http.StatusOK,
				body: `[{"number":"545454545454","status":"PROCESSED","accrual":900,"uploaded_at":"2020-12-10T15:15:45+03:00"},
{"number":"6709504120728607","status":"NEW","uploaded_at":"2020-12-10T15:12:01+03:00"}]`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, nil)
			ctx := context.WithValue(request.Context(), cookiemodels.UserID, test.req.userID)
			request = request.WithContext(ctx)
			w := httptest.NewRecorder()
			hndlr.GetOrders(w, request)

			resp := w.Result()
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, test.statusCode, resp.StatusCode)
			assert.Equal(t, test.contentType, resp.Header.Get("Content-Type"))
			if test.body != "" {
				assert.JSONEq(t, test.body, string(body))
			}
		})
	}
}
