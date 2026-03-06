package code

import "fmt"

type Code int

const (
	CodeSuccess Code = 0 // 成功
	CodeError   Code = 7 // 通用错误

	CodeUserError     Code = 1000 // 用户错误
	CodeCommonError   Code = 1001 // 通用错误
	CodeBusinessError Code = 1002 // 业务错误
	CodeAccountError  Code = 1003 // 账户错误
)

type CodeMsg struct {
	Code Code
	Msg  string
}

func (c CodeMsg) Error() string {
	return c.Msg
}

func NewError(code Code, msg string) error {
	return CodeMsg{Code: code, Msg: msg}
}

func AccError(format string, a ...any) error {
	return CodeMsg{Code: CodeAccountError, Msg: fmt.Sprintf(format, a...)}
}
