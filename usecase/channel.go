package usecase

import (
	"context"
	"fmt"
	"strconv"
	"udemy_slack_app/model"
	"udemy_slack_app/repository"
	"udemy_slack_app/transaction"
)

type ChannelUsecase interface {
	Create(ctx context.Context, channel *model.Channel) (string, error)
	GetByID(ctx context.Context, channelID string) (*model.Channel, error)
	GetAll(ctx context.Context) ([]model.Channel, error)
	Update(ctx context.Context, channel *model.Channel, channelID string) error
	Delete(ctx context.Context, channelID string) error
}

type channelUsecase struct {
	ur          repository.UserRepository
	cr          repository.ChannelRepository
	mr          repository.MessageRepository
	transaction transaction.Transaction
}

func NewChannelUsecase(ur repository.UserRepository, cr repository.ChannelRepository, mr repository.MessageRepository, transaction transaction.Transaction) ChannelUsecase {
	return &channelUsecase{ur, cr, mr, transaction}
}

func (cu *channelUsecase) Create(ctx context.Context, channel *model.Channel) (string, error) {
	// ユーザー存在確認
	userID := strconv.Itoa(channel.CreateUserID)
	user, err := cu.ur.Read(ctx, userID)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("no exist user: %s", userID)
	}

	// チャンネル作成
	createChannelID, err := cu.cr.Create(ctx, channel)
	if err != nil {
		return "", err
	}

	return createChannelID, nil
}

func (cu *channelUsecase) GetByID(ctx context.Context, channelID string) (*model.Channel, error) {
	channel, err := cu.cr.Read(ctx, channelID)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (cu *channelUsecase) GetAll(ctx context.Context) ([]model.Channel, error) {
	channels, err := cu.cr.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (cu *channelUsecase) Update(ctx context.Context, channel *model.Channel, channelID string) error {
	err := cu.cr.Update(ctx, channel, channelID)
	if err != nil {
		return err
	}

	return nil
}

func (cu *channelUsecase) Delete(ctx context.Context, channelID string) error {
	// トランザクション実行
	cu.transaction.DoInTx(ctx, func(context.Context) (any, error) {
		// チャンネル削除
		err := cu.cr.Delete(ctx, channelID)
		if err != nil {
			return nil, err
		}

		// チャンネルに紐づくメッセージを全て削除
		err = cu.mr.DeleteAll(ctx, channelID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	return nil
}
