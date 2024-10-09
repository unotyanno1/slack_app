package controller

import (
	"time"
	"udemy_slack_app/model"
)

// userリクエスト構造体
type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Email string `json:"email" validate:"required"`
}

// モデル構造体にバインドし直す
func UserToModel(req UserRequest) *model.User {
	now := time.Now()
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
