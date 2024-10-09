package controller

import (
	"time"
	"udemy_slack_app/model"
)

// channelリクエスト構造体
type ChannelRequest struct {
	ChannelName  string `json:"channel_name" validate:"required"`
	CreateUserID int    `json:"create_user_id" validate:"required"`
}

// モデル構造体にバインドし直す
func ChannelToModel(req ChannelRequest) *model.Channel {
	now := time.Now()
	return &model.Channel{
		ChannelName:  req.ChannelName,
		CreateUserID: req.CreateUserID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
