package main

import (
	"github.com/Haise240/al_tour/database"
	"github.com/Haise240/al_tour/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database connection
	_, err := database.ConnectDB("your-database-connection-string")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Routes
	e.POST("/admin/tours", handlers.CreateTour)
	e.PUT("/admin/tours/:id", handlers.UpdateTour)
	e.DELETE("/admin/tours/:id", handlers.DeleteTour)

	e.GET("/tours", handlers.GetTours)
	e.GET("/tours/:id", handlers.GetTour)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
