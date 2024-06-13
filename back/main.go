package main

import (
	"os"

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

	// Connect to database
	database.Connect()
	defer database.Close()

	// Routes
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	e.GET("/tours", handlers.GetTours)
	e.GET("/tours/:id", handlers.GetTour)
	e.POST("/tours", handlers.CreateTour)
	e.PUT("/tours/:id", handlers.UpdateTour)
	e.DELETE("/tours/:id", handlers.DeleteTour)

	e.POST("/bookings", handlers.CreateBooking)
	e.GET("/bookings", handlers.GetBookings)

	e.POST("/reviews", handlers.CreateReview)
	e.GET("/reviews/:tour_id", handlers.GetReviews)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
