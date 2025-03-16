package main

import (
	"bufio"
	"go_telegram_filter_bot/internal/bot"
	"go_telegram_filter_bot/internal/config"
	"go_telegram_filter_bot/internal/handlers"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Ошибка: не передан токен бота. Используйте: ./bot YOUR_BOT_TOKEN")
	}

	token := os.Args[1]

	forbiddenWords := readForbiddenWords()

	botConfig := config.Config{
		Token:          token,
		ForbiddenWords: forbiddenWords,
	}

	messageHandlers := []handlers.MessageHandler{
		&handlers.Logger{},
		&handlers.SpamFilter{ForbiddenWords: botConfig.ForbiddenWords},
		&handlers.LinkFilter{},
	}

	tgBot, err := bot.NewBot(botConfig, messageHandlers)
	if err != nil {
		log.Fatal(err)
	}

	tgBot.Run()
}

func readForbiddenWords() []string {
	var result []string

	file, err := os.Open("forbidden_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	wordSet := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wordSet[word] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	for word := range wordSet {
		result = append(result, word)
	}

	return result
}
