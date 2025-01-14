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

	msg := deimbotapi.NewPhoto("test002", deimbotapi.FileURL("https://s1.xx.io/apps/ak-banner.png"))
	res, err := bot.Request(msg)
	if err != nil {
		log.Fatal("Unable to send text message")
	}
	log.Printf("send photo message successfully,result: %s", string(res.Result))
}
