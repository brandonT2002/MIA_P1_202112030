package commands

type Execute struct {
	Params map[string]string
	Result string
	Line   int
	Column int
}

func NewExecute(params map[string]string, line, column int) *Execute {
	return &Execute{params, "", line, column}
}

func (exe *Execute) GetLine() int {
	return exe.Line
}

func (exe *Execute) GetColumn() int {
	return exe.Column
}

func (exe *Execute) Exec() {
	//
}

func (exe *Execute) GetResult() string { return "" }
