package service

import (
	"chat/internal/entity"
	"chat/internal/repository"
	"errors"
	"regexp"
)

type MessageService interface {
	Create(author, text string) error
	Get(limit ...int) ([]entity.Message, error)
	Delete(id int) error
}

type messageService struct {
	r repository.MessageRepository
}

var _ MessageService = (*messageService)(nil)

func NewMessageService(r repository.MessageRepository) MessageService {
	return &messageService{
		r: r,
	}
}

func (s *messageService) Create(author, text string) error {
	if len(author) > 25 {
		return errors.New("username is too long")
	}

	if len(author) < 2 {
		return errors.New("username is too short")
	}

	if len(text) > 70 {
		return errors.New("text is too long")
	}

	var (
		usernamePattern = regexp.MustCompile("^[А-Яа-яA-Za-z0-9]+$")
		textPattern     = regexp.MustCompile("^[А-Яа-яA-Za-z\\s.,;!?0-9]+$")
	)

	// if username contains special characters
	if !usernamePattern.MatchString(author) {
		return errors.New("username mustn't contain special characters")
	}

	// if text contains special characters
	if !textPattern.MatchString(text) {
		return errors.New("text contains invalid characters")
	}

	err := s.r.Create(author, text)
	if err != nil {
		return errors.New("failed to create")
	}

	return nil
}

func (s *messageService) Get(limit ...int) ([]entity.Message, error) {
	var LIMIT = 50
	if len(limit) > 0 {
		LIMIT = limit[0]
	}

	return s.r.Get(LIMIT)
}

func (s *messageService) Delete(id int) error {
	return s.r.Delete(id)
}
