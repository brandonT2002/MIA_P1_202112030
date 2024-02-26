package structs

import "bytes"

type BlockFolder struct {
	Content []*Content // 64 bytes
}

func (b *BlockFolder) Encode() []byte {
	var buf bytes.Buffer
	// Content
	for _, c := range b.Content {
		buf.Write(c.Encode())
	}
	return buf.Bytes()
}

func Decode(data []byte) *BlockFolder {
	content := []*Content{}
	for i := 0; i < 4; i++ {
		content = append(content, DecodeContent(data[i*16:16+i*16]))
	}
	return &BlockFolder{content}
}

func (b *BlockFolder) ToString() string {
	contents := ""
	for _, c := range b.Content {
		contents += c.ToString()
		if contents != "" {
			contents += "\n"
		}
	}
	return contents
}
