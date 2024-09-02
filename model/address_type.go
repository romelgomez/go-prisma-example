package model

type AddressType string

const (
	Residential AddressType = "RESIDENTIAL"
	Business    AddressType = "BUSINESS"
	Shipping    AddressType = "SHIPPING"
	Billing     AddressType = "BILLING"
	Other       AddressType = "OTHER"
)
