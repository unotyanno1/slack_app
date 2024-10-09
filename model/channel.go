package model

import "time"

type Channel struct {
	ChannelName  string    `json:"channel_name"`
	CreateUserID int       `json:"create_user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
