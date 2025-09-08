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
	XFound   = NewWithDefaultField("X_found", http.StatusFound)
	XUpdated = NewWithDefaultField("X_updated", http.StatusOK)
	XDeleted = NewWithDefaultField("X_deleted", http.StatusOK)
	XCreated = NewWithDefaultField("x_created", http.StatusCreated)
	XChanged = NewWithDefaultField("x_changed", http.StatusOK)
	XReset   = NewWithDefaultField("x_reset", http.StatusOK)
)
