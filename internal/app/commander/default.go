package commander

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "ты написал "+inputMessage.Text+" тем не менее пошел нахуй")
	msg.ReplyToMessageID = inputMessage.MessageID
	c.bot.Send(msg)
}
