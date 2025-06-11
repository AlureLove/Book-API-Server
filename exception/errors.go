package exception

import "fmt"

const (
	ERR_NOT_FOUND = 404
)

func IsNotFound(err error) bool {
	if v, ok := (err).(*ApiException); ok {
		return v.Code == ERR_NOT_FOUND
	}

	return false
}

func ErrNotFound(format string, a ...any) *ApiException {
	return NewApiException(ERR_NOT_FOUND, fmt.Sprintf(format, a...))
}
