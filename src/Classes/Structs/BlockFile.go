package structs

import (
	"bytes"
	"fmt"
)

type BlockFile struct {
	Content string // 64 bytes
}

func (b *BlockFile) Encode() []byte {
	var buf bytes.Buffer
	// Content
	if b.Content != "" {
		if len(b.Content) < 64 {
			buf.Write([]byte(b.Content)[:len(b.Content)])
			buf.Write(bytes.Repeat([]byte{0x00}, 64-len(b.Content)))
		} else {
			buf.Write([]byte(b.Content)[:64])
		}
	} else {
		buf.Write(make([]byte, 64))
	}
	return buf.Bytes()
}

func DecodeBlockFile(data []byte) *BlockFile {
	return &BlockFile{string(data[:64])}
}

func (b *BlockFile) ToString() string {
	return fmt.Sprintf("content: %s\n", b.Content)
}
