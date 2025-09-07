package success

import "net/http"

// General success
var (
	SessionRefreshed = New("session_refreshed", http.StatusOK)
	Registered       = New("registered", http.StatusCreated)
	LoggedIn         = New("logged_in", http.StatusOK)
	LoggedOut        = New("logged_out", http.StatusOK)
)

// Field-related success
var (
	XFound   = New("X_found", http.StatusFound)
	XUpdated = New("X_updated", http.StatusOK)
	XDeleted = New("X_deleted", http.StatusOK)
	XCreated = New("x_created", http.StatusCreated)
	XChanged = New("x_changed", http.StatusOK)
	XReset   = New("x_reset", http.StatusOK)
)
