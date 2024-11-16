package model

import (
	"time"

	consts "github.com/MGomed/auth/consts"
)

// MessageType declare type of message
type MessageType string

// MessageType constants
const (
	CreateType MessageType = "create"
	DeleteType MessageType = "delete"
)

var MessageTypeToTopic = map[MessageType]string{
	CreateType: consts.CreateTopic,
	DeleteType: consts.DeleteTopic,
}

// Message is MessageBus' message
type Message struct {
	Type MessageType
	Data []byte
}

// UserInfo fro MessageBus
type UserInfo struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
}
