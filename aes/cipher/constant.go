package cipher

type Type int64

const (
	CBC Type = iota //CBC 模式
	CFB             //CFB 模式
	OFB             //
)
