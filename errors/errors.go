package errors

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
)

type logger func(format string, args ...interface{})

// CustomError 自定义错误，替代go原生的err，避免层层返回层层判断，发生error时直接Bomb，写入code与message方便recover时解析
type CustomError struct {
	Cause   error  `json:"error"`   // 用来装入原始error，不影响对原始错误的判断（如if err == sql.ErrNoRows的情况），可以为nil
	Code    int    `json:"code"`    // 自定义的错误码
	Message string `json:"message"` // 自定义的错误信息
}

// Error 满足go原生error interface
func (err CustomError) Error() string {
	if err.Cause != nil {
		return err.Cause.Error()
	}

	return err.Message
}

// New 创建CustomError
func New(err error, code int, message string, args ...interface{}) *CustomError {
	msg := fmt.Sprintf(message, args...)
	return &CustomError{
		Cause:   err,
		Code:    code,
		Message: msg,
	}
}

// Bomb 直接panic，like a bomb！
// 注意：调用panic()会比 return err 慢，所以使用Bomb的error都应是不被容忍的错误，是须后续修正的错误。
func Bomb(err error, code int, format string, args ...interface{}) {
	panic(New(err, code, fmt.Sprintf(format, args...)))
}

// Recover 捕获错误并recover，须使用defer调用，handler传入nil时默认打印黄色错误信息到stdin
func Recover(handler logger) {
	if handler == nil {
		handler = color.Yellow // 使用黄色可以在调试时方便确认这是自定义的CustomError（与大多数红色打印的panic区分开）
	}

	err := recover()
	switch err.(type) {
	case *CustomError:
		stack := string(debug.Stack())
		handler("[Recovery] %s panic recovered:\n%+v\n%s",
			time.Now().Format("2006/01/02 - 15:04:05"), err, stack)
	default:
		// 除了自定义的CustomError，其余的panic维持原状
		panic(err)
	}
}
