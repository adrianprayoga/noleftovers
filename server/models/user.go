package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	ID 	uint
	Email string
	PasswordHash string
}

type NewUser struct {
	Email string
	Password string
}

type UserService struct {
	DB *sql.DB
}

type UserValidator struct {
}


func (us *UserService) Create(nu NewUser) (*User, error) {
	email := strings.ToLower(nu.Email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email: email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`INSERT INTO users (email, password_hash) 
								 VALUES ($1, $2) RETURNING id`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}

func (us *UserService) Authenticate(email string, password string) (*User, error) {
	fmt.Println(email, password)

	email = strings.ToLower(email)
	user := User{
		Email: email,
	}
	row := us.DB.QueryRow(`SELECT id, password_hash FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		fmt.Println("Error authenticating")
		return nil, fmt.Errorf("user login: %w", err)
	}

	fmt.Printf("User %+v\n", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("login issue: %w", err)
	}

	return &user, nil
}