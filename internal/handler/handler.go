package handler

import (
	"myapp/internal/bot"
	"myapp/internal/callback"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(bot *bot.Bot, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
			handleCommand(bot, update.Message)
		} else if update.CallbackQuery != nil {
			callback.HandleCallbackQuery(bot, update.CallbackQuery)
		}
	}
}
