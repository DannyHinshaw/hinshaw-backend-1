package handlers

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func (suite *HandlersTestSuite) TestHandler_POSTRegister() {

	// Valid
	payload := AuthPayload{
		Email:    "test@test.com",
		Password: "1234567890",
	}

	body := structToJSONString(payload)
	c, _ := newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err := h.POSTRegister(c)
	suite.NoError(err)

	// Invalid Email
	payload = AuthPayload{
		Email:    "",
		Password: "1234567890",
	}

	body = structToJSONString(payload)
	c, _ = newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err = h.POSTRegister(c)
	suite.Error(err)

	// Invalid Password
	payload = AuthPayload{
		Email:    "test@test.com",
		Password: "",
	}

	body = structToJSONString(payload)
	c, _ = newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err = h.POSTRegister(c)
	suite.Error(err)
}

func (suite *HandlersTestSuite) TestHandler_POSTLogin() {

	// Valid
	payload := AuthPayload{
		Email:    "test@test.com",
		Password: "1234567890",
	}

	body := structToJSONString(payload)
	c, _ := newTestRequest(echo.POST, "/login", strings.NewReader(body), "")
	err := h.POSTLogin(c)
	suite.NoError(err)
}

func (suite *HandlersTestSuite) TestHandler_POSTLogout() {
	c, _ := newTestRequest(echo.POST, "/logout", strings.NewReader(""), "TODO: TOKEN")
	err := h.POSTLogout(c)
	suite.NoError(err)
}
