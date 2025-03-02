package handlers

import "gopkg.in/telebot.v3"

type MessageHandler interface {
	HandleMessage(c telebot.Context) error
}
