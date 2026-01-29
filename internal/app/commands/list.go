package commands

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	products := c.productService.List()
	var outputMsg strings.Builder
	outputMsg.WriteString("Here are all the products: \n\n")

	for _, p := range products {
		outputMsg.WriteString(p.Title)
		outputMsg.WriteString("\n")
	}

	return tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg.String())
}
