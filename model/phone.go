package model

import "time"

type Phone struct {
	Id       string
	Number   string
	Type     PhoneType
	Tag      *string
	Created  time.Time
	Modified time.Time
	Deleted  *time.Time
}
