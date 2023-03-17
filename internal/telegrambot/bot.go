package telegrambot

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	apierror "github.com/linqcod/countries-telegram-bot/internal/error"
	"github.com/linqcod/countries-telegram-bot/internal/model"
	"github.com/linqcod/countries-telegram-bot/internal/service"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type CountriesBot struct {
	bot          *tgbotapi.BotAPI
	updateConfig tgbotapi.UpdateConfig
	service      *service.RestCountriesService
}

func NewCountriesBot(service *service.RestCountriesService) (*CountriesBot, error) {
	token := viper.GetString("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	offset := viper.GetInt("UPDATE_OFFSET")
	timeout := viper.GetInt("UPDATE_TIMEOUT")
	u := tgbotapi.UpdateConfig{
		Offset:  offset,
		Timeout: timeout,
	}

	return &CountriesBot{
		bot:          bot,
		updateConfig: u,
		service:      service,
	}, nil
}

func (b *CountriesBot) Start() {
	updates := b.bot.GetUpdatesChan(b.updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, use commands from list")
			if _, err := b.bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "get_country_by_name":
			if update.Message.CommandArguments() == "" {
				msg.Text = "Please, enter country name (E.g. /get_country_by_name Russia)"
				break
			}
			name := update.Message.CommandArguments()
			country, err := b.service.GetCountryByName(name)
			if err != nil {
				if errors.Is(err, apierror.CountryNotFound) {
					msg.Text = err.Error()
					break
				}
				log.Fatal(err)
			}
			msg.ParseMode = tgbotapi.ModeHTML

			log.Println(country.Name)
			msg.Text = createCountryInfoHtml(country)
		default:
			msg.Text = "I dont know this command"
		}

		if _, err := b.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

func createCountryInfoHtml(country *model.Country) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("<b>%s</b>\n", strings.ToUpper(country.Name.Official)))
	sb.WriteString(fmt.Sprintf("<i>%s</i>\n\n", country.Name.Common))
	sb.WriteString(fmt.Sprintf("<b> - Area:</b> %.0f km2\n<b> - Population:</b> %d\n\n",
		country.Area,
		country.Population,
	))
	if len(country.Capital) == 1 {
		sb.WriteString(fmt.Sprintf("<b> - Capital:</b> %s\n\n", country.Capital[0]))
	} else {
		sb.WriteString(fmt.Sprintf("<b> - Capitals:</b> %s\n\n", strings.Join(country.Capital, ", ")))
	}

	sb.WriteString(fmt.Sprintf("<b> - Continents:</b> %s\n", strings.Join(country.Continents, ", ")))
	sb.WriteString(fmt.Sprintf("<b> - Region:</b> %s\n", country.Region))
	sb.WriteString(fmt.Sprintf("<b> - Subregion:</b> %s\n\n", country.SubRegion))

	languages := make([]string, len(country.Languages))
	i := 0
	for _, l := range country.Languages {
		languages[i] = l
		i++
	}
	sb.WriteString(fmt.Sprintf("<b> - Languages:</b> %s\n", strings.Join(languages, ", ")))

	currencies := make([]string, len(country.Currencies))
	i = 0
	for _, c := range country.Currencies {
		currencies[i] = fmt.Sprintf("%s (%s)", c.Name, c.Symbol)
		i++
	}
	sb.WriteString(fmt.Sprintf("<b> - Currencies:</b> %s\n\n", strings.Join(currencies, ", ")))

	sb.WriteString(fmt.Sprintf("<b> - Status:</b> %s\n", country.Status))
	sb.WriteString(fmt.Sprintf("<b> - Is independent:</b> %v\n", country.Independent))
	sb.WriteString(fmt.Sprintf("<b> - Is UN member:</b> %v\n\n", country.UnMember))

	return sb.String()
}
