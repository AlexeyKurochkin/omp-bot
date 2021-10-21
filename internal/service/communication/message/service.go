package message

import (
	"errors"
	"github.com/ozonmp/omp-bot/internal/model/communication"
)

type DummyMessageService struct{}

func NewDummyMessageService() *DummyMessageService {
	return &DummyMessageService{}
}

func (d DummyMessageService) Describe(messageID uint64) (*communication.Message, error) {
	for i := 0; i < len(communication.AllMessages); i++ {
		if messageID == communication.AllMessages[i].ID {
			return &communication.AllMessages[i], nil
		}
	}

	return nil, errors.New("Message was not found")
}

func (d DummyMessageService) List(cursor uint64, limit uint64) ([]communication.Message, error) {
	if err := CheckOutOfBound(cursor); err != nil {
		return nil, err
	}

	dataCount := uint64(len(communication.AllMessages))
	endIndex := cursor + limit
	if dataCount <= endIndex {
		endIndex = dataCount
	}

	return communication.AllMessages[cursor:endIndex], nil
}

func (d DummyMessageService) Create(message communication.Message) (uint64, error) {
	message.ID = GenerateNewMessageId()
	communication.AllMessages = append(communication.AllMessages, message)
	return message.ID, nil
}

func (d DummyMessageService) Update(messageID uint64, message communication.Message) error {
	oldMessage, error := d.Describe(messageID)
	if error != nil {
		return error
	}

	oldMessage.From = message.From
	oldMessage.To = message.To
	oldMessage.Text = message.Text
	return nil
}

func (d DummyMessageService) Remove(messageID uint64) (bool, error) {
	for i := 0; i < len(communication.AllMessages); i++ {
		if messageID == communication.AllMessages[i].ID {
			communication.AllMessages = append(communication.AllMessages[:i], communication.AllMessages[i + 1:]...)
			return true, nil
		}
	}

	return false, errors.New("Message was not found")
}

func GenerateNewMessageId() uint64 {
	dataLength := uint64(len(communication.AllMessages))
	maxCurrentIndex := communication.AllMessages[dataLength - 1].ID
	return maxCurrentIndex + 1
}

func CheckOutOfBound(cursor uint64) (error) {
	dataCount := uint64(len(communication.AllMessages))
	if cursor >= dataCount {
		return errors.New(("out of range of messages array"))
	}

	return nil
}
