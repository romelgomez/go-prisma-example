package model

import "time"

type Role struct {
	Id        string
	Assigned  time.Time
	Role      RoleType
	UserId    *string
	AccountId *string
	Tag       *string
	Created   time.Time
	Modified  time.Time
	Deleted   *time.Time
}
