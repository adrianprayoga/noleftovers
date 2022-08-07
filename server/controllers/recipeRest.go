package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

//type album struct {
//	ID     string  `json:"id"`
//	Title  string  `json:"title"`
//	Artist string  `json:"artist"`
//	Price  float64 `json:"price"`
//}
//
//var albums = []album{
//	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}

type RecipeResource struct {
	RecipeService *models.RecipeService
}

func (rs RecipeResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.List)    // GET /posts - Read a list of posts.
	r.Post("/", rs.Create) // POST /posts - Create a new post.

	r.Route("/{id}", func(r chi.Router) {
		r.Use(GetCtx)      // Use the same middleware
		r.Get("/", rs.Get) // GET /posts/{id} - Read a single post by :id.
	})

	return r
}

func GetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rs RecipeResource) List(w http.ResponseWriter, r *http.Request) {
	recipes, err := rs.RecipeService.GetRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(recipes)
	_, err = w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rs RecipeResource) Create(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(recipe)

	_, err = rs.RecipeService.CreateRecipe(recipe)
	if err != nil {
		fmt.Println("Something went wrong when creating an object")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (rs RecipeResource) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.Context().Value("id").(string), 10, 31)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	recipe, err := rs.RecipeService.GetRecipeById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, _ := json.Marshal(recipe)
	_, err = w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
