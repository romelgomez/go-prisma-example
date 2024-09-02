package model

import "time"

type UserMedia struct {
	Id       string
	UserId   string
	MediaId  string
	Tag      *string
	Created  time.Time
	Modified time.Time
	Deleted  *time.Time
}
