package models

type Status string

const (
	Active   Status = "active"
	InActive Status = "inactive"
	Deleted  Status = "deleted"
	Locked   Status = "locked"
)

type UserType string

const (
	Buyer    UserType = "buyer"
	Supplier UserType = "supplier"
)
