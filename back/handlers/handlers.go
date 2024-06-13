package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Haise240/al_tour/database"
	"github.com/Haise240/al_tour/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()

	_, err = database.DB.Exec(context.Background(),
		"INSERT INTO users (username, password, email, created_at) VALUES ($1, $2, $3, $4)",
		user.Username, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User registered successfully"})
}

func LoginUser(c echo.Context) error {
	var creds Credentials
	if err := c.Bind(&creds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	var user models.User
	err := database.DB.QueryRow(context.Background(),
		"SELECT id, username, password FROM users WHERE username=$1",
		creds.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid username or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid username or password"})
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}

func GetTours(c echo.Context) error {
	rows, err := database.DB.Query(context.Background(), "SELECT id, title, description, price, duration, created_at FROM tours")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch tours"})
	}
	defer rows.Close()

	var tours []models.Tour
	for rows.Next() {
		var tour models.TTour
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

func CreateTour(c echo.Context) error {
	var tour models.Tour
	if err := c.Bind(&tour); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	tour.CreatedAt = time.Now()
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO tours (title, description, price, duration, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		tour.Title, tour.Description, tour.Price, tour.Duration, tour.CreatedAt).Scan(&tour.ID)
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
		tour.Title, tour.Description, tour.Price, tour.Duration, id)
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

func CreateBooking(c echo.Context) error {
	var booking models.Booking
	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	booking.BookingDate = time.Now()
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO bookings (user_id, tour_id, booking_date, status, total_price) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		booking.UserID, booking.TourID, booking.BookingDate, booking.Status, booking.TotalPrice).Scan(&booking.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create booking"})
	}

	return c.JSON(http.StatusOK, booking)
}

func GetBookings(c echo.Context) error {
	userID := c.QueryParam("user_id")
	rows, err := database.DB.Query(context.Background(),
		"SELECT id, user_id, tour_id, booking_date, status, total_price FROM bookings WHERE user_id=$1",
		userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch bookings"})
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.ID, &booking.UserID, &booking.TourID, &booking.BookingDate, &booking.Status, &booking.TotalPrice); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan booking"})
		}
		bookings = append(bookings, booking)
	}

	return c.JSON(http.StatusOK, bookings)
}

func CreateReview(c echo.Context) error {
	var review models.Review
	if err := c.Bind(&review); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	review.CreatedAt = time.Now()
	err := database.DB.QueryRow(context.Background(),
		"INSERT INTO reviews (user_id, tour_id, rating, comment, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		review.UserID, review.TourID, review.Rating, review.Comment, review.CreatedAt).Scan(&review.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create review"})
	}

	return c.JSON(http.StatusOK, review)
}

func GetReviews(c echo.Context) error {
	tourID := c.Param("tour_id")
	rows, err := database.DB.Query(context.Background(),
		"SELECT id, user_id, tour_id, rating, comment, created_at FROM reviews WHERE tour_id=$1",
		tourID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch reviews"})
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.TourID, &review.Rating, &review.Comment, &review.CreatedAt); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to scan review"})
		}
		reviews = append(reviews, review)
	}

	return c.JSON(http.StatusOK, reviews)
}
