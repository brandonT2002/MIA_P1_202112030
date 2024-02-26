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

	// EnterMkdisk is called when entering the mkdisk production.
	EnterMkdisk(c *MkdiskContext)

	// EnterMkdiskparams is called when entering the mkdiskparams production.
	EnterMkdiskparams(c *MkdiskparamsContext)

	// EnterMkdiskparam is called when entering the mkdiskparam production.
	EnterMkdiskparam(c *MkdiskparamContext)

	// EnterRmdisk is called when entering the rmdisk production.
	EnterRmdisk(c *RmdiskContext)

	// EnterRep is called when entering the rep production.
	EnterRep(c *RepContext)

	// EnterRepparams is called when entering the repparams production.
	EnterRepparams(c *RepparamsContext)

	// EnterRepparam is called when entering the repparam production.
	EnterRepparam(c *RepparamContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterCommentary is called when entering the commentary production.
	EnterCommentary(c *CommentaryContext)

	// EnterExit is called when entering the exit production.
	EnterExit(c *ExitContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)

	// ExitCommands is called when exiting the commands production.
	ExitCommands(c *CommandsContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitMkdisk is called when exiting the mkdisk production.
	ExitMkdisk(c *MkdiskContext)

	// ExitMkdiskparams is called when exiting the mkdiskparams production.
	ExitMkdiskparams(c *MkdiskparamsContext)

	// ExitMkdiskparam is called when exiting the mkdiskparam production.
	ExitMkdiskparam(c *MkdiskparamContext)

	// ExitRmdisk is called when exiting the rmdisk production.
	ExitRmdisk(c *RmdiskContext)

	// ExitRep is called when exiting the rep production.
	ExitRep(c *RepContext)

	// ExitRepparams is called when exiting the repparams production.
	ExitRepparams(c *RepparamsContext)

	// ExitRepparam is called when exiting the repparam production.
	ExitRepparam(c *RepparamContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitCommentary is called when exiting the commentary production.
	ExitCommentary(c *CommentaryContext)

	// ExitExit is called when exiting the exit production.
	ExitExit(c *ExitContext)
}
