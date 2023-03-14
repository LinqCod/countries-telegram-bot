package telegrambot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/linqcod/countries-telegram-bot/internal/restcountriesapi"
	"github.com/spf13/viper"
	"log"
)

type CountriesBot struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	countriesApi *restcountriesapi.CountriesApi
	//db           *sql.DB
}

func NewCountriesBot(countriesApi *restcountriesapi.CountriesApi) (*CountriesBot, error) {
	token := viper.GetString("telegram-bot.token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	offset := viper.GetInt("telegram-bot.update-offset")
	timeout := viper.GetInt("telegram-bot.update-timeout")
	u := tgbotapi.UpdateConfig{
		Offset:  offset,
		Timeout: timeout,
	}

	return &CountriesBot{
		bot:          bot,
		updateConfig: u,
		countriesApi: countriesApi,
		//db:           db,
	}, nil
}

func (b *CountriesBot) Start() {
	updates := b.bot.GetUpdatesChan(b.updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		if update.Message.CommandArguments() == "" {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "get_country_by_name":
			name := update.Message.CommandArguments()
			country, err := b.countriesApi.GetCountryByName(name)
			if err != nil {
				log.Fatal(err)
			}
			msg.Text = country.SubRegion
		default:
			msg.Text = "I dont know this command"
		}

		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
