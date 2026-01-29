package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			
			var msg tgbotapi.MessageConfig

			switch update.Message.Command() {
				case "help":
					msg = processHelpCommand(update.Message)
				case "list":
					msg = processListCommand(update.Message)
				default:
					msg = processDefaultBehavior(update.Message)
			}
		
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func processHelpCommand(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID,
		 "/help - help\n"+
		 	"/list - list products",
		)

}

func processDefaultBehavior(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "luv u")
}

func processListCommand(inputMessage *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(inputMessage.Chat.ID, "in process")
}