package main

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/controllers"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "this is user"+userID)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var resource controllers.RecipeResource
	resource.RecipeService = &models.RecipeService{
		DB: db,
	}

	r.Mount("/recipe", resource.Routes())

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
