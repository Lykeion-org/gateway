package models


type ErrorResponse struct {
	Error 					ErrorBody	`json:"error"`
}

type ErrorBody struct {
	Code					ErrorCode	`json:"code"`
	Message					string		`json:"message"`
}

type ErrorCode string
const (
	ErrInvalidRequest 		ErrorCode = "INVALID_REQUEST"
	ErrInternal				ErrorCode = "INTERNAL_SERVER_ERROR"
	ErrUnauthorized			ErrorCode = "UNAUTHORIZED_REQUEST"
	ErrForbidden			ErrorCode = "NO_PERMISSION"
	ErrConflict				ErrorCode = "CONFLICT"
)