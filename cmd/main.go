package main

import (
	countriesRepository "github.com/linqcod/countries-telegram-bot/internal/countries/repository"
	countriesService "github.com/linqcod/countries-telegram-bot/internal/countries/service"
	"github.com/linqcod/countries-telegram-bot/internal/telegrambot"
	statsRepository "github.com/linqcod/countries-telegram-bot/internal/userstatistics/repository"
	statsService "github.com/linqcod/countries-telegram-bot/internal/userstatistics/service"
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

	countriesRepo := countriesRepository.NewRepository(db)
	statsRepo := statsRepository.NewRepository(db)

	countriesServ := countriesService.NewRestCountriesService(countriesRepo)
	statsServ := statsService.NewService(statsRepo)

	bot, err := telegrambot.NewCountriesBot(countriesServ, statsServ)
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}
