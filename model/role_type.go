package model

type RoleType string

const (
	Owner   RoleType = "OWNER"
	Admin   RoleType = "ADMIN"
	Manager RoleType = "MANAGER"
	Viewer  RoleType = "VIEWER"
)
