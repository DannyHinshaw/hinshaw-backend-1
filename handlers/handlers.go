package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hinshaw-backend-1/cache"
	"hinshaw-backend-1/db"
	mw "hinshaw-backend-1/middleware"
	"log"
	"net/http"
)

type Handler struct {
	RedisService cache.IService
	DBService    db.IService
	Context      context.Context
	UserId       string
	Token        string
}

// Create a new Handler with DI for services.
func NewHandler(dbService db.IService, redisService cache.IService) *Handler {
	return &Handler{
		RedisService: redisService,
		DBService:    dbService,
	}
}

// Handler middleware to assign handler values centrally.
func (h *Handler) HandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.Context = c.Request().Context()
		// TODO: If useful
		//h.UserId = mw.GetUserId(c)

		return next(c)
	}
}

// Checks the JWT in request is also in redis (valid).
func (h *Handler) JWTRedisMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := mw.ExtractToken(c.Request())
		isValid := h.RedisService.IsKeyInRedis(token)
		if !isValid {
			log.Println("attempted unauthorized token use in redis middleware::" + token)
			return echo.NewHTTPError(http.StatusUnauthorized, "Session expired.")
		}

		h.Token = token
		return next(c)
	}
}

// Util to convert structs to their JSON string (*Reader) counterpart for POST request tests.
func structToJSONString(i interface{}) string {
	e, err := json.Marshal(i)
	if err != nil {
		fmt.Println("error converting struct to JSON::", err)
		panic(err)
	}

	return string(e)
}

// Register REST API endpoints.
func (h *Handler) RegisterRouteHandlers(v1 *echo.Echo) {

	/** 	Unrestricted Endpoints
	===================================*/

	// API Health check
	v1.GET("/health", h.GETHealth)

	// Auth
	v1.POST("/register", h.POSTRegister)
	v1.POST("/login", h.POSTLogin)

	/** 	Restricted Endpoints
	===================================*/

	// Restricted group
	r := v1.Group("")
	r.Use(middleware.JWTWithConfig(mw.JWTConf))
	r.Use(h.JWTRedisMiddleware)

	// Auth
	r.POST("/validate", h.POSTValidateToken)
	r.POST("/logout", h.POSTLogout)

	// Customers
	r.GET("/customers", h.GETCustomers)
}
