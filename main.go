package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// Setup Echo framework
	e := echo.New()
	h := handlers.NewHandler(&dbs)

	// Setup middlewares
	e.Use(middleware.CORS())
	e.Use(mw.RouteLogger)
	e.Use(h.HandlerMiddleware)

	// Register the URL endpoints to Handler
	h.RegisterRoutes(e)

	// Serve
	e.Logger.Fatal(e.Start(":8080"))
}
