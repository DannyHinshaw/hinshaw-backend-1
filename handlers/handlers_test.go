package handlers

import (
	"hinshaw-backend-1/db"
	t "hinshaw-backend-1/test"

	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	h *Handler
	e *echo.Echo
)

type HandlersTestSuite struct {
	suite.Suite
}

// Reset mock service and mock data.
func (suite *HandlersTestSuite) SetupTest() {
	mockDB := &db.MockService{}
	err := mockDB.ParseDB(os.Getenv("DATABASE_URL"))
	suite.NoError(err)
	err = mockDB.Init()
	suite.NoError(err)

	h = NewHandler(mockDB)
	h.UserId = t.UserUUID
}

// Util function to bootstrap a test authed API request with boilerplate.
func newTestRequest(method string, url string, body io.Reader, token string) (echo.Context, *httptest.ResponseRecorder) {

	// Setup echo framework
	e = echo.New()

	// Register URL endpoints (skip middleware for these tests as they read the request body and ruin it)
	h.RegisterRoutes(e)

	req := httptest.NewRequest(method, url, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	if token != "" {
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	}

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return c, rec
}

// Test the main server init.
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

// Run the Handlers test suite.
func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersTestSuite))
}
