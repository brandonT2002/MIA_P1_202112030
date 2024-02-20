package main

import (
	"fmt"
	"io/ioutil"
	listener "mia/Language"
	parser "mia/Language/Parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	filePath := "C:\\Users\\bryan\\Documents\\USAC\\Organización de Lenguajes y Compiladores 2\\Laboratorio\\Proyecto1\\Calificación\\Hanoi.swift"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	input := string(fileData)
	inputStream := antlr.NewInputStream(input)
	scanner := parser.NewScanner(inputStream)
	tokens := antlr.NewCommonTokenStream(scanner, antlr.TokenDefaultChannel)
	parser := parser.NewParserParser(tokens)
	parser.RemoveErrorListeners()
	parserErrors := listener.NewMIAErrorListener()
	parser.AddErrorListener(parserErrors)

	parser.BuildParseTrees = true
	tree := parser.Init()
	var listener *listener.MIAListener = listener.NewMIAListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}
