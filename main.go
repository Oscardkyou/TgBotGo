package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

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

func main() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Panic("Ошибка загрузки файла .env")
	}

	// Создание бота с использованием переменной окружения TELEGRAM_APITOKEN
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	// Включение режима отладки бота
	bot.Debug = true

	log.Printf("Авторизован под учетной записью %s", bot.Self.UserName)

	// Настройка канала для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Обработка входящих обновлений
	for update := range updates {
		if update.Message != nil {
			handleMessage(bot, update.Message)
			handleCommand(bot, update.Message)
		} else if update.CallbackQuery != nil {
			handleCallbackQuery(bot, update.CallbackQuery)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
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

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
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

func handleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
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
