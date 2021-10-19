package message

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (m MessageCommander) Delete(inputMsg *tgbotapi.Message) {
	arguments := inputMsg.CommandArguments()
	index, error := strconv.ParseUint(arguments, 0, 64)
	if error != nil {
		log.Println("Incorrect message number", error)
	}

	successfullyDeleted, error := m.messageService.Remove(index - 1)
	if error != nil {
		log.Println("Error appeared during deletion of message", error)
	}

	if successfullyDeleted {
		messageText := fmt.Sprintf("Successfully deleted message with index %v", index)
		message := tgbotapi.NewMessage(inputMsg.Chat.ID, messageText)
		_, error := m.bot.Send(message)
		if error != nil {
			log.Printf("Error sending message to chat %v", error)
		}
	}

}
