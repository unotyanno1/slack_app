package controller

import (
	"time"
	"udemy_slack_app/model"
)

// messageリクエスト構造体
type MessageRequest struct {
	ChannelID int    `json:"channel_id" validate:"required"`
	UserID    int    `json:"user_id" validate:"required"`
	Message   string `json:"message" validate:"required"`
}

// モデル構造体にバインドし直す
func MessageToModel(req MessageRequest) *model.Message {
	now := time.Now()
	return &model.Message{
		ChannelID: req.ChannelID,
		UserID:    req.UserID,
		Message:   req.Message,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
