package model

import "time"

type Publication struct {
	Id          string
	Title       string
	Description string
	Visibility  PublicationVisibility
	ListingId   *string
	Tag         *string
	Created     time.Time
	Modified    time.Time
	Deleted     *time.Time
}
