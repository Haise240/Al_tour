package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at`
}

type Tour struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"descriprion"`
	Price       float64   `json:"price"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at`
}

// Booking represents a booking made by a user
type Booking struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TourID      int       `json:"tour_id"`
	BookingDate time.Time `json:"booking_date"`
	Status      string    `json:"status"`
	TotalPrice  float64   `json:"total_price"`
}

// Review represents a review made by a user
type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	TourID    int       `json:"tour_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
