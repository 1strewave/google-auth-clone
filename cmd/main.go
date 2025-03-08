package main

import (
	"os"
	"time"

	"github.com/1strewave/google-auth-clone/internal/handlers"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setupLogger() (*zap.SugaredLogger, *os.File) {
	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open log file")
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	consoleOutput := zapcore.AddSync(os.Stdout)
	fileOutput := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleOutput, zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, fileOutput, zapcore.InfoLevel),
	)

	logger := zap.New(core).Sugar()
	return logger, logFile
}

func main() {
	e := echo.New()

	logger, logFile := setupLogger()
	defer logFile.Close()
	defer logger.Sync()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Infof("Request: %s %s from %s, time: %s", c.Request().Method, c.Request().URL.Path, c.RealIP(), time.Now().Format("2006-01-02 15:04:05"))
			return next(c)
		}
	})

	e.GET("/status", handlers.StatusCheck)
	logger.Infof("Started server - Time: %s", time.Now().Format("2006-01-02 15:04:05"))
	e.Logger.Fatal(e.Start(":8080"))
}
