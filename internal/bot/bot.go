package bot

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	"myapp/internal/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("info.kz", "info.kz"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

type Bot struct {
	*tgbotapi.BotAPI
}

func NewBot() *Bot {
	if err := godotenv.Load(); err != nil {
		log.Panic("Ошибка загрузки файла .env")
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован под учетной записью %s", bot.Self.UserName)

	return &Bot{BotAPI: bot}
}

func (b *Bot) Run() {
	updates := b.GetUpdatesChan(tgbotapi.NewUpdate(0))
	handler.HandleUpdates(b.BotAPI, updates)
}

func handleMessage(bot *Bot, message *tgbotapi.Message) {
	// Создание нового сообщения с полученным текстом
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	// Если сообщение "open", добавляем клавиатуру
	if message.Text == "open" {
		msg.ReplyMarkup = numericKeyboard
	}

	// Отправка сообщения
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func handleCommand(bot *Bot, message *tgbotapi.Message) {
	if !message.IsCommand() {
		return
	}

	// Создание нового сообщения. Пока текст пустой.
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	// Извлечение команды из сообщения
	switch message.Command() {
	case "help":
		msg.Text = "Я понимаю /sayhi, /status и /random."
	case "sayhi":
		msg.Text = "Привет :)"
	case "status":
		msg.Text = "Всё в порядке."
	case "random":
		msg.Text = strconv.Itoa(rand.Intn(100)) // Генерация случайного числа от 0 до 99
	default:
		msg.Text = "Неизвестная команда"
	}

	// Отправка сообщения с ответом на команду
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func handleCallbackQuery(bot *Bot, callbackQuery *tgbotapi.CallbackQuery) {
	// Ответ на коллбек-запрос
	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		log.Panic(err)
	}

	// Отправка сообщения с данными из коллбек-запроса
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, callbackQuery.Data)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
