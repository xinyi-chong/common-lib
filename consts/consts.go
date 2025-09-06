package consts

const (
	Localizer = "localizer"
)

// Context keys used in Gin/Go contexts
const (
	CtxAccessToken = "access_token"
	CtxUserID      = "user_id"
	CtxUsername    = "username"
	CtxUserEmail   = "email"
)

// Standardized headers for forwarding user claims to downstream services
const (
	HeaderUserID    = "X-User-Id"
	HeaderUsername  = "X-User-Name"
	HeaderUserEmail = "X-User-Email"
)
