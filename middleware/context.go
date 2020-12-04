package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

type ContextKeys struct {
	UserId string `json:"userId"`
}

var Keys = ContextKeys{
	UserId: "userId",
}

// Util function for retrieving strings from echo context.
func HandleGetterString(c echo.Context, key string) string {
	val := c.Get(key)
	if val == nil {
		log.Printf("key %s unavailable in context::", key)
		return ""
	}

	return val.(string)
}

// Util function to retrieve requests user id from context.
func GetUserId(c echo.Context) string {
	return HandleGetterString(c, Keys.UserId)
}
