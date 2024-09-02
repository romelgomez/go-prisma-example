package model

import "time"

type Account struct {
	Id       string
	Name     *string
	Disabled bool
	OwnerId  *string
	Slug     *string
	Tag      *string
	Created  time.Time
	Modified time.Time
	Deleted  *time.Time
}
