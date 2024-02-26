package commands

import (
	"fmt"
	env "mia/Classes/Env"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"strconv"
	"strings"
)

type Mkdisk struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewMkdisk(line, column int, params map[string]string) *Mkdisk {
	return &Mkdisk{Type: utils.MKDISK, Line: line, Column: column, Params: params}
}

func (m *Mkdisk) GetLine() int {
	return m.Line
}

func (m *Mkdisk) GetColumn() int {
	return m.Column
}

func (m *Mkdisk) GetType() utils.Type {
	return m.Type
}

func (m *Mkdisk) Exec() {
	if m.ValidateParams() {
		file, _ := os.Create(fmt.Sprintf("../../Discos/%c.dsk", env.Asciiletter))
		units := m.RecalculateUnits()
		size, _ := strconv.Atoi(m.Params["size"])
		for i := 0; i < size; i++ {
			byte := make([]byte, units)
			file.Write(byte)
		}
		file.Close()
		file, _ = os.OpenFile(fmt.Sprintf("../../Discos/%c.dsk", env.Asciiletter), os.O_RDWR, 0644)
		file.Seek(0, 0)
		mbr := structs.NewMBR(units*size, m.GetFit())
		file.Write(mbr.Encode())
		file.Close()
		env.Asciiletter++
		fmt.Printf("\033[96m-> %v. [%v:%v]\033[0m\n", "mkdisk: Disco creado", m.Line, m.Column+1)
	} else {
		fmt.Printf("\033[31m-> Error mkdisk: Faltan parámetros obligatorios. [%v:%v]\033[0m\n", m.Line, m.Column+1)
	}
}

func (m *Mkdisk) ValidateParams() bool {
	if _, exist := m.Params["size"]; exist {
		return true
	}
	return false
}

func (m *Mkdisk) RecalculateUnits() int {
	if _, exist := m.Params["unit"]; exist {
		if strings.ToUpper(m.Params["unit"]) == "K" {
			return 1024
		}
	}
	return 1024 * 1024
}

func (m *Mkdisk) GetFit() rune {
	if strings.ToUpper(m.Params["fit"]) == "FF" {
		return 'F'
	}
	if strings.ToUpper(m.Params["fit"]) == "BF" {
		return 'B'
	}
	return 'W'
}

func (m *Mkdisk) GetResult() string { return "" }
