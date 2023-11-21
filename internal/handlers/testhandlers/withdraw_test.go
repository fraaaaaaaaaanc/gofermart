package testhandlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofermart/internal/handlers/allhandlers"
	cookiemodels "gofermart/internal/models/cookie_models"
	"gofermart/internal/models/handlers_models"
	"gofermart/internal/storage/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWithDraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	hndlrs := allhandlers.NewHandlers(mockStorage, "test")

	gomock.InOrder(
		mockStorage.EXPECT().CheckOrderNumber(gomock.Any(), gomock.Any()).Return(handlers_models.ErrDuplicateOrderNumber),
		mockStorage.EXPECT().CheckOrderNumber(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().WithdrawBalance(gomock.Any()).Return(handlers_models.ErrNegativeBalanceValue),
		mockStorage.EXPECT().CheckOrderNumber(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().WithdrawBalance(gomock.Any()).Return(handlers_models.ErrDuplicateOrderNumberHistoryBalance),
		mockStorage.EXPECT().CheckOrderNumber(gomock.Any(), gomock.Any()).Return(nil),
		mockStorage.EXPECT().WithdrawBalance(gomock.Any()).Return(nil),
	)

	method := http.MethodPost
	url := "http://localhost:8080/api/user/balance/withdraw"

	type req struct {
		body        string
		contentType string
		userID      int
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
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with an invalid " +
				"request body, should return the status code 400",
			req: req{
				body:        `{"login":"test"}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with an incomplete " +
				"request body, should return the status code 400",
			req: req{
				body:        `{"sum":100}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusBadRequest,
			},
		}, {
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with the correct " +
				"request body,but the luhn algorithm does not pass on the order, should return the status code 422",
			req: req{
				body:        `{"order":"123","sum":100}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusUnprocessableEntity,
			},
		}, {
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with the correct " +
				"request body, but there is already such a warrant, should return the status code 422",
			req: req{
				body:        `{"order":"6011964086036747","sum":1000}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusUnprocessableEntity,
			},
		}, {
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with the correct " +
				"request body,but there are not enough funds to write off, should return the status code 402",
			req: req{
				body:        `{"order":"5460262971178544","sum":100}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusPaymentRequired,
			},
		}, {
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with the correct " +
				"request body,but the order number is already in the balance history, should return the " +
				"status code 422",
			req: req{
				body:        `{"order":"5460262971178544","sum":100}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusUnprocessableEntity,
			},
		},
		{
			name: "POST request was sent to \"http://localhost:8080/api/user/balance/withdraw\", with the correct " +
				"request body, should return the status code 200",
			req: req{
				body:        `{"order":"5460262971178544","sum":100}`,
				contentType: "application/json",
				userID:      1,
			},
			resp: resp{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		request := httptest.NewRequest(method, url, strings.NewReader(test.body))
		request.Header.Set("Content-Type", test.contentType)
		ctx := context.WithValue(request.Context(), cookiemodels.UserID, test.userID)
		request = request.WithContext(ctx)
		w := httptest.NewRecorder()
		hndlrs.WithDraw(w, request)

		resp := w.Result()
		defer resp.Body.Close()

		assert.Equal(t, test.statusCode, resp.StatusCode)
	}
}
