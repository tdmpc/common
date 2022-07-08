package errors

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	unknown = "unknown\n"
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
	case nil:
		return
	case *CustomError:
		stack := getStack()
		handler("[Recovery] %s panic recovered:\n%+v\n%s",
			time.Now().Format("2006/01/02 - 15:04:05"), err, stack)
	default:
		// 除了自定义的CustomError，其余的panic维持原状
		panic(err)
	}
}

func getStack() string {
	const depth = 6 // 回溯最近6个函数栈信息
	var pcs [depth]uintptr
	var buf strings.Builder

	n := runtime.Callers(3, pcs[:]) // skip = 3，跳过Recover的函数调用栈信息
	// 获取文件名、函数所在行、函数名信息
	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pcs[i])
		if fn == nil {
			buf.WriteString(unknown)
			continue
		}

		filename, line := fn.FileLine(pcs[i])
		funcname := getFuncname(fn.Name())
		fmt.Fprintf(&buf, "%s:%d (0x%x)\n\t%s\n", filename, line, pcs[i], funcname)
	}

	return buf.String()
}

func getFuncname(name string) string {
	// 不打印函数所在的文件的路径信息
	index := strings.LastIndex(name, "/")
	name = name[index+1:]
	// 不打印包名
	index = strings.Index(name, ".")
	return name[index+1:]
}
