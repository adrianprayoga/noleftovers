package models

import (
	"database/sql"
	"fmt"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	Id    uint `json:"id"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	AuthMethod string `json:"auth_method"`
	OauthId string `json:"oauth_id"`
	LastLogin string `json:"last_login"`
	Picture string `json:"picture"`
}

type NewUser struct {
	Email string
	Password string
	AuthMethod string
	OauthId string
	Picture string
}

type UserService struct {
	DB *sql.DB
}

type UserValidator struct {
}

func (us *UserService) CreateOrUpdateByOauth(nu NewUser) (*User, error) {
	email := strings.ToLower(nu.Email)

	user := User{
		Email: email,
		AuthMethod: nu.AuthMethod,
		OauthId: nu.OauthId,
		Picture: nu.Picture,
	}

	exists := true
	err := us.DB.QueryRow(`SELECT id FROM users WHERE email=$1`, user.Email).Scan(&user.Id)

	if err != nil && err != sql.ErrNoRows {
		logger.Log.Error("Error checking whether user exists")
		logger.Log.Error("", zap.Error(err))
	} else if err != nil && err == sql.ErrNoRows {
		exists = false
	}

	if exists {
		logger.Log.Info("updating last login time")
		_, err = us.DB.Exec(`UPDATE users SET last_login = $1 WHERE email=$2`, time.Now(), &user.Email)
		if err != nil {
			logger.Log.Error("Error when updating user", zap.Error(err))
			return nil, fmt.Errorf("update user: %w", err)
		}
	} else {
		logger.Log.Info("creating user")
		row := us.DB.QueryRow(`INSERT INTO users (email, auth_method, oauth_id, picture) 
								 VALUES ($1, $2, $3, $4) RETURNING id`,
			user.Email, user.AuthMethod, user.OauthId, user.Picture)
		err = row.Scan(&user.Id)
		if err != nil {
			logger.Log.Error("Error when creating user", zap.Error(err))
			return nil, fmt.Errorf("create user: %w", err)
		}
	}

	return &user, nil
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
	err = row.Scan(&user.Id)
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
	err := row.Scan(&user.Id, &user.PasswordHash)
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