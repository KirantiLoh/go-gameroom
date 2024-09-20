package users

import "github.com/kirantiloh/gameroom/pkg/auth"

type User struct {
	auth.UserData
	Password   string
  IsVerified bool `json:"is_verified"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
  LoginDto
  Name string `json:"name"`
}

func (u *User) toUserData() *auth.UserData {
	return &u.UserData
}
