package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "incorrect args")
	}

	product, ok := c.productService.Get(arg)

	if !ok {
		return tgbotapi.NewMessage(inputMessage.Chat.ID, "this product does not exist")
	}

	return tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
}
