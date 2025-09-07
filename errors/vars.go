package apperrors

import (
	"net/http"
)

// General errors
var (
	ErrUnauthorized        = New("unauthorized", http.StatusUnauthorized)
	ErrRequestFailed       = New("request_failed", http.StatusBadGateway)
	ErrSessionExpired      = New("session_expired", http.StatusUnauthorized)
	ErrBadRequest          = New("bad_request", http.StatusBadRequest)
	ErrInternalServerError = New("internal_server_error", http.StatusInternalServerError)
	ErrTooManyRequests     = New("too_many_requests", http.StatusTooManyRequests)
	ErrRegistrationFailed  = New("registration_failed", http.StatusInternalServerError)
	ErrLoginFailed         = New("login_failed", http.StatusInternalServerError)
)

// Field-related errors
var (
	ErrXNotFound   = NewWithDefaultField("x_not_found", http.StatusNotFound)
	ErrXIsRequired = NewWithDefaultField("x_is_required", http.StatusBadRequest)
	ErrXConflict   = NewWithDefaultField("x_already_exists", http.StatusConflict)
	ErrInvalidX    = NewWithDefaultField("invalid_x", http.StatusBadRequest)
	ErrIncorrectX  = NewWithDefaultField("incorrect_x", http.StatusBadRequest)
)
