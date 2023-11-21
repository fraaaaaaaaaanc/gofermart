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

func TestWithdrawals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mock.NewMockStorage(ctrl)
	hndlrs := allhandlers.NewHandlers(mockStorage, "test")

	gomock.InOrder(
		mockStorage.EXPECT().GetAllHistoryBalance(1).Return(nil, handlers_models.ErrTheAreNoWithdraw),
		mockStorage.EXPECT().GetAllHistoryBalance(2).Return(
			[]handlers_models.RespWithdrawalsHistory{
				{
					OrderNumber: "2377225624",
					SumWithdraw: 500,
					ProcessedAt: "2020-12-09T16:09:57+03:00",
				},
			}, nil),
	)

	method := http.MethodPost
	url := "http://localhost:8080/api/user/withdrawals"

	type req struct {
		userID int
	}
	type resp struct {
		contentType string
		statusCode  int
		body        string
	}
	tests := []struct {
		name string
		req
		resp
	}{
		{
			name: "",
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
			name: "",
			req: req{
				userID: 2,
			},
			resp: resp{
				statusCode:  http.StatusOK,
				body:        `[{"order":"2377225624","sum":500,"processed_at":"2020-12-09T16:09:57+03:00"}]`,
				contentType: "application/json",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(method, url, nil)
			ctx := context.WithValue(request.Context(), cookiemodels.UserID, test.userID)
			request = request.WithContext(ctx)
			w := httptest.NewRecorder()
			hndlrs.Withdrawals(w, request)

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
