package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"hinshaw-backend-1/cache"
	"hinshaw-backend-1/db"
)

type Handler struct {
	RedisService cache.IService
	DBService    db.IService
	Context      context.Context
	UserId       string
}

// Create a new Handler with option for DI.
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

// Util to convert structs to their JSON string (*Reader) counterpart for POST request tests.
func structToJSONString(i interface{}) string {
	e, err := json.Marshal(i)
	if err != nil {
		fmt.Println("error converting struct to JSON::", err)
		panic(err)
	}

	return string(e)
}
