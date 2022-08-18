package main

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/controllers"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	var recipeResource controllers.RecipeResource
	recipeResource.Service = &models.RecipeService{
		DB: db,
	}

	var measureResource controllers.MeasureResource
	measureResource.Service = &models.MeasureService{
		DB: db,
	}

	var imageResource controllers.ImageResource
	imageResource.Service = &models.ImageService{
		DB: db,
	}

	r.Mount("/recipe", recipeResource.Routes())
	r.Mount("/measures", measureResource.Routes())
	r.Mount("/images", imageResource.Routes())

	fmt.Println("Starting the server on :7171...")
	http.ListenAndServe(":7171", r)
}
