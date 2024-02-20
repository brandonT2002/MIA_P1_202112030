package interfaces

type Command interface {
	GetLine() int
	GetColumn() int
	Exec()
	GetResult() string
}
