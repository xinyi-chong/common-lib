package consts

const (
	Localizer = "localizer"
)

// Context keys
const (
	CtxAccessToken = "access_token"
	CtxUserID      = "user_id"
	CtxUsername    = "username"
	CtxUserEmail   = "email"
)

// HTTP headers
const (
	HeaderUserID    = "X-User-ID"
	HeaderUsername  = "X-User-Name"
	HeaderUserEmail = "X-User-Email"
)

// Redis prefixes
const (
	RedisAuthBlacklistPrefix = "auth:blacklist:"
)

// Fields
const (
	DefaultField         = "default"
	UserField            = "user"
	EmailField           = "email"
	PasswordField        = "password"
	UsernameField        = "username"
	EmailOrUsernameField = "email_or_username"
)
