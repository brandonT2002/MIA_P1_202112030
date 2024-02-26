package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
	"unicode/utf8"
)

type Inode struct {
	Uid   int       // 4 bytes
	Gid   int       // 4 bytes
	Size  int       // 4 bytes
	Atime time.Time // 4 bytes
	Ctime time.Time // 4 bytes
	Mtime time.Time // 4 bytes
	Block []int     // 60 bytes
	Type  rune      // 1 byte
	Perm  []rune    // 3 bytes
} // 88 bytes

func (i *Inode) Encode() []byte {
	var buf bytes.Buffer
	// Uid
	uidBytes := make([]byte, 4)
	if i.Uid != 0 {
		binary.BigEndian.PutUint32(uidBytes, uint32(i.Uid))
	}
	buf.Write(uidBytes)
	// Gid
	gidBytes := make([]byte, 4)
	if i.Gid != 0 {
		binary.BigEndian.PutUint32(gidBytes, uint32(i.Gid))
	}
	buf.Write(gidBytes)
	// Size
	sizeBytes := make([]byte, 4)
	if i.Size != 0 {
		binary.BigEndian.PutUint32(sizeBytes, uint32(i.Size))
	}
	buf.Write(sizeBytes)
	// Atime
	atimeBytes := make([]byte, 4)
	if !i.Atime.IsZero() {
		binary.BigEndian.PutUint32(atimeBytes, uint32(i.Atime.Unix()))
	}
	buf.Write(atimeBytes)
	// Ctime
	ctimeBytes := make([]byte, 4)
	if !i.Ctime.IsZero() {
		binary.BigEndian.PutUint32(ctimeBytes, uint32(i.Ctime.Unix()))
	}
	buf.Write(ctimeBytes)
	// Mtime
	mtimeBytes := make([]byte, 4)
	if !i.Mtime.IsZero() {
		binary.BigEndian.PutUint32(mtimeBytes, uint32(i.Mtime.Unix()))
	}
	buf.Write(mtimeBytes)
	// Block
	for _, block := range i.Block {
		blockBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(blockBytes, uint32(block))
		buf.Write(blockBytes)
	}
	// Type
	buf.WriteByte(byte(i.Type))
	// Perm
	for _, p := range i.Perm {
		buf.WriteByte(byte(p))
	}
	return buf.Bytes()
}

func DecodeInode(data []byte) *Inode {
	// Uid
	uid := int(binary.BigEndian.Uint32(data[:4]))
	// Gid
	gid := int(binary.BigEndian.Uint32(data[4:8]))
	// Size
	size := int(binary.BigEndian.Uint32(data[8:12]))
	// Atime
	unixTime := binary.BigEndian.Uint32(data[12:16])
	atime := time.Unix(int64(unixTime), 0)
	// Ctime
	unixTime = binary.BigEndian.Uint32(data[16:20])
	ctime := time.Unix(int64(unixTime), 0)
	// Mtime
	unixTime = binary.BigEndian.Uint32(data[20:24])
	mtime := time.Unix(int64(unixTime), 0)
	// Block
	block := []int{}
	for i := 0; i < 15; i++ {
		block = append(block, int(binary.BigEndian.Uint32(data[24+i*4:28+i*4])))
	}
	// Type
	type_, _ := utf8.DecodeRune(data[84:85])
	// Perm
	p, _ := utf8.DecodeRune(data[85:86])
	perm := []rune{p}
	p, _ = utf8.DecodeRune(data[86:87])
	perm = append(perm, p)
	p, _ = utf8.DecodeRune(data[87:])
	perm = append(perm, p)
	return &Inode{uid, gid, size, atime, ctime, mtime, block, type_, perm}
}

func (i *Inode) ToString() string {
	return fmt.Sprintf(`uid: %v
gid: %v
size: %v
atime: %v
ctime: %v
mtime: %v
block: %v
type: %v
perm: %v`,
		i.Uid,
		i.Gid,
		i.Size,
		i.Atime,
		i.Ctime,
		i.Mtime,
		i.Block,
		i.Type,
		i.Perm,
	)
}
