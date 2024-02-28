package utils

type Type int

const (
	EXECUTE Type = iota
	MKDISK
	RMDISK
	FDISK
	REP
	COMMENTARY
)
