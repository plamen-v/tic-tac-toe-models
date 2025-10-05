package models

import (
	"fmt"

	"github.com/pkg/errors"
)

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
	return errors.WithStack(&NotFoundError{
		Message: msg})
}

func NewNotFoundErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return errors.WithStack(&NotFoundError{
		Message: msg})
}

func IsNotFoundError(err error) bool {
	var testError *NotFoundError
	return errors.As(err, &testError)
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidationError(msg string) error {
	return errors.WithStack(&ValidationError{
		Message: msg})
}

func NewValidationErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return errors.WithStack(&ValidationError{
		Message: msg})
}

func IsValidationErrorf(err error) bool {
	var testError *ValidationError
	return errors.As(err, &testError)
}

type GenericError struct {
	Message string
	Code    ErrorCode
}

func (e *GenericError) Error() string {
	return e.Message
}

func NewGenericError(msg string) error {
	return errors.WithStack(&GenericError{
		Message: msg})
}

func NewGenericErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return errors.WithStack(&GenericError{
		Message: msg})
}

func IsGenericError(err error) bool {
	var testError *GenericError
	return errors.As(err, &testError)
}

type AuthorizationError struct {
	Message string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

func NewAuthorizationError(msg string) error {
	return errors.WithStack(&AuthorizationError{
		Message: msg})
}

func NewAuthorizationErrorf(format string, arg ...any) error {
	msg := fmt.Sprintf(format, arg...)
	return errors.WithStack(&AuthorizationError{
		Message: msg})
}

func IsAuthorizationError(err error) bool {
	var testError *AuthorizationError
	return errors.As(err, &testError)
}
