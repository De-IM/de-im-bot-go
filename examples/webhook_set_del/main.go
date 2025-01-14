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

	wh, _ := deimbotapi.NewWebhook("https://www.example.com/" + botToken)

	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("get webhook successfully,result: %v", info)

	bot.Request(deimbotapi.DeleteWebhookConfig{})

	info, err = bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("get webhook successfully,result: %v", info)
}
