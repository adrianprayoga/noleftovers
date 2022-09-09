package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type favoritesRequest struct {
	RecipeId uint `json:"recipe_id" validate:"required"`
}

type FavoritesResource struct {
	Service *models.FavoritesService
}

func (rs FavoritesResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /posts - Read a list of posts.
	r.Post("/", rs.Create) // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(GetFavCtx)         // Use the same middleware
		r.Delete("/", rs.Delete) // Delete /posts/{id} - Read a single post by :id.
	})

	return r
}

func GetFavCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs FavoritesResource) List(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	logger.Log.Info("userId", zap.Uint("userId", userId))
	if err != nil {
		http.Error(w, "user is not authenticated", http.StatusUnauthorized)
	}

	recipeIds, err := rs.Service.GetFavoritesByUserId(userId)
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(recipeIds)

	res, _ := json.Marshal(recipeIds)
	_, err = w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rs FavoritesResource) Create(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		http.Error(w, "user is not authenticated", http.StatusUnauthorized)
	}

	var req favoritesRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to unmarshal request", http.StatusBadRequest)
		return
	}

	err = rs.Service.AddFavoritesForUserId(userId, req.RecipeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (rs FavoritesResource) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	logger.Log.Info("userId", zap.Uint("userId", userId))
	if err != nil {
		http.Error(w, "user is not authenticated", http.StatusUnauthorized)
		return
	}

	recipeId, err := strconv.ParseUint(r.Context().Value("id").(string), 10, 31)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = rs.Service.RemoveFavoritesForUserId(userId, uint(recipeId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
