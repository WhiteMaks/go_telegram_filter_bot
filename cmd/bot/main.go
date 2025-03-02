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
	forbiddenWords := readForbiddenWords()

	botConfig := config.Config{
		Token:          "",
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
