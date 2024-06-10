package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func checkBookingStatus(phone string, bookingID string) (string, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return "", err
	}
	defer db.Close()

	var status string
	query := "SELECT status FROM bookings WHERE phone = ? AND id = ?"
	err = db.QueryRow(query, phone, bookingID).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Бронь не найдена. Пожалуйста, проверьте введенные данные.", nil
		}
		return "", err
	}
	return fmt.Sprintf("Информация о вашей брони: %s", status), nil
}
