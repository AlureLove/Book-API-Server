package exception

import "fmt"

type ApiException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

func (e *ApiException) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func (e *ApiException) IsException(target *ApiException) bool {
	return e.Code == target.Code
}
