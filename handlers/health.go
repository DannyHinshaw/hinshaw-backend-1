package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Health check endpoint for API QA.
func (h *Handler) GETHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
