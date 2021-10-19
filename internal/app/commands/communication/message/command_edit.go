package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/communication"
	"strconv"
)

func (m MessageCommander) Edit(inputMsg *tgbotapi.Message) {
	messageData, error := CheckMessageInput(inputMsg.CommandArguments(), editArgumentRowsCount)
	text := ""
	if error != nil {
		text = fmt.Sprintf("Less then %v rows of values were provided", editArgumentRowsCount)
	} else {
		message := CreateMessage(messageData[1:])
		lookupMessageNumber, error := strconv.ParseUint(messageData[0], 0, 64)
		if error != nil {
			text = "Incorrect index provided"
		} else {
			error := UpdateMessage(lookupMessageNumber - 1, message, m.messageService)
			if error != nil {
				text = fmt.Sprintf("Message for update was not found")
			} else {
				text = fmt.Sprintf("Message was successfully updated")
			}
		}
	}

	newBotMessage := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	m.bot.Send(newBotMessage)
}

func UpdateMessage(index uint64, updateMessageData communication.Message, messageService IMessageService) error {
	return messageService.Update(index, updateMessageData)
}
