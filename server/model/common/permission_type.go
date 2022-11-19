package common

type PermissionType int

const (
	ADD PermissionType = iota
	DELETE
	UPDATE
	GET
)
