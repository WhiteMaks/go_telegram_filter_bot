package handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"log"
	"strings"
)

type SpamFilter struct {
	ForbiddenWords []string
}

func (sf *SpamFilter) HandleMessage(c telebot.Context) error {
	messageText := c.Text()
	sender := c.Sender()

	for _, word := range sf.ForbiddenWords {
		if strings.Contains(strings.ToLower(messageText), word) {
			err := c.Delete()
			if err != nil {
				log.Printf("[ERROR] Ошибка удаления сообщения: %v\n", err)
			} else {
				fmt.Printf("[FORBIDDEN WORD] Удалено сообщение от %s: %s\n", sender.Username, messageText)
			}

			return nil
		}
	}

	return nil
}
