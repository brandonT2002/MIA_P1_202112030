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
		"mkdiskparam", "rmdisk", "fdisk", "fdiskparams", "fdiskparam", "mount",
		"mountparams", "mountparam", "unmount", "mkfs", "mkfsparams", "mkfsparam",
		"login", "loginparams", "loginparam", "logout", "pause", "rep", "repparams",
		"repparam", "name", "commentary", "exit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 74, 417, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		3, 0, 65, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 75,
		8, 1, 10, 1, 12, 1, 78, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 117, 8, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 126, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 134, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 144, 8, 5, 10, 5, 12, 5, 147, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 161, 8, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 170, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 178, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 188, 8, 9, 10, 9, 12, 9, 191, 9, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 225, 8, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 233, 8, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 243, 8, 12, 10, 12,
		12, 12, 246, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 256, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 265, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 273,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 283,
		8, 16, 10, 16, 12, 16, 286, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 300, 8, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 308, 8, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 318, 8, 19, 10, 19, 12, 19,
		321, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 3, 20, 339, 8, 20,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 353, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 5, 24, 363, 8, 24, 10, 24, 12, 24, 366, 9, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 385, 8, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 409,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 0, 7, 2, 10, 18,
		24, 32, 38, 48, 29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0, 0, 444,
		0, 64, 1, 0, 0, 0, 2, 66, 1, 0, 0, 0, 4, 116, 1, 0, 0, 0, 6, 125, 1, 0,
		0, 0, 8, 133, 1, 0, 0, 0, 10, 135, 1, 0, 0, 0, 12, 160, 1, 0, 0, 0, 14,
		169, 1, 0, 0, 0, 16, 177, 1, 0, 0, 0, 18, 179, 1, 0, 0, 0, 20, 224, 1,
		0, 0, 0, 22, 232, 1, 0, 0, 0, 24, 234, 1, 0, 0, 0, 26, 255, 1, 0, 0, 0,
		28, 264, 1, 0, 0, 0, 30, 272, 1, 0, 0, 0, 32, 274, 1, 0, 0, 0, 34, 299,
		1, 0, 0, 0, 36, 307, 1, 0, 0, 0, 38, 309, 1, 0, 0, 0, 40, 338, 1, 0, 0,
		0, 42, 340, 1, 0, 0, 0, 44, 343, 1, 0, 0, 0, 46, 352, 1, 0, 0, 0, 48, 354,
		1, 0, 0, 0, 50, 384, 1, 0, 0, 0, 52, 408, 1, 0, 0, 0, 54, 410, 1, 0, 0,
		0, 56, 413, 1, 0, 0, 0, 58, 59, 3, 2, 1, 0, 59, 60, 5, 0, 0, 1, 60, 61,
		6, 0, -1, 0, 61, 65, 1, 0, 0, 0, 62, 63, 5, 0, 0, 1, 63, 65, 6, 0, -1,
		0, 64, 58, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 1, 1, 0, 0, 0, 66, 67, 6,
		1, -1, 0, 67, 68, 3, 4, 2, 0, 68, 69, 6, 1, -1, 0, 69, 76, 1, 0, 0, 0,
		70, 71, 10, 2, 0, 0, 71, 72, 3, 4, 2, 0, 72, 73, 6, 1, -1, 0, 73, 75, 1,
		0, 0, 0, 74, 70, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 3, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 3, 6, 3,
		0, 80, 81, 6, 2, -1, 0, 81, 117, 1, 0, 0, 0, 82, 83, 3, 8, 4, 0, 83, 84,
		6, 2, -1, 0, 84, 117, 1, 0, 0, 0, 85, 86, 3, 14, 7, 0, 86, 87, 6, 2, -1,
		0, 87, 117, 1, 0, 0, 0, 88, 89, 3, 16, 8, 0, 89, 90, 6, 2, -1, 0, 90, 117,
		1, 0, 0, 0, 91, 92, 3, 22, 11, 0, 92, 93, 6, 2, -1, 0, 93, 117, 1, 0, 0,
		0, 94, 95, 3, 28, 14, 0, 95, 96, 6, 2, -1, 0, 96, 117, 1, 0, 0, 0, 97,
		98, 3, 30, 15, 0, 98, 99, 6, 2, -1, 0, 99, 117, 1, 0, 0, 0, 100, 101, 3,
		36, 18, 0, 101, 102, 6, 2, -1, 0, 102, 117, 1, 0, 0, 0, 103, 104, 3, 42,
		21, 0, 104, 105, 6, 2, -1, 0, 105, 117, 1, 0, 0, 0, 106, 107, 3, 44, 22,
		0, 107, 108, 6, 2, -1, 0, 108, 117, 1, 0, 0, 0, 109, 110, 3, 46, 23, 0,
		110, 111, 6, 2, -1, 0, 111, 117, 1, 0, 0, 0, 112, 113, 3, 54, 27, 0, 113,
		114, 6, 2, -1, 0, 114, 117, 1, 0, 0, 0, 115, 117, 3, 56, 28, 0, 116, 79,
		1, 0, 0, 0, 116, 82, 1, 0, 0, 0, 116, 85, 1, 0, 0, 0, 116, 88, 1, 0, 0,
		0, 116, 91, 1, 0, 0, 0, 116, 94, 1, 0, 0, 0, 116, 97, 1, 0, 0, 0, 116,
		100, 1, 0, 0, 0, 116, 103, 1, 0, 0, 0, 116, 106, 1, 0, 0, 0, 116, 109,
		1, 0, 0, 0, 116, 112, 1, 0, 0, 0, 116, 115, 1, 0, 0, 0, 117, 5, 1, 0, 0,
		0, 118, 119, 5, 28, 0, 0, 119, 120, 5, 56, 0, 0, 120, 121, 5, 70, 0, 0,
		121, 122, 5, 69, 0, 0, 122, 126, 6, 3, -1, 0, 123, 124, 5, 28, 0, 0, 124,
		126, 6, 3, -1, 0, 125, 118, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 7, 1,
		0, 0, 0, 127, 128, 5, 1, 0, 0, 128, 129, 3, 10, 5, 0, 129, 130, 6, 4, -1,
		0, 130, 134, 1, 0, 0, 0, 131, 132, 5, 1, 0, 0, 132, 134, 6, 4, -1, 0, 133,
		127, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 9, 1, 0, 0, 0, 135, 136, 6,
		5, -1, 0, 136, 137, 3, 12, 6, 0, 137, 138, 6, 5, -1, 0, 138, 145, 1, 0,
		0, 0, 139, 140, 10, 2, 0, 0, 140, 141, 3, 12, 6, 0, 141, 142, 6, 5, -1,
		0, 142, 144, 1, 0, 0, 0, 143, 139, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145,
		143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 11, 1, 0, 0, 0, 147, 145, 1,
		0, 0, 0, 148, 149, 5, 43, 0, 0, 149, 150, 5, 70, 0, 0, 150, 151, 5, 67,
		0, 0, 151, 161, 6, 6, -1, 0, 152, 153, 5, 44, 0, 0, 153, 154, 5, 70, 0,
		0, 154, 155, 5, 63, 0, 0, 155, 161, 6, 6, -1, 0, 156, 157, 5, 45, 0, 0,
		157, 158, 5, 70, 0, 0, 158, 159, 5, 64, 0, 0, 159, 161, 6, 6, -1, 0, 160,
		148, 1, 0, 0, 0, 160, 152, 1, 0, 0, 0, 160, 156, 1, 0, 0, 0, 161, 13, 1,
		0, 0, 0, 162, 163, 5, 2, 0, 0, 163, 164, 5, 46, 0, 0, 164, 165, 5, 70,
		0, 0, 165, 166, 5, 68, 0, 0, 166, 170, 6, 7, -1, 0, 167, 168, 5, 2, 0,
		0, 168, 170, 6, 7, -1, 0, 169, 162, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170,
		15, 1, 0, 0, 0, 171, 172, 5, 3, 0, 0, 172, 173, 3, 18, 9, 0, 173, 174,
		6, 8, -1, 0, 174, 178, 1, 0, 0, 0, 175, 176, 5, 3, 0, 0, 176, 178, 6, 8,
		-1, 0, 177, 171, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 17, 1, 0, 0, 0,
		179, 180, 6, 9, -1, 0, 180, 181, 3, 20, 10, 0, 181, 182, 6, 9, -1, 0, 182,
		189, 1, 0, 0, 0, 183, 184, 10, 2, 0, 0, 184, 185, 3, 20, 10, 0, 185, 186,
		6, 9, -1, 0, 186, 188, 1, 0, 0, 0, 187, 183, 1, 0, 0, 0, 188, 191, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 19, 1, 0, 0, 0,
		191, 189, 1, 0, 0, 0, 192, 193, 5, 43, 0, 0, 193, 194, 5, 70, 0, 0, 194,
		195, 5, 67, 0, 0, 195, 225, 6, 10, -1, 0, 196, 197, 5, 46, 0, 0, 197, 198,
		5, 70, 0, 0, 198, 199, 5, 68, 0, 0, 199, 225, 6, 10, -1, 0, 200, 201, 5,
		47, 0, 0, 201, 202, 5, 70, 0, 0, 202, 203, 5, 68, 0, 0, 203, 225, 6, 10,
		-1, 0, 204, 205, 5, 45, 0, 0, 205, 206, 5, 70, 0, 0, 206, 207, 5, 64, 0,
		0, 207, 225, 6, 10, -1, 0, 208, 209, 5, 48, 0, 0, 209, 210, 5, 70, 0, 0,
		210, 211, 5, 65, 0, 0, 211, 225, 6, 10, -1, 0, 212, 213, 5, 44, 0, 0, 213,
		214, 5, 70, 0, 0, 214, 215, 5, 63, 0, 0, 215, 225, 6, 10, -1, 0, 216, 217,
		5, 49, 0, 0, 217, 218, 5, 70, 0, 0, 218, 219, 5, 42, 0, 0, 219, 225, 6,
		10, -1, 0, 220, 221, 5, 50, 0, 0, 221, 222, 5, 70, 0, 0, 222, 223, 5, 67,
		0, 0, 223, 225, 6, 10, -1, 0, 224, 192, 1, 0, 0, 0, 224, 196, 1, 0, 0,
		0, 224, 200, 1, 0, 0, 0, 224, 204, 1, 0, 0, 0, 224, 208, 1, 0, 0, 0, 224,
		212, 1, 0, 0, 0, 224, 216, 1, 0, 0, 0, 224, 220, 1, 0, 0, 0, 225, 21, 1,
		0, 0, 0, 226, 227, 5, 4, 0, 0, 227, 228, 3, 24, 12, 0, 228, 229, 6, 11,
		-1, 0, 229, 233, 1, 0, 0, 0, 230, 231, 5, 4, 0, 0, 231, 233, 6, 11, -1,
		0, 232, 226, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 23, 1, 0, 0, 0, 234,
		235, 6, 12, -1, 0, 235, 236, 3, 26, 13, 0, 236, 237, 6, 12, -1, 0, 237,
		244, 1, 0, 0, 0, 238, 239, 10, 2, 0, 0, 239, 240, 3, 26, 13, 0, 240, 241,
		6, 12, -1, 0, 241, 243, 1, 0, 0, 0, 242, 238, 1, 0, 0, 0, 243, 246, 1,
		0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 25, 1, 0, 0,
		0, 246, 244, 1, 0, 0, 0, 247, 248, 5, 46, 0, 0, 248, 249, 5, 70, 0, 0,
		249, 250, 5, 68, 0, 0, 250, 256, 6, 13, -1, 0, 251, 252, 5, 47, 0, 0, 252,
		253, 5, 70, 0, 0, 253, 254, 5, 68, 0, 0, 254, 256, 6, 13, -1, 0, 255, 247,
		1, 0, 0, 0, 255, 251, 1, 0, 0, 0, 256, 27, 1, 0, 0, 0, 257, 258, 5, 5,
		0, 0, 258, 259, 5, 51, 0, 0, 259, 260, 5, 70, 0, 0, 260, 261, 5, 68, 0,
		0, 261, 265, 6, 14, -1, 0, 262, 263, 5, 5, 0, 0, 263, 265, 6, 14, -1, 0,
		264, 257, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 265, 29, 1, 0, 0, 0, 266, 267,
		5, 6, 0, 0, 267, 268, 3, 32, 16, 0, 268, 269, 6, 15, -1, 0, 269, 273, 1,
		0, 0, 0, 270, 271, 5, 6, 0, 0, 271, 273, 6, 15, -1, 0, 272, 266, 1, 0,
		0, 0, 272, 270, 1, 0, 0, 0, 273, 31, 1, 0, 0, 0, 274, 275, 6, 16, -1, 0,
		275, 276, 3, 34, 17, 0, 276, 277, 6, 16, -1, 0, 277, 284, 1, 0, 0, 0, 278,
		279, 10, 2, 0, 0, 279, 280, 3, 34, 17, 0, 280, 281, 6, 16, -1, 0, 281,
		283, 1, 0, 0, 0, 282, 278, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282,
		1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 33, 1, 0, 0, 0, 286, 284, 1, 0,
		0, 0, 287, 288, 5, 51, 0, 0, 288, 289, 5, 70, 0, 0, 289, 290, 5, 68, 0,
		0, 290, 300, 6, 17, -1, 0, 291, 292, 5, 48, 0, 0, 292, 293, 5, 70, 0, 0,
		293, 294, 5, 42, 0, 0, 294, 300, 6, 17, -1, 0, 295, 296, 5, 52, 0, 0, 296,
		297, 5, 70, 0, 0, 297, 298, 5, 66, 0, 0, 298, 300, 6, 17, -1, 0, 299, 287,
		1, 0, 0, 0, 299, 291, 1, 0, 0, 0, 299, 295, 1, 0, 0, 0, 300, 35, 1, 0,
		0, 0, 301, 302, 5, 7, 0, 0, 302, 303, 3, 38, 19, 0, 303, 304, 6, 18, -1,
		0, 304, 308, 1, 0, 0, 0, 305, 306, 5, 7, 0, 0, 306, 308, 6, 18, -1, 0,
		307, 301, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 308, 37, 1, 0, 0, 0, 309, 310,
		6, 19, -1, 0, 310, 311, 3, 40, 20, 0, 311, 312, 6, 19, -1, 0, 312, 319,
		1, 0, 0, 0, 313, 314, 10, 2, 0, 0, 314, 315, 3, 40, 20, 0, 315, 316, 6,
		19, -1, 0, 316, 318, 1, 0, 0, 0, 317, 313, 1, 0, 0, 0, 318, 321, 1, 0,
		0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 39, 1, 0, 0, 0,
		321, 319, 1, 0, 0, 0, 322, 323, 5, 53, 0, 0, 323, 324, 5, 70, 0, 0, 324,
		325, 5, 68, 0, 0, 325, 339, 6, 20, -1, 0, 326, 327, 5, 54, 0, 0, 327, 328,
		5, 70, 0, 0, 328, 329, 5, 68, 0, 0, 329, 339, 6, 20, -1, 0, 330, 331, 5,
		54, 0, 0, 331, 332, 5, 70, 0, 0, 332, 333, 5, 67, 0, 0, 333, 339, 6, 20,
		-1, 0, 334, 335, 5, 51, 0, 0, 335, 336, 5, 70, 0, 0, 336, 337, 5, 68, 0,
		0, 337, 339, 6, 20, -1, 0, 338, 322, 1, 0, 0, 0, 338, 326, 1, 0, 0, 0,
		338, 330, 1, 0, 0, 0, 338, 334, 1, 0, 0, 0, 339, 41, 1, 0, 0, 0, 340, 341,
		5, 8, 0, 0, 341, 342, 6, 21, -1, 0, 342, 43, 1, 0, 0, 0, 343, 344, 5, 25,
		0, 0, 344, 345, 6, 22, -1, 0, 345, 45, 1, 0, 0, 0, 346, 347, 5, 29, 0,
		0, 347, 348, 3, 48, 24, 0, 348, 349, 6, 23, -1, 0, 349, 353, 1, 0, 0, 0,
		350, 351, 5, 29, 0, 0, 351, 353, 6, 23, -1, 0, 352, 346, 1, 0, 0, 0, 352,
		350, 1, 0, 0, 0, 353, 47, 1, 0, 0, 0, 354, 355, 6, 24, -1, 0, 355, 356,
		3, 50, 25, 0, 356, 357, 6, 24, -1, 0, 357, 364, 1, 0, 0, 0, 358, 359, 10,
		2, 0, 0, 359, 360, 3, 50, 25, 0, 360, 361, 6, 24, -1, 0, 361, 363, 1, 0,
		0, 0, 362, 358, 1, 0, 0, 0, 363, 366, 1, 0, 0, 0, 364, 362, 1, 0, 0, 0,
		364, 365, 1, 0, 0, 0, 365, 49, 1, 0, 0, 0, 366, 364, 1, 0, 0, 0, 367, 368,
		5, 47, 0, 0, 368, 369, 5, 70, 0, 0, 369, 370, 3, 52, 26, 0, 370, 371, 6,
		25, -1, 0, 371, 385, 1, 0, 0, 0, 372, 373, 5, 56, 0, 0, 373, 374, 5, 70,
		0, 0, 374, 375, 5, 69, 0, 0, 375, 385, 6, 25, -1, 0, 376, 377, 5, 62, 0,
		0, 377, 378, 5, 70, 0, 0, 378, 379, 5, 69, 0, 0, 379, 385, 6, 25, -1, 0,
		380, 381, 5, 51, 0, 0, 381, 382, 5, 70, 0, 0, 382, 383, 5, 68, 0, 0, 383,
		385, 6, 25, -1, 0, 384, 367, 1, 0, 0, 0, 384, 372, 1, 0, 0, 0, 384, 376,
		1, 0, 0, 0, 384, 380, 1, 0, 0, 0, 385, 51, 1, 0, 0, 0, 386, 387, 5, 31,
		0, 0, 387, 409, 6, 26, -1, 0, 388, 389, 5, 32, 0, 0, 389, 409, 6, 26, -1,
		0, 390, 391, 5, 33, 0, 0, 391, 409, 6, 26, -1, 0, 392, 393, 5, 34, 0, 0,
		393, 409, 6, 26, -1, 0, 394, 395, 5, 35, 0, 0, 395, 409, 6, 26, -1, 0,
		396, 397, 5, 36, 0, 0, 397, 409, 6, 26, -1, 0, 398, 399, 5, 37, 0, 0, 399,
		409, 6, 26, -1, 0, 400, 401, 5, 38, 0, 0, 401, 409, 6, 26, -1, 0, 402,
		403, 5, 39, 0, 0, 403, 409, 6, 26, -1, 0, 404, 405, 5, 40, 0, 0, 405, 409,
		6, 26, -1, 0, 406, 407, 5, 41, 0, 0, 407, 409, 6, 26, -1, 0, 408, 386,
		1, 0, 0, 0, 408, 388, 1, 0, 0, 0, 408, 390, 1, 0, 0, 0, 408, 392, 1, 0,
		0, 0, 408, 394, 1, 0, 0, 0, 408, 396, 1, 0, 0, 0, 408, 398, 1, 0, 0, 0,
		408, 400, 1, 0, 0, 0, 408, 402, 1, 0, 0, 0, 408, 404, 1, 0, 0, 0, 408,
		406, 1, 0, 0, 0, 409, 53, 1, 0, 0, 0, 410, 411, 5, 72, 0, 0, 411, 412,
		6, 27, -1, 0, 412, 55, 1, 0, 0, 0, 413, 414, 5, 30, 0, 0, 414, 415, 6,
		28, -1, 0, 415, 57, 1, 0, 0, 0, 25, 64, 76, 116, 125, 133, 145, 160, 169,
		177, 189, 224, 232, 244, 255, 264, 272, 284, 299, 307, 319, 338, 352, 364,
		384, 408,
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
	ParserParserRULE_fdisk        = 8
	ParserParserRULE_fdiskparams  = 9
	ParserParserRULE_fdiskparam   = 10
	ParserParserRULE_mount        = 11
	ParserParserRULE_mountparams  = 12
	ParserParserRULE_mountparam   = 13
	ParserParserRULE_unmount      = 14
	ParserParserRULE_mkfs         = 15
	ParserParserRULE_mkfsparams   = 16
	ParserParserRULE_mkfsparam    = 17
	ParserParserRULE_login        = 18
	ParserParserRULE_loginparams  = 19
	ParserParserRULE_loginparam   = 20
	ParserParserRULE_logout       = 21
	ParserParserRULE_pause        = 22
	ParserParserRULE_rep          = 23
	ParserParserRULE_repparams    = 24
	ParserParserRULE_repparam     = 25
	ParserParserRULE_name         = 26
	ParserParserRULE_commentary   = 27
	ParserParserRULE_exit         = 28
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
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mkdisk, ParserParserRW_rmdisk, ParserParserRW_fdisk, ParserParserRW_mount, ParserParserRW_unmount, ParserParserRW_mkfs, ParserParserRW_login, ParserParserRW_logout, ParserParserRW_pause, ParserParserRW_execute, ParserParserRW_rep, ParserParserRW_exit, ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)

			var _x = p.commands(0)

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(59)
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
			p.SetState(62)
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
		p.SetState(67)

		var _x = p.Command()

		localctx.(*CommandsContext).c = _x
	}
	localctx.(*CommandsContext).result = []interfaces.Command{localctx.(*CommandsContext).GetC().GetResult()}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(76)
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
			p.SetState(70)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(71)

				var _x = p.Command()

				localctx.(*CommandsContext).c = _x
			}
			localctx.(*CommandsContext).SetResult(localctx.(*CommandsContext).GetL().GetResult())
			localctx.(*CommandsContext).result = append(localctx.(*CommandsContext).result, localctx.(*CommandsContext).GetC().GetResult())

		}
		p.SetState(78)
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

	// GetC4 returns the c4 rule contexts.
	GetC4() IFdiskContext

	// GetC5 returns the c5 rule contexts.
	GetC5() IMountContext

	// GetC6 returns the c6 rule contexts.
	GetC6() IUnmountContext

	// GetC7 returns the c7 rule contexts.
	GetC7() IMkfsContext

	// GetC8 returns the c8 rule contexts.
	GetC8() ILoginContext

	// GetC9 returns the c9 rule contexts.
	GetC9() ILogoutContext

	// GetC19 returns the c19 rule contexts.
	GetC19() IPauseContext

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

	// SetC4 sets the c4 rule contexts.
	SetC4(IFdiskContext)

	// SetC5 sets the c5 rule contexts.
	SetC5(IMountContext)

	// SetC6 sets the c6 rule contexts.
	SetC6(IUnmountContext)

	// SetC7 sets the c7 rule contexts.
	SetC7(IMkfsContext)

	// SetC8 sets the c8 rule contexts.
	SetC8(ILoginContext)

	// SetC9 sets the c9 rule contexts.
	SetC9(ILogoutContext)

	// SetC19 sets the c19 rule contexts.
	SetC19(IPauseContext)

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
	Fdisk() IFdiskContext
	Mount() IMountContext
	Unmount() IUnmountContext
	Mkfs() IMkfsContext
	Login() ILoginContext
	Logout() ILogoutContext
	Pause() IPauseContext
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
	c4     IFdiskContext
	c5     IMountContext
	c6     IUnmountContext
	c7     IMkfsContext
	c8     ILoginContext
	c9     ILogoutContext
	c19    IPauseContext
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

func (s *CommandContext) GetC4() IFdiskContext { return s.c4 }

func (s *CommandContext) GetC5() IMountContext { return s.c5 }

func (s *CommandContext) GetC6() IUnmountContext { return s.c6 }

func (s *CommandContext) GetC7() IMkfsContext { return s.c7 }

func (s *CommandContext) GetC8() ILoginContext { return s.c8 }

func (s *CommandContext) GetC9() ILogoutContext { return s.c9 }

func (s *CommandContext) GetC19() IPauseContext { return s.c19 }

func (s *CommandContext) GetC20() IRepContext { return s.c20 }

func (s *CommandContext) GetC21() ICommentaryContext { return s.c21 }

func (s *CommandContext) SetC1(v IExecuteContext) { s.c1 = v }

func (s *CommandContext) SetC2(v IMkdiskContext) { s.c2 = v }

func (s *CommandContext) SetC3(v IRmdiskContext) { s.c3 = v }

func (s *CommandContext) SetC4(v IFdiskContext) { s.c4 = v }

func (s *CommandContext) SetC5(v IMountContext) { s.c5 = v }

func (s *CommandContext) SetC6(v IUnmountContext) { s.c6 = v }

func (s *CommandContext) SetC7(v IMkfsContext) { s.c7 = v }

func (s *CommandContext) SetC8(v ILoginContext) { s.c8 = v }

func (s *CommandContext) SetC9(v ILogoutContext) { s.c9 = v }

func (s *CommandContext) SetC19(v IPauseContext) { s.c19 = v }

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

func (s *CommandContext) Fdisk() IFdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskContext)
}

func (s *CommandContext) Mount() IMountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountContext)
}

func (s *CommandContext) Unmount() IUnmountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnmountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnmountContext)
}

func (s *CommandContext) Mkfs() IMkfsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsContext)
}

func (s *CommandContext) Login() ILoginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginContext)
}

func (s *CommandContext) Logout() ILogoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogoutContext)
}

func (s *CommandContext) Pause() IPauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPauseContext)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_execute:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)

			var _x = p.Execute()

			localctx.(*CommandContext).c1 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC1().GetResult()

	case ParserParserRW_mkdisk:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)

			var _x = p.Mkdisk()

			localctx.(*CommandContext).c2 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC2().GetResult()

	case ParserParserRW_rmdisk:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(85)

			var _x = p.Rmdisk()

			localctx.(*CommandContext).c3 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC3().GetResult()

	case ParserParserRW_fdisk:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)

			var _x = p.Fdisk()

			localctx.(*CommandContext).c4 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC4().GetResult()

	case ParserParserRW_mount:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)

			var _x = p.Mount()

			localctx.(*CommandContext).c5 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC5().GetResult()

	case ParserParserRW_unmount:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(94)

			var _x = p.Unmount()

			localctx.(*CommandContext).c6 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC6().GetResult()

	case ParserParserRW_mkfs:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(97)

			var _x = p.Mkfs()

			localctx.(*CommandContext).c7 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC7().GetResult()

	case ParserParserRW_login:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)

			var _x = p.Login()

			localctx.(*CommandContext).c8 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC8().GetResult()

	case ParserParserRW_logout:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(103)

			var _x = p.Logout()

			localctx.(*CommandContext).c9 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC9().GetResult()

	case ParserParserRW_pause:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(106)

			var _x = p.Pause()

			localctx.(*CommandContext).c19 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC19().GetResult()

	case ParserParserRW_rep:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(109)

			var _x = p.Rep()

			localctx.(*CommandContext).c20 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC20().GetResult()

	case ParserParserCOMMENTARY:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(112)

			var _x = p.Commentary()

			localctx.(*CommandContext).c21 = _x
		}
		localctx.(*CommandContext).result = localctx.(*CommandContext).GetC21().GetResult()

	case ParserParserRW_exit:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(115)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)

			var _m = p.Match(ParserParserRW_execute)

			localctx.(*ExecuteContext).e = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)

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
			p.SetState(123)

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
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)

			var _m = p.Match(ParserParserRW_mkdisk)

			localctx.(*MkdiskContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)

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
			p.SetState(131)

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
		p.SetState(136)

		var _x = p.Mkdiskparam()

		localctx.(*MkdiskparamsContext).p = _x
	}
	localctx.(*MkdiskparamsContext).result = map[string]string{"fit": "FF", "unit": "M", localctx.(*MkdiskparamsContext).GetP().GetResult()[0]: localctx.(*MkdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(145)
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
			p.SetState(139)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(140)

				var _x = p.Mkdiskparam()

				localctx.(*MkdiskparamsContext).p = _x
			}
			localctx.(*MkdiskparamsContext).SetResult(localctx.(*MkdiskparamsContext).GetL().GetResult())
			localctx.(*MkdiskparamsContext).result[localctx.(*MkdiskparamsContext).GetP().GetResult()[0]] = localctx.(*MkdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(147)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)

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
			p.SetState(152)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)

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
			p.SetState(156)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)

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
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)

			var _m = p.Match(ParserParserRW_rmdisk)

			localctx.(*RmdiskContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)

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
			p.SetState(167)

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

// IFdiskContext is an interface to support dynamic dispatch.
type IFdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF returns the f token.
	GetF() antlr.Token

	// SetF sets the f token.
	SetF(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IFdiskparamsContext

	// SetP sets the p rule contexts.
	SetP(IFdiskparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Fdisk

	// SetResult sets the result attribute.
	SetResult(*commands.Fdisk)

	// Getter signatures
	RW_fdisk() antlr.TerminalNode
	Fdiskparams() IFdiskparamsContext

	// IsFdiskContext differentiates from other interfaces.
	IsFdiskContext()
}

type FdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Fdisk
	f      antlr.Token
	p      IFdiskparamsContext
}

func NewEmptyFdiskContext() *FdiskContext {
	var p = new(FdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
	return p
}

func InitEmptyFdiskContext(p *FdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdisk
}

func (*FdiskContext) IsFdiskContext() {}

func NewFdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskContext {
	var p = new(FdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdisk

	return p
}

func (s *FdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskContext) GetF() antlr.Token { return s.f }

func (s *FdiskContext) SetF(v antlr.Token) { s.f = v }

func (s *FdiskContext) GetP() IFdiskparamsContext { return s.p }

func (s *FdiskContext) SetP(v IFdiskparamsContext) { s.p = v }

func (s *FdiskContext) GetResult() *commands.Fdisk { return s.result }

func (s *FdiskContext) SetResult(v *commands.Fdisk) { s.result = v }

func (s *FdiskContext) RW_fdisk() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fdisk, 0)
}

func (s *FdiskContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdisk(s)
	}
}

func (s *FdiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdisk(s)
	}
}

func (p *ParserParser) Fdisk() (localctx IFdiskContext) {
	localctx = NewFdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserParserRULE_fdisk)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _x = p.fdiskparams(0)

			localctx.(*FdiskContext).p = _x
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), localctx.(*FdiskContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)

			var _m = p.Match(ParserParserRW_fdisk)

			localctx.(*FdiskContext).f = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskContext).result = commands.NewFdisk((func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetLine()
			}
		}()), (func() int {
			if localctx.(*FdiskContext).GetF() == nil {
				return 0
			} else {
				return localctx.(*FdiskContext).GetF().GetColumn()
			}
		}()), map[string]string{"unit": "K", "fit": "WF", "type": "P"})

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

// IFdiskparamsContext is an interface to support dynamic dispatch.
type IFdiskparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IFdiskparamsContext

	// GetP returns the p rule contexts.
	GetP() IFdiskparamContext

	// SetL sets the l rule contexts.
	SetL(IFdiskparamsContext)

	// SetP sets the p rule contexts.
	SetP(IFdiskparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Fdiskparam() IFdiskparamContext
	Fdiskparams() IFdiskparamsContext

	// IsFdiskparamsContext differentiates from other interfaces.
	IsFdiskparamsContext()
}

type FdiskparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IFdiskparamsContext
	p      IFdiskparamContext
}

func NewEmptyFdiskparamsContext() *FdiskparamsContext {
	var p = new(FdiskparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
	return p
}

func InitEmptyFdiskparamsContext(p *FdiskparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparams
}

func (*FdiskparamsContext) IsFdiskparamsContext() {}

func NewFdiskparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamsContext {
	var p = new(FdiskparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparams

	return p
}

func (s *FdiskparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamsContext) GetL() IFdiskparamsContext { return s.l }

func (s *FdiskparamsContext) GetP() IFdiskparamContext { return s.p }

func (s *FdiskparamsContext) SetL(v IFdiskparamsContext) { s.l = v }

func (s *FdiskparamsContext) SetP(v IFdiskparamContext) { s.p = v }

func (s *FdiskparamsContext) GetResult() map[string]string { return s.result }

func (s *FdiskparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *FdiskparamsContext) Fdiskparam() IFdiskparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamContext)
}

func (s *FdiskparamsContext) Fdiskparams() IFdiskparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskparamsContext)
}

func (s *FdiskparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdiskparams(s)
	}
}

func (s *FdiskparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdiskparams(s)
	}
}

func (p *ParserParser) Fdiskparams() (localctx IFdiskparamsContext) {
	return p.fdiskparams(0)
}

func (p *ParserParser) fdiskparams(_p int) (localctx IFdiskparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFdiskparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFdiskparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, ParserParserRULE_fdiskparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)

		var _x = p.Fdiskparam()

		localctx.(*FdiskparamsContext).p = _x
	}
	localctx.(*FdiskparamsContext).result = map[string]string{"unit": "K", "fit": "WF", "type": "P", localctx.(*FdiskparamsContext).GetP().GetResult()[0]: localctx.(*FdiskparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(189)
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
			localctx = NewFdiskparamsContext(p, _parentctx, _parentState)
			localctx.(*FdiskparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_fdiskparams)
			p.SetState(183)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(184)

				var _x = p.Fdiskparam()

				localctx.(*FdiskparamsContext).p = _x
			}
			localctx.(*FdiskparamsContext).SetResult(localctx.(*FdiskparamsContext).GetL().GetResult())
			localctx.(*FdiskparamsContext).result[localctx.(*FdiskparamsContext).GetP().GetResult()[0]] = localctx.(*FdiskparamsContext).GetP().GetResult()[1]

		}
		p.SetState(191)
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

// IFdiskparamContext is an interface to support dynamic dispatch.
type IFdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// GetV5 returns the v5 token.
	GetV5() antlr.Token

	// GetV6 returns the v6 token.
	GetV6() antlr.Token

	// GetV7 returns the v7 token.
	GetV7() antlr.Token

	// GetV8 returns the v8 token.
	GetV8() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// SetV5 sets the v5 token.
	SetV5(antlr.Token)

	// SetV6 sets the v6 token.
	SetV6(antlr.Token)

	// SetV7 sets the v7 token.
	SetV7(antlr.Token)

	// SetV8 sets the v8 token.
	SetV8(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_size() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_driveletter() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_name() antlr.TerminalNode
	RW_unit() antlr.TerminalNode
	TK_unit() antlr.TerminalNode
	RW_type() antlr.TerminalNode
	TK_type() antlr.TerminalNode
	RW_fit() antlr.TerminalNode
	TK_fit() antlr.TerminalNode
	RW_delete() antlr.TerminalNode
	RW_full() antlr.TerminalNode
	RW_add() antlr.TerminalNode

	// IsFdiskparamContext differentiates from other interfaces.
	IsFdiskparamContext()
}

type FdiskparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
	v5     antlr.Token
	v6     antlr.Token
	v7     antlr.Token
	v8     antlr.Token
}

func NewEmptyFdiskparamContext() *FdiskparamContext {
	var p = new(FdiskparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
	return p
}

func InitEmptyFdiskparamContext(p *FdiskparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_fdiskparam
}

func (*FdiskparamContext) IsFdiskparamContext() {}

func NewFdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamContext {
	var p = new(FdiskparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_fdiskparam

	return p
}

func (s *FdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamContext) GetV1() antlr.Token { return s.v1 }

func (s *FdiskparamContext) GetV2() antlr.Token { return s.v2 }

func (s *FdiskparamContext) GetV3() antlr.Token { return s.v3 }

func (s *FdiskparamContext) GetV4() antlr.Token { return s.v4 }

func (s *FdiskparamContext) GetV5() antlr.Token { return s.v5 }

func (s *FdiskparamContext) GetV6() antlr.Token { return s.v6 }

func (s *FdiskparamContext) GetV7() antlr.Token { return s.v7 }

func (s *FdiskparamContext) GetV8() antlr.Token { return s.v8 }

func (s *FdiskparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *FdiskparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *FdiskparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *FdiskparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *FdiskparamContext) SetV5(v antlr.Token) { s.v5 = v }

func (s *FdiskparamContext) SetV6(v antlr.Token) { s.v6 = v }

func (s *FdiskparamContext) SetV7(v antlr.Token) { s.v7 = v }

func (s *FdiskparamContext) SetV8(v antlr.Token) { s.v8 = v }

func (s *FdiskparamContext) GetResult() []string { return s.result }

func (s *FdiskparamContext) SetResult(v []string) { s.result = v }

func (s *FdiskparamContext) RW_size() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_size, 0)
}

func (s *FdiskparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *FdiskparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *FdiskparamContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *FdiskparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *FdiskparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *FdiskparamContext) RW_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unit, 0)
}

func (s *FdiskparamContext) TK_unit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_unit, 0)
}

func (s *FdiskparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *FdiskparamContext) TK_type() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_type, 0)
}

func (s *FdiskparamContext) RW_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fit, 0)
}

func (s *FdiskparamContext) TK_fit() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fit, 0)
}

func (s *FdiskparamContext) RW_delete() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_delete, 0)
}

func (s *FdiskparamContext) RW_full() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_full, 0)
}

func (s *FdiskparamContext) RW_add() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_add, 0)
}

func (s *FdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFdiskparam(s)
	}
}

func (s *FdiskparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFdiskparam(s)
	}
}

func (p *ParserParser) Fdiskparam() (localctx IFdiskparamContext) {
	localctx = NewFdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_fdiskparam)
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_size:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(ParserParserRW_size)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"size", (func() string {
			if localctx.(*FdiskparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_driveletter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*FdiskparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"driveletter", (func() string {
			if localctx.(*FdiskparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*FdiskparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"name", (func() string {
			if localctx.(*FdiskparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV3().GetText()
			}
		}())}

	case ParserParserRW_unit:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(ParserParserRW_unit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)

			var _m = p.Match(ParserParserTK_unit)

			localctx.(*FdiskparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"unit", (func() string {
			if localctx.(*FdiskparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV4().GetText()
			}
		}())}

	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(208)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)

			var _m = p.Match(ParserParserTK_type)

			localctx.(*FdiskparamContext).v5 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"type", (func() string {
			if localctx.(*FdiskparamContext).GetV5() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV5().GetText()
			}
		}())}

	case ParserParserRW_fit:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(212)
			p.Match(ParserParserRW_fit)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)

			var _m = p.Match(ParserParserTK_fit)

			localctx.(*FdiskparamContext).v6 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"fit", (func() string {
			if localctx.(*FdiskparamContext).GetV6() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV6().GetText()
			}
		}())}

	case ParserParserRW_delete:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(216)
			p.Match(ParserParserRW_delete)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)

			var _m = p.Match(ParserParserRW_full)

			localctx.(*FdiskparamContext).v7 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"delete", (func() string {
			if localctx.(*FdiskparamContext).GetV7() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV7().GetText()
			}
		}())}

	case ParserParserRW_add:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(220)
			p.Match(ParserParserRW_add)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*FdiskparamContext).v8 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FdiskparamContext).result = []string{"add", (func() string {
			if localctx.(*FdiskparamContext).GetV8() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetV8().GetText()
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

// IMountContext is an interface to support dynamic dispatch.
type IMountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMountparamsContext

	// SetP sets the p rule contexts.
	SetP(IMountparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mount

	// SetResult sets the result attribute.
	SetResult(*commands.Mount)

	// Getter signatures
	RW_mount() antlr.TerminalNode
	Mountparams() IMountparamsContext

	// IsMountContext differentiates from other interfaces.
	IsMountContext()
}

type MountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mount
	m      antlr.Token
	p      IMountparamsContext
}

func NewEmptyMountContext() *MountContext {
	var p = new(MountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
	return p
}

func InitEmptyMountContext(p *MountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mount
}

func (*MountContext) IsMountContext() {}

func NewMountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountContext {
	var p = new(MountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mount

	return p
}

func (s *MountContext) GetParser() antlr.Parser { return s.parser }

func (s *MountContext) GetM() antlr.Token { return s.m }

func (s *MountContext) SetM(v antlr.Token) { s.m = v }

func (s *MountContext) GetP() IMountparamsContext { return s.p }

func (s *MountContext) SetP(v IMountparamsContext) { s.p = v }

func (s *MountContext) GetResult() *commands.Mount { return s.result }

func (s *MountContext) SetResult(v *commands.Mount) { s.result = v }

func (s *MountContext) RW_mount() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mount, 0)
}

func (s *MountContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMount(s)
	}
}

func (s *MountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMount(s)
	}
}

func (p *ParserParser) Mount() (localctx IMountContext) {
	localctx = NewMountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParserParserRULE_mount)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)

			var _x = p.mountparams(0)

			localctx.(*MountContext).p = _x
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
			}
		}()), localctx.(*MountContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)

			var _m = p.Match(ParserParserRW_mount)

			localctx.(*MountContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountContext).result = commands.NewMount((func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MountContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MountContext).GetM().GetColumn()
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

// IMountparamsContext is an interface to support dynamic dispatch.
type IMountparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMountparamsContext

	// GetP returns the p rule contexts.
	GetP() IMountparamContext

	// SetL sets the l rule contexts.
	SetL(IMountparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMountparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mountparam() IMountparamContext
	Mountparams() IMountparamsContext

	// IsMountparamsContext differentiates from other interfaces.
	IsMountparamsContext()
}

type MountparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMountparamsContext
	p      IMountparamContext
}

func NewEmptyMountparamsContext() *MountparamsContext {
	var p = new(MountparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
	return p
}

func InitEmptyMountparamsContext(p *MountparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparams
}

func (*MountparamsContext) IsMountparamsContext() {}

func NewMountparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamsContext {
	var p = new(MountparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparams

	return p
}

func (s *MountparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamsContext) GetL() IMountparamsContext { return s.l }

func (s *MountparamsContext) GetP() IMountparamContext { return s.p }

func (s *MountparamsContext) SetL(v IMountparamsContext) { s.l = v }

func (s *MountparamsContext) SetP(v IMountparamContext) { s.p = v }

func (s *MountparamsContext) GetResult() map[string]string { return s.result }

func (s *MountparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MountparamsContext) Mountparam() IMountparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamContext)
}

func (s *MountparamsContext) Mountparams() IMountparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountparamsContext)
}

func (s *MountparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMountparams(s)
	}
}

func (s *MountparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMountparams(s)
	}
}

func (p *ParserParser) Mountparams() (localctx IMountparamsContext) {
	return p.mountparams(0)
}

func (p *ParserParser) mountparams(_p int) (localctx IMountparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMountparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMountparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ParserParserRULE_mountparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)

		var _x = p.Mountparam()

		localctx.(*MountparamsContext).p = _x
	}
	localctx.(*MountparamsContext).result = map[string]string{localctx.(*MountparamsContext).GetP().GetResult()[0]: localctx.(*MountparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMountparamsContext(p, _parentctx, _parentState)
			localctx.(*MountparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mountparams)
			p.SetState(238)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(239)

				var _x = p.Mountparam()

				localctx.(*MountparamsContext).p = _x
			}
			localctx.(*MountparamsContext).SetResult(localctx.(*MountparamsContext).GetL().GetResult())
			localctx.(*MountparamsContext).result[localctx.(*MountparamsContext).GetP().GetResult()[0]] = localctx.(*MountparamsContext).GetP().GetResult()[1]

		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// IMountparamContext is an interface to support dynamic dispatch.
type IMountparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_driveletter() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_name() antlr.TerminalNode

	// IsMountparamContext differentiates from other interfaces.
	IsMountparamContext()
}

type MountparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
}

func NewEmptyMountparamContext() *MountparamContext {
	var p = new(MountparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
	return p
}

func InitEmptyMountparamContext(p *MountparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mountparam
}

func (*MountparamContext) IsMountparamContext() {}

func NewMountparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamContext {
	var p = new(MountparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mountparam

	return p
}

func (s *MountparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MountparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MountparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MountparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MountparamContext) GetResult() []string { return s.result }

func (s *MountparamContext) SetResult(v []string) { s.result = v }

func (s *MountparamContext) RW_driveletter() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_driveletter, 0)
}

func (s *MountparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MountparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MountparamContext) RW_name() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_name, 0)
}

func (s *MountparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMountparam(s)
	}
}

func (s *MountparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMountparam(s)
	}
}

func (p *ParserParser) Mountparam() (localctx IMountparamContext) {
	localctx = NewMountparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParserParserRULE_mountparam)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_driveletter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(ParserParserRW_driveletter)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MountparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"driveletter", (func() string {
			if localctx.(*MountparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(253)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MountparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MountparamContext).result = []string{"name", (func() string {
			if localctx.(*MountparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetV2().GetText()
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

// IUnmountContext is an interface to support dynamic dispatch.
type IUnmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetU returns the u token.
	GetU() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetU sets the u token.
	SetU(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Unmount

	// SetResult sets the result attribute.
	SetResult(*commands.Unmount)

	// Getter signatures
	RW_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	RW_unmount() antlr.TerminalNode
	TK_id() antlr.TerminalNode

	// IsUnmountContext differentiates from other interfaces.
	IsUnmountContext()
}

type UnmountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Unmount
	u      antlr.Token
	p      antlr.Token
}

func NewEmptyUnmountContext() *UnmountContext {
	var p = new(UnmountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_unmount
	return p
}

func InitEmptyUnmountContext(p *UnmountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_unmount
}

func (*UnmountContext) IsUnmountContext() {}

func NewUnmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnmountContext {
	var p = new(UnmountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_unmount

	return p
}

func (s *UnmountContext) GetParser() antlr.Parser { return s.parser }

func (s *UnmountContext) GetU() antlr.Token { return s.u }

func (s *UnmountContext) GetP() antlr.Token { return s.p }

func (s *UnmountContext) SetU(v antlr.Token) { s.u = v }

func (s *UnmountContext) SetP(v antlr.Token) { s.p = v }

func (s *UnmountContext) GetResult() *commands.Unmount { return s.result }

func (s *UnmountContext) SetResult(v *commands.Unmount) { s.result = v }

func (s *UnmountContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *UnmountContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *UnmountContext) RW_unmount() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_unmount, 0)
}

func (s *UnmountContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *UnmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnmountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterUnmount(s)
	}
}

func (s *UnmountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitUnmount(s)
	}
}

func (p *ParserParser) Unmount() (localctx IUnmountContext) {
	localctx = NewUnmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParserParserRULE_unmount)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)

			var _m = p.Match(ParserParserRW_unmount)

			localctx.(*UnmountContext).u = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*UnmountContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*UnmountContext).result = commands.NewUnmount((func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetLine()
			}
		}()), (func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetColumn()
			}
		}()), map[string]string{"id": (func() string {
			if localctx.(*UnmountContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*UnmountContext).GetP().GetText()
			}
		}())})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)

			var _m = p.Match(ParserParserRW_unmount)

			localctx.(*UnmountContext).u = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*UnmountContext).result = commands.NewUnmount((func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetLine()
			}
		}()), (func() int {
			if localctx.(*UnmountContext).GetU() == nil {
				return 0
			} else {
				return localctx.(*UnmountContext).GetU().GetColumn()
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

// IMkfsContext is an interface to support dynamic dispatch.
type IMkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetM returns the m token.
	GetM() antlr.Token

	// SetM sets the m token.
	SetM(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() IMkfsparamsContext

	// SetP sets the p rule contexts.
	SetP(IMkfsparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Mkfs

	// SetResult sets the result attribute.
	SetResult(*commands.Mkfs)

	// Getter signatures
	RW_mkfs() antlr.TerminalNode
	Mkfsparams() IMkfsparamsContext

	// IsMkfsContext differentiates from other interfaces.
	IsMkfsContext()
}

type MkfsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Mkfs
	m      antlr.Token
	p      IMkfsparamsContext
}

func NewEmptyMkfsContext() *MkfsContext {
	var p = new(MkfsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
	return p
}

func InitEmptyMkfsContext(p *MkfsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfs
}

func (*MkfsContext) IsMkfsContext() {}

func NewMkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsContext {
	var p = new(MkfsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfs

	return p
}

func (s *MkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsContext) GetM() antlr.Token { return s.m }

func (s *MkfsContext) SetM(v antlr.Token) { s.m = v }

func (s *MkfsContext) GetP() IMkfsparamsContext { return s.p }

func (s *MkfsContext) SetP(v IMkfsparamsContext) { s.p = v }

func (s *MkfsContext) GetResult() *commands.Mkfs { return s.result }

func (s *MkfsContext) SetResult(v *commands.Mkfs) { s.result = v }

func (s *MkfsContext) RW_mkfs() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mkfs, 0)
}

func (s *MkfsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfs(s)
	}
}

func (s *MkfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfs(s)
	}
}

func (p *ParserParser) Mkfs() (localctx IMkfsContext) {
	localctx = NewMkfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ParserParserRULE_mkfs)
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)

			var _x = p.mkfsparams(0)

			localctx.(*MkfsContext).p = _x
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), localctx.(*MkfsContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)

			var _m = p.Match(ParserParserRW_mkfs)

			localctx.(*MkfsContext).m = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsContext).result = commands.NewMkfs((func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetLine()
			}
		}()), (func() int {
			if localctx.(*MkfsContext).GetM() == nil {
				return 0
			} else {
				return localctx.(*MkfsContext).GetM().GetColumn()
			}
		}()), map[string]string{"fs": "2fs"})

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

// IMkfsparamsContext is an interface to support dynamic dispatch.
type IMkfsparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IMkfsparamsContext

	// GetP returns the p rule contexts.
	GetP() IMkfsparamContext

	// SetL sets the l rule contexts.
	SetL(IMkfsparamsContext)

	// SetP sets the p rule contexts.
	SetP(IMkfsparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Mkfsparam() IMkfsparamContext
	Mkfsparams() IMkfsparamsContext

	// IsMkfsparamsContext differentiates from other interfaces.
	IsMkfsparamsContext()
}

type MkfsparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      IMkfsparamsContext
	p      IMkfsparamContext
}

func NewEmptyMkfsparamsContext() *MkfsparamsContext {
	var p = new(MkfsparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
	return p
}

func InitEmptyMkfsparamsContext(p *MkfsparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparams
}

func (*MkfsparamsContext) IsMkfsparamsContext() {}

func NewMkfsparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamsContext {
	var p = new(MkfsparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparams

	return p
}

func (s *MkfsparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamsContext) GetL() IMkfsparamsContext { return s.l }

func (s *MkfsparamsContext) GetP() IMkfsparamContext { return s.p }

func (s *MkfsparamsContext) SetL(v IMkfsparamsContext) { s.l = v }

func (s *MkfsparamsContext) SetP(v IMkfsparamContext) { s.p = v }

func (s *MkfsparamsContext) GetResult() map[string]string { return s.result }

func (s *MkfsparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *MkfsparamsContext) Mkfsparam() IMkfsparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamContext)
}

func (s *MkfsparamsContext) Mkfsparams() IMkfsparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsparamsContext)
}

func (s *MkfsparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfsparams(s)
	}
}

func (s *MkfsparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfsparams(s)
	}
}

func (p *ParserParser) Mkfsparams() (localctx IMkfsparamsContext) {
	return p.mkfsparams(0)
}

func (p *ParserParser) mkfsparams(_p int) (localctx IMkfsparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewMkfsparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMkfsparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, ParserParserRULE_mkfsparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)

		var _x = p.Mkfsparam()

		localctx.(*MkfsparamsContext).p = _x
	}
	localctx.(*MkfsparamsContext).result = map[string]string{"fs": "2fs", localctx.(*MkfsparamsContext).GetP().GetResult()[0]: localctx.(*MkfsparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMkfsparamsContext(p, _parentctx, _parentState)
			localctx.(*MkfsparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_mkfsparams)
			p.SetState(278)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(279)

				var _x = p.Mkfsparam()

				localctx.(*MkfsparamsContext).p = _x
			}
			localctx.(*MkfsparamsContext).SetResult(localctx.(*MkfsparamsContext).GetL().GetResult())
			localctx.(*MkfsparamsContext).result[localctx.(*MkfsparamsContext).GetP().GetResult()[0]] = localctx.(*MkfsparamsContext).GetP().GetResult()[1]

		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
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

// IMkfsparamContext is an interface to support dynamic dispatch.
type IMkfsparamContext interface {
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
	RW_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_type() antlr.TerminalNode
	RW_full() antlr.TerminalNode
	RW_fs() antlr.TerminalNode
	TK_fs() antlr.TerminalNode

	// IsMkfsparamContext differentiates from other interfaces.
	IsMkfsparamContext()
}

type MkfsparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
}

func NewEmptyMkfsparamContext() *MkfsparamContext {
	var p = new(MkfsparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
	return p
}

func InitEmptyMkfsparamContext(p *MkfsparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_mkfsparam
}

func (*MkfsparamContext) IsMkfsparamContext() {}

func NewMkfsparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamContext {
	var p = new(MkfsparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_mkfsparam

	return p
}

func (s *MkfsparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamContext) GetV1() antlr.Token { return s.v1 }

func (s *MkfsparamContext) GetV2() antlr.Token { return s.v2 }

func (s *MkfsparamContext) GetV3() antlr.Token { return s.v3 }

func (s *MkfsparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *MkfsparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *MkfsparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *MkfsparamContext) GetResult() []string { return s.result }

func (s *MkfsparamContext) SetResult(v []string) { s.result = v }

func (s *MkfsparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *MkfsparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *MkfsparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *MkfsparamContext) RW_type() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_type, 0)
}

func (s *MkfsparamContext) RW_full() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_full, 0)
}

func (s *MkfsparamContext) RW_fs() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_fs, 0)
}

func (s *MkfsparamContext) TK_fs() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_fs, 0)
}

func (s *MkfsparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterMkfsparam(s)
	}
}

func (s *MkfsparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitMkfsparam(s)
	}
}

func (p *ParserParser) Mkfsparam() (localctx IMkfsparamContext) {
	localctx = NewMkfsparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParserParserRULE_mkfsparam)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_id:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*MkfsparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"id", (func() string {
			if localctx.(*MkfsparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV1().GetText()
			}
		}())}

	case ParserParserRW_type:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Match(ParserParserRW_type)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)

			var _m = p.Match(ParserParserRW_full)

			localctx.(*MkfsparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"type", (func() string {
			if localctx.(*MkfsparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV2().GetText()
			}
		}())}

	case ParserParserRW_fs:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(ParserParserRW_fs)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)

			var _m = p.Match(ParserParserTK_fs)

			localctx.(*MkfsparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*MkfsparamContext).result = []string{"fs", (func() string {
			if localctx.(*MkfsparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).GetV3().GetText()
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

// ILoginContext is an interface to support dynamic dispatch.
type ILoginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// GetP returns the p rule contexts.
	GetP() ILoginparamsContext

	// SetP sets the p rule contexts.
	SetP(ILoginparamsContext)

	// GetResult returns the result attribute.
	GetResult() *commands.Login

	// SetResult sets the result attribute.
	SetResult(*commands.Login)

	// Getter signatures
	RW_login() antlr.TerminalNode
	Loginparams() ILoginparamsContext

	// IsLoginContext differentiates from other interfaces.
	IsLoginContext()
}

type LoginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Login
	l      antlr.Token
	p      ILoginparamsContext
}

func NewEmptyLoginContext() *LoginContext {
	var p = new(LoginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
	return p
}

func InitEmptyLoginContext(p *LoginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_login
}

func (*LoginContext) IsLoginContext() {}

func NewLoginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginContext {
	var p = new(LoginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_login

	return p
}

func (s *LoginContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginContext) GetL() antlr.Token { return s.l }

func (s *LoginContext) SetL(v antlr.Token) { s.l = v }

func (s *LoginContext) GetP() ILoginparamsContext { return s.p }

func (s *LoginContext) SetP(v ILoginparamsContext) { s.p = v }

func (s *LoginContext) GetResult() *commands.Login { return s.result }

func (s *LoginContext) SetResult(v *commands.Login) { s.result = v }

func (s *LoginContext) RW_login() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_login, 0)
}

func (s *LoginContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLogin(s)
	}
}

func (s *LoginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLogin(s)
	}
}

func (p *ParserParser) Login() (localctx ILoginContext) {
	localctx = NewLoginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ParserParserRULE_login)
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)

			var _x = p.loginparams(0)

			localctx.(*LoginContext).p = _x
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
			}
		}()), localctx.(*LoginContext).GetP().GetResult())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)

			var _m = p.Match(ParserParserRW_login)

			localctx.(*LoginContext).l = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginContext).result = commands.NewLogin((func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetLine()
			}
		}()), (func() int {
			if localctx.(*LoginContext).GetL() == nil {
				return 0
			} else {
				return localctx.(*LoginContext).GetL().GetColumn()
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

// ILoginparamsContext is an interface to support dynamic dispatch.
type ILoginparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() ILoginparamsContext

	// GetP returns the p rule contexts.
	GetP() ILoginparamContext

	// SetL sets the l rule contexts.
	SetL(ILoginparamsContext)

	// SetP sets the p rule contexts.
	SetP(ILoginparamContext)

	// GetResult returns the result attribute.
	GetResult() map[string]string

	// SetResult sets the result attribute.
	SetResult(map[string]string)

	// Getter signatures
	Loginparam() ILoginparamContext
	Loginparams() ILoginparamsContext

	// IsLoginparamsContext differentiates from other interfaces.
	IsLoginparamsContext()
}

type LoginparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result map[string]string
	l      ILoginparamsContext
	p      ILoginparamContext
}

func NewEmptyLoginparamsContext() *LoginparamsContext {
	var p = new(LoginparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
	return p
}

func InitEmptyLoginparamsContext(p *LoginparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparams
}

func (*LoginparamsContext) IsLoginparamsContext() {}

func NewLoginparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamsContext {
	var p = new(LoginparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparams

	return p
}

func (s *LoginparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamsContext) GetL() ILoginparamsContext { return s.l }

func (s *LoginparamsContext) GetP() ILoginparamContext { return s.p }

func (s *LoginparamsContext) SetL(v ILoginparamsContext) { s.l = v }

func (s *LoginparamsContext) SetP(v ILoginparamContext) { s.p = v }

func (s *LoginparamsContext) GetResult() map[string]string { return s.result }

func (s *LoginparamsContext) SetResult(v map[string]string) { s.result = v }

func (s *LoginparamsContext) Loginparam() ILoginparamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamContext)
}

func (s *LoginparamsContext) Loginparams() ILoginparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginparamsContext)
}

func (s *LoginparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoginparams(s)
	}
}

func (s *LoginparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoginparams(s)
	}
}

func (p *ParserParser) Loginparams() (localctx ILoginparamsContext) {
	return p.loginparams(0)
}

func (p *ParserParser) loginparams(_p int) (localctx ILoginparamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLoginparamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILoginparamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, ParserParserRULE_loginparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)

		var _x = p.Loginparam()

		localctx.(*LoginparamsContext).p = _x
	}
	localctx.(*LoginparamsContext).result = map[string]string{localctx.(*LoginparamsContext).GetP().GetResult()[0]: localctx.(*LoginparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLoginparamsContext(p, _parentctx, _parentState)
			localctx.(*LoginparamsContext).l = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_loginparams)
			p.SetState(313)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(314)

				var _x = p.Loginparam()

				localctx.(*LoginparamsContext).p = _x
			}
			localctx.(*LoginparamsContext).SetResult(localctx.(*LoginparamsContext).GetL().GetResult())
			localctx.(*LoginparamsContext).result[localctx.(*LoginparamsContext).GetP().GetResult()[0]] = localctx.(*LoginparamsContext).GetP().GetResult()[1]

		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
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

// ILoginparamContext is an interface to support dynamic dispatch.
type ILoginparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetV1 returns the v1 token.
	GetV1() antlr.Token

	// GetV2 returns the v2 token.
	GetV2() antlr.Token

	// GetV3 returns the v3 token.
	GetV3() antlr.Token

	// GetV4 returns the v4 token.
	GetV4() antlr.Token

	// SetV1 sets the v1 token.
	SetV1(antlr.Token)

	// SetV2 sets the v2 token.
	SetV2(antlr.Token)

	// SetV3 sets the v3 token.
	SetV3(antlr.Token)

	// SetV4 sets the v4 token.
	SetV4(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() []string

	// SetResult sets the result attribute.
	SetResult([]string)

	// Getter signatures
	RW_user() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_pass() antlr.TerminalNode
	TK_number() antlr.TerminalNode
	RW_id() antlr.TerminalNode

	// IsLoginparamContext differentiates from other interfaces.
	IsLoginparamContext()
}

type LoginparamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []string
	v1     antlr.Token
	v2     antlr.Token
	v3     antlr.Token
	v4     antlr.Token
}

func NewEmptyLoginparamContext() *LoginparamContext {
	var p = new(LoginparamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
	return p
}

func InitEmptyLoginparamContext(p *LoginparamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loginparam
}

func (*LoginparamContext) IsLoginparamContext() {}

func NewLoginparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamContext {
	var p = new(LoginparamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loginparam

	return p
}

func (s *LoginparamContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamContext) GetV1() antlr.Token { return s.v1 }

func (s *LoginparamContext) GetV2() antlr.Token { return s.v2 }

func (s *LoginparamContext) GetV3() antlr.Token { return s.v3 }

func (s *LoginparamContext) GetV4() antlr.Token { return s.v4 }

func (s *LoginparamContext) SetV1(v antlr.Token) { s.v1 = v }

func (s *LoginparamContext) SetV2(v antlr.Token) { s.v2 = v }

func (s *LoginparamContext) SetV3(v antlr.Token) { s.v3 = v }

func (s *LoginparamContext) SetV4(v antlr.Token) { s.v4 = v }

func (s *LoginparamContext) GetResult() []string { return s.result }

func (s *LoginparamContext) SetResult(v []string) { s.result = v }

func (s *LoginparamContext) RW_user() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_user, 0)
}

func (s *LoginparamContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *LoginparamContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *LoginparamContext) RW_pass() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pass, 0)
}

func (s *LoginparamContext) TK_number() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_number, 0)
}

func (s *LoginparamContext) RW_id() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_id, 0)
}

func (s *LoginparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoginparam(s)
	}
}

func (s *LoginparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoginparam(s)
	}
}

func (p *ParserParser) Loginparam() (localctx ILoginparamContext) {
	localctx = NewLoginparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParserParserRULE_loginparam)
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(322)
			p.Match(ParserParserRW_user)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"user", (func() string {
			if localctx.(*LoginparamContext).GetV1() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV1().GetText()
			}
		}())}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", (func() string {
			if localctx.(*LoginparamContext).GetV2() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV2().GetText()
			}
		}())}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Match(ParserParserRW_pass)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)

			var _m = p.Match(ParserParserTK_number)

			localctx.(*LoginparamContext).v3 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"pass", (func() string {
			if localctx.(*LoginparamContext).GetV3() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV3().GetText()
			}
		}())}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(334)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*LoginparamContext).v4 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*LoginparamContext).result = []string{"id", (func() string {
			if localctx.(*LoginparamContext).GetV4() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetV4().GetText()
			}
		}())}

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

// ILogoutContext is an interface to support dynamic dispatch.
type ILogoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Logout

	// SetResult sets the result attribute.
	SetResult(*commands.Logout)

	// Getter signatures
	RW_logout() antlr.TerminalNode

	// IsLogoutContext differentiates from other interfaces.
	IsLogoutContext()
}

type LogoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Logout
	p      antlr.Token
}

func NewEmptyLogoutContext() *LogoutContext {
	var p = new(LogoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
	return p
}

func InitEmptyLogoutContext(p *LogoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_logout
}

func (*LogoutContext) IsLogoutContext() {}

func NewLogoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogoutContext {
	var p = new(LogoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_logout

	return p
}

func (s *LogoutContext) GetParser() antlr.Parser { return s.parser }

func (s *LogoutContext) GetP() antlr.Token { return s.p }

func (s *LogoutContext) SetP(v antlr.Token) { s.p = v }

func (s *LogoutContext) GetResult() *commands.Logout { return s.result }

func (s *LogoutContext) SetResult(v *commands.Logout) { s.result = v }

func (s *LogoutContext) RW_logout() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_logout, 0)
}

func (s *LogoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLogout(s)
	}
}

func (s *LogoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLogout(s)
	}
}

func (p *ParserParser) Logout() (localctx ILogoutContext) {
	localctx = NewLogoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ParserParserRULE_logout)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)

		var _m = p.Match(ParserParserRW_logout)

		localctx.(*LogoutContext).p = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*LogoutContext).result = commands.NewLogout((func() int {
		if localctx.(*LogoutContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetP().GetLine()
		}
	}()), (func() int {
		if localctx.(*LogoutContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*LogoutContext).GetP().GetColumn()
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

// IPauseContext is an interface to support dynamic dispatch.
type IPauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetResult returns the result attribute.
	GetResult() *commands.Pause

	// SetResult sets the result attribute.
	SetResult(*commands.Pause)

	// Getter signatures
	RW_pause() antlr.TerminalNode

	// IsPauseContext differentiates from other interfaces.
	IsPauseContext()
}

type PauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Pause
	p      antlr.Token
}

func NewEmptyPauseContext() *PauseContext {
	var p = new(PauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_pause
	return p
}

func InitEmptyPauseContext(p *PauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_pause
}

func (*PauseContext) IsPauseContext() {}

func NewPauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PauseContext {
	var p = new(PauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_pause

	return p
}

func (s *PauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PauseContext) GetP() antlr.Token { return s.p }

func (s *PauseContext) SetP(v antlr.Token) { s.p = v }

func (s *PauseContext) GetResult() *commands.Pause { return s.result }

func (s *PauseContext) SetResult(v *commands.Pause) { s.result = v }

func (s *PauseContext) RW_pause() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_pause, 0)
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterPause(s)
	}
}

func (s *PauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitPause(s)
	}
}

func (p *ParserParser) Pause() (localctx IPauseContext) {
	localctx = NewPauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ParserParserRULE_pause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)

		var _m = p.Match(ParserParserRW_pause)

		localctx.(*PauseContext).p = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*PauseContext).result = commands.NewPause((func() int {
		if localctx.(*PauseContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*PauseContext).GetP().GetLine()
		}
	}()), (func() int {
		if localctx.(*PauseContext).GetP() == nil {
			return 0
		} else {
			return localctx.(*PauseContext).GetP().GetColumn()
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
	p.EnterRule(localctx, 46, ParserParserRULE_rep)
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)

			var _m = p.Match(ParserParserRW_rep)

			localctx.(*RepContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(347)

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
			p.SetState(350)

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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, ParserParserRULE_repparams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)

		var _x = p.Repparam()

		localctx.(*RepparamsContext).p = _x
	}
	localctx.(*RepparamsContext).result = map[string]string{localctx.(*RepparamsContext).GetP().GetResult()[0]: localctx.(*RepparamsContext).GetP().GetResult()[1]}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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
			p.SetState(358)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(359)

				var _x = p.Repparam()

				localctx.(*RepparamsContext).p = _x
			}
			localctx.(*RepparamsContext).SetResult(localctx.(*RepparamsContext).GetL().GetResult())
			localctx.(*RepparamsContext).result[localctx.(*RepparamsContext).GetP().GetResult()[0]] = localctx.(*RepparamsContext).GetP().GetResult()[1]

		}
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 50, ParserParserRULE_repparam)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_name:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(ParserParserRW_name)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(368)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)

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
			p.SetState(372)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)

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
			p.SetState(376)
			p.Match(ParserParserRW_ruta)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(378)

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
			p.SetState(380)
			p.Match(ParserParserRW_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(382)

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
	p.EnterRule(localctx, 52, ParserParserRULE_name)
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_mbr:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(386)

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
			p.SetState(388)

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
			p.SetState(390)

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
			p.SetState(392)

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
			p.SetState(394)

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
			p.SetState(396)

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
			p.SetState(398)

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
			p.SetState(400)

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
			p.SetState(402)

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
			p.SetState(404)

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
			p.SetState(406)

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
	p.EnterRule(localctx, 54, ParserParserRULE_commentary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)

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
	p.EnterRule(localctx, 56, ParserParserRULE_exit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
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
		var t *FdiskparamsContext = nil
		if localctx != nil {
			t = localctx.(*FdiskparamsContext)
		}
		return p.Fdiskparams_Sempred(t, predIndex)

	case 12:
		var t *MountparamsContext = nil
		if localctx != nil {
			t = localctx.(*MountparamsContext)
		}
		return p.Mountparams_Sempred(t, predIndex)

	case 16:
		var t *MkfsparamsContext = nil
		if localctx != nil {
			t = localctx.(*MkfsparamsContext)
		}
		return p.Mkfsparams_Sempred(t, predIndex)

	case 19:
		var t *LoginparamsContext = nil
		if localctx != nil {
			t = localctx.(*LoginparamsContext)
		}
		return p.Loginparams_Sempred(t, predIndex)

	case 24:
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

func (p *ParserParser) Fdiskparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mountparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Mkfsparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Loginparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Repparams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
