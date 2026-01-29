package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *CommandRouter) List(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	products := c.productService.List()
	outputMsg := "Here are all the products: \n\n"

	for _, p := range products {
		outputMsg += p.Title
		outputMsg += "\n"
	}

	return tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)
}