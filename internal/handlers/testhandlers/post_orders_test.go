package testhandlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofermart/internal/handlers/allhandlers"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/models/orderstatuses"
	"gofermart/internal/storage/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	hndlrs := allhandlers.NewHandlers(mockStorage, "test")

	gomock.InOrder(
		mockStorage.EXPECT().AddNewOrder(&handlers_models.ReqOrder{
			OrderStatus: orderstatuses.NEW,
			OrderNumber: "54545454",
			UserID:      1,
			Ctx:         context.WithValue(context.Background(), cookiemodels.UserID, 1),
		}).Return(nil),
		mockStorage.EXPECT().AddNewOrder(&handlers_models.ReqOrder{
			OrderStatus: orderstatuses.NEW,
			OrderNumber: "54545454",
			UserID:      1,
			Ctx:         context.WithValue(context.Background(), cookiemodels.UserID, 1),
		}).Return(handlers_models.ErrConflictOrderNumberSameUser),
		mockStorage.EXPECT().AddNewOrder(&handlers_models.ReqOrder{
			OrderStatus: orderstatuses.NEW,
			OrderNumber: "54545454",
			UserID:      2,
			Ctx:         context.WithValue(context.Background(), cookiemodels.UserID, 2),
		}).Return(handlers_models.ErrConflictOrderNumberAnotherUser),
	)

	method := http.MethodPost
	url := "http://localhost:8080/api/user/orders"
	contentType := "text/plain"

	type req struct {
		body   string
		userID int
	}
	type resp struct {
		statusCode int
	}
	tests := []struct {
		name string
		req
		resp
	}{
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/orders\", with an order number that " +
				"does not pass the luhn algorithm in the request body, should return the status code 422",
			req: req{
				body:   "123",
				userID: 1,
			},
			resp: resp{
				statusCode: http.StatusUnprocessableEntity,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/orders\", with the correct order " +
				"number in the request body, should return the status code 202",
			req: req{
				body:   "54545454",
				userID: 1,
			},
			resp: resp{
				statusCode: http.StatusAccepted,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/orders\", with the order number in the " +
				"request body already added by this user, should return the status code 200",
			req: req{
				body:   "54545454",
				userID: 1,
			},
			resp: resp{
				statusCode: http.StatusOK,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/orders\", with the order number in the " +
				"request body already added by another user, should return the status code 409",
			req: req{
				body:   "54545454",
				userID: 2,
			},
			resp: resp{
				statusCode: http.StatusConflict,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, strings.NewReader(test.body))
			request.Header.Set("Content-Type", contentType)
			ctx := context.WithValue(request.Context(), cookiemodels.UserID, test.userID)
			request = request.WithContext(ctx)
			w := httptest.NewRecorder()
			hndlrs.PostOrders(w, request)

			resp := w.Result()
			defer resp.Body.Close()

			assert.Equal(t, test.statusCode, resp.StatusCode)
		})
	}
}
