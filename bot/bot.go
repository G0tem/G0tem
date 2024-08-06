package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func RunBot() {
	bot, err := tgbotapi.NewBotAPI("ВАШ_ТОКЕН_БОТА")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Авторизован на аккаунте %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // пропускаем любые несообщения обновления
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
