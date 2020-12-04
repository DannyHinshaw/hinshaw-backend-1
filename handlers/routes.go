package handlers

import (
	"github.com/labstack/echo/v4"
)

// Register REST API endpoints.
func (h *Handler) RegisterRoutes(v1 *echo.Echo) {

	// API Health check
	v1.GET("/health", h.GETHealth)

	// Auth
	v1.POST("/register", h.POSTRegister)
	v1.POST("/login", h.POSTLogin)
	v1.POST("/logout", h.POSTLogout)

	// TODO: Users (for data demo)
	v1.GET("/users/:userId", h.GETHealth)
}
