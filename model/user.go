package model

import "time"

type User struct {
	Id             string
	FirstName      *string
	LastName       *string
	Email          string
	SecondName     *string
	SecondLastName *string
	Document       *string
	DocumentType   *string
	Nationality    *string
	Birthday       *time.Time
	Tag            *string
	Created        time.Time
	Modified       time.Time
	Deleted        *time.Time
}
