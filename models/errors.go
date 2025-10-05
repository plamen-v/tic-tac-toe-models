package models

import "fmt"

type ErrorCode string

const (
	UnauthorizedErrorCode        ErrorCode = "UNAUTHORIZED"
	InternalServerErrorErrorCode ErrorCode = "INTERNAL_SERVER_ERROR"
	NotFoundErrorCode            ErrorCode = "NOT_FOUND"
	BadRequestErrorCode          ErrorCode = "BAD_REQUEST"
)

const InternalServerErrorMessage string = "Internal server error"
const AuthorizationErrorMessage string = "Access Denied"

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(msg string) error {
	return &NotFoundError{
		Message: msg}
}

func NewNotFoundErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return &NotFoundError{
		Message: msg}
}

func IsNotFoundError(err error) bool {
	if _, ok := err.(*NotFoundError); ok {
		return true
	}
	return false
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(msg string) error {
	return &ValidationError{
		Message: msg}
}

func NewValidationErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return &ValidationError{
		Message: msg}
}

func IsValidationErrorf(err error) bool {
	if _, ok := err.(*ValidationError); ok {
		return true
	}
	return false
}

type GenericError struct {
	Message string
	Code    ErrorCode
}

func (e *GenericError) Error() string {
	return e.Message
}

func NewGenericError(msg string) error {
	return &GenericError{
		Message: msg}
}

func NewGenericErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return &GenericError{
		Message: msg}
}

func IsGenericError(err error) bool {
	if _, ok := err.(*GenericError); ok {
		return true
	}
	return false
}

type AuthorizationError struct {
	Message string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

func NewAuthorizationError(msg string) error {
	return &AuthorizationError{
		Message: msg}
}

func NewAuthorizationErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return &AuthorizationError{
		Message: msg}
}

func IsAuthorizationError(err error) bool {
	if _, ok := err.(*AuthorizationError); ok {
		return true
	}
	return false
}
