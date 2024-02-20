package language

import (
	interfaces "mia/Classes/Interfaces"
	parser "mia/Language/Parser"
)

type MIAListener struct {
	*parser.BaseParserListener
	Code []interfaces.Command
}

func NewMIAListener() *MIAListener {
	return new(MIAListener)
}

func (l *MIAListener) ExitInit(ctx *parser.InitContext) {
	l.Code = ctx.GetResult()
}
