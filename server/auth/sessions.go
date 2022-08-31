package auth

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore
//= sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func InitSessions() {
	key := "key"// securecookie.GenerateRandomKey(32)

	//logger.Log.Info("cookie",zap.ByteString("cookie", key))
	Store = sessions.NewCookieStore([]byte(key))
}
