package schemas

import (
	"github.com/labstack/echo/v4"
	"log"
	"reflect"
)

// Util function to handle Echo data binding with proper error messages.
func HandleBindData(c echo.Context, target interface{}) error {
	if err := c.Bind(target); err != nil {
		targetType := reflect.TypeOf(target).String()
		log.Printf("error binding payload to %s schema:: %s\n", targetType, err)
		return err
	}

	return nil
}
