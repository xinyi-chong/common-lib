package apperrors

import (
	"net/http"
)

var (
	ErrUnauthorized             = New("unauthorized", http.StatusUnauthorized)
	ErrRequestFailed            = New("request_failed", http.StatusBadGateway)
	ErrSessionExpired           = New("session_expired", http.StatusUnauthorized)
	ErrBadRequest               = New("bad_request", http.StatusBadRequest)
	ErrInternalServerError      = New("internal_server_error", http.StatusInternalServerError)
	ErrTooManyRequests          = New("too_many_requests", http.StatusTooManyRequests)
	ErrRegistrationFailed       = New("registration_failed", http.StatusInternalServerError)
	ErrLoginFailed              = New("login_failed", http.StatusInternalServerError)
	ErrInvalidEmail             = New("invalid_email", http.StatusBadRequest)
	ErrIncorrectEmailOrPassword = New("incorrect_email_or_password", http.StatusUnauthorized)
	ErrIncorrectPassword        = New("incorrect_password", http.StatusUnauthorized)
	ErrEmailConflict            = New("email_already_exists", http.StatusConflict)
	ErrUserNotFound             = New("user_not_found", http.StatusNotFound)
	ErrUserConflict             = New("user_already_exists", http.StatusConflict)
	ErrUsernameConflict         = New("username_already_exists", http.StatusConflict)
)
