package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
)

// TestHandler_GETHealth return a simple response when healthy.
func (suite *HandlersTestSuite) Test_GETCustomers() {

	// Setup echo framework
	e = echo.New()

	// Register URL endpoints (skip middleware for these tests as they read the request body and ruin it)
	h.RegisterRouteHandlers(e)

	req := httptest.NewRequest(echo.GET, "/customers", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if suite.NoError(h.GETCustomers(c)) {
		suite.Equal(http.StatusOK, rec.Code)
		suite.Contains(rec.Body.String(), "equifax")
	}
}
