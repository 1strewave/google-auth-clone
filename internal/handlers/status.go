package handlers

import (
	"net/http"
	//	"go.uber.org/zap"
	//	"github.com/1strewave/google-auth-clone/cmd/main"
	"github.com/labstack/echo/v4"
)

func StatusCheck(c echo.Context) error {
	//	Sugar.Info("Handling root route")
	return c.JSON(http.StatusOK, map[string]string{"message": "Server works fine!"})
}
