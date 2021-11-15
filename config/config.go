package config

type Config struct {
	Addr     string // "REDIS_HOST"
	Password string // "REDIS_PASSWORD"
	Db       int    // "REDIS_DATABASE"

	SessionCookieName string // SESSION_COOKIE_NAME
	SessionHashKey    string // SESSION_HASH_KEY
}
