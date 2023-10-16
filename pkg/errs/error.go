package errs

import "fmt"

type Attribute struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code       uint16      `json:"error_code"`
	Message    string      `json:"error_message"`
	Err        error       `json:"-"`
	Attributes []Attribute `json:"attributes,omitempty"`
}

func New(code uint16, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

func NewWithAttribute(code uint16, msg string, attrs []Attribute) *Error {
	return &Error{Code: code, Message: msg, Attributes: attrs}
}

func NewWithErr(code uint16, msg string, err error) *Error {
	return &Error{Code: code, Message: msg, Err: err}
}

func NewFull(code uint16, msg string, err error, attrs []Attribute) *Error {
	return &Error{Code: code, Message: msg, Attributes: attrs}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
