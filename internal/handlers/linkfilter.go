package handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"log"
	"regexp"
)

type LinkFilter struct {
}

func (sf *LinkFilter) HandleMessage(c telebot.Context) error {
	messageText := c.Text()
	sender := c.Sender()

	if containsLink(messageText) {
		err := c.Delete()
		if err != nil {
			log.Printf("[ERROR] Ошибка удаления сообщения: %v\n", err)
		} else {
			fmt.Printf("[LINK] Удалено сообщение с ссылкой от %s: %s\n", sender.Username, messageText)
		}

		return nil
	}

	return nil
}

func containsLink(text string) bool {
	re := regexp.MustCompile(`(?:https?://|www\.)?\S+\.\S+`)
	return re.MatchString(text)
}
