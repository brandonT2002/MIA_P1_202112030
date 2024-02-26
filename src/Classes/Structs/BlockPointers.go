package structs

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BlockPointers struct {
	Pointers []int // 64 bytes
}

func (b *BlockPointers) Encode() []byte {
	var buf bytes.Buffer
	// Pointers
	for _, pointer := range b.Pointers {
		pointerBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(pointerBytes, uint32(pointer))
		buf.Write(pointerBytes)
	}
	return buf.Bytes()
}

func DecodeBlockPointers(data []byte) *BlockPointers {
	// Pointers
	pointers := []int{}
	for i := 0; i < 16; i++ {
		pointers = append(pointers, int(binary.BigEndian.Uint32(data[i*4:4+i*4])))
	}
	return &BlockPointers{pointers}
}

func (b *BlockPointers) ToString() string {
	return fmt.Sprintf("pointers: %v\n", b.Pointers)
}
