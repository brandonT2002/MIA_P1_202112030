// Code generated from Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Parser

import "github.com/antlr4-go/antlr/v4"

// ParserListener is a complete listener for a parse tree produced by ParserParser.
type ParserListener interface {
	antlr.ParseTreeListener

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// EnterCommands is called when entering the commands production.
	EnterCommands(c *CommandsContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)
}
