package main

import (
	"log"

	deimbotapi "github.com/De-IM/de-im-bot-go"
)

func main() {
	botToken := "APITOKEN"
	bot, err := deimbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := deimbotapi.NewMessage("test001", "hello world")
	res, err := bot.Request(msg)
	if err != nil {
		log.Fatal("Unable to send text message")
	}
	log.Printf("send text message successfully,result: %s", string(res.Result))
}
