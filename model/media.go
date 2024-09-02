package model

import "time"

type Media struct {
	Id      string
	Type    string
	Size    int
	Version int
	Name    string
	Tag     *string
	Created time.Time
	Deleted *time.Time
}
