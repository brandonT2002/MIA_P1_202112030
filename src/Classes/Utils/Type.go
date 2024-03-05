package utils

type Type int

const (
	EXECUTE Type = iota
	MKDISK
	RMDISK
	FDISK
	MOUNT
	UNMOUNT
	MKFS
	LOGIN
	LOGOUT
	PAUSE
	REP
	COMMENTARY
)
