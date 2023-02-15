package controllers

import (
	"expresso/configs/core"
	"expresso/configs/core/initializers"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	core.App
	server *initializers.HttpServer
}

func NewRoutes(app core.App) *Routes {
	return &Routes{app, app.Server().Provider()}
}

func (r *Routes) SetupRoutes() {
	r.server.Logger.Info()
	statusController := &statusController{NewController(r.App)}
	r.server.GET("/health_check", statusController.healthCheck)
	r.server.GET("/health_check_db", statusController.healthCheckDB)
}

func (r *Routes) SetupMiddleware() {
	r.server.Use(echozap.ZapLogger(r.Logger().With().(*initializers.FileLogger).ZapLogger()))
	r.server.Use(middleware.Recover())
	r.server.Use(middleware.RequestID())
	r.server.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		r.Logger().With("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).Info("Request: ", string(reqBody))
		r.Logger().With("request_id", c.Response().Header().Get(echo.HeaderXRequestID)).Info("Response: ", string(resBody))
	}))
}
