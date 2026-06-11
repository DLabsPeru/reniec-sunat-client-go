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
