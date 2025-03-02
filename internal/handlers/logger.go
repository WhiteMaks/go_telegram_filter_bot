package handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"time"
)

type Logger struct{}

func (l *Logger) HandleMessage(c telebot.Context) error {
	fmt.Printf("[%s] Получено сообщение от %s: %s\n", time.Now().Format("02-01-2006::15:04:05"), c.Sender().Username, c.Text())
	return nil
}
