package models

type User struct {
	Base
	AvatarURL string
	Email     string
	Username  string
}
