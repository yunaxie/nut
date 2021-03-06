package iterf

type Iterator interface {
	Foreach(func(...interface{}))
	ForeachBreak(func(...interface{}) bool, func(...interface{})) interface{}
}
