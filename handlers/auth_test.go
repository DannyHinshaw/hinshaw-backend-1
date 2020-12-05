package handlers

import (
	"github.com/labstack/echo/v4"
	"hinshaw-backend-1/db"
	mw "hinshaw-backend-1/middleware"
	"hinshaw-backend-1/schemas"
	td "hinshaw-backend-1/test"
	"strings"
)

func (suite *HandlersTestSuite) TestGenerateJWT() {
	jwtPayload, err := mw.GenerateJWT(td.UserUUID)
	suite.NoError(err)
	suite.NotNil(jwtPayload)
	suite.NotNil(jwtPayload.AccessToken)

	valid, err := mw.ValidateJWT(jwtPayload.AccessToken)
	suite.NoError(err)
	suite.True(valid)
}

func (suite *HandlersTestSuite) TestHandler_POSTRegister() {

	// Invalid Email
	payload := AuthPayload{
		Email:    "",
		Password: "1234567890",
	}

	body := structToJSONString(payload)
	c, _ := newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err := h.POSTRegister(c)
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

	// Valid
	payload = AuthPayload{
		Email:    "test@test.com",
		Password: "1234567890",
	}

	body = structToJSONString(payload)
	c, _ = newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err = h.POSTRegister(c)
	suite.NoError(err)

	// Invalid (email already registered)
	body = structToJSONString(payload)
	c, _ = newTestRequest(echo.POST, "/register", strings.NewReader(body), "")
	err = h.POSTRegister(c)
	suite.Error(err)
}

func (suite *HandlersTestSuite) TestHandler_POSTLogin() {
	payload := AuthPayload{
		Email:    "test@test.com",
		Password: "1234567890",
	}

	// Make sure the user exist in mockDB
	newUser, err := db.CreateNewUser(payload.Email, payload.Password)
	suite.NoError(err)
	db.MockDB.Users = []*schemas.AppUser{newUser}

	body := structToJSONString(payload)
	c, _ := newTestRequest(echo.POST, "/login", strings.NewReader(body), "")
	err = h.POSTLogin(c)
	suite.NoError(err)
}

func (suite *HandlersTestSuite) TestHandler_POSTLogout() {
	c, _ := newTestRequest(echo.POST, "/logout", strings.NewReader(""), "JWT")
	err := h.POSTLogout(c)
	suite.NoError(err)
}
