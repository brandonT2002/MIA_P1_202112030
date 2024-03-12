package structs

import (
	"fmt"
	utils "mia/Classes/Utils"
	"os"
	"strings"
)

type DataBlock struct {
	I             int32
	BlockFile     *BlockFile
	BlockFolder   *BlockFolder
	BlockPointers *BlockPointers
}

type Tree struct {
	SuperBlock *SuperBlock
	File       os.File
	Blocks     []*DataBlock
	FileBlocks []*SuperBlock
}

type BlockFileTuple struct {
	Block int
	File  *BlockFile
}

func NewTree(SuperBlock SuperBlock, file os.File) *Tree {
	return &Tree{SuperBlock: &SuperBlock, File: file}
}

// ============================== GRAPH TREE ================================

func (t *Tree) GetDot(diskname, partName string) string {
	dot := "digraph Tree{\n\tnode [shape=plaintext];\n\trankdir=LR;\n\t"
	dot += fmt.Sprintf("label=\"%v: %v\";\n\tlabelloc=t;\n\t", diskname, partName)
	dot += t.getDotInode(0)
	dot += "\n}"
	return dot
}

func (t *Tree) getDotInode(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i*88), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero:", err)
	}
	readedBytes := make([]byte, 88)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	inode := DecodeInode(readedBytes)
	dot := inode.getDot(i)
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					dot += "\n\t" + t.getDotBlockFolder(inode.Block[p])
				} else {
					dot += "\n\t" + t.getDotBlockFile(inode.Block[p])
				}
			} else if p == 12 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 1)
			} else if p == 13 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 2)
			} else if p == 14 {
				dot += "\n\t" + t.getDotBlockPointers(inode.Block[p], inode.Type, 3)
			}
			dot += fmt.Sprintf("\n\tinode%v:A%v -> block%v:B%v;", i, p, inode.Block[p], inode.Block[p])
		}
	}
	return dot
}

func (t *Tree) getDotBlockPointers(i int32, inodeType rune, simplicity int) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockPointers := DecodeBlockPointers(readedBytes)
	dot := blockPointers.GetDot(int(i))
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					dot += "\n\t" + t.getDotBlockFolder(int32(blockPointers.Pointers[p]))
				} else {
					dot += "\n\t" + t.getDotBlockFile(int32(blockPointers.Pointers[p]))
				}
			} else {
				dot += "\n\t" + t.getDotBlockPointers(int32(blockPointers.Pointers[p]), inodeType, simplicity-1)
			}
			dot += fmt.Sprintf("\n\tblock%v:A%v -> block%v:B%v;", i, p, blockPointers.Pointers[p], blockPointers.Pointers[p])
		}
	}
	return dot
}

func (t *Tree) getDotBlockFolder(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockFolder := DecodeBlockFolder(readedBytes)
	dot := blockFolder.GetDot(int(i))
	for p := 0; p < len(blockFolder.Content); p++ {
		if strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != "." && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != ".." && blockFolder.Content[p].Inode != -1 {
			dot += "\n\t" + t.getDotInode(blockFolder.Content[p].Inode)
			dot += fmt.Sprintf("\n\tblock%v:A%v -> inode%v:I%v;", i, p, blockFolder.Content[p].Inode, blockFolder.Content[p].Inode)
		}
	}
	return dot
}

func (t *Tree) getDotBlockFile(i int32) string {
	_, err := t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*64, 0)
	if err != nil {
		fmt.Println("Error al posicionar el puntero seek:", err)
	}
	readedBytes := make([]byte, 64)
	_, err = t.File.Read(readedBytes)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("-> Error al leer el archivo:", err)
	}
	blockFile := DecodeBlockFile(readedBytes)
	return blockFile.GetDot(int(i))
}

// =================================== GET BLOCKS ==================================

func (t *Tree) GetBlocks() []*DataBlock {
	t.searchInInodes(0)
	if len(t.Blocks) > 1 {
		for i := 1; i < len(t.Blocks); i++ {
			for j := i; j > 0; j-- {
				if t.Blocks[j].I < t.Blocks[j-1].I {
					t.Blocks[j], t.Blocks[j-1] = t.Blocks[j-1], t.Blocks[j]
					continue
				}
				break
			}
		}
	}
	return t.Blocks
}

func (t *Tree) searchInInodes(i int32) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i*88), 0)
	inodeData := make([]byte, 88)
	_, err := t.File.Read(inodeData)
	if err != nil {
		fmt.Println("Error al leer el inode:", err)
		return
	}
	inode := DecodeInode(inodeData)
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					t.searchInBlockFolder(inode.Block[p])
				} else {
					t.searchInBlockFile(inode.Block[p])
				}
			} else if p == 12 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 1)
			} else if p == 13 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 2)
			} else if p == 14 {
				t.searchInBlockPointers(inode.Block[p], inode.Type, 3)
			}
		}
	}
}

func (t *Tree) searchInBlockPointers(i int32, inodeType rune, simplicity int) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockPointersData := make([]byte, 64)
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
		return
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockPointers: blockPointers})
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					t.searchInBlockFolder(int32(blockPointers.Pointers[p]))
				} else {
					t.searchInBlockFile(int32(blockPointers.Pointers[p]))
				}
			} else {
				t.searchInBlockPointers(int32(blockPointers.Pointers[p]), inodeType, simplicity-1)
			}
		}
	}
}

func (t *Tree) searchInBlockFolder(i int32) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*64), 0)
	blockFolderData := make([]byte, 64)
	_, err := t.File.Read(blockFolderData)
	if err != nil {
		fmt.Println("Error al leer el block folder:", err)
		return
	}
	blockFolder := DecodeBlockFolder(blockFolderData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockFolder: blockFolder})
	for p := 0; p < len(blockFolder.Content); p++ {
		if strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != "." && strings.TrimSpace(strings.ReplaceAll(blockFolder.Content[p].Name, "\x00", "")) != ".." && blockFolder.Content[p].Inode != -1 {
			t.searchInInodes(int32(blockFolder.Content[p].Inode))
		}
	}
}

func (t *Tree) searchInBlockFile(i int32) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockFileData := make([]byte, 64)
	_, err := t.File.Read(blockFileData)
	if err != nil {
		fmt.Println("Error al leer el block file:", err)
		return
	}
	blockFile := DecodeBlockFile(blockFileData)
	t.Blocks = append(t.Blocks, &DataBlock{I: i, BlockFile: blockFile})
}

// ================================== READ CONTENT ==================================

func (t *Tree) ReadFile(path string) (string, bool) {
	dir := strings.Split(path, "/")
	var cleanDir []string
	for _, d := range dir {
		if d != "" {
			cleanDir = append(cleanDir, d)
		}
	}
	return t.readFileInInodes(0, cleanDir)
}

func (t *Tree) readFileInInodes(i int32, path []string) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i)*int64(SizeOfInode()), 0)
	inodeData := make([]byte, SizeOfInode())
	_, err := t.File.Read(inodeData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el inode (READ FILE):", err)
		return "", false
	}
	inode := DecodeInode(inodeData)
	content := ""
	founded := false
	for p := 0; p < len(inode.Block); p++ {
		if inode.Block[p] == -1 {
			break
		}
		if inode.Block[p] != -1 {
			if p < 12 {
				if inode.Type == '0' {
					content, founded = t.readFileInBlockFolder(inode.Block[p], path)
					if founded {
						return content, founded
					}
				} else {
					cont := ""
					cont, founded = t.readFileInBlockFile(inode.Block[p])
					content += cont
				}
			}
		} else if p == 12 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 1)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 1)
				content += cont
			}
		} else if p == 13 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 2)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 2)
				content += cont
			}
		} else if p == 14 {
			if inode.Type == '0' {
				content, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 3)
			} else {
				cont := ""
				cont, founded = t.readFileInBlockPointers(inode.Block[p], path, inode.Type, 3)
				content += cont
			}
		}
	}
	return content, founded
}

func (t *Tree) readFileInBlockPointers(i int32, path []string, inodeType rune, simplicity int) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(SizeOfBlockPointers()), 0)
	blockPointersData := make([]byte, SizeOfBlockPointers())
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block pointer (READ FILE):", err)
		return "", false
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	content := ""
	founded := false
	for p := 0; p < len(blockPointers.Pointers); p++ {
		if blockPointers.Pointers[p] != -1 {
			if simplicity == 1 {
				if inodeType == '0' {
					content, founded = t.readFileInBlockFolder(blockPointers.Pointers[p], path)
					if founded {
						return content, founded
					}
				} else {
					cont := ""
					cont, founded = t.readFileInBlockFile(blockPointers.Pointers[p])
					content += cont
				}
			} else {
				if inodeType == '0' {
					content, founded = t.readFileInBlockPointers(blockPointers.Pointers[p], path, inodeType, simplicity-1)
				} else {
					count := ""
					count, founded = t.readFileInBlockPointers(blockPointers.Pointers[p], path, inodeType, simplicity-1)
					content += count
				}
			}
		} else {
			break
		}
	}
	return content, founded
}

func (t *Tree) readFileInBlockFolder(i int32, path []string) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i*64), 0)
	blockFolderData := make([]byte, 64)
	_, err := t.File.Read(blockFolderData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block folder (READ FILE):", err)
		return "", false
	}
	blockFolder := DecodeBlockFolder(blockFolderData)
	for p := 0; p < len(blockFolder.Content); p++ {
		if !(utils.CompareStrings(blockFolder.Content[p].Name, ".") || utils.CompareStrings(blockFolder.Content[p].Name, "..")) && blockFolder.Content[p].Inode != -1 && utils.CompareStrings(blockFolder.Content[p].Name, path[0]) {
			path = path[1:]
			return t.readFileInInodes(blockFolder.Content[p].Inode, path)
		}
	}
	return "", false
}

func (t *Tree) readFileInBlockFile(i int32) (string, bool) {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(i)*int64(64), 0)
	blockFileData := make([]byte, 64)
	_, err := t.File.Read(blockFileData)
	if err != nil {
		if err.Error() == "EOF" {
			return "", false
		}
		fmt.Println("Error al leer el block file (READ FILE):", err)
		return "", false
	}
	blockFile := DecodeBlockFile(blockFileData)
	return blockFile.Content, true
}

// ================================= WRITE FILE ===================================

func (t *Tree) writeFile(path, diskpath string, partstart int, newContent string) {
	dir := strings.Split(path, "/")
	var cleanDir []string
	for _, d := range dir {
		if d != "" {
			cleanDir = append(cleanDir, d)
		}
	}
	t.writeFileInInodes(0, cleanDir, diskpath, newContent, partstart)
	t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
	t.SuperBlock.First_ino = t.findNextFreeInode(1)[0]
}

func (t *Tree) writeFileInInodes(i int32, path []string, pathdsk, newContent string, partstart int) {
	t.File.Seek(int64(t.SuperBlock.Inode_start)+int64(i)*int64(88), 0)
	inodeData := make([]byte, 88)
	_, err := t.File.Read(inodeData)
	if err != nil {
		fmt.Println("Error al leer el inode:", err)
	}
	inode := DecodeInode(inodeData)
	if inode.Type == '0' {
		for p := 0; p < len(inode.Block); p++ {
			if inode.Block[p] != -1 {
				writed := t.writeFileInBlockFolder(inode.Block[p], path, pathdsk, newContent, int32(partstart))
				if writed {
					return
				}
			}
		}
	} else {
		blocksFile := &BlockFileTuple{Block: -1, File: nil}
		for p := 0; p < len(inode.Block); p++ {
			if inode.Block[p] != -1 {
				if p < 12 {
					blocksFile = t.writeFileInBlockFile(inode.Block[p])
				} else if p == 12 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 1)
				} else if p == 13 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 2)
				} else if p == 14 {
					blocksFile = t.writeFileInBlockPointers3(pathdsk, inode.Block[p], 3)
				}
			}
		}
		num, block := blocksFile.Block, blocksFile.File
		if block != nil {
			block = &BlockFile{Content: ""}
			num = t.findNextFreeBlock(1)[0]
			inode.Block[0] = int32(num)
			t.rewriteInode(pathdsk, i, inode)
		}
		contents := [][]rune{{}}
		for _, z := range newContent {
			if len(contents[len(contents)-1]) < 64 {
				contents[len(contents)-1] = append(contents[len(contents)-1], rune(z))
			} else {
				contents = append(contents, []rune{z})
			}
		}
		blockContent := contents[0]
		t.writeInDisk(pathdsk, int32(t.SuperBlock.Block_start+num*64), []byte(string(blockContent)))
		newSizeInode := (num-1)*64 + len(blockContent)
		for len(contents) > 0 {
			newBlock := BlockFile{Content: ""}
			contenew := contents[0]
			newSizeInode = (num-1)*64 + len(contenew)
			if len(newBlock.Content) == len(contenew) {
				newBlock.Content = string(contenew)
			} else {
				for z := 0; z < len(contenew); z++ {
					contentRunes := []rune(newBlock.Content)
					contentRunes[z] = contenew[z]
					newBlock.Content = string(contentRunes)
				}
			}
			nextFreeBitBlock := t.findNextFreeBlock(1)
			for h := 0; h < len(inode.Block); h++ {
				if inode.Block[h] == -1 {
					if h < 12 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeInDisk(pathdsk, int32(t.SuperBlock.Bm_block_start+nextFreeBitBlock[0]), []byte{1})
						t.writeInDisk(pathdsk, int32(t.SuperBlock.Block_start+nextFreeBitBlock[0]*64), newBlock.Encode())
						t.SuperBlock.First_blo = t.findNextFreeBlock(1)[0]
						t.SuperBlock.Free_blocks_count -= 1
					} else if h == 12 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 1)
						t.SuperBlock.Free_blocks_count -= 1
						t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					} else if h == 13 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 2)
						t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					} else if h == 14 {
						inode.Block[h] = int32(nextFreeBitBlock[0])
						t.writeFileInBlockPointers(pathdsk, newBlock, int32(nextFreeBitBlock[0]), 3)
						t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					}
					break
				} else if h == 12 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 1) {
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 1)
					t.SuperBlock.Free_blocks_count -= 1
					t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				} else if h == 13 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 2) {
					posiblePointer := t.searchPointer(inode.Block[h], 2)
					if posiblePointer[0] != -1 {
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock, int32(posiblePointer[1]))
						break
					}
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 2)
					t.SuperBlock.Free_blocks_count -= 1
					t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				} else if h == 14 && inode.Block[h] != -1 && t.validateSpacePointers(inode.Block[h], 3) {
					posiblePointer := t.searchPointer(inode.Block[h], 3)
					if posiblePointer[0] != -1 {
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock, int32(posiblePointer[1]))
						break
					}
					posiblePointer = t.searchPointer(inode.Block[h], 2)
					if posiblePointer[0] != -1 {
						// t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), BlockPointers{Pointers: []int32{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}}, int32(posiblePointer[1]))
						posiblePointer = t.searchPointer(inode.Block[h], 3)
						t.writeNewBlockInIndirect(pathdsk, int32(posiblePointer[0]), newBlock, int32(posiblePointer[1]))
						break
					}
					t.writeFileInBlockPointers(pathdsk, newBlock, inode.Block[h], 3)
					t.SuperBlock.Free_blocks_count -= 1
					t.writeInDisk(pathdsk, int32(partstart), t.SuperBlock.Encode())
					break
				}
			}
			inode.Size = newSizeInode
			t.writeInDisk(pathdsk, int32(t.SuperBlock.Inode_start)+i*64, inode.Encode())
		}
	}
}

func (t *Tree) validateSpacePointers(numBlock, simplicity int32) bool {
	t.File.Seek(int64(t.SuperBlock.Block_start)+int64(numBlock)*int64(64), 0)
	blockPointersData := make([]byte, 64)
	_, err := t.File.Read(blockPointersData)
	if err != nil {
		fmt.Println("Error al leer el block pointer:", err)
	}
	blockPointers := DecodeBlockPointers(blockPointersData)
	resultado := false
	for i := 0; i < len(blockPointers.Pointers); i++ {
		if simplicity == 1 {
			if blockPointers.Pointers[i] == -1 {
				return true
			} else {
				if blockPointers.Pointers[i] != -1 {
					resultado = t.validateSpacePointers(blockPointers.Pointers[i], simplicity-1)
				} else {
					return true
				}
			}
		}
	}
	return resultado
}

func (t *Tree) writeFileInBlockPointers(pathdsk string, newBlockFile BlockFile, nextFreeBitBlock, simplicity int32) {
}
func (t *Tree) writeFileInBlockPointers2(pathdsk string, newBlockFile BlockFile, nextFreeBitBlock, simplicity int32) {
}
func (t *Tree) writeNewBlockInIndirect(pathdsk string, numBlock int32, newBlockFile BlockFile, numPtr int32) {
}
func (t *Tree) searchPointer(numBlock, simplicity int32) []int { return []int{} }
func (t *Tree) writeFileInBlockPointers3(pathdsk string, i, simplicity int32) *BlockFileTuple {
	return nil
}
func (t *Tree) writeFileInBlockFolder(i int32, path []string, pathdsk, content string, partstart int32) bool {
	return true
}
func (t *Tree) writeFileInBlockFile(i int32) *BlockFileTuple { return nil }

// =========================================== MKDIR AND MKFILE =============================================

func (t *Tree) mkdir()                                                    {}
func (t *Tree) mkfile()                                                   {}
func (t *Tree) mkdirInInode()                                             {}
func (t *Tree) mkdirInBlockFolder()                                       {}
func (t *Tree) searchPointerInodeFolder()                                 {}
func (t *Tree) searchPointerFolder()                                      {}
func (t *Tree) searchPointerBlockFolder()                                 {}
func (t *Tree) writeNewBlock()                                            {}
func (t *Tree) writeNewInode()                                            {}
func (t *Tree) rewriteBlock()                                             {}
func (t *Tree) rewriteInode(pathdsk string, numInode int32, inode *Inode) {}

// =========================================== SEARCH DIRECTORY =============================================

func (t *Tree) searchdir()                {}
func (t *Tree) searchdirInInode()         {}
func (t *Tree) searchdirInBlockPointers() {}
func (t *Tree) searchdirInBlockFolder()   {}

// ========================================= FIND NEXT BIT =============================================

func (t *Tree) findNextFreeInode(count int) []int { return []int{} }
func (t *Tree) findNextFreeBlock(count int) []int { return []int{} }

// ========================================= USERS AND GROUPS ===============================================

func (t *Tree) getUsers()  {}
func (t *Tree) getGroups() {}

// ========================================== WRITE IN DISK =================================================

func (t *Tree) writeInDisk(path string, seek int32, content []byte) {}
