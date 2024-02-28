// Code generated from Parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import (
	commands "mia/Classes/Commands"
	interfaces "mia/Classes/Interfaces"

	"os"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParserParser struct {
	*antlr.BaseParser
}

var ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserParserInit() {
	staticData := &ParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'unmount'", "'mkfs'",
		"'login'", "'logout'", "'mkgrp'", "'rmgrp'", "'mkusr'", "'rmusr'", "'mkfile'",
		"'cat'", "'remove'", "'edit'", "'rename'", "'mkdir'", "'copy'", "'move'",
		"'find'", "'chown'", "'chgrp'", "'chmod'", "'pause'", "'recovery'",
		"'loss'", "'execute'", "'rep'", "'exit'", "'mbr'", "'disk'", "'inode'",
		"'journaling'", "'block'", "'bm_inode'", "'bm_block'", "'tree'", "'sb'",
		"'file'", "'ls'", "'full'", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"'='", "", "", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "RW_mkdisk", "RW_rmdisk", "RW_fdisk", "RW_mount", "RW_unmount",
		"RW_mkfs", "RW_login", "RW_logout", "RW_mkgrp", "RW_rmgrp", "RW_mkusr",
		"RW_rmusr", "RW_mkfile", "RW_cat", "RW_remove", "RW_edit", "RW_rename",
		"RW_mkdir", "RW_copy", "RW_move", "RW_find", "RW_chown", "RW_chgrp",
		"RW_chmod", "RW_pause", "RW_recovery", "RW_loss", "RW_execute", "RW_rep",
		"RW_exit", "RW_mbr", "RW_disk", "RW_inode", "RW_journaling", "RW_block",
		"RW_bm_inode", "RW_bm_block", "RW_tree", "RW_sb", "RW_file", "RW_ls",
		"RW_full", "RW_size", "RW_fit", "RW_unit", "RW_driveletter", "RW_name",
		"RW_type", "RW_delete", "RW_add", "RW_id", "RW_fs", "RW_user", "RW_pass",
		"RW_grp", "RW_path", "RW_r", "RW_cont", "RW_fileN", "RW_destino", "RW_ugo",
		"RW_ruta", "TK_fit", "TK_unit", "TK_type", "TK_fs", "TK_number", "TK_id",
		"TK_path", "TK_equ", "TK_sym", "COMMENTARY", "NEWLINE", "UNUSED_",
	}
	staticData.RuleNames = []string{
		"init", "commands", "command", "execute", "mkdisk", "mkdiskparams",
		"mkdiskparam", "rmdisk", "rep", "repparams", "repparam", "name", "commentary",
		"exit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 191, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 3, 0, 35, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 66,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 75, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 83, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 5, 5, 93, 8, 5, 10, 5, 12, 5, 96, 9, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 110, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 119, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 127, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 5, 9, 137, 8, 9, 10, 9, 12, 9, 140, 9, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 159, 8, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 183,
		8, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 0, 3, 2, 10, 18,
		14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 0, 204, 0, 34,
		1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 65, 1, 0, 0, 0, 6, 74, 1, 0, 0, 0, 8,
		82, 1, 0, 0, 0, 10, 84, 1, 0, 0, 0, 12, 109, 1, 0, 0, 0, 14, 118, 1, 0,
		0, 0, 16, 126, 1, 0, 0, 0, 18, 128, 1, 0, 0, 0, 20, 158, 1, 0, 0, 0, 22,
		182, 1, 0, 0, 0, 24, 184, 1, 0, 0, 0, 26, 187, 1, 0, 0, 0, 28, 29, 3, 2,
		1, 0, 29, 30, 5, 0, 0, 1, 30, 31, 6, 0, -1, 0, 31, 35, 1, 0, 0, 0, 32,
		33, 5, 0, 0, 1, 33, 35, 6, 0, -1, 0, 34, 28, 1, 0, 0, 0, 34, 32, 1, 0,
		0, 0, 35, 1, 1, 0, 0, 0, 36, 37, 6, 1, -1, 0, 37, 38, 3, 4, 2, 0, 38, 39,
		6, 1, -1, 0, 39, 46, 1, 0, 0, 0, 40, 41, 10, 2, 0, 0, 41, 42, 3, 4, 2,
		0, 42, 43, 6, 1, -1, 0, 43, 45, 1, 0, 0, 0, 44, 40, 1, 0, 0, 0, 45, 48,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0,
		48, 46, 1, 0, 0, 0, 49, 50, 3, 6, 3, 0, 50, 51, 6, 2, -1, 0, 51, 66, 1,
		0, 0, 0, 52, 53, 3, 8, 4, 0, 53, 54, 6, 2, -1, 0, 54, 66, 1, 0, 0, 0, 55,
		56, 3, 14, 7, 0, 56, 57, 6, 2, -1, 0, 57, 66, 1, 0, 0, 0, 58, 59, 3, 16,
		8, 0, 59, 60, 6, 2, -1, 0, 60, 66, 1, 0, 0, 0, 61, 62, 3, 24, 12, 0, 62,
		63, 6, 2, -1, 0, 63, 66, 1, 0, 0, 0, 64, 66, 3, 26, 13, 0, 65, 49, 1, 0,
		0, 0, 65, 52, 1, 0, 0, 0, 65, 55, 1, 0, 0, 0, 65, 58, 1, 0, 0, 0, 65, 61,
		1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 5, 1, 0, 0, 0, 67, 68, 5, 28, 0, 0,
		68, 69, 5, 56, 0, 0, 69, 70, 5, 70, 0, 0, 70, 71, 5, 69, 0, 0, 71, 75,
		6, 3, -1, 0, 72, 73, 5, 28, 0, 0, 73, 75, 6, 3, -1, 0, 74, 67, 1, 0, 0,
		0, 74, 72, 1, 0, 0, 0, 75, 7, 1, 0, 0, 0, 76, 77, 5, 1, 0, 0, 77, 78, 3,
		10, 5, 0, 78, 79, 6, 4, -1, 0, 79, 83, 1, 0, 0, 0, 80, 81, 5, 1, 0, 0,
		81, 83, 6, 4, -1, 0, 82, 76, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 9, 1,
		0, 0, 0, 84, 85, 6, 5, -1, 0, 85, 86, 3, 12, 6, 0, 86, 87, 6, 5, -1, 0,
		87, 94, 1, 0, 0, 0, 88, 89, 10, 2, 0, 0, 89, 90, 3, 12, 6, 0, 90, 91, 6,
		5, -1, 0, 91, 93, 1, 0, 0, 0, 92, 88, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 11, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 97, 98, 5, 43, 0, 0, 98, 99, 5, 70, 0, 0, 99, 100, 5, 67, 0, 0, 100,
		110, 6, 6, -1, 0, 101, 102, 5, 44, 0, 0, 102, 103, 5, 70, 0, 0, 103, 104,
		5, 63, 0, 0, 104, 110, 6, 6, -1, 0, 105, 106, 5, 45, 0, 0, 106, 107, 5,
		70, 0, 0, 107, 108, 5, 64, 0, 0, 108, 110, 6, 6, -1, 0, 109, 97, 1, 0,
		0, 0, 109, 101, 1, 0, 0, 0, 109, 105, 1, 0, 0, 0, 110, 13, 1, 0, 0, 0,
		111, 112, 5, 2, 0, 0, 112, 113, 5, 46, 0, 0, 113, 114, 5, 70, 0, 0, 114,
		115, 5, 68, 0, 0, 115, 119, 6, 7, -1, 0, 116, 117, 5, 2, 0, 0, 117, 119,
		6, 7, -1, 0, 118, 111, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 15, 1, 0,
		0, 0, 120, 121, 5, 29, 0, 0, 121, 122, 3, 18, 9, 0, 122, 123, 6, 8, -1,
		0, 123, 127, 1, 0, 0, 0, 124, 125, 5, 29, 0, 0, 125, 127, 6, 8, -1, 0,
		126, 120, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 17, 1, 0, 0, 0, 128, 129,
		6, 9, -1, 0, 129, 130, 3, 20, 10, 0, 130, 131, 6, 9, -1, 0, 131, 138, 1,
		0, 0, 0, 132, 133, 10, 2, 0, 0, 133, 134, 3, 20, 10, 0, 134, 135, 6, 9,
		-1, 0, 135, 137, 1, 0, 0, 0, 136, 132, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0,
		138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 19, 1, 0, 0, 0, 140, 138,
		1, 0, 0, 0, 141, 142, 5, 47, 0, 0, 142, 143, 5, 70, 0, 0, 143, 144, 3,
		22, 11, 0, 144, 145, 6, 10, -1, 0, 145, 159, 1, 0, 0, 0, 146, 147, 5, 56,
		0, 0, 147, 148, 5, 70, 0, 0, 148, 149, 5, 69, 0, 0, 149, 159, 6, 10, -1,
		0, 150, 151, 5, 62, 0, 0, 151, 152, 5, 70, 0, 0, 152, 153, 5, 69, 0, 0,
		153, 159, 6, 10, -1, 0, 154, 155, 5, 51, 0, 0, 155, 156, 5, 70, 0, 0, 156,
		157, 5, 68, 0, 0, 157, 159, 6, 10, -1, 0, 158, 141, 1, 0, 0, 0, 158, 146,
		1, 0, 0, 0, 158, 150, 1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 159, 21, 1, 0,
		0, 0, 160, 161, 5, 31, 0, 0, 161, 183, 6, 11, -1, 0, 162, 163, 5, 32, 0,
		0, 163, 183, 6, 11, -1, 0, 164, 165, 5, 33, 0, 0, 165, 183, 6, 11, -1,
		0, 166, 167, 5, 34, 0, 0, 167, 183, 6, 11, -1, 0, 168, 169, 5, 35, 0, 0,
		169, 183, 6, 11, -1, 0, 170, 171, 5, 36, 0, 0, 171, 183, 6, 11, -1, 0,
		172, 173, 5, 37, 0, 0, 173, 183, 6, 11, -1, 0, 174, 175, 5, 38, 0, 0, 175,
		183, 6, 11, -1, 0, 176, 177, 5, 39, 0, 0, 177, 183, 6, 11, -1, 0, 178,
		179, 5, 40, 0, 0, 179, 183, 6, 11, -1, 0, 180, 181, 5, 41, 0, 0, 181, 183,
		6, 11, -1, 0, 182, 160, 1, 0, 0, 0, 182, 162, 1, 0, 0, 0, 182, 164, 1,
		0, 0, 0, 182, 166, 1, 0, 0, 0, 182, 168, 1, 0, 0, 0, 182, 170, 1, 0, 0,
		0, 182, 172, 1, 0, 0, 0, 182, 174, 1, 0, 0, 0, 182, 176, 1, 0, 0, 0, 182,
		178, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 5,
		72, 0, 0, 185, 186, 6, 12, -1, 0, 186, 25, 1, 0, 0, 0, 187, 188, 5, 30,
		0, 0, 188, 189, 6, 13, -1, 0, 189, 27, 1, 0, 0, 0, 12, 34, 46, 65, 74,
		82, 94, 109, 118, 126, 138, 158, 182,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ParserParserInit initializes any static state used to implement ParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserParserInit() {
	staticData := &ParserParserStaticData
	staticData.once.Do(parserParserInit)
}

// NewParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParserParser(input antlr.TokenStream) *ParserParser {
	ParserParserInit()
	this := new(ParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Parser.g4"

	return this
}

// ParserParser tokens.
const (
	ParserParserEOF            = antlr.TokenEOF
	ParserParserRW_mkdisk      = 1
	ParserParserRW_rmdisk      = 2
	ParserParserRW_fdisk       = 3
	ParserParserRW_mount       = 4
	ParserParserRW_unmount     = 5
	ParserParserRW_mkfs        = 6
	ParserParserRW_login       = 7
	ParserParserRW_logout      = 8
	ParserParserRW_mkgrp       = 9
	ParserParserRW_rmgrp       = 10
	ParserParserRW_mkusr       = 11
	ParserParserRW_rmusr       = 12
	ParserParserRW_mkfile      = 13
	ParserParserRW_cat         = 14
	ParserParserRW_remove      = 15
	ParserParserRW_edit        = 16
	ParserParserRW_rename      = 17
	ParserParserRW_mkdir       = 18
	ParserParserRW_copy        = 19
	ParserParserRW_move        = 20
	ParserParserRW_find        = 21
	ParserParserRW_chown       = 22
	ParserParserRW_chgrp       = 23
	ParserParserRW_chmod       = 24
	ParserParserRW_pause       = 25
	ParserParserRW_recovery    = 26
	ParserParserRW_loss        = 27
	ParserParserRW_execute     = 28
	ParserParserRW_rep         = 29
	ParserParserRW_exit        = 30
	ParserParserRW_mbr         = 31
	ParserParserRW_disk        = 32
	ParserParserRW_inode       = 33
	ParserParserRW_journaling  = 34
	ParserParserRW_block       = 35
	ParserParserRW_bm_inode    = 36
	ParserParserRW_bm_block    = 37
	ParserParserRW_tree        = 38
	ParserParserRW_sb          = 39
	ParserParserRW_file        = 40
	ParserParserRW_ls          = 41
	ParserParserRW_full        = 42
	ParserParserRW_size        = 43
	ParserParserRW_fit         = 44
	ParserParserRW_unit        = 45
	ParserParserRW_driveletter = 46
	ParserParserRW_name        = 47
	ParserParserRW_type        = 48
	ParserParserRW_delete      = 49
	ParserParserRW_add         = 50
	ParserParserRW_id          = 51
	ParserParserRW_fs          = 52
	ParserParserRW_user        = 53
	ParserParserRW_pass        = 54
	ParserParserRW_grp         = 55
	ParserParserRW_path        = 56
	ParserParserRW_r           = 57
	ParserParserRW_cont        = 58
	ParserParserRW_fileN       = 59
	ParserParserRW_destino     = 60
	ParserParserRW_ugo         = 61
	ParserParserRW_ruta        = 62
	ParserParserTK_fit         = 63
	ParserParserTK_unit        = 64
	ParserParserTK_type        = 65
	ParserParserTK_fs          = 66
	ParserParserTK_number      = 67
	ParserParserTK_id          = 68
	ParserParserTK_path        = 69
	ParserParserTK_equ         = 70
	ParserParserTK_sym         = 71
	ParserParserCOMMENTARY     = 72
	ParserParserNEWLINE        = 73
	ParserParserUNUSED_        = 74
)

// ParserParser rules.
const (
	ParserParserRULE_init         = 0
	ParserParserRULE_commands     = 1
	ParserParserRULE_command      = 2
	ParserParserRULE_execute      = 3
	ParserParserRULE_mkdisk       = 4
	ParserParserRULE_mkdiskparams = 5
	ParserParserRULE_mkdiskparam  = 6
	ParserParserRULE_rmdisk       = 7
	ParserParserRULE_rep          = 8
	ParserParserRULE_repparams    = 9
	ParserParserRULE_repparam     = 10
	ParserParserRULE_name         = 11
	ParserParserRULE_commentary   = 12
	ParserParserRULE_exit         = 13
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() ICommandsContext

	// SetC sets the c rule contexts.
	SetC(ICommandsContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	EOF() antlr.TerminalNode
	Commands() ICommandsContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	c      ICommandsContext
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) GetC() ICommandsContext { return s.c }

func (s *InitContext) SetC(v ICommandsContext) { s.c = v }

func (s *InitContext) GetResult() []interfaces.Command { return s.result }

func (s *InitContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserParserEOF, 0)
}

func (s *InitContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ParserParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserParserRULE_init)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk, ParserParserRW_rmdisk, ParserParserRW_execute, ParserParserRW_rep, ParserParserRW_exit, ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)

			var _x = p.commands(0)

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(29)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = localctx.(*InitContext).GetC().GetResult()

	case ParserParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = []interfaces.Command{}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ICommandsContext

	// GetC returns the c rule contexts.
	GetC() ICommandContext

	// SetL sets the l rule contexts.
	SetL(ICommandsContext)

	// SetC sets the c rule contexts.
	SetC(ICommandContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// Getter signatures
	Command() ICommandContext
	Commands() ICommandsContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	l      ICommandsContext
	c      ICommandContext
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) GetL() ICommandsContext { return s.l }

func (s *CommandsContext) GetC() ICommandContext { return s.c }

func (s *CommandsContext) SetL(v ICommandsContext) { s.l = v }

func (s *CommandsContext) SetC(v ICommandContext) { s.c = v }

func (s *CommandsContext) GetResult() []interfaces.Command { return s.result }

func (s *CommandsContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *CommandsContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) Commands() ICommandsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandsContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCommands(s)
	}
}

func (s *CommandsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCommands(s)
	}
}

func (p *ParserParser) Commands() (localctx ICommandsContext) {
	return p.commands(0)
}

func (p *ParserParser) commands(_p int) (localctx ICommandsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICommandsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ParserParserRULE_commands, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)

		var _x = p.Command()

		localctx.(*CommandsContext).c = _x
	}
	localctx.(*CommandsContext).result = []interfaces.Command{localctx.(*CommandsContext).GetC().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCommandsContext(p, _parentctx, _parentState)
			localctx.(*CommandsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_commands)
			p.SetState(40)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(41)

				var _x = p.Command()

				localctx.(*CommandsContext).c = _x
			}
			localctx.(*CommandsContext).SetResult(localctx.(*CommandsContext).GetL().GetResult())
			localctx.(*CommandsContext).result = append(localctx.(*CommandsContext).result, localctx.(*CommandsContext).GetC().GetResult())

		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() IExecuteContext

	// GetC2 returns the c2 rule contexts.
	GetC2() IMkdiskContext

	// GetC3 returns the c3 rule contexts.
	GetC3() IRmdiskContext

	// GetC20 returns the c20 rule contexts.
	GetC20() IRepContext

	// GetC21 returns the c21 rule contexts.
	GetC21() ICommentaryContext

	// SetC1 sets the c1 rule contexts.
	SetC1(IExecuteContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(IMkdiskContext)

	// SetC3 sets the c3 rule contexts.
	SetC3(IRmdiskContext)

	// SetC20 sets the c20 rule contexts.
	SetC20(IRepContext)

	// SetC21 sets the c21 rule contexts.
	SetC21(ICommentaryContext)

	// GetResult returns the result attribute.
	GetResult() interfaces.Command

	// SetResult sets the result attribute.
	SetResult(interfaces.Command)

	// Getter signatures
	Execute() IExecuteContext
	Mkdisk() IMkdiskContext
	Rmdisk() IRmdiskContext
	Rep() IRepContext
	Commentary() ICommentaryContext
	Exit() IExitContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result interfaces.Command
	c1     IExecuteContext
	c2     IMkdiskContext
	c3     IRmdiskContext
	c20    IRepContext
	c21    ICommentaryContext
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) GetC1() IExecuteContext { return s.c1 }

func (s *CommandContext) GetC2() IMkdiskContext { return s.c2 }

func (s *CommandContext) GetC3() IRmdiskContext { return s.c3 }

func (s *CommandContext) GetC20() IRepContext { return s.c20 }

func (s *CommandContext) GetC21() ICommentaryContext { return s.c21 }

func (s *CommandContext) SetC1(v IExecuteContext) { s.c1 = v }

func (s *CommandContext) SetC2(v IMkdiskContext) { s.c2 = v }

func (s *CommandContext) SetC3(v IRmdiskContext) { s.c3 = v }

func (s *CommandContext) SetC20(v IRepContext) { s.c20 = v }

func (s *CommandContext) SetC21(v ICommentaryContext) { s.c21 = v }

func (s *CommandContext) GetResult() interfaces.Command { return s.result }

func (s *CommandContext) SetResult(v interfaces.Command) { s.result = v }

func (s *CommandContext) Execute() IExecuteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *CommandContext) Mkdisk() IMkdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *CommandContext) Rmdisk() IRmdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdiskContext)
}

func (s *CommandContext) Rep() IRepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepContext)
}

func (s *CommandContext) Commentary() ICommentaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentaryContext)
}

func (s *CommandContext) Exit() IExitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExitContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *ParserParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserParserRULE_command)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_execute:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _x = p.Execute()

			localctx.(*CommandContext).c1 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC1().GetResult()

	case ParserParserRW_mkdisk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)

			var _x = p.Mkdisk()

			localctx.(*CommandContext).c2 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC2().GetResult()

	case ParserParserRW_rmdisk:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)

			var _x = p.Rmdisk()

			localctx.(*CommandContext).c3 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC3().GetResult()

	case ParserParserRW_rep:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)

			var _x = p.Rep()

			localctx.(*CommandContext).c20 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC20().GetResult()

	case ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)

			var _x = p.Commentary()

			localctx.(*CommandContext).c21 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC21().GetResult()

	case ParserParserRW_exit:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Exit()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e token.
	GetE() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetE sets the e token.
	SetE(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Execute

	// SetResult sets the result attribute.
	SetResult(*commands.Execute)

	// Getter signatures
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_execute() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Execute
	e      antlr.Token
	p      antlr.Token
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_execute
	return p
}

func InitEmptyExecuteContext(p *ExecuteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_execute
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) GetE() antlr.Token { return s.e }

func (s *ExecuteContext) GetP() antlr.Token { return s.p }

func (s *ExecuteContext) SetE(v antlr.Token) { s.e = v }

func (s *ExecuteContext) SetP(v antlr.Token) { s.p = v }

func (s *ExecuteContext) GetResult() *commands.Execute { return s.result }

func (s *ExecuteContext) SetResult(v *commands.Execute) { s.result = v }

func (s *ExecuteContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *ExecuteContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *ExecuteContext) RW_execute() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_execute, 0)
}

func (s *ExecuteContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (p *ParserParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserParserRULE_execute)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)

			var _m = p.Match(ParserParserRW_execute)

			localctx.(*ExecuteContext).e = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*ExecuteContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExecuteContext).result = commands.NewExecute((func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetColumn()
			}
		}()), map[string]string{"path": (func() string {
			if localctx.(*ExecuteContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*ExecuteContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)

			var _m = p.Match(ParserParserRW_execute)

			localctx.(*ExecuteContext).e = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExecuteContext).result = commands.NewExecute((func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExecuteContext).GetE() == nil {
				return 0
			} else {
				return localctx.(*ExecuteContext).GetE().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskContext is an interface to support dynamic dispatch.
type IMkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Mkdisk)

	// Getter signatures
	RW_mkdisk() antlr.TerminalNode
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskContext differentiates from other interfaces.
	IsMkdiskContext()
}

type MkdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkdisk
	m      antlr.Token
	p      IMkdiskparamsContext
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
	return p
}

func InitEmptyMkdiskContext(p *MkdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdisk
}

func (*MkdiskContext) IsMkdiskContext() {}

func NewMkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskContext {
	var p = new(MkdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdisk

	return p
}

func (s *MkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskContext) GetM() antlr.Token { return s.m }

func (s *MkdiskContext) SetM(v antlr.Token) { s.m = v }

func (s *MkdiskContext) GetP() IMkdiskparamsContext { return s.p }

func (s *MkdiskContext) SetP(v IMkdiskparamsContext) { s.p = v }

func (s *MkdiskContext) GetResult() *commands.Mkdisk { return s.result }

func (s *MkdiskContext) SetResult(v *commands.Mkdisk) { s.result = v }

func (s *MkdiskContext) RW_mkdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkdisk, 0)
}

func (s *MkdiskContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdisk(s)
	}
}

func (s *MkdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdisk(s)
	}
}

func (p *ParserParser) Mkdisk() (localctx IMkdiskContext) {
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserParserRULE_mkdisk)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)

			var _x = p.mkdiskparams(0)

			localctx.(*MkdiskContext).p = _x
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), localctx.(*MkdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskContext).result = commands.NewMkdisk((func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkdiskContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkdiskContext).GetM().GetColumn()
			}
		}()), map[string]string{"fit": "FF", "unit": "M"})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamsContext is an interface to support dynamic dispatch.
type IMkdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IMkdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkdiskparam() IMkdiskparamContext
	Mkdiskparams() IMkdiskparamsContext

	// IsMkdiskparamsContext differentiates from other interfaces.
	IsMkdiskparamsContext()
}

type MkdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkdiskparamsContext
	p      IMkdiskparamContext
}

func NewEmptyMkdiskparamsContext() *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
	return p
}

func InitEmptyMkdiskparamsContext(p *MkdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparams
}

func (*MkdiskparamsContext) IsMkdiskparamsContext() {}

func NewMkdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamsContext {
	var p = new(MkdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparams

	return p
}

func (s *MkdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamsContext) GetL() IMkdiskparamsContext { return s.l }

func (s *MkdiskparamsContext) GetP() IMkdiskparamContext { return s.p }

func (s *MkdiskparamsContext) SetL(v IMkdiskparamsContext) { s.l = v }

func (s *MkdiskparamsContext) SetP(v IMkdiskparamContext) { s.p = v }

func (s *MkdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *MkdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkdiskparamsContext) Mkdiskparam() IMkdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamContext)
}

func (s *MkdiskparamsContext) Mkdiskparams() IMkdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskparamsContext)
}

func (s *MkdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdiskparams(s)
	}
}

func (s *MkdiskparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdiskparams(s)
	}
}

func (p *ParserParser) Mkdiskparams() (localctx IMkdiskparamsContext) {
	return p.mkdiskparams(0)
}

func (p *ParserParser) mkdiskparams(_p int) (localctx IMkdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ParserParserRULE_mkdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)

		var _x = p.Mkdiskparam()

		localctx.(*MkdiskparamsContext).p = _x
	}
	localctx.(*MkdiskparamsContext).result = map[string]string{"fit": "FF", "unit": "M", localctx.(*MkdiskparamsContext).GetP().GetResult()[0]: localctx.(*MkdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*MkdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkdiskparams)
			p.SetState(88)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(89)

				var _x = p.Mkdiskparam()

				localctx.(*MkdiskparamsContext).p = _x
			}
			localctx.(*MkdiskparamsContext).SetResult(localctx.(*MkdiskparamsContext).GetL().GetResult())
			localctx.(*MkdiskparamsContext).result[localctx.(*MkdiskparamsContext).GetP().GetResult()[0]] = localctx.(*MkdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskparamContext is an interface to support dynamic dispatch.
type IMkdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	TK_unit() antlr.TerminalNode

	// IsMkdiskparamContext differentiates from other interfaces.
	IsMkdiskparamContext()
}

type MkdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkdiskparamContext() *MkdiskparamContext {
	var p = new(MkdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
	return p
}

func InitEmptyMkdiskparamContext(p *MkdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkdiskparam
}

func (*MkdiskparamContext) IsMkdiskparamContext() {}

func NewMkdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskparamContext {
	var p = new(MkdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkdiskparam

	return p
}

func (s *MkdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkdiskparamContext) GetResult() []string { return s.result }

func (s *MkdiskparamContext) SetResult(v []string) { s.result = v }

func (s *MkdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *MkdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *MkdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *MkdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *MkdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *MkdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *MkdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdiskparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkdiskparam(s)
	}
}

func (s *MkdiskparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkdiskparam(s)
	}
}

func (p *ParserParser) Mkdiskparam() (localctx IMkdiskparamContext) {
	localctx = NewMkdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserParserRULE_mkdiskparam)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*MkdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*MkdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*MkdiskparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*MkdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*MkdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*MkdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkdiskparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmdiskContext is an interface to support dynamic dispatch.
type IRmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Rmdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Rmdisk)

	// Getter signatures
	RW_driveletter() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_rmdisk() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsRmdiskContext differentiates from other interfaces.
	IsRmdiskContext()
}

type RmdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rmdisk
	r      antlr.Token
	p      antlr.Token
}

func NewEmptyRmdiskContext() *RmdiskContext {
	var p = new(RmdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
	return p
}

func InitEmptyRmdiskContext(p *RmdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rmdisk
}

func (*RmdiskContext) IsRmdiskContext() {}

func NewRmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmdiskContext {
	var p = new(RmdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rmdisk

	return p
}

func (s *RmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *RmdiskContext) GetR() antlr.Token { return s.r }

func (s *RmdiskContext) GetP() antlr.Token { return s.p }

func (s *RmdiskContext) SetR(v antlr.Token) { s.r = v }

func (s *RmdiskContext) SetP(v antlr.Token) { s.p = v }

func (s *RmdiskContext) GetResult() *commands.Rmdisk { return s.result }

func (s *RmdiskContext) SetResult(v *commands.Rmdisk) { s.result = v }

func (s *RmdiskContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *RmdiskContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RmdiskContext) RW_rmdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rmdisk, 0)
}

func (s *RmdiskContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *RmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RmdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRmdisk(s)
	}
}

func (s *RmdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRmdisk(s)
	}
}

func (p *ParserParser) Rmdisk() (localctx IRmdiskContext) {
	localctx = NewRmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserParserRULE_rmdisk)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*RmdiskContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{"driveletter": (func() string {
			if localctx.(*RmdiskContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*RmdiskContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RmdiskContext).result = commands.NewRmdisk((func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RmdiskContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RmdiskContext).GetR().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepContext is an interface to support dynamic dispatch.
type IRepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r token.
	GetR() antlr.Token

	// SetR sets the r token.
	SetR(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IRepparamsContext

	// SetP sets the p rule contexts.
	SetP(IRepparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Rep

	// SetResult sets the result attribute.
	SetResult(*commands.Rep)

	// Getter signatures
	RW_rep() antlr.TerminalNode
	Repparams() IRepparamsContext

	// IsRepContext differentiates from other interfaces.
	IsRepContext()
}

type RepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Rep
	r      antlr.Token
	p      IRepparamsContext
}

func NewEmptyRepContext() *RepContext {
	var p = new(RepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rep
	return p
}

func InitEmptyRepContext(p *RepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_rep
}

func (*RepContext) IsRepContext() {}

func NewRepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepContext {
	var p = new(RepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_rep

	return p
}

func (s *RepContext) GetParser() antlr.Parser { return s.parser }

func (s *RepContext) GetR() antlr.Token { return s.r }

func (s *RepContext) SetR(v antlr.Token) { s.r = v }

func (s *RepContext) GetP() IRepparamsContext { return s.p }

func (s *RepContext) SetP(v IRepparamsContext) { s.p = v }

func (s *RepContext) GetResult() *commands.Rep { return s.result }

func (s *RepContext) SetResult(v *commands.Rep) { s.result = v }

func (s *RepContext) RW_rep() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_rep, 0)
}

func (s *RepContext) Repparams() IRepparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamsContext)
}

func (s *RepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRep(s)
	}
}

func (s *RepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRep(s)
	}
}

func (p *ParserParser) Rep() (localctx IRepContext) {
	localctx = NewRepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserParserRULE_rep)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)

			var _m = p.Match(ParserParserRW_rep)

			localctx.(*RepContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)

			var _x = p.repparams(0)

			localctx.(*RepContext).p = _x
		}
		localctx.(*RepContext).result = commands.NewRep((func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetColumn()
			}
		}()), localctx.(*RepContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)

			var _m = p.Match(ParserParserRW_rep)

			localctx.(*RepContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepContext).result = commands.NewRep((func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetLine()
			}
		}()), (func() int {
			if localctx.(*RepContext).GetR() == nil {
				return 0
			} else {
				return localctx.(*RepContext).GetR().GetColumn()
			}
		}()), map[string]string{})

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepparamsContext is an interface to support dynamic dispatch.
type IRepparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IRepparamsContext

	// GetP returns the p rule contexts.
	GetP() IRepparamContext

	// SetL sets the l rule contexts.
	SetL(IRepparamsContext)

	// SetP sets the p rule contexts.
	SetP(IRepparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Repparam() IRepparamContext
	Repparams() IRepparamsContext

	// IsRepparamsContext differentiates from other interfaces.
	IsRepparamsContext()
}

type RepparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IRepparamsContext
	p      IRepparamContext
}

func NewEmptyRepparamsContext() *RepparamsContext {
	var p = new(RepparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparams
	return p
}

func InitEmptyRepparamsContext(p *RepparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparams
}

func (*RepparamsContext) IsRepparamsContext() {}

func NewRepparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepparamsContext {
	var p = new(RepparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_repparams

	return p
}

func (s *RepparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *RepparamsContext) GetL() IRepparamsContext { return s.l }

func (s *RepparamsContext) GetP() IRepparamContext { return s.p }

func (s *RepparamsContext) SetL(v IRepparamsContext) { s.l = v }

func (s *RepparamsContext) SetP(v IRepparamContext) { s.p = v }

func (s *RepparamsContext) GetResult() map[string]string { return s.result }

func (s *RepparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *RepparamsContext) Repparam() IRepparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamContext)
}

func (s *RepparamsContext) Repparams() IRepparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepparamsContext)
}

func (s *RepparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRepparams(s)
	}
}

func (s *RepparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRepparams(s)
	}
}

func (p *ParserParser) Repparams() (localctx IRepparamsContext) {
	return p.repparams(0)
}

func (p *ParserParser) repparams(_p int) (localctx IRepparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRepparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRepparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ParserParserRULE_repparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)

		var _x = p.Repparam()

		localctx.(*RepparamsContext).p = _x
	}
	localctx.(*RepparamsContext).result = map[string]string{localctx.(*RepparamsContext).GetP().GetResult()[0]: localctx.(*RepparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRepparamsContext(p, _parentctx, _parentState)
			localctx.(*RepparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_repparams)
			p.SetState(132)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(133)

				var _x = p.Repparam()

				localctx.(*RepparamsContext).p = _x
			}
			localctx.(*RepparamsContext).SetResult(localctx.(*RepparamsContext).GetL().GetResult())
			localctx.(*RepparamsContext).result[localctx.(*RepparamsContext).GetP().GetResult()[0]] = localctx.(*RepparamsContext).GetP().GetResult()[1]

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepparamContext is an interface to support dynamic dispatch.
type IRepparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// GetV1 returns the v1 rule contexts.
	GetV1() INameContext

	// SetV1 sets the v1 rule contexts.
	SetV1(INameContext)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_name() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	Name() INameContext
	RW_path() antlr.TerminalNode
	TK_path() antlr.TerminalNode
	RW_ruta() antlr.TerminalNode
	RW_id() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsRepparamContext differentiates from other interfaces.
	IsRepparamContext()
}

type RepparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     INameContext
	v2     antlr.Token
	v4     antlr.Token
	v3     antlr.Token
}

func NewEmptyRepparamContext() *RepparamContext {
	var p = new(RepparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparam
	return p
}

func InitEmptyRepparamContext(p *RepparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_repparam
}

func (*RepparamContext) IsRepparamContext() {}

func NewRepparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepparamContext {
	var p = new(RepparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_repparam

	return p
}

func (s *RepparamContext) GetParser() antlr.Parser { return s.parser }

func (s *RepparamContext) GetV2() antlr.Token { return s.v2 }

func (s *RepparamContext) GetV4() antlr.Token { return s.v4 }

func (s *RepparamContext) GetV3() antlr.Token { return s.v3 }

func (s *RepparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *RepparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *RepparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *RepparamContext) GetV1() INameContext { return s.v1 }

func (s *RepparamContext) SetV1(v INameContext) { s.v1 = v }

func (s *RepparamContext) GetResult() []string { return s.result }

func (s *RepparamContext) SetResult(v []string) { s.result = v }

func (s *RepparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *RepparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *RepparamContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *RepparamContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *RepparamContext) TK_path() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_path, 0)
}

func (s *RepparamContext) RW_ruta() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_ruta, 0)
}

func (s *RepparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *RepparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *RepparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRepparam(s)
	}
}

func (s *RepparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRepparam(s)
	}
}

func (p *ParserParser) Repparam() (localctx IRepparamContext) {
	localctx = NewRepparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_repparam)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)

			var _x = p.Name()

			localctx.(*RepparamContext).v1 = _x
		}
		localctx.(*RepparamContext).result = []string{"name", (func() string {
			if localctx.(*RepparamContext).GetV1() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*RepparamContext).GetV1().GetStart(), localctx.(*RepparamContext).v1.GetStop())
			}
		}())}

	case ParserParserRW_path:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RepparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"path", (func() string {
			if localctx.(*RepparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_ruta:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Match(ParserParserRW_ruta)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*RepparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"ruta", (func() string {
			if localctx.(*RepparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV4().GetText()
			}
		}())}

	case ParserParserRW_id:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*RepparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*RepparamContext).result = []string{"id", (func() string {
			if localctx.(*RepparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).GetV3().GetText()
			}
		}())}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() string

	// SetResult sets the result attribute.
	SetResult(string)

	// Getter signatures
	RW_mbr() antlr.TerminalNode
	RW_disk() antlr.TerminalNode
	RW_inode() antlr.TerminalNode
	RW_journaling() antlr.TerminalNode
	RW_block() antlr.TerminalNode
	RW_bm_inode() antlr.TerminalNode
	RW_bm_block() antlr.TerminalNode
	RW_tree() antlr.TerminalNode
	RW_sb() antlr.TerminalNode
	RW_file() antlr.TerminalNode
	RW_ls() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result string
	n      antlr.Token
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) GetN() antlr.Token { return s.n }

func (s *NameContext) SetN(v antlr.Token) { s.n = v }

func (s *NameContext) GetResult() string { return s.result }

func (s *NameContext) SetResult(v string) { s.result = v }

func (s *NameContext) RW_mbr() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mbr, 0)
}

func (s *NameContext) RW_disk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_disk, 0)
}

func (s *NameContext) RW_inode() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_inode, 0)
}

func (s *NameContext) RW_journaling() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_journaling, 0)
}

func (s *NameContext) RW_block() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_block, 0)
}

func (s *NameContext) RW_bm_inode() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_bm_inode, 0)
}

func (s *NameContext) RW_bm_block() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_bm_block, 0)
}

func (s *NameContext) RW_tree() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_tree, 0)
}

func (s *NameContext) RW_sb() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_sb, 0)
}

func (s *NameContext) RW_file() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_file, 0)
}

func (s *NameContext) RW_ls() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_ls, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *ParserParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParserParserRULE_name)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mbr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)

			var _m = p.Match(ParserParserRW_mbr)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_disk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)

			var _m = p.Match(ParserParserRW_disk)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_inode:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)

			var _m = p.Match(ParserParserRW_inode)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_journaling:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)

			var _m = p.Match(ParserParserRW_journaling)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_block:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(168)

			var _m = p.Match(ParserParserRW_block)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_bm_inode:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(170)

			var _m = p.Match(ParserParserRW_bm_inode)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_bm_block:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(172)

			var _m = p.Match(ParserParserRW_bm_block)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_tree:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(174)

			var _m = p.Match(ParserParserRW_tree)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_sb:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(176)

			var _m = p.Match(ParserParserRW_sb)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_file:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(178)

			var _m = p.Match(ParserParserRW_file)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	case ParserParserRW_ls:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(180)

			var _m = p.Match(ParserParserRW_ls)

			localctx.(*NameContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*NameContext).result = (func() string {
			if localctx.(*NameContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*NameContext).GetN().GetText()
			}
		}())

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentaryContext is an interface to support dynamic dispatch.
type ICommentaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c token.
	GetC() antlr.Token

	// SetC sets the c token.
	SetC(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Commentary

	// SetResult sets the result attribute.
	SetResult(*commands.Commentary)

	// Getter signatures
	COMMENTARY() antlr.TerminalNode

	// IsCommentaryContext differentiates from other interfaces.
	IsCommentaryContext()
}

type CommentaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Commentary
	c      antlr.Token
}

func NewEmptyCommentaryContext() *CommentaryContext {
	var p = new(CommentaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commentary
	return p
}

func InitEmptyCommentaryContext(p *CommentaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_commentary
}

func (*CommentaryContext) IsCommentaryContext() {}

func NewCommentaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentaryContext {
	var p = new(CommentaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_commentary

	return p
}

func (s *CommentaryContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentaryContext) GetC() antlr.Token { return s.c }

func (s *CommentaryContext) SetC(v antlr.Token) { s.c = v }

func (s *CommentaryContext) GetResult() *commands.Commentary { return s.result }

func (s *CommentaryContext) SetResult(v *commands.Commentary) { s.result = v }

func (s *CommentaryContext) COMMENTARY() antlr.TerminalNode {
	return s.GetToken(ParserParserCOMMENTARY, 0)
}

func (s *CommentaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCommentary(s)
	}
}

func (s *CommentaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCommentary(s)
	}
}

func (p *ParserParser) Commentary() (localctx ICommentaryContext) {
	localctx = NewCommentaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParserParserRULE_commentary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)

		var _m = p.Match(ParserParserCOMMENTARY)

		localctx.(*CommentaryContext).c = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CommentaryContext).result = commands.NewCommentary((func() int {
		if localctx.(*CommentaryContext).GetC() == nil {
			return 0
		} else {
			return localctx.(*CommentaryContext).GetC().GetLine()
		}
	}()), (func() int {
		if localctx.(*CommentaryContext).GetC() == nil {
			return 0
		} else {
			return localctx.(*CommentaryContext).GetC().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CommentaryContext).GetC() == nil {
			return ""
		} else {
			return localctx.(*CommentaryContext).GetC().GetText()
		}
	}()))

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExitContext is an interface to support dynamic dispatch.
type IExitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_exit() antlr.TerminalNode

	// IsExitContext differentiates from other interfaces.
	IsExitContext()
}

type ExitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExitContext() *ExitContext {
	var p = new(ExitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exit
	return p
}

func InitEmptyExitContext(p *ExitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exit
}

func (*ExitContext) IsExitContext() {}

func NewExitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExitContext {
	var p = new(ExitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_exit

	return p
}

func (s *ExitContext) GetParser() antlr.Parser { return s.parser }

func (s *ExitContext) RW_exit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_exit, 0)
}

func (s *ExitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterExit(s)
	}
}

func (s *ExitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitExit(s)
	}
}

func (p *ParserParser) Exit() (localctx IExitContext) {
	localctx = NewExitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParserParserRULE_exit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(ParserParserRW_exit)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	os.Exit(1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CommandsContext = nil
		if localctx != nil {
			t = localctx.(*CommandsContext)
		}
		return p.Commands_Sempred(t, predIndex)

	case 5:
		var t *MkdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkdiskparamsContext)
		}
		return p.Mkdiskparams_Sempred(t, predIndex)

	case 9:
		var t *RepparamsContext = nil
		if localctx != nil {
			t = localctx.(*RepparamsContext)
		}
		return p.Repparams_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserParser) Commands_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Repparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
