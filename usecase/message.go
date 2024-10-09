package usecase

import (
	"context"
	"udemy_slack_app/model"
	"udemy_slack_app/repository"
)

type MessageUsecase interface {
	Create(ctx context.Context, message *model.Message) (string, error)
	GetAll(ctx context.Context) ([]model.Message, error)
	Update(ctx context.Context, message *model.Message, messageID string) error
}

type messageUsecase struct {
	mr repository.MessageRepository
}

func NewMessageUsecase(mr repository.MessageRepository) MessageUsecase {
	return &messageUsecase{mr}
}

func (mu *messageUsecase) Create(ctx context.Context, message *model.Message) (string, error) {
	id, err := mu.mr.Create(ctx, message)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (mu *messageUsecase) GetAll(ctx context.Context) ([]model.Message, error) {
	messages, err := mu.mr.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (mu *messageUsecase) Update(ctx context.Context, message *model.Message, messageID string) error {
	err := mu.mr.Update(ctx, message, messageID)
	if err != nil {
		return err
	}

	return nil
}
