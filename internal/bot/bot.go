package bot

import (
	"go_telegram_filter_bot/internal/config"
	"go_telegram_filter_bot/internal/handlers"
	"gopkg.in/telebot.v3"
	"time"
)

type Bot struct {
	TelegramBot *telebot.Bot
	Handlers    []handlers.MessageHandler
}

func NewBot(config config.Config, handlers []handlers.MessageHandler) (*Bot, error) {
	settings := telebot.Settings{
		Token:  config.Token,
		Poller: &telebot.LongPoller{Timeout: 3 * time.Second},
	}

	telegramBot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, err
	}

	bot := &Bot{
		TelegramBot: telegramBot,
		Handlers:    handlers,
	}

	bot.TelegramBot.Handle(telebot.OnText, bot.HandleMessage)

	return bot, nil
}

func (bot *Bot) HandleMessage(c telebot.Context) error {
	for _, handler := range bot.Handlers {
		if err := handler.HandleMessage(c); err != nil {
			return err
		}
	}
	return nil
}

func (bot *Bot) Run() {
	bot.TelegramBot.Start()
}
