package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Default(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "luv u")
}
