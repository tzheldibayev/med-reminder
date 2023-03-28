package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal("Error loading .env file: " + err.Error())
	//}

	botToken := "6117441992:AAF1gwFr2SuT2yHhY9ojWR73qYuuvJzSReM"

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(bot, update.Message)
		}
	}
}

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		handleStartCommand(bot, message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не найдена. Введите /start для начала работы с ботом.")
		bot.Send(msg)
	}
}

func handleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	response := "Привет! Я бот 'Дед пей таблетки', и я здесь, чтобы помочь тебе не забывать о приеме таблеток. " +
		"Чтобы добавить напоминание о приеме таблеток, используй команду /add, например: /add Парацетамол. " +
		"После этого я задам несколько вопросов, чтобы настроить напоминания для тебя."

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	bot.Send(msg)
}
