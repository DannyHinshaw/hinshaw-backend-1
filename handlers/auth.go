package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	mw "hinshaw-backend-1/middleware"
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
		return err
	}

	email := payload.Email
	if email == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Registration missing email.")
	}

	userExists, err := h.DBService.QueryUserEmailExists(email, h.Context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Check a user with email doesn't exist yet.
	if userExists {
		return echo.NewHTTPError(http.StatusConflict, "A user with that email already exists.")
	}

	password := payload.Password
	if len(password) < 5 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Password must have at least 5 characters.")
	}
	// TODO: Any other password restriction checks here.

	// Add user to db
	err = h.DBService.AddNewUser(email, password, h.Context)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, "Registered new user successfully.")
}

// Handles user login with email/password..
func (h *Handler) POSTLogin(c echo.Context) error {

	// Extract request payload data.
	payload := new(AuthPayload)
	if err := schemas.HandleBindData(c, payload); err != nil {
		return err
	}

	email := payload.Email
	if email == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Login payload missing email.")
	}

	password := payload.Password
	if password == "" {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Login payload missing password.")
	}

	// Look up users stored password by email and compare against request password.
	bytes := []byte(password)
	userAuth, err := h.DBService.QueryUserAuth(email, h.Context)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userAuth.Password), bytes)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials.")
	}

	jwtPayload, err := mw.GenerateJWT(userAuth.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	err = h.RedisService.SetJWTRedis(jwtPayload.AccessToken, userAuth.UserId)
	if err != nil {
		log.Println("error saving jwt in redis on login::", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, jwtPayload)
}

// Handles user logout by removing key from redis.
func (h *Handler) POSTLogout(c echo.Context) error {
	h.RedisService.ExpireKey(h.Token)
	return c.JSON(http.StatusOK, "User logout successful.")
}

func (h *Handler) POSTValidateToken(c echo.Context) error {
	return c.JSON(http.StatusOK, "Token validated successfully.")
}
