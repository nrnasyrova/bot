package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nrnasyrova/bot/internal/app/commands"
	"github.com/nrnasyrova/bot/internal/service/product"
) 

func main() {
	token := os.Getenv("TOKEN")
	
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)



	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()


	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			
			var msg tgbotapi.MessageConfig
			var commander = commands.NewCommander(bot, productService)

			switch update.Message.Command() {
				case "help":
					msg = commander.Help(update.Message)
				case "list":
					msg = commander.List(update.Message)
				default:
					msg = commander.Default(update.Message)
			}
		
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}