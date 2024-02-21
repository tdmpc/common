package errors

// Wrap code为自定义错误码
func Wrap(err error, code int) *Error {
	return &Error{
		Cause:   err,
		Code:    code,
		Message: message[code],
		ChMsg:   messageCh[code],
	}
}

func Wrap404Error(err error) *Error            { return Wrap(err, CodeResourcesNotFount) }
func WrapInternalServerError(err error) *Error { return Wrap(err, CodeInternalServerError) }
func WrapInvalidParamsError(err error) *Error  { return Wrap(err, CodeInvalidParams) }
func WrapUnauthorizedError(err error) *Error   { return Wrap(err, CodeUnauthorized) }
func WrapForBiddenError(err error) *Error      { return Wrap(err, CodeForbidden) }
