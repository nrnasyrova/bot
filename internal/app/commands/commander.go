package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nrnasyrova/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: service,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {

	if update.Message != nil { // If we got a message
		var msg tgbotapi.MessageConfig

		switch update.Message.Command() {
		case "help":
			msg = c.Help(update.Message)
		case "list":
			msg = c.List(update.Message)
		case "get":
			msg = c.Get(update.Message)
		default:
			msg = c.Default(update.Message)
		}

		msg.ReplyToMessageID = update.Message.MessageID

		c.bot.Send(msg)
	}
}
