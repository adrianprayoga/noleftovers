package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adrianprayoga/noleftovers/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type createRecipeRequest struct {
	Name        string                    `json:"name" validate:"required"`
	Description string                    `json:"description" validate:"required,min=1,max=400"`
	Author      uint                      `json:"author"`
	Steps       []createStepsRequest      `json:"steps" validate:"required,dive,min=1"`
	Ingredients []createIngredientRequest `json:"ingredients" validate:"required,dive,min=1"`
}

type createStepsRequest struct {
	Text string `json:"text" validate:"required,min=1"`
}

type createIngredientRequest struct {
	Name    string      `json:"name" validate:"required"`
	Amount  json.Number `json:"amount"`
	Measure json.Number `json:"measure"`
}

type ApiError struct {
	Field        string `json:"field"`
	ErrorMessage string `json:"errorMessage"`
}

type RecipeResource struct {
	Service      *models.RecipeService
	ImageService *models.ImageService
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
	recipes, err := rs.Service.GetRecipes()
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

func (r createRecipeRequest) IsValid() error {
	// Note: https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835

	v := validator.New()
	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	err := v.Struct(r)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
	}

	return err
}

func (rs RecipeResource) Create(w http.ResponseWriter, r *http.Request) {

	var recipe createRecipeRequest
	decoder := json.NewDecoder(strings.NewReader(r.FormValue("recipe")))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = recipe.IsValid(); err != nil {
		// TODO: refactor out
		var out []ApiError
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out = make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), fe.Error()}
			}
		}
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(out)
		_, err = w.Write(res)
		return
	}

	var maxUploadSize int64 = 10 << 20 // 10 binary shifted 20 times. 10 * 2^20 = 10 MB

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "The uploaded image is too big. Please use an image less than 10MB in size", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", fileHeader.Filename)
	fmt.Printf("File Size: %+v\n", fileHeader.Size)

	errCode, err := rs.ImageService.ValidateImage(file, fileHeader)
	if err != nil {
		http.Error(w, err.Error(), errCode)
		return
	}

	id := uuid.New().String()
	fileName := fmt.Sprintf("%s%s", id, filepath.Ext(fileHeader.Filename))
	if err != rs.ImageService.UploadImage(file, fileName) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Successfully Uploaded File")

	// Map to domain entity
	rec := models.Recipe{
		Name:        recipe.Name,
		Description: recipe.Description,
		Author:      recipe.Author,
		ImageLink:   fileName,
	}

	ingredients := make([]models.Ingredient, len(recipe.Ingredients))
	for i, ing := range recipe.Ingredients {
		ingredients[i] = models.Ingredient{
			Name: ing.Name,
		}

		// TODO: handle error
		amt, err := ing.Amount.Float64()
		if err == nil {
			ingredients[i].Amount = float32(amt)
		}

		m, err := ing.Measure.Float64()
		if err == nil {
			ingredients[i].Measure = uint(m)
		}
	}

	steps := make([]models.RecipeStep, len(recipe.Steps))
	for i, s := range recipe.Steps {
		steps[i] = models.RecipeStep{
			Text: s.Text,
		}
	}

	newRecipe, err := rs.Service.CreateRecipe(rec, steps, ingredients)
	if err != nil {
		fmt.Println("Something went wrong when creating an object")
		rs.ImageService.RemoveImage(fileName)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(newRecipe)
	_, err = w.Write(res)
	w.WriteHeader(http.StatusOK)
}

func (rs RecipeResource) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.Context().Value("id").(string), 10, 31)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	recipe, err := rs.Service.GetRecipeById(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(recipe)

	_, err = w.Write(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
