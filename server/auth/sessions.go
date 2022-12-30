package auth

import (
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

var Store *sessions.CookieStore

//= sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func InitSessions() {
	key := viper.GetString("session_key")

	//logger.Log.Info("cookie",zap.ByteString("cookie", key))
	Store = sessions.NewCookieStore([]byte(key))
}
