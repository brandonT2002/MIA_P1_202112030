package commands

import (
	"fmt"
	"io"
	structs "mia/Classes/Structs"
	"os"
	"strconv"
	"strings"
)

type Fdisk struct {
	Line   int
	Column int
	Params map[string]string
}

func thereAreFour(partitions []*structs.Partition) bool {
	for _, i := range partitions {
		if i.Status != '1' {
			return false
		}
	}
	return true
}

func thereIsNameR(name string, partitions []*structs.Partition) bool {
	for _, i := range partitions {
		if i.Status == '1' && strings.TrimSpace(i.Name) == name {
			return true
		}
	}
	return false
}

func thereIsNameRL(name string, partitions []*structs.EBR) bool {
	for _, i := range partitions {
		if i.Mount == '1' && strings.TrimSpace(i.Name) == name {
			return true
		}
	}
	return false
}

func getListEBR(start, size int, file *os.File) *structs.ListEBR {
	listEBR := structs.NewListEBR(start, size)
	_, err := file.Seek(int64(start), io.SeekStart)
	if err != nil {
		return nil
	}
	ebr := structs.DecodeEBR(make([]byte, 30))
	listEBR.Insert(ebr)
	for ebr.Next != -1 {
		_, err := file.Seek(int64(ebr.Next), io.SeekStart)
		if err != nil {
			return nil
		}
		ebr = structs.DecodeEBR(make([]byte, 30))
		listEBR.Insert(ebr)
	}
	return listEBR
}

func getExtended(partitions []*structs.Partition) int {
	for i, partition := range partitions {
		if partition.Type == 'E' {
			return i
		}
	}
	return -1
}

func sortBestFit(disponible [][]int) [][]int {
	if len(disponible) > 1 {
		for i := range disponible {
			for j := i; j > 0; j-- {
				if disponible[j][1] < disponible[j-1][1] {
					disponible[j], disponible[j-1] = disponible[j-1], disponible[j]
					continue
				}
				break
			}
		}
	}
	return disponible
}

func sortWorsFit(disponible [][]int) [][]int {
	if len(disponible) > 1 {
		for i := 1; i < len(disponible); i++ {
			for j := i; j > 0; j-- {
				if disponible[j][1] > disponible[j-1][1] {
					disponible[j], disponible[j-1] = disponible[j-1], disponible[j]
					continue
				}
				break
			}
		}
	}
	return disponible
}

func sortOrder(partitions []*structs.Partition) []*structs.Partition {
	for i := 1; i < len(partitions); i++ {
		if partitions[i].Start != 0 {
			for j := i; j > 0; j-- {
				if partitions[j].Start < partitions[j-1].Start {
					partitions[j], partitions[j-1] = partitions[j-1], partitions[j]
					continue
				}
				break
			}
			continue
		}
		break
	}
	return partitions
}

func (f *Fdisk) isDelete() bool {
	for k := range f.Params {
		if k == "delete" {
			return true
		}
	}
	return false
}

func (f *Fdisk) validateParamsDelete() bool {
	_, pathExists := f.Params["driveletter"]
	_, nameExists := f.Params["name"]
	return pathExists && nameExists
}

func (f *Fdisk) isAdd() bool {
	addValue, ok := f.Params["add"]
	if ok {
		_, err := strconv.Atoi(addValue)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (f *Fdisk) validateParamsAdd() bool {
	_, pathExists := f.Params["driveletter"]
	_, nameExists := f.Params["name"]
	unitValue, unitExists := f.Params["unit"]
	if pathExists && nameExists && unitExists {
		f.Params["unit"] = strings.ToUpper(unitValue)
		return true
	}
	return false
}

func (f *Fdisk) validateParams() bool {
	sizeStr, sizeExists := f.Params["size"]
	pathExists := f.Params["driveletter"] != ""
	nameExists := f.Params["name"] != ""
	if sizeExists && pathExists && nameExists {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			return false
		}
		f.Params["size"] = strconv.Itoa(size)
		return true
	}
	return false
}

func (f *Fdisk) printError(text string) {
	fmt.Printf("\033[31m-> %s. [%v:%v]\033[0m\n", text, f.Line, f.Column+1)
}

func (f *Fdisk) printSuccessDelete(text string) {
	fmt.Printf("\033[32m-> %s. [%v:%v]\033[0m\n", text, f.Line, f.Column+1)
}

func (f *Fdisk) printSuccessAdd(text, name, sign string, size int, unit, partitionType string) {
	partitionType = func() string {
		switch partitionType {
		case "P":
			return "Primaria"
		case "E":
			return "Extendida"
		default:
			return "Logica"
		}
	}()

	unit = func() string {
		if unit == "K" || unit == "M" {
			return unit
		}
		return ""
	}()
	fmt.Printf("\033[32m%s %s (%s %s%d %sB) [%d:%d]\033[0m\n", text, partitionType, name, sign, size, unit, f.Line, f.Column)
}

func (f *Fdisk) printSuccessCreate(diskname, name, partitionType string, size int, unit string) {
	partitionType = func() string {
		switch partitionType {
		case "P":
			return "Primaria"
		case "E":
			return "Extendida"
		default:
			return "Logica"
		}
	}()

	unit = func() string {
		if unit == "K" || unit == "M" {
			return unit
		}
		return ""
	}()
	fmt.Printf("\033[32m -> fdisk: Partición creada exitosamente en %s. %-9s (%s: %d %sB) [%d:%d]\033[0m\n", diskname, partitionType, name, size, unit, f.Line, f.Column)
}
