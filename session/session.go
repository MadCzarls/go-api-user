package session

import "github.com/gin-gonic/contrib/sessions"

func SetUpSession() sessions.CookieStore {
	//@TODO put params in ENVs
	store := sessions.NewCookieStore([]byte("secret_for_session_hashing"))

	return store
}
