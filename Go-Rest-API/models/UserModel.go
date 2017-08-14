package models

import "time"

type UserModel struct {
	UserId      int    `json:"userId"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AuthToken   string `json:"authToken"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *UserModel) SetUserId(uid int) {
	u.UserId = uid
}

func (u *UserModel) GetUserId() int {
	return u.UserId
}

func (u *UserModel) SetPassword(password string) {
	u.Password = password
}

func (u UserModel) GetPassword() string {
	return u.Password
}

func (u *UserModel) SetOldPassword(oldpassword string) {
	u.OldPassword = oldpassword
}

func (u UserModel) GetOldPassword() string {
	return u.OldPassword
}

func (u *UserModel) SetNewPassword(newpassword string) {
	u.NewPassword = newpassword
}

func (u UserModel) GetNewPassword() string {
	return u.NewPassword
}

func (u *UserModel) SetAuthToken(authToken string) {
	u.AuthToken = authToken
}

func (u UserModel) GetAuthToken() string {
	return u.AuthToken
}

func (u *UserModel) SetEmail(Email string) {
	u.Email = Email
}

func (u UserModel) GetEmail() string {
	return u.Email
}
func (u *UserModel) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u UserModel) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u *UserModel) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u UserModel) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
