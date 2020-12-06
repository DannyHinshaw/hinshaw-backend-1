package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hinshaw-backend-1/cache"
	"hinshaw-backend-1/db"
	"hinshaw-backend-1/handlers"
	mw "hinshaw-backend-1/middleware"
	"log"
	"os"
)

func main() {

	// Init DB and prep for handlers DI.
	dbs := db.DatabaseService
	err := dbs.ParseDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	err = dbs.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer dbs.Close()

	// Init redis client
	rdc := cache.NewRedisClient()

	// Setup Echo framework
	e := echo.New()
	h := handlers.NewHandler(&dbs, rdc)

	// Setup middlewares
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8000"},
		AllowHeaders: []string{
			echo.HeaderAuthorization, echo.HeaderOrigin,
			echo.HeaderContentType, echo.HeaderAccept,
		},
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XFrameOptions:         "DENY",
		ContentSecurityPolicy: "default-src 'self'",
	}))
	e.Use(mw.RouteLogger)
	e.Use(h.HandlerMiddleware)

	// Register the URL endpoints to Handler
	h.RegisterRouteHandlers(e)

	// Serve
	e.Logger.Fatal(e.Start(":8080"))
}
