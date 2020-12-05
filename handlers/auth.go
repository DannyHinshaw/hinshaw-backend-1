package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"hinshaw-backend-1/schemas"
	"log"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User register with an email and password..
func (h *Handler) POSTRegister(c echo.Context) error {

	// Extract request payload data.
	payload := new(AuthPayload)
	if err := schemas.HandleBindData(c, payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Registration payload was malformed.")
	}

	email := payload.Email
	if email == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Registration missing email.")
	}
	// TODO: Check email doesn't exist yet

	password := payload.Password
	if len(password) < 5 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Password must have at least 5 characters.")
	}
	// TODO: Any other password restriction checks here.

	// Hash/salt the password to prep for db.
	bytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("error occurred while hashing password::", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	fmt.Println(string(hashedPassword))

	return c.JSON(http.StatusOK, "Registered user successfully")
}

// Handles user login with email/password..
func (h *Handler) POSTLogin(c echo.Context) error {

	// Extract request payload data.
	payload := new(AuthPayload)
	if err := schemas.HandleBindData(c, payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Login payload was malformed.")
	}

	email := payload.Email
	if email == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Login payload missing email.")
	}
	// TODO: Check email doesn't exist yet

	password := payload.Password
	if password == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Login payload missing password.")
	}

	// TODO: Look up user by email in PG and compare password
	bytes := []byte(password)
	hashedPasswordDB, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println("error occurred while hashing password::", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword(hashedPasswordDB, bytes)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials.")
	}

	// TODO: JWT and user id

	return c.JSON(http.StatusOK, "User login successful.")
}

// Handles user logout.
func (h *Handler) POSTLogout(c echo.Context) error {
	// TODO: Get JWT from header and parse out user id. Then remove token.
	return c.JSON(http.StatusOK, "User logout successful.")
}

func (h *Handler) POSTValidateToken(c echo.Context) error {
	// TODO: Get JWT from header and parse out user id. Then remove token.
	return c.JSON(http.StatusOK, "User logout successful.")
}
func (h *Handler) POSTRefreshToken(c echo.Context) error {
	// TODO: Get JWT from header and parse out user id. Then remove token.
	return c.JSON(http.StatusOK, "User logout successful.")
}
