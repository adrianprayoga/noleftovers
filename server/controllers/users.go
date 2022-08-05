package controllers

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"
	"net/http"
)

type Users struct{
	Templates struct {
		New Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	u.Templates.SignIn.Execute(w, r, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Create(models.NewUser{
		Email: r.PostForm.Get("email"),
		Password: r.PostForm.Get("password")},
	)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) Auth(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form submission.", http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Authenticate(
		r.PostForm.Get("email"),
		r.PostForm.Get("password"))

	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,
		Path:  "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User account authenticated: %v", user)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "The email cookie could not be read.")
		return
	}
	fmt.Fprintf(w, "Email cookie: %s\n", email.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}