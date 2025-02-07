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

	setCommands := deimbotapi.NewSetMyCommands(deimbotapi.BotCommand{
		Command:     "test",
		Description: "a test command",
	})

	res, err := bot.Request(setCommands)
	if err != nil {
		log.Fatal("Unable to set commands")
	}
	log.Printf("set command successfully,result: %s", string(res.Result))

	commands, err := bot.GetMyCommands()
	if err != nil {
		log.Fatal("Unable to get commands")
	}

	log.Printf("get command successfully,result: %v", commands)

	if len(commands) != 1 {
		log.Fatal("Incorrect number of commands returned")
	}

	if commands[0].Command != "test" || commands[0].Description != "a test command" {
		log.Fatal("Commands were incorrectly set")
	}

	delCommands := deimbotapi.NewDeleteMyCommands()

	res, err = bot.Request(delCommands)
	if err != nil {
		log.Fatal("Unable to del commands")
	}

	log.Printf("del command successfully,result: %v", res.Description)

	commands, err = bot.GetMyCommands()
	if err != nil {
		log.Fatal("Unable to get commands")
	}

	log.Printf("get command successfully,result: %v", commands)

}
