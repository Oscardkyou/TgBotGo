package command

import (
	"log"
	"math/rand"
	"my_project/internal/bot"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *bot.Bot, message *tgbotapi.Message) {
	if !message.IsCommand() {
		return
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	switch message.Command() {
	case "help":
		msg.Text = "Я понимаю /sayhi, /status и /random."
	case "sayhi":
		msg.Text = "Привет :)"
	case "status":
		msg.Text = "Всё в порядке."
	case "random":
		msg.Text = strconv.Itoa(rand.Intn(100))
	default:
		msg.Text = "Неизвестная команда"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
