package success

import "net/http"

var (
	// Auth
	SessionRefreshed = New("session_refreshed", http.StatusOK)
	Registered       = New("registered", http.StatusCreated)
	LoggedIn         = New("logged_in", http.StatusOK)
	LoggedOut        = New("logged_out", http.StatusOK)
	PasswordChanged  = New("password_changed", http.StatusOK)
	PasswordReset    = New("password_reset", http.StatusOK)

	// User
	UserUpdated = New("user_updated", http.StatusOK)
	UserDeleted = New("user_deleted", http.StatusOK)
	UserFound   = New("user_found", http.StatusFound)
)
