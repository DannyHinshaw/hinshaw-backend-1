package handlers

import (
	mw "hinshaw-backend-1/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register REST API endpoints.
func (h *Handler) RegisterRoutes(v1 *echo.Echo) {

	/** 	Unrestricted Endpoints
	===================================*/

	// API Health check
	v1.GET("/health", h.GETHealth)

	// Auth
	v1.POST("/register", h.POSTRegister)
	v1.POST("/login", h.POSTLogin)

	// Restricted group
	r := v1.Group("")
	r.Use(middleware.JWTWithConfig(mw.JWTConf))

	/** 	Restricted Endpoints
	===================================*/

	// Auth
	r.POST("/logout", h.POSTLogout)
	r.POST("/refresh_token", h.POSTRefreshToken)
	r.POST("/validate", h.POSTValidateToken)

	// Customers
	r.GET("/customers", h.GETCustomers)
}
