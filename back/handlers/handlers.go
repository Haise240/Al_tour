package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Haise240/al_tour/database"
	"github.com/Haise240/al_tour/models"
	"github.com/labstack/echo/v4"
)

func CreateBooking(c echo.Context) error {
	var booking models.Booking
	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	booking.BookingDate = time.Now()
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO bookings (tour_id, user_contact, check_in_date, check_out_date, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		booking.TourID, booking.UserContact, booking.CheckInDate, booking.CheckOutDate, booking.CreatedAt).Scan(&booking.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create booking"})
	}

	return c.JSON(http.StatusOK, booking)
}

func GetBookings(c echo.Context) error {
	userContact := c.QueryParam("user_contact")
	rows, err := database.DB.Query(context.Background(),
		"SELECT id, tour_id, user_contact, check_in_date, check_out_date, created_at FROM bookings WHERE user_contact=$1",
		userContact)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch bookings"})
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.TourID, &booking.UserContact, &booking.CheckInDate, &booking.CheckOutDate, &booking.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan booking"})
		}
		bookings = append(bookings, booking)
	}

	return c.JSON(http.StatusOK, bookings)
}

func CreateTour(c echo.Context) error {
	var tour models.Tour
	if err := c.Bind(&tour); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	tour.CreatedAt = time.Now()
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO tours (title, description, price, duration, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		tour.Title, tour.Description, tour.Price, tour.Duration, tour.Seats).Scan(&tour.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create tour"})
	}

	return c.JSON(http.StatusOK, tour)
}

func UpdateTour(c echo.Context) error {
	id := c.Param("id")
	var tour models.Tour
	if err := c.Bind(&tour); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	_, err := database.DB.Exec(context.Background(),
		"UPDATE tours SET title=$1, description=$2, price=$3, duration=$4 WHERE id=$5",
		tour.Title, tour.Description, tour.Price, tour.Duration, id, tour.Seats)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update tour"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Tour updated successfully"})
}

func DeleteTour(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DB.Exec(context.Background(), "DELETE FROM tours WHERE id=$1", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete tour"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Tour deleted successfully"})
}

func GetTours(c echo.Context) error {
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, description, price, duration, created_at FROM tours")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch tours"})
	}
	defer rows.Close()

	var tours []models.Tour
	for rows.Next() {
		var tour models.Tour
		if err := rows.Scan(&tour.ID, &tour.Title, &tour.Description, &tour.Price, &tour.Duration, &tour.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan tour"})
		}
		tours = append(tours, tour)
	}

	return c.JSON(http.StatusOK, tours)
}

func GetTour(c echo.Context) error {
	id := c.Param("id")
	var tour models.Tour
	err := database.DB.QueryRow(context.Background(),
		"SELECT id, title, description, price, duration, created_at FROM tours WHERE id=$1",
		id).Scan(&tour.ID, &tour.Title, &tour.Description, &tour.Price, &tour.Duration, &tour.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch tour"})
	}
	return c.JSON(http.StatusOK, tour)
}
