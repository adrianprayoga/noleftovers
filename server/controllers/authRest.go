package controllers

import (
	"encoding/json"
	"github.com/adrianprayoga/noleftovers/server/auth"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthResource struct {
	Service *models.UserService
}

func (rs AuthResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/login-gl", rs.HandleGoogleLogin)
	r.Get("/callback-gl", rs.CallBackFromGoogle)
	r.Get("/success", rs.HandleSuccess)
	r.Get("/failed", rs.HandleFailure)

	return r
}

type AuthResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	User models.User `json:"user"`
	Cookie string `json:"cookie"`
}

func (rs AuthResource) HandleSuccess(w http.ResponseWriter, r *http.Request) {
	// if r.user exist, return true


	logger.Log.Info("", zap.String("agent", r.UserAgent()))

	session, err := auth.Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["authenticated"].(bool); ok{
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(AuthResponse{
			Success: true,
			Message: "user has successfully authenticated",
			User: models.User{
				Email: session.Values["email"].(string),
			},
		})
		_, err := w.Write(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "user is unauthenticated", http.StatusUnauthorized)
	}

}

func (rs AuthResource) HandleFailure(w http.ResponseWriter, r *http.Request) {
	// if r.user exist, return true

	if true {
		w.Header().Set("Content-Type", "application/json")
		res, _ := json.Marshal(AuthResponse{
			Success: false,
			Message: "failed to authenticate user",
		})
		_, err := w.Write(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}


func (rs AuthResource) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	auth.HandleLogin(w, r, auth.OauthConfGl, auth.OauthStateStringGl)
}

func (rs AuthResource) CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Callback-gl..")

	state := r.FormValue("state")
	logger.Log.Info(state)
	if state != auth.OauthStateStringGl {
		logger.Log.Info("invalid oauth state, expected " + auth.OauthStateStringGl + ", got " + state + "\n")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	logger.Log.Info(code)

	if code == "" {
		logger.Log.Warn("Code not found..")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("User has denied Permission.."))
		}
		// User has denied access..
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		token, err := auth.OauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			logger.Log.Error("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		logger.Log.Info("TOKEN>> AccessToken>> " + token.AccessToken)
		logger.Log.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
		logger.Log.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)

		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			logger.Log.Error("Get: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Log.Error("ReadAll: " + err.Error() + "\n")
			// TODO change redirect
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		logger.Log.Info("parseResponseBody: " + string(response) + "\n")

		//w.Write([]byte("Hello, I'm protected\n"))
		//w.Write()
		var userResponse auth.GoogleCallbackResponse
		err = json.Unmarshal([]byte(string(response)), &userResponse)
		newUser := models.NewUser{
			Email: userResponse.Email,
			AuthMethod: "Google",
			OauthId: userResponse.Id,
			Picture: userResponse.Picture,
		}

		logger.Log.Info("Creating / updating user based on oauth input")
		_, err = rs.Service.CreateOrUpdateByOauth(newUser)



		cookie := http.Cookie{
			Name:  "email",
			Value: newUser.Email,
			Path:  "/",
		}
		http.SetCookie(w, &cookie)

		session, _ := auth.Store.Get(r, "session-name")
		// Set some session values.
		session.Values["authenticated"] = true
		session.Values["email"] = newUser.Email
		session.Values["picture"] = newUser.Picture

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "http://localhost:3000", http.StatusTemporaryRedirect)

		return
	}
}