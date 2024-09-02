package model

import "time"

type PublicationMedia struct {
	Id            string
	PublicationId string
	MediaId       string
	Tag           *string
	Created       time.Time
	Modified      time.Time
	Deleted       *time.Time
}
