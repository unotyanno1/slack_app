package model

import "time"

type Message struct {
	ChannelID int       `json:"channel_id"`
	UserID    int       `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
