package errors

import "errors"

// Common errors
var (
	InternalServerError     = errors.New("внутрення ошибка сервера")
	ErrInvalidRequestFormat = errors.New("неверный формат запроса")
)

// ErrTaskNotFound Task errors
var (
	ErrTaskNotFound = errors.New("задача не найдена")
)

// ErrUserNotFound User errors
var (
	ErrUserNotFound      = errors.New("пользователь не найден")
	ErrUserAlreadyExists = errors.New("пользовательн уже существует")
)
