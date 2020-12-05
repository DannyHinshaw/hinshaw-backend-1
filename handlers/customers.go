package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

// Endpoint to handle retrieving all customers along with their scores in db.
func (h *Handler) GETCustomers(c echo.Context) error {
	ctx := c.Request().Context()
	customers, err := h.DBService.QueryAllCustomers(ctx)
	if err != nil {
		log.Println("error querying customers::", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, customers)
}
