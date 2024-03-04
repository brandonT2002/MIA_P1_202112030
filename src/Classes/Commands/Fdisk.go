package commands

import (
	"fmt"
	"io"
	structs "mia/Classes/Structs"
	utils "mia/Classes/Utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Fdisk struct {
	Result string
	Type   utils.Type
	Line   int
	Column int
	Params map[string]string
}

func NewFdisk(line, column int, params map[string]string) *Fdisk {
	return &Fdisk{Type: utils.FDISK, Line: line, Column: column, Params: params}
}

func (f *Fdisk) GetLine() int {
	return f.Line
}

func (f *Fdisk) GetColumn() int {
	return f.Column
}

func (f *Fdisk) GetType() utils.Type {
	return f.Type
}

func (f *Fdisk) Exec() {
	if f.isDelete() {
		if f.validateParamsDelete() {
			// Eliminar partición
			fmt.Println("Eliminar partición")
		} else {
			f.printError("Error fdisk: Faltan parámetros obligatorios para eliminar la partición.")
		}
		return
	}
	if f.isAdd() {
		if f.validateParamsAdd() {
			// Agregar espacio de partición
			fmt.Println("Agregar espacio de partición")
		} else {
			f.printError("Error fdisk: Faltan parámetros obligatorios para agregar espacio a la partición.")
		}
		return
	}
	if f.validateParams() {
		// Crear Partición
		fmt.Println("Agregar partición")
		f.createPartition()
	} else {
		f.printError("Error fdisk: Faltan parámetros obligatorios para crear la partición.")
	}
}

func (f *Fdisk) createPartition() {
	f.Params["driveletter"] = strings.Replace(f.Params["driveletter"], `"`, "", -1)
	f.Params["fit"] = strings.ToUpper(f.Params["fit"])
	f.Params["type"] = strings.ToUpper(f.Params["type"])
	absolutePath, _ := filepath.Abs(fmt.Sprintf("../../Discos/%s.dsk", f.Params["driveletter"]))
	_, err := os.Stat(absolutePath)
	if err != nil {
		if os.IsNotExist(err) {
			f.printError(fmt.Sprintf("Error fdisk: No existe disco \"%s\" para particionar", strings.Split(filepath.Base(absolutePath), ".")[0]))
			return
		}
	}
	f.Params["unit"] = strings.ToUpper(f.Params["unit"])
	size, _ := strconv.Atoi(f.Params["size"])
	if size <= 0 {
		f.printError("Error: El tamaño de partición debe ser mayor que cero")
		return
	}
	units := 1
	if f.Params["unit"] == "M" {
		units = 1024 * 1024
	} else if f.Params["unit"] == "K" {
		units = 1024
	} else if f.Params["unit"] == "B" {
		units = 1
	} else {
		f.printError("Eror fdisk: Unidad de Bytes incorrecta")
		return
	}
	file, err := os.Open(absolutePath)
	if err != nil {
		fmt.Print("-> No se pudo abrir el archivo")
		return
	}
	defer file.Close()
	readedBytes := make([]byte, 153)
	_, err = file.Read(readedBytes)
	if err != nil {
		return
	}
	mbr := structs.DecodeMBR(readedBytes)
	if f.Params["type"] != "L" && f.thereAreFour(mbr.Partitions) {
		baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
		f.printError(fmt.Sprintf("Error fdisk: Ya existen 4 particiones en el disco %s.", baseName))
		return
	}
	if f.thereIsNameR(f.Params["name"], mbr.Partitions) {
		baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
		f.printError(fmt.Sprintf("Error fdisk: Ya existe una partición con el nombre %s en el disco \"%s\".", f.Params["name"], baseName))
		return
	}
	f.Params["fit"] = f.Params["fit"][:1]
	if f.Params["type"] == "P" || f.Params["type"] == "E" {
		var disponible [][]int
		lastNoEmptyByte := 126
		size, _ := strconv.Atoi(f.Params["size"])
		for i := 0; i < len(mbr.Partitions); i++ {
			if mbr.Partitions[i].Status == '1' {
				if mbr.Partitions[i].Start-lastNoEmptyByte > 2 && mbr.Partitions[i].Start-lastNoEmptyByte >= size*units+1 {
					disponible = append(disponible, []int{lastNoEmptyByte + 1, mbr.Partitions[i].Start - lastNoEmptyByte})
				}
				lastNoEmptyByte = mbr.Partitions[i].Start + mbr.Partitions[i].Size - 1
			}
		}
		if mbr.Size-lastNoEmptyByte > 2 && mbr.Size-lastNoEmptyByte >= size*units+1 {
			disponible = append(disponible, []int{lastNoEmptyByte + 1, mbr.Size - lastNoEmptyByte - 1})
		}
		if len(disponible) > 0 {
			if mbr.Fit == 'B' {
				disponible = f.sortBestFit(disponible)
			} else if mbr.Fit == 'W' {
				disponible = f.sortWorsFit(disponible)
			}
			if f.Params["type"] == "E" && f.getExtended(mbr.Partitions) != -1 {
				baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
				f.printError(fmt.Sprintf("Error fdisk: Ya existe una partición extendida en el disco \"%s.\"", baseName))
				return
			}
			for i := 0; i < len(mbr.Partitions); i++ {
				if mbr.Partitions[i].Status != '1' {
					mbr.Partitions[i] = &structs.Partition{
						Status:      '0',
						Type:        f.GetTypeP(),
						Fit:         f.GetFit(),
						Start:       disponible[0][0],
						Size:        size * units,
						Name:        f.Params["name"],
						Correlative: 0,
					}
					mbr.Partitions = f.sortOrder(mbr.Partitions)
					file, err := os.OpenFile(absolutePath, os.O_RDWR, 0644)
					if err != nil {
						fmt.Print("-> No es puedo abrir el archivo para escritura")
						return
					}
					defer file.Close()
					_, err = file.Seek(0, 0)
					if err != nil {
						fmt.Print("-> Ha ocurrido un error (Seek)(1)")
					}
					_, err = file.Write(mbr.Encode())
					if err != nil {
						fmt.Print("-> No se pudo escribir en el archivo")
					}
					if f.Params["type"] == "E" {
						_, err = file.Seek(int64(mbr.Partitions[i].Start), 0)
						if err != nil {
							fmt.Print("-> Error al posicionar el Seek (Partición Extendida)")
						}
						_, err = file.Write(structs.NewEBR().Encode())
						if err != nil {
							fmt.Print("-> No se pudo escribir en el archivo (EBR)")
						}
					}
					f.printSuccessCreate(
						strings.Split(filepath.Base(absolutePath), ".")[0],
						f.Params["name"],
						f.Params["type"],
						size,
						f.Params["unit"],
					)
					return
				}
			}
			baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
			f.printError(fmt.Sprintf("Error fdisk: No pueden crearse mas particiones en el disco \"%s\".", baseName))
		} else {
			baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
			f.printError(fmt.Sprintf("Error fdisk: No hay espacio suficiente para la nueva partición en el disco \"%s\".", baseName))
		}
	} else if f.Params["type"] == "L" {
		i := f.getExtended(mbr.Partitions)
		size, _ := strconv.Atoi(f.Params["size"])
		if i != -1 {
			listEbr := f.getListEBR(mbr.Partitions[i].Start, mbr.Partitions[i].Size, file)
			if f.thereIsNameRL(f.Params["name"], listEbr.GetIterable()) {
				baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
				f.printError(fmt.Sprintf("Error fdisk: Ya existe una partición con el nombre %s en disco \"%s\".", f.Params["name"], baseName))
				return
			}
			disponible := listEbr.SearchEmptySpace(size * units)
			if len(disponible) > 0 {
				if mbr.Partitions[i].Fit == 'B' {
					disponible = f.sortBestFit(disponible)
				} else if mbr.Partitions[i].Fit == 'W' {
					disponible = f.sortWorsFit(disponible)
				}
				name := f.padLeft(f.Params["name"], 16)[:16]
				ebr := &structs.EBR{
					Mount: '0',
					Fit:   f.GetFit(),
					Start: disponible[0][0],
					Size:  size * units,
					Next:  -1,
					Name:  name,
				}
				listEbr.Insert(ebr)
				file, err := os.OpenFile(absolutePath, os.O_RDWR, 0644)
				if err != nil {
					fmt.Print("-> No es puedo abrir el archivo para escritura")
					return
				} else {
					defer file.Close()
					var pos int
					for _, e := range listEbr.GetIterable() {
						pos = e.Start
						if pos == 0 {
							pos = mbr.Partitions[i].Start
						}
						_, err := file.Seek(int64(pos), 0)
						if err != nil {
							fmt.Print("-> Ha ocurrido un error (Seek)(2)")
						}
						_, err = file.Write(e.Encode())
						if err != nil {
							fmt.Print("-> Ha ocurrido al escribir los datos en el archivo (2)")
						}
					}
					f.printSuccessCreate(strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath)), f.Params["name"], f.Params["type"], size, f.Params["unit"])
					return
				}
			}
			baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
			f.printError(fmt.Sprintf("Error fdisk: No hay espacio suficiente para la nueva partición en el disco \"%s\".", baseName))
			return
		}
		baseName := strings.TrimSuffix(filepath.Base(absolutePath), filepath.Ext(absolutePath))
		f.printError(fmt.Sprintf("Error fdisk: No existe una partición extendida en el disco \"%s\" para crear la partición lógica..", baseName))
	}
}

func (f *Fdisk) padLeft(s string, width int) string {
	padLength := width - len(s)
	if padLength > 0 {
		paddingStr := strings.Repeat(" ", padLength)
		return paddingStr + s
	}
	return s
}

func (f *Fdisk) thereAreFour(partitions []*structs.Partition) bool {
	for _, i := range partitions {
		if i.Status != '1' {
			return false
		}
	}
	return true
}

func (f *Fdisk) thereIsNameR(name string, partitions []*structs.Partition) bool {
	for _, i := range partitions {
		if i.Status == '1' && strings.TrimSpace(i.Name) == name {
			return true
		}
	}
	return false
}

func (f *Fdisk) thereIsNameRL(name string, partitions []*structs.EBR) bool {
	for _, i := range partitions {
		if i.Mount == '1' && strings.TrimSpace(i.Name) == name {
			return true
		}
	}
	return false
}

func (f *Fdisk) getListEBR(start, size int, file *os.File) *structs.ListEBR {
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

func (f *Fdisk) getExtended(partitions []*structs.Partition) int {
	for i, partition := range partitions {
		if partition.Type == 'E' {
			return i
		}
	}
	return -1
}

func (f *Fdisk) sortBestFit(disponible [][]int) [][]int {
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

func (f *Fdisk) sortWorsFit(disponible [][]int) [][]int {
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

func (f *Fdisk) sortOrder(partitions []*structs.Partition) []*structs.Partition {
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
	_, sizeExists := f.Params["size"]
	_, pathExists := f.Params["driveletter"]
	_, nameExists := f.Params["name"]
	return sizeExists && pathExists && nameExists
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
		fmt.Println("-> SI FUNKA")
		f.Params["unit"] = strings.ToUpper(unitValue)
		return true
	}
	return false
}

func (f *Fdisk) validateParams() bool {
	sizeStr, sizeExists := f.Params["size"]
	_, pathExists := f.Params["driveletter"]
	_, nameExists := f.Params["name"]
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

func (f *Fdisk) GetFit() rune {
	f.Params["fit"] = strings.ToUpper(f.Params["fit"])
	if f.Params["fit"] == "FF" {
		return 'F'
	}
	if f.Params["fit"] == "BF" {
		return 'B'
	}
	return 'W'
}

func (f *Fdisk) GetTypeP() rune {
	f.Params["type"] = strings.ToUpper(f.Params["type"])
	if f.Params["type"] == "P" {
		return 'P'
	}
	if f.Params["type"] == "E" {
		return 'E'
	}
	return 'W'
}

func (f *Fdisk) GetResult() string { return "" }
