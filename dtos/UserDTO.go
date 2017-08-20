package dtos

import "time"

type User struct {
	UserId      int    `json:"userId,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	AuthToken   string `json:"authtoken,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
