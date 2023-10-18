package model

const (
	ECodeInternal      = 50001 // for error when internal server error
	ECodeBadRequest    = 40001 // for error when failed decode request
	ECodeNotFound      = 40002 // for error when resource not found
	ECodeValidateFail  = 40003 // for error when request is failed validation
	ECodeMethodFail    = 40004 // for error when request is method not allowed
	ECodeDataExists    = 40005 // for error when data is duplicate or exists
	ECodeAuthorization = 40006 // for error when request need authorization
	ECodeForbidden     = 40007 // for error when request cannot be access because some reason
)
