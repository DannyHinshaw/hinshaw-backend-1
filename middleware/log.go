package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

// Basic util logger to capture data points for every request.
func RouteLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		log.Println(c.RealIP(), req.Method, req.RequestURI)
		return next(c)
	}
}
