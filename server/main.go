package main

import (
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/auth"
	"github.com/adrianprayoga/noleftovers/server/internals/configs"
	logger "github.com/adrianprayoga/noleftovers/server/internals/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	//"github.com/adrianprayoga/noleftovers/server/configs"
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
	configs.InitializeViper()
	logger.InitializeZapCustomLogger()
	auth.InitializeOAuthGoogle()
	auth.InitSessions()

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.Log.Info("Connected to database db cfg", zap.Any("cfg", cfg))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{viper.GetString("client_host")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	var recipeResource controllers.RecipeResource
	recipeResource.Service = &models.RecipeService{
		DB: db,
	}
	recipeResource.ImageService = &models.ImageService{
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

	var authResource controllers.AuthResource
	authResource.Service = &models.UserService{
		DB: db,
	}

	var favoritesResource controllers.FavoritesResource
	favoritesResource.Service = &models.FavoritesService{
		DB: db,
	}

	r.Mount("api/recipe", recipeResource.Routes())
	r.Mount("api/measures", measureResource.Routes())
	r.Mount("api/images", imageResource.Routes())
	r.Mount("api/auth", authResource.Routes())
	r.Mount("api/favorites", favoritesResource.Routes())

	fmt.Println("Starting the server on :" + viper.GetString("port"))
	http.ListenAndServe(":"+viper.GetString("port"), r)
}
