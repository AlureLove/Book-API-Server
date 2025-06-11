package exception

import "fmt"

type ApiException struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

func (e *ApiException) WithHttpCode(httpCode int) *ApiException {
	e.HttpCode = httpCode
	return e
}

func (e *ApiException) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
