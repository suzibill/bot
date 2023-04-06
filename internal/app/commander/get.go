package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong arg:", arg)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Success parce argument")
	c.bot.Send(msg)

}
