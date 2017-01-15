package client

import (
	"fmt"
	"time"
)

type Consumer struct {
	ID           string            `json:"id"`
	App          string            `json:"app"`
	Roles        []string          `json:"roles" bson:"roles"`
	Username     string            `json:"username" bson:"username"`
	CustomID     string            `json:"custom_id" bson:"custom_id"`
	CustomFields map[string]string `json:"custom_fields" bson:"custom_fields"`
	UpdatedAt    time.Time         `json:"updated_at" bson:"updated_at"`
	CreatedAt    time.Time         `json:"created_at" bson:"created_at"`
}

type Token struct {
	ID         string    `json:"id" bson:"_id"`
	Source     string    `json:"source" bson:"source"`
	ConsumerID string    `json:"consumer_id" bson:"consumer_id"`
	IPAddress  string    `json:"ip_address" bson:"ip_address"`
	ExpiresIn  int64     `json:"expires_in" bson:"-"`
	Expiration time.Time `json:"Expiration" bson:"expiration"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

type AppError struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%d - %s", e.StatusCode, e.Message)
}
