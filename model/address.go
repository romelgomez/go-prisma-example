package model

import (
	"time"
)

type Address struct {
	Id       string
	Location string
	District string
	Type     AddressType
	Tag      *string
	Created  time.Time
	Modified time.Time
	Deleted  *time.Time
}
