package controllers

import (
	"expresso/pkg/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type statusController struct {
	*Controller
}

func (s *statusController) healthCheck(c echo.Context) error {
	srv := services.InitStatusService(s.App, nil)
	return c.JSONPretty(http.StatusOK, srv.Status(), "")
}

func (s *statusController) healthCheckDB(c echo.Context) error {
	srv := services.InitStatusService(s.App, nil)
	return c.JSONPretty(http.StatusOK, srv.DBSetGetStatus(), "")
}
