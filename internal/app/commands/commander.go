package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nrnasyrova/bot/internal/service/product"
)

type CommandRouter struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) * CommandRouter {
	return &CommandRouter{
		bot: bot,
		productService: service,
	}
}