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
		"'loss'", "'execute'", "'rep'", "'mbr'", "'disk'", "'inode'", "'journaling'",
		"'block'", "'bm_inode'", "'bm_block'", "'tree'", "'sb'", "'file'", "'ls'",
		"'full'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "'2fs'", "'3fs'", "", "", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "RW_mkdisk", "RW_rmdisk", "RW_fdisk", "RW_mount", "RW_unmount",
		"RW_mkfs", "RW_login", "RW_logout", "RW_mkgrp", "RW_rmgrp", "RW_mkusr",
		"RW_rmusr", "RW_mkfile", "RW_cat", "RW_remove", "RW_edit", "RW_rename",
		"RW_mkdir", "RW_copy", "RW_move", "RW_find", "RW_chown", "RW_chgrp",
		"RW_chmod", "RW_pause", "RW_recovery", "RW_loss", "RW_execute", "RW_rep",
		"RW_mbr", "RW_disk", "RW_inode", "RW_journaling", "RW_block", "RW_bm_inode",
		"RW_bm_block", "RW_tree", "RW_sb", "RW_file", "RW_ls", "RW_full", "RW_size",
		"RW_path", "RW_fit", "RW_unit", "RW_name", "RW_type", "RW_delete", "RW_add",
		"RW_id", "RW_fs", "RW_user", "RW_pass", "RW_grp", "RW_r", "RW_cont",
		"RW_fileN", "RW_destino", "RW_ugo", "RW_ruta", "TK_fit", "TK_unit",
		"TK_type", "TK_2fs", "TK_3fs", "TK_number", "TK_path", "TK_equ", "COMMENTARY",
	}
	staticData.RuleNames = []string{
		"init", "commands", "command", "execute",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 69, 39, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 15, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 20,
		8, 1, 10, 1, 12, 1, 23, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 35, 8, 3, 1, 3, 1, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6,
		0, 0, 37, 0, 14, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 26, 1, 0, 0, 0, 6, 29,
		1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1, 10, 11, 6, 0, -1, 0, 11,
		15, 1, 0, 0, 0, 12, 13, 5, 0, 0, 1, 13, 15, 6, 0, -1, 0, 14, 8, 1, 0, 0,
		0, 14, 12, 1, 0, 0, 0, 15, 1, 1, 0, 0, 0, 16, 17, 3, 4, 2, 0, 17, 18, 6,
		1, -1, 0, 18, 20, 1, 0, 0, 0, 19, 16, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21,
		19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 24, 1, 0, 0, 0, 23, 21, 1, 0, 0,
		0, 24, 25, 6, 1, -1, 0, 25, 3, 1, 0, 0, 0, 26, 27, 3, 6, 3, 0, 27, 28,
		6, 2, -1, 0, 28, 5, 1, 0, 0, 0, 29, 34, 5, 28, 0, 0, 30, 31, 5, 43, 0,
		0, 31, 32, 5, 68, 0, 0, 32, 33, 5, 67, 0, 0, 33, 35, 6, 3, -1, 0, 34, 30,
		1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 6, 3, -1, 0,
		37, 7, 1, 0, 0, 0, 3, 14, 21, 34,
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
	ParserParserEOF           = antlr.TokenEOF
	ParserParserRW_mkdisk     = 1
	ParserParserRW_rmdisk     = 2
	ParserParserRW_fdisk      = 3
	ParserParserRW_mount      = 4
	ParserParserRW_unmount    = 5
	ParserParserRW_mkfs       = 6
	ParserParserRW_login      = 7
	ParserParserRW_logout     = 8
	ParserParserRW_mkgrp      = 9
	ParserParserRW_rmgrp      = 10
	ParserParserRW_mkusr      = 11
	ParserParserRW_rmusr      = 12
	ParserParserRW_mkfile     = 13
	ParserParserRW_cat        = 14
	ParserParserRW_remove     = 15
	ParserParserRW_edit       = 16
	ParserParserRW_rename     = 17
	ParserParserRW_mkdir      = 18
	ParserParserRW_copy       = 19
	ParserParserRW_move       = 20
	ParserParserRW_find       = 21
	ParserParserRW_chown      = 22
	ParserParserRW_chgrp      = 23
	ParserParserRW_chmod      = 24
	ParserParserRW_pause      = 25
	ParserParserRW_recovery   = 26
	ParserParserRW_loss       = 27
	ParserParserRW_execute    = 28
	ParserParserRW_rep        = 29
	ParserParserRW_mbr        = 30
	ParserParserRW_disk       = 31
	ParserParserRW_inode      = 32
	ParserParserRW_journaling = 33
	ParserParserRW_block      = 34
	ParserParserRW_bm_inode   = 35
	ParserParserRW_bm_block   = 36
	ParserParserRW_tree       = 37
	ParserParserRW_sb         = 38
	ParserParserRW_file       = 39
	ParserParserRW_ls         = 40
	ParserParserRW_full       = 41
	ParserParserRW_size       = 42
	ParserParserRW_path       = 43
	ParserParserRW_fit        = 44
	ParserParserRW_unit       = 45
	ParserParserRW_name       = 46
	ParserParserRW_type       = 47
	ParserParserRW_delete     = 48
	ParserParserRW_add        = 49
	ParserParserRW_id         = 50
	ParserParserRW_fs         = 51
	ParserParserRW_user       = 52
	ParserParserRW_pass       = 53
	ParserParserRW_grp        = 54
	ParserParserRW_r          = 55
	ParserParserRW_cont       = 56
	ParserParserRW_fileN      = 57
	ParserParserRW_destino    = 58
	ParserParserRW_ugo        = 59
	ParserParserRW_ruta       = 60
	ParserParserTK_fit        = 61
	ParserParserTK_unit       = 62
	ParserParserTK_type       = 63
	ParserParserTK_2fs        = 64
	ParserParserTK_3fs        = 65
	ParserParserTK_number     = 66
	ParserParserTK_path       = 67
	ParserParserTK_equ        = 68
	ParserParserCOMMENTARY    = 69
)

// ParserParser rules.
const (
	ParserParserRULE_init     = 0
	ParserParserRULE_commands = 1
	ParserParserRULE_command  = 2
	ParserParserRULE_execute  = 3
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
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(8)

			var _x = p.Commands()

			localctx.(*InitContext).c = _x
		}
		{
			p.SetState(9)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = localctx.(*InitContext).GetC().GetResult()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(12)
			p.Match(ParserParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*InitContext).result = []interfaces.Command{}

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

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() ICommandContext

	// SetC sets the c rule contexts.
	SetC(ICommandContext)

	// GetResult returns the result attribute.
	GetResult() []interfaces.Command

	// GetCmds returns the cmds attribute.
	GetCmds() []interfaces.Command

	// SetResult sets the result attribute.
	SetResult([]interfaces.Command)

	// SetCmds sets the cmds attribute.
	SetCmds([]interfaces.Command)

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result []interfaces.Command
	cmds   []interfaces.Command // TODO = []interfaces.Command{}
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

func (s *CommandsContext) GetC() ICommandContext { return s.c }

func (s *CommandsContext) SetC(v ICommandContext) { s.c = v }

func (s *CommandsContext) GetResult() []interfaces.Command { return s.result }

func (s *CommandsContext) GetCmds() []interfaces.Command { return s.cmds }

func (s *CommandsContext) SetResult(v []interfaces.Command) { s.result = v }

func (s *CommandsContext) SetCmds(v []interfaces.Command) { s.cmds = v }

func (s *CommandsContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
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
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParserParserRULE_commands)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParserParserRW_execute {
		{
			p.SetState(16)

			var _x = p.Command()

			localctx.(*CommandsContext).c = _x
		}
		localctx.(*CommandsContext).cmds = append(localctx.(*CommandsContext).cmds, localctx.(*CommandsContext).GetC().GetResult())

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	localctx.(*CommandsContext).result = localctx.(*CommandsContext).cmds

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

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC returns the c rule contexts.
	GetC() IExecuteContext

	// SetC sets the c rule contexts.
	SetC(IExecuteContext)

	// GetResult returns the result attribute.
	GetResult() interfaces.Command

	// SetResult sets the result attribute.
	SetResult(interfaces.Command)

	// Getter signatures
	Execute() IExecuteContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result interfaces.Command
	c      IExecuteContext
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

func (s *CommandContext) GetC() IExecuteContext { return s.c }

func (s *CommandContext) SetC(v IExecuteContext) { s.c = v }

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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)

		var _x = p.Execute()

		localctx.(*CommandContext).c = _x
	}
	localctx.(*CommandContext).result = localctx.(*CommandContext).GetC().GetResult()

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

	// GetParams returns the params attribute.
	GetParams() map[string]string

	// SetResult sets the result attribute.
	SetResult(*commands.Execute)

	// SetParams sets the params attribute.
	SetParams(map[string]string)

	// Getter signatures
	RW_execute() antlr.TerminalNode
	RW_path() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_path() antlr.TerminalNode

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	result *commands.Execute
	params map[string]string // TODO = map[string]string{}
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

func (s *ExecuteContext) GetParams() map[string]string { return s.params }

func (s *ExecuteContext) SetResult(v *commands.Execute) { s.result = v }

func (s *ExecuteContext) SetParams(v map[string]string) { s.params = v }

func (s *ExecuteContext) RW_execute() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_execute, 0)
}

func (s *ExecuteContext) RW_path() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_path, 0)
}

func (s *ExecuteContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)

		var _m = p.Match(ParserParserRW_execute)

		localctx.(*ExecuteContext).e = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserRW_path {
		{
			p.SetState(30)
			p.Match(ParserParserRW_path)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)

			var _m = p.Match(ParserParserTK_path)

			localctx.(*ExecuteContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExecuteContext).params["path"] = (func() string {
			if localctx.(*ExecuteContext).GetP() == nil {
				return ""
			} else {
				return localctx.(*ExecuteContext).GetP().GetText()
			}
		}())

	}
	localctx.(*ExecuteContext).result = commands.NewExecute(localctx.(*ExecuteContext).params, (func() int {
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
