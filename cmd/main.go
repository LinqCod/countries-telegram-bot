package main

import (
	"github.com/linqcod/countries-telegram-bot/internal/restcountriesapi"
	"github.com/linqcod/countries-telegram-bot/internal/telegrambot"
	"github.com/linqcod/countries-telegram-bot/pkg/config"
	"log"
)

func init() {
	config.LoadConfig()
}

func main() {
	//db, err := database.InitDB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	countriesApi := restcountriesapi.NewCountriesApi()

	bot, err := telegrambot.NewCountriesBot(countriesApi)
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()
}
