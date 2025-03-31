package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1strewave/google-auth-clone/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	e := echo.New()

	e.GET("/status", handlers.StatusCheck)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	expected := `{"message":"Server works fine!"}`
	assert.JSONEq(t, expected, rec.Body.String())
}
