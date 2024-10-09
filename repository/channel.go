package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"udemy_slack_app/model"
)

type ChannelRepository interface {
	Create(ctx context.Context, channel *model.Channel) (string, error)
	Read(ctx context.Context, channelID string) (*model.Channel, error)
	ReadAll(ctx context.Context) ([]model.Channel, error)
	Update(ctx context.Context, channel *model.Channel, channelID string) error
	Delete(ctx context.Context, channelID string) error
}

type channelRepository struct {
	db *sql.DB
}

func NewChannelRepository(db *sql.DB) ChannelRepository {
	return &channelRepository{db}
}

func (r *channelRepository) Create(ctx context.Context, channel *model.Channel) (string, error) {
	result, err := r.db.Exec("INSERT INTO channel (channel_name, create_user_id, created_at) VALUES (?, ?, ?)", channel.ChannelName, channel.CreateUserID, channel.CreatedAt)

	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	idStr := strconv.FormatInt(id, 10)
	return idStr, nil
}

func (r *channelRepository) Read(ctx context.Context, channelID string) (*model.Channel, error) {
	var channel model.Channel
	err := r.db.QueryRow("SELECT * FROM channel WHERE id = ?", channelID).Scan(&channel.ChannelName, &channel.CreateUserID)
	if err != nil {
		return nil, err
	}

	return &channel, nil
}

func (r *channelRepository) ReadAll(ctx context.Context) ([]model.Channel, error) {
	channels := []model.Channel{}
	rows, err := r.db.Query("SELECT channel_name, create_user_id FROM channel ORDER BY id DESC")
	if err != nil {
		return channels, err
	}

	for rows.Next() {
		var channel model.Channel
		if err := rows.Scan(&channel.ChannelName, &channel.CreateUserID); err != nil {
			return channels, err
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

func (r *channelRepository) Update(ctx context.Context, channel *model.Channel, channelID string) error {
	result, err := r.db.Exec("UPDATE channel SET channel_name = ?, updated_at = ? WHERE id = ?", channel.ChannelName, channel.UpdatedAt, channelID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", channelID)
	}

	return nil
}

func (r *channelRepository) Delete(ctx context.Context, channelID string) error {
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("transaction not found")
	}

	result, err := db.Exec("DELETE FROM channel WHERE id = ?", channelID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %s", channelID)
	}

	return nil
}
