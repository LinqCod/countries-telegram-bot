package main

import (
	"github.com/linqcod/countries-telegram-bot/internal/repository"
	"github.com/linqcod/countries-telegram-bot/internal/service"
	"github.com/linqcod/countries-telegram-bot/internal/telegrambot"
	"github.com/linqcod/countries-telegram-bot/pkg/config"
	"github.com/linqcod/countries-telegram-bot/pkg/database"
	"log"
)

func init() {
	config.LoadConfig()
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewCountriesRepository(db)

	s := service.NewRestCountriesService(repo)

	bot, err := telegrambot.NewCountriesBot(s)
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}
