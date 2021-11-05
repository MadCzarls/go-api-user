package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/mad-czarls/go-api-user/container"
)

func SetUpSession() *sessions.CookieStore {
	envManager := container.GetEnvManager()
	store := sessions.NewCookieStore([]byte(*envManager.GetEnvString("SESSION_HASH_KEY")))

	return &store
}
