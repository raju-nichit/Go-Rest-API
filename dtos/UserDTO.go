package dtos

import "time"

type UserDTO struct {
	UserId      int    `json:"userId,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	AuthToken   string `json:"authtoken,omitempty"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *UserDTO) SetUserId(uid int) {
	u.UserId = uid
}

func (u UserDTO) GetUserId() int {
	return u.UserId
}

func (u *UserDTO) SetPassword(password string) {
	u.Password = password
}

func (u UserDTO) GetPassword() string {
	return u.Password
}
func (u *UserDTO) SetOldPassword(oldpassword string) {
	u.OldPassword = oldpassword
}

func (u UserDTO) GetOldPassword() string {
	return u.OldPassword
}

func (u *UserDTO) SetNewPassword(newpassword string) {
	u.NewPassword = newpassword
}

func (u UserDTO) GetNewPassword() string {
	return u.NewPassword
}

func (u *UserDTO) SetAuthToken(authToken string) {
	u.AuthToken = authToken
}

func (u UserDTO) GetAuthToken() string {
	return u.AuthToken
}

func (u *UserDTO) SetEmail(Email string) {
	u.Email = Email
}

func (u UserDTO) GetEmail() string {
	return u.Email
}

func (u *UserDTO) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u UserDTO) GetCreatedAt() time.Time {
	return u.CreatedAt
}
func (u *UserDTO) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}

func (u UserDTO) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
