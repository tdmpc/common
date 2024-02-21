package errors

var (
	message = map[int]string{
		CodeOK:                  "ok",
		CodeInternalServerError: "internal server error",
		CodeUnauthorized:        "header authorization is empty",
		CodeForbidden:           "token is forbidden",
		CodeInvalidParams:       "invalid param",
		CodeResourcesNotFount:   "the resource was not found",
		CodeResourcesHasExist:   "the resource has already exists",
		CodeResourcesConflict:   "the resource has conflict expect status value",
		CodeUnknownError:        "unknown error",
		CodeNoRight2Modify:      "have no right to modify resource",
	}

	messageCh = map[int]string{
		CodeOK:                  "成功",
		CodeUnauthorized:        "认证token不能为空",
		CodeInvalidParams:       "无效参数",
		CodeForbidden:           "鉴权失败，token无效或者过期",
		CodeResourcesNotFount:   "资源未找到",
		CodeResourcesHasExist:   "资源已存在",
		CodeResourcesConflict:   "资源冲突",
		CodeUnknownError:        "未知错误",
		CodeInternalServerError: "系统错误",
	}
)

func ErrorMessage(code int) string {
	return message[code]
}

func ErrorMessageCh(code int) string {
	return messageCh[code]
}
