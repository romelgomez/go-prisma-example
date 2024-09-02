package model

import "time"

type Listing struct {
	Id          string
	Name        *string
	Visibility  ListingVisibility
	AccountId   *string
	Description *string
	Tag         *string
	Created     time.Time
	Modified    time.Time
	Deleted     *time.Time
}
