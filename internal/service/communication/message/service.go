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
	if err := CheckOutOfBound(messageID); err != nil {
		return nil, err
	}

	return &communication.AllMessages[messageID], nil
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
	communication.AllMessages = append(communication.AllMessages, message)
	return uint64(len(communication.AllMessages)), nil
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
	if error := CheckOutOfBound(messageID); error != nil {
		return false, error
	}

	communication.AllMessages = append(communication.AllMessages[:messageID], communication.AllMessages[messageID + 1:]...)
	return true, nil
}

func CheckOutOfBound(cursor uint64) (error) {
	dataCount := uint64(len(communication.AllMessages))
	if cursor >= dataCount {
		return errors.New(("out of range of messages array"))
	}

	return nil
}
