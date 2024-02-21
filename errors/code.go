package errors

import "net/http"

// 自定义错误类型与http status code是多对一的关系，所以要自定义错误码
const (
	CodeOK                  = 0     // 200 成功ok
	CodeInternalServerError = 10000 // 500 服务内部失败
	CodeUnauthorized        = 10001 // 401 用户未传token
	CodeForbidden           = 10002 // 403 鉴权失败，如token无效或者过期
	CodeInvalidParams       = 10003 // 400 参数错误
	CodeResourcesNotFount   = 10004 // 404 资源未找到
	CodeResourcesHasExist   = 10005 // 409 资源已存在
	CodeResourcesConflict   = 10006 // 409 状态冲突
	CodeUnknownError        = 10007 // 500 未知异常
	CodeNoRight2Modify      = 10008 // 403 用户没权限修改相关资源
)

var (
	// 自定义code和http status code的对应关系
	statusCode = map[int]int{
		CodeOK:                  http.StatusOK,
		CodeInternalServerError: http.StatusInternalServerError,
		CodeUnauthorized:        http.StatusUnauthorized,
		CodeForbidden:           http.StatusForbidden,
		CodeInvalidParams:       http.StatusBadRequest,
		CodeResourcesNotFount:   http.StatusNotFound,
		CodeResourcesHasExist:   http.StatusConflict,
		CodeResourcesConflict:   http.StatusConflict,
		CodeUnknownError:        http.StatusInternalServerError,
		CodeNoRight2Modify:      http.StatusForbidden,
	}
)

// StatusCode 使用自定义code取得http status code
func StatusCode(code int) int {
	return statusCode[code]
}
