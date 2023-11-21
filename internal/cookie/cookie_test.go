package cookie

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMiddlewareCheckCookie(t *testing.T) {
	// Mock handler for testing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	secretKey := os.Getenv("SECRET_KEY_FOR_COOKIE_TOKEN")
	// Secret key for JWT token
	cookie, _ := NewCookie(1, secretKey)

	// Test case 1: Valid Authorization cookie_models
	t.Run("ValidAuthorizationCookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: cookie.Value})
		recorder := httptest.NewRecorder()

		handler := MiddlewareCheckCookie(secretKey)(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	// Test case 2: Missing Authorization cookie_models
	t.Run("MissingAuthorizationCookie", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		recorder := httptest.NewRecorder()

		handler := MiddlewareCheckCookie(secretKey)(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "there is no authorization cookie_models")
	})

	// Test case 3: Error working with Authorization token
	t.Run("ErrorWorkingWithToken", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: "invalid_token"})
		recorder := httptest.NewRecorder()

		handler := MiddlewareCheckCookie(secretKey)(mockHandler)
		handler.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "error working with the authorization token")
	})
}
