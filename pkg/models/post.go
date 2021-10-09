package models

import (
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Caption   string    `json:"caption,omitempty"`
	ImgUrl    string    `json:"img_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
