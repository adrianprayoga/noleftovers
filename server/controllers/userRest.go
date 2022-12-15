package controllers

import (
	"encoding/json"
	"github.com/adrianprayoga/noleftovers/server/auth"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type UserResource struct {
	Service *models.UserService
}

func (rs UserResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Put("/", rs.Update)

	return r
}

type updateUserRequest struct {
	FullName string `json:"full_name" validate:"required"`
}

func (rs UserResource) Update(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		http.Error(w, "user is not authenticated", http.StatusUnauthorized)
	}

	var req updateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Log.Error("error", zap.Error(err))
		http.Error(w, "Unable to unmarshal request", http.StatusBadRequest)
		return
	}

	_, err = rs.Service.UpdateUserDetails(models.User{
		Id: userId,
		FullName: req.FullName,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, _ := auth.Store.Get(r, "session-name")
	session.Values["fullName"] = req.FullName
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
