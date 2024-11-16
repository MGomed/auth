package converters

import (
	"encoding/json"
	"fmt"

	service_model "github.com/MGomed/auth/internal/model"
	msg_bus_errors "github.com/MGomed/auth/internal/storage/message_bus/errors"
	msg_bus_model "github.com/MGomed/auth/internal/storage/message_bus/model"
)

// ToMessageTypeFromUserInfo converts service_modelUserinfo to msg_bus_model.Message
func ToMessageTypeFromUserInfo(user *service_model.UserInfo) (*msg_bus_model.Message, error) {
	if user == nil {
		return nil, nil
	}

	info := &msg_bus_model.UserInfo{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("%w: %w", msg_bus_errors.ErrMarshal, err)
	}

	return &msg_bus_model.Message{
		Type: msg_bus_model.CreateType,
		Data: data,
	}, nil
}

// ToMessageTypeFromID converts id to msg_bus_model.Message
func ToMessageTypeFromID(id int64) (*msg_bus_model.Message, error) {
	type ID struct {
		ID int64 `json:"id"`
	}

	data, err := json.MarshalIndent(&ID{
		ID: id,
	}, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("%w: %w", msg_bus_errors.ErrMarshal, err)
	}

	return &msg_bus_model.Message{
		Type: msg_bus_model.CreateType,
		Data: data,
	}, nil
}
