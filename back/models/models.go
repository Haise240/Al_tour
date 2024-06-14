package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type Tour struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"descriprion"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	Seats       int     `json:"suration"`
}

// Booking represents a booking made by a user
type Booking struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	TourID       uint      `json:"tour_id" gorm:"not null"`
	UserName     string    `json:"user_name" gorm:"not null"`
	UserContact  string    `json:"user_contact" gorm:"not null"` // Telegram или WhatsApp идентификатор
	CheckInDate  time.Time `json:"check_in_date" gorm:"not null"`
	CheckOutDate time.Time `json:"check_out_date" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
