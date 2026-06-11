package reniecsunatclient

import "fmt"

type Error struct {
	StatusCode int
	Code       string
	Message    string
	RequestID  string
	Path       string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.Code != "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}

	return e.Message
}

func (e *Error) IsNotFound() bool {
	return e != nil && e.StatusCode == 404
}

func (e *Error) IsBadRequest() bool {
	return e != nil && e.StatusCode == 400
}

func (e *Error) IsBadGateway() bool {
	return e != nil && e.StatusCode == 502
}
