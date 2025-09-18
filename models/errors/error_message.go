package errors

type ErrorCode string

const (
	UnauthorizedErrorCode        ErrorCode = "UNAUTHORIZED"
	InternalServerErrorErrorCode ErrorCode = "INTERNAL_SERVER_ERROR"
	NotFoundErrorCode            ErrorCode = "NOT_FOUND"
	BadRequestErrorCode          ErrorCode = "BAD_REQUEST"
)

const InternalServerErrorMessage string = "Internal server error"
const AuthorizationErrorMessage string = "Access Denied"

type ErrorMessage struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

func NewErrorMessage(code string, message string) ErrorMessage {
	return ErrorMessage{
		Code:    code,
		Message: message,
	}
}
