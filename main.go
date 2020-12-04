package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hinshaw-backend-1/handlers"
	mw "hinshaw-backend-1/middleware"
	"log"
)

func main() {

	// Initialize environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("godotenv.Load: %v\n", err)
	}

	// Setup Echo framework
	e := echo.New()
	h := handlers.NewHandler()

	// Setup middlewares
	e.Use(middleware.CORS())
	e.Use(mw.RouteLogger)
	e.Use(h.HandlerMiddleware)

	// Register the URL endpoints to Handler
	h.RegisterRoutes(e)

	// Serve
	e.Logger.Fatal(e.Start(":8080"))
}
