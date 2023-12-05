package cookie

import (
	"github.com/stretchr/testify/assert"
	"gofermart/internal/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMiddlewareCheckCookie(t *testing.T) {
	// Mock handler for testing

	secretKey := os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN")
	cookies := NewCookie(secretKey)
	userCookie, _ := cookies.NewUserCookie(1)

	_ = logger.NewZapLogger("", "local")
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Test case 1: Valid Authorization cookie_models
	t.Run("ValidAuthorizationCookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: userCookie.Value})
		recorder := httptest.NewRecorder()

		handler := cookies.MiddlewareCheckCookie()(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	// Test case 2: Missing Authorization cookie_models
	t.Run("MissingAuthorizationCookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		recorder := httptest.NewRecorder()

		handler := cookies.MiddlewareCheckCookie()(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "there is no authorization cookie_models")
	})

	// Test case 3: Error working with Authorization token
	t.Run("ErrorWorkingWithToken", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: "invalid_token"})
		recorder := httptest.NewRecorder()

		handler := cookies.MiddlewareCheckCookie()(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "error working with the authorization token")
	})
}
