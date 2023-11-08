package usecase

import (
	"chat/internal/entity"
	"chat/internal/service"
	"encoding/json"
	"errors"
)

type MessageUseCase interface {
	Create(body []byte) error
	Get(body []byte) ([]entity.Message, error)
	Delete(id int) error
}

type messageUseCase struct {
	s service.MessageService
}

var _ MessageUseCase = (*messageUseCase)(nil)

func NewMessageUseCase(s service.MessageService) *messageUseCase {
	return &messageUseCase{
		s: s,
	}
}

func (uc *messageUseCase) Create(body []byte) error {
	item := &entity.Message{}
	err := json.Unmarshal(body, item)
	if err != nil {
		return err
	}

	err = uc.s.Create(item.Author, item.Text)
	if err != nil {
		return err
	}

	return nil
}

func (uc *messageUseCase) Get(body []byte) ([]entity.Message, error) {
	// if body will be empty, set default value for limit
	limit := map[string]int{"limit": 50}

	// try to unmarshal body if not empty
	if len(body) > 0 {
		err := json.Unmarshal(body, &limit)
		if err != nil {
			return nil, errors.New("incorrect body")
		}
	}

	// get limit from map where saved default value
	value, _ := limit["limit"]

	// get items
	items, err := uc.s.Get(value)
	if err != nil {
		return nil, errors.New("failed to get")
	}

	return items, nil
}

func (uc *messageUseCase) Delete(id int) error {
	err := uc.s.Delete(id)
	if err != nil {
		return errors.New("failed to delete")
	}

	return nil
}
