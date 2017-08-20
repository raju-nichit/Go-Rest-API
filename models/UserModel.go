package models

import "time"

//UserModel --
type UserModel struct {
	UserID      int    `json:"userId"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AuthToken   string `json:"authToken"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
