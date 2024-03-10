package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logout struct {
	Result string
	Type   utils.Type
	Line   int
	Column int
}

func NewLogout(line, column int) *Logout {
	return &Logout{Type: utils.LOGOUT, Line: line, Column: column}
}

func (l *Logout) GetLine() int {
	return l.Line
}

func (l *Logout) GetColumn() int {
	return l.Column
}

func (l *Logout) GetType() utils.Type {
	return l.Type
}

func (l *Logout) Exec() {
	if env.CurrentLogged.User != nil {
		driveletter := env.CurrentLogged.Driveletter
		absolutePath, _ := filepath.Abs(fmt.Sprintf("/home/jefferson/Escritorio/MIA/P1/%s.dsk", *driveletter))
		file, err := os.Open(absolutePath)
		if err != nil {
			fmt.Println("Error al abrir el archivo:", err)
			return
		}
		defer file.Close()
		readedBytes := make([]byte, 153)
		_, err = file.Read(readedBytes)
		if err != nil {
			fmt.Println("-> Error al leer el archivo:", err)
			return
		}
		mbr := structs.DecodeMBR(readedBytes)
		for i := 0; i < len(mbr.Partitions); i++ {
			if mbr.Partitions[i].Size != 0 && utils.CompareStrings(mbr.Partitions[i].Name, *env.CurrentLogged.Partition) {
				_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
				if err != nil {
					fmt.Println("Error al mover el puntero del archivo:", err)
					return
				}
				superBlockData := make([]byte, 68)
				_, err = file.Read(superBlockData)
				if err != nil {
					fmt.Println("Error al leer el super block:", err)
					return
				}
				superBlock := structs.DecodeSuperBlock(superBlockData)
				if superBlock.Filesystem_type == 3 {
					_, err = file.Seek(int64(mbr.Partitions[i].Start+68), 0)
					if err != nil {
						fmt.Println("Error al mover el puntero del archivo:", err)
						return
					}
					for r := 0; r < superBlock.Inodes_count; r++ {
						journalData := make([]byte, 212)
						_, err = file.Read(journalData)
						if err != nil {
							fmt.Println("Error al leer el super block:", err)
							return
						}
						if string(journalData) == strings.Repeat("\x00", 212) {
							var file1, err = os.OpenFile(absolutePath, os.O_WRONLY|os.O_CREATE, 0644)
							_, err = file1.Seek(int64(mbr.Partitions[i].Start+structs.SizeOfSuperBlock()+r*structs.SizeOfJournal()), 0)
							if err != nil {
								fmt.Println("Error al mover el puntero del archivo:", err)
								return
							}
							defer file1.Close()
							_, err = file1.Write(structs.NewJournal("logout", "", "", time.Now()).Encode())
							if err != nil {
								fmt.Println("Error al escribir SuperBlock:", err)
								return
							}
							break
						}
					}
				}
				break
			}
		}
		fmt.Printf("\033[32m-> logout: Sesión finalizada exitosamente. (%s) [%d:%d]\033[0m\n", env.CurrentLogged.User.Name, l.Line, l.Column)
		env.CurrentLogged.User = nil
		env.CurrentLogged.Partition = nil
		env.CurrentLogged.Driveletter = nil
		env.CurrentLogged.IDPart = nil
	} else {
		fmt.Printf("\033[31m-> Error logout: No hay ningún usuario loggeado actualmente. [%d:%d]\033[0m\n", l.Line, l.Column)
	}
}

func (l *Logout) GetResult() string { return "" }
