package main

import (
	"log"
	"myapp/internal/bot"
	"myapp/internal/command"
	"myapp/internal/service"
)

func main() {
	command.ServiceCovnverter = &service.ConverterService{}
	b := bot.NewBot()
	b.Run()
	log.Println("Bot stopped.")
}
