# Keyboard

This bot shows a numeric keyboard when you send a "open" message and hides it
when you send "close" message.

```go
package main

import (
	"log"
	"os"

	"github.com/De-IM/de-im-bot-go"
)

var numericKeyboard = deimbotapi.NewReplyKeyboard(
	deimbotapi.NewKeyboardButtonRow(
		deimbotapi.NewKeyboardButton("1"),
		deimbotapi.NewKeyboardButton("2"),
		deimbotapi.NewKeyboardButton("3"),
	),
	deimbotapi.NewKeyboardButtonRow(
		deimbotapi.NewKeyboardButton("4"),
		deimbotapi.NewKeyboardButton("5"),
		deimbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	bot, err := deimbotapi.NewBotAPI(os.Getenv("DE_IM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := deimbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := deimbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "open":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = deimbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
```
