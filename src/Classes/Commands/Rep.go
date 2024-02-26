package commands

import (
	"fmt"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
)

type Rep struct {
	Result string
	Type   utils.Type
	Params map[string]string
	Line   int
	Column int
}

func NewRep(line, column int, params map[string]string) *Rep {
	return &Rep{Type: utils.REP, Line: line, Column: column, Params: params}
}

func (r *Rep) GetLine() int {
	return r.Line
}

func (r *Rep) GetColumn() int {
	return r.Column
}

func (r *Rep) GetType() utils.Type {
	return r.Type
}

func (r *Rep) Exec() {
	file, _ := os.Open("../../Discos/Disco.dsk")
	readedBytes := make([]byte, 153)
	file.Read(readedBytes)
	fmt.Printf("\033[32m%v\033[0m\n", structs.DecodeMBR(readedBytes).ToString())
}

func (r *Rep) GetResult() string { return "" }
