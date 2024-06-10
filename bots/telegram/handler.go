package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	PhoneState = iota
	BookingIDState
)

var userStates = make(map[int64]int)
var userPhones = make(map[int64]string)

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	switch userStates[chatID] {
	case PhoneState:
		userPhones[chatID] = text
		msg := tgbotapi.NewMessage(chatID, "Спасибо! Теперь введите номер вашей брони:")
		bot.Send(msg)
		userStates[chatID] = BookingIDState

	case BookingIDState:
		phone := userPhones[chatID]
		bookingID := text
		status, err := checkBookingStatus(phone, bookingID)
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Произошла ошибка при проверке брони. Попробуйте позже.")
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, status)
			bot.Send(msg)
		}
		delete(userStates, chatID)
		delete(userPhones, chatID)

	default:
		msg := tgbotapi.NewMessage(chatID, "Здравствуйте! Введите ваш номер телефона:")
		bot.Send(msg)
		userStates[chatID] = PhoneState
	}
}
